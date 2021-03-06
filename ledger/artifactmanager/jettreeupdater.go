//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package artifactmanager

import (
	"context"
	"sync"

	"github.com/insolar/insolar/ledger/storage/node"
	"github.com/pkg/errors"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/jet"
	"github.com/insolar/insolar/insolar/message"
	"github.com/insolar/insolar/insolar/reply"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/instrumentation/instracer"
)

type seqEntry struct {
	ch   chan struct{}
	once sync.Once
}

type seqKey struct {
	pulse insolar.PulseNumber
	jet   insolar.ID
}

type fetchResult struct {
	jet *insolar.ID
	err error
}

type jetTreeUpdater struct {
	Nodes          node.Accessor
	JetStorage     jet.Storage
	MessageBus     insolar.MessageBus
	JetCoordinator insolar.JetCoordinator

	seqMutex  sync.Mutex
	sequencer map[seqKey]*seqEntry
}

func newJetTreeUpdater(
	ans node.Accessor,
	js jet.Storage,
	mb insolar.MessageBus,
	jc insolar.JetCoordinator,
) *jetTreeUpdater {
	return &jetTreeUpdater{
		Nodes:          ans,
		JetStorage:     js,
		MessageBus:     mb,
		JetCoordinator: jc,
		sequencer:      map[seqKey]*seqEntry{},
	}
}

func (jtu *jetTreeUpdater) fetchJet(
	ctx context.Context, target insolar.ID, pulse insolar.PulseNumber,
) (*insolar.ID, error) {
	ctx, span := instracer.StartSpan(ctx, "jet_tree_updater.fetch_jet")
	defer span.End()

	// Look in the local tree. Return if the actual jet found.
	jetID, actual := jtu.JetStorage.ForID(ctx, pulse, target)
	if actual {
		return (*insolar.ID)(&jetID), nil
	}

	// Not actual in our tree, asking neighbors for jet.
	span.Annotate(nil, "tree in DB is not actual")
	key := seqKey{pulse, insolar.ID(jetID)}

	executing := false

	jtu.seqMutex.Lock()
	if _, ok := jtu.sequencer[key]; !ok {
		jtu.sequencer[key] = &seqEntry{ch: make(chan struct{})}
		executing = true
	}
	entry := jtu.sequencer[key]
	jtu.seqMutex.Unlock()

	span.Annotate(nil, "got sequencer entry")

	if !executing {
		<-entry.ch

		// Tree was updated in another thread, rechecking.
		span.Annotate(nil, "somebody else updated actuality")
		return jtu.fetchJet(ctx, target, pulse)
	}

	defer func() {
		entry.once.Do(func() {
			close(entry.ch)
		})

		jtu.seqMutex.Lock()
		delete(jtu.sequencer, key)
		jtu.seqMutex.Unlock()
	}()

	resJet, err := jtu.fetchActualJetFromOtherNodes(ctx, target, pulse)
	if err != nil {
		return nil, err
	}

	jtu.JetStorage.Update(ctx, pulse, true, insolar.JetID(*resJet))

	return resJet, nil
}

func (jtu *jetTreeUpdater) releaseJet(ctx context.Context, jetID insolar.ID, pulse insolar.PulseNumber) {
	jtu.seqMutex.Lock()
	defer jtu.seqMutex.Unlock()

	depth := insolar.JetID(jetID).Depth()
	for {
		key := seqKey{pulse, jetID}
		if v, ok := jtu.sequencer[key]; ok {
			v.once.Do(func() {
				close(v.ch)
			})

			delete(jtu.sequencer, key)
		}

		if depth == 0 {
			break
		}
		jetID = insolar.ID(jet.Parent(insolar.JetID(jetID)))
		depth--
	}
}

func (jtu *jetTreeUpdater) fetchActualJetFromOtherNodes(
	ctx context.Context, target insolar.ID, pulse insolar.PulseNumber,
) (*insolar.ID, error) {
	ctx, span := instracer.StartSpan(ctx, "jet_tree_updater.fetch_jet_from_other_nodes")
	defer span.End()

	ch := make(chan fetchResult, 1)

	go func() {
		nodes, err := jtu.otherNodesForPulse(ctx, pulse)
		if err != nil {
			ch <- fetchResult{nil, err}
			return
		}

		num := len(nodes)

		wg := sync.WaitGroup{}
		wg.Add(num)

		once := sync.Once{}

		replies := make([]*reply.Jet, num)
		for i, node := range nodes {
			go func(i int, node insolar.Node) {
				ctx, span := instracer.StartSpan(ctx, "jet_tree_updater.one_node_get_jet")
				defer span.End()

				defer wg.Done()

				nodeID := node.ID
				rep, err := jtu.MessageBus.Send(
					ctx,
					&message.GetJet{Object: target, Pulse: pulse},
					&insolar.MessageSendOptions{Receiver: &nodeID},
				)
				if err != nil {
					inslogger.FromContext(ctx).Error(
						errors.Wrap(err, "couldn't get jet"),
					)
					return
				}

				r, ok := rep.(*reply.Jet)
				if !ok {
					inslogger.FromContext(ctx).Errorf("middleware.fetchActualJetFromOtherNodes: unexpected reply: %#v\n", rep)
					return
				}

				if !r.Actual {
					return
				}

				once.Do(func() {
					jetID := r.ID
					ch <- fetchResult{&jetID, nil}
					close(ch)
				})

				replies[i] = r
			}(i, node)
		}
		wg.Wait()

		seen := make(map[insolar.ID]struct{})
		res := make([]*insolar.ID, 0)
		for _, r := range replies {
			if r == nil {
				continue
			}
			if _, ok := seen[r.ID]; ok {
				continue
			}

			seen[r.ID] = struct{}{}
			res = append(res, &r.ID)
		}

		if len(res) == 0 {
			inslogger.FromContext(ctx).WithFields(map[string]interface{}{
				"pulse":  pulse,
				"object": target.DebugString(),
			}).Error("all lights for pulse have no actual jet for object")
			ch <- fetchResult{nil, errors.New("impossible situation")}
			close(ch)
		} else if len(res) > 1 {
			inslogger.FromContext(ctx).WithFields(map[string]interface{}{
				"pulse":  pulse,
				"object": target.DebugString(),
			}).Error("lights said different actual jet for object")
		}
	}()

	res := <-ch
	return res.jet, res.err
}

func (jtu *jetTreeUpdater) otherNodesForPulse(
	ctx context.Context, pulse insolar.PulseNumber,
) ([]insolar.Node, error) {
	ctx, span := instracer.StartSpan(ctx, "jet_tree_updater.other_nodes_for_pulse")
	defer span.End()

	res, err := jtu.Nodes.InRole(pulse, insolar.StaticRoleLightMaterial)
	if err != nil {
		return nil, err
	}

	me := jtu.JetCoordinator.Me()
	for i := range res {
		if res[i].ID == me {
			res = append(res[:i], res[i+1:]...)
			break
		}
	}

	num := len(res)
	if num == 0 {
		inslogger.FromContext(ctx).Error(
			"This shouldn't happen. We're solo active light material",
		)

		return nil, errors.New("impossible situation")
	}

	return res, nil
}
