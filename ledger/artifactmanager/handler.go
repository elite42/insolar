/*
 *    Copyright 2019 Insolar Technologies
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package artifactmanager

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"

	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/instrumentation/insmetrics"
	"github.com/insolar/insolar/ledger/recentstorage"
	"github.com/insolar/insolar/ledger/storage/jet"

	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/core/message"
	"github.com/insolar/insolar/core/reply"
	"github.com/insolar/insolar/instrumentation/hack"
	"github.com/insolar/insolar/ledger/storage"
	"github.com/insolar/insolar/ledger/storage/index"
	"github.com/insolar/insolar/ledger/storage/record"
)

// MessageHandler processes messages for local storage interaction.
type MessageHandler struct {
	RecentStorageProvider      recentstorage.Provider          `inject:""`
	Bus                        core.MessageBus                 `inject:""`
	PlatformCryptographyScheme core.PlatformCryptographyScheme `inject:""`
	JetCoordinator             core.JetCoordinator             `inject:""`
	CryptographyService        core.CryptographyService        `inject:""`
	DelegationTokenFactory     core.DelegationTokenFactory     `inject:""`
	HeavySync                  core.HeavySync                  `inject:""`
	PulseStorage               core.PulseStorage               `inject:""`
	JetStorage                 storage.JetStorage              `inject:""`
	DropStorage                storage.DropStorage             `inject:""`
	ObjectStorage              storage.ObjectStorage           `inject:""`
	NodeStorage                storage.NodeStorage             `inject:""`
	PulseTracker               storage.PulseTracker            `inject:""`
	DBContext                  storage.DBContext               `inject:""`
	HotDataWaiter              HotDataWaiter                   `inject:""`

	certificate    core.Certificate
	replayHandlers map[core.MessageType]core.MessageHandler
	conf           *configuration.Ledger
	middleware     *middleware
	jetTreeUpdater *jetTreeUpdater
	isHeavy        bool
}

// NewMessageHandler creates new handler.
func NewMessageHandler(conf *configuration.Ledger, certificate core.Certificate) *MessageHandler {
	return &MessageHandler{
		certificate:    certificate,
		replayHandlers: map[core.MessageType]core.MessageHandler{},
		conf:           conf,
	}
}

func instrumentHandler(name string) Handler {
	return func(handler core.MessageHandler) core.MessageHandler {
		return func(ctx context.Context, p core.Parcel) (core.Reply, error) {
			// TODO: add tags to log
			inslog := inslogger.FromContext(ctx)
			start := time.Now()
			code := "2xx"
			ctx = insmetrics.InsertTag(ctx, tagMethod, name)

			repl, err := handler(ctx, p)

			latency := time.Since(start)
			if err != nil {
				code = "5xx"
				inslog.Errorf("AM's handler %v returns error: %v", name, err)
			}
			inslog.Debugf("measured time of AM method %v is %v", name, latency)

			ctx = insmetrics.ChangeTags(
				ctx,
				tag.Insert(tagMethod, name),
				tag.Insert(tagResult, code),
			)
			stats.Record(ctx, statCalls.M(1), statLatency.M(latency.Nanoseconds()/1e6))

			return repl, err
		}
	}
}

// Init initializes handlers and middleware.
func (h *MessageHandler) Init(ctx context.Context) error {
	m := newMiddleware(h)
	h.middleware = m

	h.jetTreeUpdater = newJetTreeUpdater(h.NodeStorage, h.JetStorage, h.Bus, h.JetCoordinator)

	h.isHeavy = h.certificate.GetRole() == core.StaticRoleHeavyMaterial

	// core.StaticRoleUnknown - genesis
	if h.certificate.GetRole() == core.StaticRoleLightMaterial || h.certificate.GetRole() == core.StaticRoleUnknown {
		h.setHandlersForLight(m)
		h.setReplayHandlers(m)
	}

	if h.isHeavy {
		h.setHandlersForHeavy(m)
	}

	return nil
}

func (h *MessageHandler) setHandlersForLight(m *middleware) {
	// Generic.
	h.Bus.MustRegister(core.TypeGetCode, BuildMiddleware(h.handleGetCode))

	h.Bus.MustRegister(core.TypeGetObject,
		BuildMiddleware(h.handleGetObject,
			instrumentHandler("handleGetObject"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(core.TypeGetDelegate,
		BuildMiddleware(h.handleGetDelegate,
			instrumentHandler("handleGetDelegate"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(core.TypeGetChildren,
		BuildMiddleware(h.handleGetChildren,
			instrumentHandler("handleGetChildren"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(core.TypeSetRecord,
		BuildMiddleware(h.handleSetRecord,
			instrumentHandler("handleSetRecord"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(core.TypeUpdateObject,
		BuildMiddleware(h.handleUpdateObject,
			instrumentHandler("handleUpdateObject"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(core.TypeRegisterChild,
		BuildMiddleware(h.handleRegisterChild,
			instrumentHandler("handleRegisterChild"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(core.TypeSetBlob,
		BuildMiddleware(h.handleSetBlob,
			instrumentHandler("handleSetBlob"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(core.TypeGetObjectIndex,
		BuildMiddleware(h.handleGetObjectIndex,
			instrumentHandler("handleGetObjectIndex"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(core.TypeGetPendingRequests,
		BuildMiddleware(h.handleHasPendingRequests,
			instrumentHandler("handleHasPendingRequests"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(core.TypeGetJet,
		BuildMiddleware(h.handleGetJet,
			instrumentHandler("handleGetJet")))

	h.Bus.MustRegister(core.TypeHotRecords,
		BuildMiddleware(h.handleHotRecords,
			instrumentHandler("handleHotRecords"),
			m.releaseHotDataWaiters))

	h.Bus.MustRegister(
		core.TypeGetRequest,
		BuildMiddleware(
			h.handleGetRequest,
			instrumentHandler("handleGetRequest"),
			m.checkJet,
		),
	)

	h.Bus.MustRegister(
		core.TypeGetPendingRequestID,
		BuildMiddleware(
			h.handleGetPendingRequestID,
			instrumentHandler("handleGetPendingRequestID"),
			m.checkJet,
		),
	)

	// Validation.
	h.Bus.MustRegister(core.TypeValidateRecord,
		BuildMiddleware(h.handleValidateRecord,
			m.addFieldsToLogger,
			m.checkJet))

	h.Bus.MustRegister(core.TypeValidationCheck,
		BuildMiddleware(h.handleValidationCheck,
			m.addFieldsToLogger,
			m.checkJet))

	h.Bus.MustRegister(core.TypeJetDrop,
		BuildMiddleware(h.handleJetDrop,
			m.addFieldsToLogger,
			m.checkJet))
}
func (h *MessageHandler) setReplayHandlers(m *middleware) {
	// Generic.
	h.replayHandlers[core.TypeGetCode] = BuildMiddleware(h.handleGetCode, m.addFieldsToLogger)
	h.replayHandlers[core.TypeGetObject] = BuildMiddleware(h.handleGetObject, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[core.TypeGetDelegate] = BuildMiddleware(h.handleGetDelegate, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[core.TypeGetChildren] = BuildMiddleware(h.handleGetChildren, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[core.TypeSetRecord] = BuildMiddleware(h.handleSetRecord, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[core.TypeUpdateObject] = BuildMiddleware(h.handleUpdateObject, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[core.TypeRegisterChild] = BuildMiddleware(h.handleRegisterChild, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[core.TypeSetBlob] = BuildMiddleware(h.handleSetBlob, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[core.TypeGetObjectIndex] = BuildMiddleware(h.handleGetObjectIndex, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[core.TypeGetPendingRequests] = BuildMiddleware(h.handleHasPendingRequests, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[core.TypeGetJet] = BuildMiddleware(h.handleGetJet)

	// Validation.
	h.replayHandlers[core.TypeValidateRecord] = BuildMiddleware(h.handleValidateRecord, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[core.TypeValidationCheck] = BuildMiddleware(h.handleValidationCheck, m.addFieldsToLogger, m.checkJet)
}
func (h *MessageHandler) setHandlersForHeavy(m *middleware) {
	// Heavy.
	h.Bus.MustRegister(core.TypeHeavyStartStop,
		BuildMiddleware(h.handleHeavyStartStop,
			instrumentHandler("handleHeavyStartStop")))

	h.Bus.MustRegister(core.TypeHeavyReset,
		BuildMiddleware(h.handleHeavyReset,
			instrumentHandler("handleHeavyReset")))

	h.Bus.MustRegister(core.TypeHeavyPayload,
		BuildMiddleware(h.handleHeavyPayload,
			instrumentHandler("handleHeavyPayload")))

	// Generic.
	h.Bus.MustRegister(core.TypeGetCode,
		BuildMiddleware(h.handleGetCode))

	h.Bus.MustRegister(core.TypeGetObject,
		BuildMiddleware(h.handleGetObject,
			instrumentHandler("handleGetObject"),
			m.zeroJetForHeavy))

	h.Bus.MustRegister(core.TypeGetDelegate,
		BuildMiddleware(h.handleGetDelegate,
			instrumentHandler("handleGetDelegate"),
			m.zeroJetForHeavy))

	h.Bus.MustRegister(core.TypeGetChildren,
		BuildMiddleware(h.handleGetChildren,
			instrumentHandler("handleGetChildren"),
			m.zeroJetForHeavy))

	h.Bus.MustRegister(core.TypeGetObjectIndex,
		BuildMiddleware(h.handleGetObjectIndex,
			instrumentHandler("handleGetObjectIndex"),
			m.zeroJetForHeavy))
}

func (h *MessageHandler) handleSetRecord(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	msg := parcel.Message().(*message.SetRecord)
	rec := record.DeserializeRecord(msg.Record)
	jetID := jetFromContext(ctx)

	id := record.NewRecordIDFromRecord(h.PlatformCryptographyScheme, parcel.Pulse(), rec)

	if !h.isHeavy {
		recentStorage := h.RecentStorageProvider.GetStorage(ctx, jetID)
		if request, ok := rec.(record.Request); ok {
			recentStorage.AddPendingRequest(ctx, request.GetObject(), *id)
		}
		if result, ok := rec.(*record.ResultRecord); ok {
			recentStorage.RemovePendingRequest(ctx, result.Object, *result.Request.Record())
		}
	}

	id, err := h.ObjectStorage.SetRecord(ctx, jetID, parcel.Pulse(), rec)
	if err != nil {
		return nil, err
	}

	return &reply.ID{ID: *id}, nil
}

func (h *MessageHandler) handleSetBlob(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	msg := parcel.Message().(*message.SetBlob)
	jetID := jetFromContext(ctx)
	calculatedID := record.CalculateIDForBlob(h.PlatformCryptographyScheme, parcel.Pulse(), msg.Memory)

	_, err := h.ObjectStorage.GetBlob(ctx, jetID, calculatedID)
	if err == nil {
		return &reply.ID{ID: *calculatedID}, nil
	}
	if err != nil && err != storage.ErrNotFound {
		return nil, err
	}

	id, err := h.ObjectStorage.SetBlob(ctx, jetID, parcel.Pulse(), msg.Memory)
	if err == nil {
		return &reply.ID{ID: *id}, nil
	}
	if err == storage.ErrOverride {
		return &reply.ID{ID: *calculatedID}, nil
	}
	return nil, err
}

func (h *MessageHandler) handleGetCode(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	logger := inslogger.FromContext(ctx)
	logger.Debug("CALL handleGetCode")

	msg := parcel.Message().(*message.GetCode)
	jetID := *jet.NewID(0, nil)

	codeRec, err := h.getCode(ctx, msg.Code.Record())
	if err == storage.ErrNotFound {
		// We don't have code record. Must be on another node.
		node, err := h.JetCoordinator.NodeForJet(ctx, jetID, parcel.Pulse(), msg.Code.Record().Pulse())
		if err != nil {
			return nil, err
		}
		return reply.NewGetCodeRedirect(h.DelegationTokenFactory, parcel, node)
	}
	if err != nil {
		return nil, err
	}
	code, err := h.ObjectStorage.GetBlob(ctx, jetID, codeRec.Code)
	if err != nil {
		return nil, err
	}

	rep := reply.Code{
		Code:        code,
		MachineType: codeRec.MachineType,
	}

	return &rep, nil
}

func (h *MessageHandler) handleGetObject(
	ctx context.Context, parcel core.Parcel,
) (core.Reply, error) {
	logger := inslogger.FromContext(ctx)
	logger.Debug("CALL handleGetObject")

	msg := parcel.Message().(*message.GetObject)
	jetID := jetFromContext(ctx)

	if !h.isHeavy {
		h.RecentStorageProvider.GetStorage(ctx, jetID).AddObject(ctx, *msg.Head.Record())
	}

	// Fetch object index. If not found redirect.
	idx, err := h.ObjectStorage.GetObjectIndex(ctx, jetID, msg.Head.Record(), false)
	if err == storage.ErrNotFound {
		if h.isHeavy {
			return nil, fmt.Errorf("failed to fetch index for %s", msg.Head.Record().String())
		}

		logger.Errorf(
			"failed to fetch index (going to heavy). jet: %v, obj: %v",
			jetID.DebugString(),
			msg.Head.Record().DebugString(),
		)
		node, err := h.JetCoordinator.Heavy(ctx, parcel.Pulse())
		if err != nil {
			return nil, err
		}
		idx, err = h.saveIndexFromHeavy(ctx, jetID, msg.Head, node)
		if err != nil {
			return nil, errors.Wrap(err, "failed to fetch index from heavy")
		}
	} else if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch object index %s", msg.Head.Record().String())
	}

	// Determine object state id.
	var stateID *core.RecordID
	if msg.State != nil {
		stateID = msg.State
	} else {
		if msg.Approved {
			stateID = idx.LatestStateApproved
		} else {
			stateID = idx.LatestState
		}
	}
	if stateID == nil {
		return &reply.Error{ErrType: reply.ErrStateNotAvailable}, nil
	}

	var stateJet *core.RecordID
	if h.isHeavy {
		stateJet = &jetID
	} else {
		var actual bool
		onHeavy, err := h.JetCoordinator.IsBeyondLimit(ctx, parcel.Pulse(), stateID.Pulse())
		if err != nil {
			return nil, err
		}
		if onHeavy {
			node, err := h.JetCoordinator.Heavy(ctx, parcel.Pulse())
			if err != nil {
				return nil, err
			}
			logger.Debugf(
				"redirect (on heavy). pulse: %v, id: %v, state: %v, to: %v",
				parcel.Pulse(),
				msg.Head.Record().DebugString(),
				stateID.DebugString(),
				node.String(),
			)
			return reply.NewGetObjectRedirectReply(h.DelegationTokenFactory, parcel, node, stateID)
		}

		stateTree, err := h.JetStorage.GetJetTree(ctx, stateID.Pulse())
		if err != nil {
			return nil, err
		}
		stateJet, actual = stateTree.Find(*msg.Head.Record())
		if !actual {
			actualJet, err := h.jetTreeUpdater.fetchJet(ctx, *msg.Head.Record(), stateID.Pulse())
			if err != nil {
				return nil, err
			}
			stateJet = actualJet
		}
	}

	// Fetch state record.
	rec, err := h.ObjectStorage.GetRecord(ctx, *stateJet, stateID)
	if err == storage.ErrNotFound {
		if h.isHeavy {
			return nil, fmt.Errorf("failed to fetch state for %v. jet: %v, state: %v", msg.Head.Record(), stateJet.DebugString(), stateID.DebugString())
		}
		// The record wasn't found on the current node. Return redirect to the node that contains it.
		// We get Jet tree for pulse when given state was added.
		node, err := h.JetCoordinator.NodeForJet(ctx, *stateJet, parcel.Pulse(), stateID.Pulse())
		if err != nil {
			return nil, err
		}

		logger.Debugf(
			"redirect (record not found). jet: %v, id: %v, state: %v, to: %v",
			stateJet.DebugString(),
			msg.Head.Record().DebugString(),
			stateID.DebugString(),
			node.String(),
		)
		return reply.NewGetObjectRedirectReply(h.DelegationTokenFactory, parcel, node, stateID)
	}
	if err != nil {
		return nil, err
	}
	state, ok := rec.(record.ObjectState)
	if !ok {
		return nil, errors.New("invalid object record")
	}
	if state.State() == record.StateDeactivation {
		return &reply.Error{ErrType: reply.ErrDeactivated}, nil
	}

	var childPointer *core.RecordID
	if idx.ChildPointer != nil {
		childPointer = idx.ChildPointer
	}
	rep := reply.Object{
		Head:         msg.Head,
		State:        *stateID,
		Prototype:    state.GetImage(),
		IsPrototype:  state.GetIsPrototype(),
		ChildPointer: childPointer,
		Parent:       idx.Parent,
	}

	if state.GetMemory() != nil {
		rep.Memory, err = h.ObjectStorage.GetBlob(ctx, *stateJet, state.GetMemory())
		if err != nil {
			logger.Errorf(
				"failed to fetch blob. pulse: %v, jet: %v, id: %v",
				parcel.Pulse(),
				stateJet.DebugString(),
				state.GetMemory().DebugString(),
			)
			return nil, errors.Wrap(err, "failed to fetch blob")
		}
	}

	return &rep, nil
}

func (h *MessageHandler) handleHasPendingRequests(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	msg := parcel.Message().(*message.GetPendingRequests)
	jetID := jetFromContext(ctx)

	for _, reqID := range h.RecentStorageProvider.GetStorage(ctx, jetID).GetRequestsForObject(*msg.Object.Record()) {
		if reqID.Pulse() < parcel.Pulse() {
			return &reply.HasPendingRequests{Has: true}, nil
		}
	}

	return &reply.HasPendingRequests{Has: false}, nil
}

func (h *MessageHandler) handleGetJet(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	msg := parcel.Message().(*message.GetJet)
	tree, err := h.JetStorage.GetJetTree(ctx, msg.Pulse)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch jet tree")
	}
	jetID, actual := tree.Find(msg.Object)
	if err != nil {
		return nil, err
	}

	return &reply.Jet{ID: *jetID, Actual: actual}, nil
}

func (h *MessageHandler) handleGetDelegate(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	logger := inslogger.FromContext(ctx)
	logger.Debug("CALL handleGetDelegate")

	msg := parcel.Message().(*message.GetDelegate)
	jetID := jetFromContext(ctx)

	if !h.isHeavy {
		h.RecentStorageProvider.GetStorage(ctx, jetID).AddObject(ctx, *msg.Head.Record())
	}

	idx, err := h.ObjectStorage.GetObjectIndex(ctx, jetID, msg.Head.Record(), false)
	if err == storage.ErrNotFound {
		if h.isHeavy {
			return nil, fmt.Errorf("failed to fetch index for %v", msg.Head.Record())
		}

		heavy, err := h.JetCoordinator.Heavy(ctx, parcel.Pulse())
		if err != nil {
			return nil, err
		}
		idx, err = h.saveIndexFromHeavy(ctx, jetID, msg.Head, heavy)
		if err != nil {
			return nil, errors.Wrap(err, "failed to fetch index from heavy")
		}
	} else if err != nil {
		return nil, errors.Wrap(err, "failed to fetch object index")
	}

	delegateRef, ok := idx.Delegates[msg.AsType]
	if !ok {
		return nil, errors.New("the object has no delegate for this type")
	}

	rep := reply.Delegate{
		Head: delegateRef,
	}

	return &rep, nil
}

func (h *MessageHandler) handleGetChildren(
	ctx context.Context, parcel core.Parcel,
) (core.Reply, error) {
	logger := inslogger.FromContext(ctx)
	logger.Debug("CALL handleGetChildren")

	msg := parcel.Message().(*message.GetChildren)
	jetID := jetFromContext(ctx)

	if !h.isHeavy {
		h.RecentStorageProvider.GetStorage(ctx, jetID).AddObject(ctx, *msg.Parent.Record())
	}

	idx, err := h.ObjectStorage.GetObjectIndex(ctx, jetID, msg.Parent.Record(), false)
	if err == storage.ErrNotFound {
		if h.isHeavy {
			return nil, fmt.Errorf("failed to fetch index for %v", msg.Parent.Record())
		}

		heavy, err := h.JetCoordinator.Heavy(ctx, parcel.Pulse())
		if err != nil {
			return nil, err
		}
		idx, err = h.saveIndexFromHeavy(ctx, jetID, msg.Parent, heavy)
		if err != nil {
			return nil, errors.Wrap(err, "failed to fetch index from heavy")
		}
		if idx.ChildPointer == nil {
			return &reply.Children{Refs: nil, NextFrom: nil}, nil
		}
	} else if err != nil {
		return nil, errors.Wrap(err, "failed to fetch object index")
	}

	var (
		refs         []core.RecordRef
		currentChild *core.RecordID
	)

	// Counting from specified child or the latest.
	if msg.FromChild != nil {
		currentChild = msg.FromChild
	} else {
		currentChild = idx.ChildPointer
	}

	// The object has no children.
	if currentChild == nil {
		return &reply.Children{Refs: nil, NextFrom: nil}, nil
	}

	var childJet *core.RecordID
	if h.isHeavy {
		childJet = &jetID
	} else {
		var actual bool
		onHeavy, err := h.JetCoordinator.IsBeyondLimit(ctx, parcel.Pulse(), currentChild.Pulse())
		if err != nil {
			return nil, err
		}
		if onHeavy {
			node, err := h.JetCoordinator.Heavy(ctx, parcel.Pulse())
			if err != nil {
				return nil, err
			}
			return reply.NewGetChildrenRedirect(h.DelegationTokenFactory, parcel, node, *currentChild)
		}

		childTree, err := h.JetStorage.GetJetTree(ctx, currentChild.Pulse())
		if err != nil {
			return nil, err
		}
		childJet, actual = childTree.Find(*msg.Parent.Record())
		if !actual {
			actualJet, err := h.jetTreeUpdater.fetchJet(ctx, *msg.Parent.Record(), currentChild.Pulse())
			if err != nil {
				return nil, err
			}
			childJet = actualJet
		}
	}

	// Try to fetch the first child.
	_, err = h.ObjectStorage.GetRecord(ctx, *childJet, currentChild)
	if err == storage.ErrNotFound {
		if h.isHeavy {
			return nil, fmt.Errorf("failed to fetch child for %v. jet: %v, state: %v", msg.Parent.Record(), childJet.DebugString(), currentChild.DebugString())
		}
		node, err := h.JetCoordinator.NodeForJet(ctx, *childJet, parcel.Pulse(), currentChild.Pulse())
		if err != nil {
			return nil, err
		}
		return reply.NewGetChildrenRedirect(h.DelegationTokenFactory, parcel, node, *currentChild)
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch child")
	}

	counter := 0
	for currentChild != nil {
		// We have enough results.
		if counter >= msg.Amount {
			return &reply.Children{Refs: refs, NextFrom: currentChild}, nil
		}
		counter++

		rec, err := h.ObjectStorage.GetRecord(ctx, *childJet, currentChild)
		// We don't have this child reference. Return what was collected.
		if err == storage.ErrNotFound {
			return &reply.Children{Refs: refs, NextFrom: currentChild}, nil
		}
		if err != nil {
			return nil, errors.New("failed to retrieve children")
		}

		childRec, ok := rec.(*record.ChildRecord)
		if !ok {
			return nil, errors.New("failed to retrieve children")
		}
		currentChild = childRec.PrevChild

		// Skip records later than specified pulse.
		recPulse := childRec.Ref.Record().Pulse()
		if msg.FromPulse != nil && recPulse > *msg.FromPulse {
			continue
		}
		refs = append(refs, childRec.Ref)
	}

	return &reply.Children{Refs: refs, NextFrom: nil}, nil
}

func (h *MessageHandler) handleGetRequest(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	jetID := jetFromContext(ctx)
	msg := parcel.Message().(*message.GetRequest)

	rec, err := h.ObjectStorage.GetRecord(ctx, jetID, &msg.Request)
	if err != nil {
		return nil, errors.New("failed to fetch request")
	}

	req, ok := rec.(*record.RequestRecord)
	if !ok {
		return nil, errors.New("failed to decode request")
	}

	rep := reply.Request{
		ID:     msg.Request,
		Record: record.SerializeRecord(req),
	}

	return &rep, nil
}

func (h *MessageHandler) handleGetPendingRequestID(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	jetID := jetFromContext(ctx)
	msg := parcel.Message().(*message.GetPendingRequestID)

	requests := h.RecentStorageProvider.GetStorage(ctx, jetID).GetRequestsForObject(msg.ObjectID)
	if len(requests) == 0 {
		return &reply.Error{ErrType: reply.ErrNoPendingRequests}, nil
	}

	rep := reply.ID{
		ID: requests[0],
	}

	return &rep, nil
}

func (h *MessageHandler) handleUpdateObject(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	logger := inslogger.FromContext(ctx)

	msg := parcel.Message().(*message.UpdateObject)
	jetID := jetFromContext(ctx)

	rec := record.DeserializeRecord(msg.Record)
	state, ok := rec.(record.ObjectState)
	if !ok {
		return nil, errors.New("wrong object state record")
	}

	if !h.isHeavy {
		h.RecentStorageProvider.GetStorage(ctx, jetID).AddObject(ctx, *msg.Object.Record())
	}

	// FIXME: temporary fix. If we calculate blob id on the client, pulse can change before message sending and this
	//  id will not match the one calculated on the server.
	blobID, err := h.ObjectStorage.SetBlob(ctx, jetID, parcel.Pulse(), msg.Memory)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set blob")
	}
	logger.Debugf("save blob. pulse: %v, jet: %v, id: %v", parcel.Pulse(), jetID.DebugString(), blobID.DebugString())

	switch s := state.(type) {
	case *record.ObjectActivateRecord:
		s.Memory = blobID
	case *record.ObjectAmendRecord:
		s.Memory = blobID
	}

	var idx *index.ObjectLifeline
	err = h.DBContext.Update(ctx, func(tx *storage.TransactionManager) error {
		var err error
		logger.Debugf("Get index for: %v, jet: %v", msg.Object.Record(), jetID.DebugString())
		idx, err = tx.GetObjectIndex(ctx, jetID, msg.Object.Record(), true)
		// No index on our node.
		if err == storage.ErrNotFound {
			if state.State() == record.StateActivation {
				// We are activating the object. There is no index for it anywhere.
				idx = &index.ObjectLifeline{State: record.StateUndefined}
			} else {
				logger.Debugf("Not found index for: %v, jet: %v", msg.Object.Record(), jetID.DebugString())
				// We are updating object. Index should be on the heavy executor.
				heavy, err := h.JetCoordinator.Heavy(ctx, parcel.Pulse())
				if err != nil {
					return err
				}
				idx, err = h.saveIndexFromHeavy(ctx, jetID, msg.Object, heavy)
				if err != nil {
					return errors.Wrap(err, "failed to fetch index from heavy")
				}
			}
		} else if err != nil {
			return err
		}

		if err = validateState(idx.State, state.State()); err != nil {
			return err
		}
		// Index exists and latest record id does not match (preserving chain consistency).
		if idx.LatestState != nil && !state.PrevStateID().Equal(idx.LatestState) {
			return errors.New("invalid state record")
		}

		id, err := tx.SetRecord(ctx, jetID, parcel.Pulse(), rec)
		if err != nil {
			return err
		}
		idx.LatestState = id
		idx.State = state.State()
		if state.State() == record.StateActivation {
			idx.Parent = state.(*record.ObjectActivateRecord).Parent
		}

		logger.WithFields(map[string]interface{}{"jet": jetID.DebugString()}).Debugf("saved object. jet: %v, id: %v, state: %v", jetID.DebugString(), msg.Object.Record().DebugString(), id.DebugString())

		idx.LatestUpdate = parcel.Pulse()
		return tx.SetObjectIndex(ctx, jetID, msg.Object.Record(), idx)
	})
	if err != nil {
		if err == ErrObjectDeactivated {
			return &reply.Error{ErrType: reply.ErrDeactivated}, nil
		}
		return nil, err
	}

	rep := reply.Object{
		Head:         msg.Object,
		State:        *idx.LatestState,
		Prototype:    state.GetImage(),
		IsPrototype:  state.GetIsPrototype(),
		ChildPointer: idx.ChildPointer,
		Parent:       idx.Parent,
	}
	return &rep, nil
}

func (h *MessageHandler) handleRegisterChild(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	msg := parcel.Message().(*message.RegisterChild)
	jetID := jetFromContext(ctx)
	rec := record.DeserializeRecord(msg.Record)
	childRec, ok := rec.(*record.ChildRecord)
	if !ok {
		return nil, errors.New("wrong child record")
	}

	if !h.isHeavy {
		h.RecentStorageProvider.GetStorage(ctx, jetID).AddObject(ctx, *msg.Parent.Record())
	}

	var child *core.RecordID
	err := h.DBContext.Update(ctx, func(tx *storage.TransactionManager) error {
		idx, err := h.ObjectStorage.GetObjectIndex(ctx, jetID, msg.Parent.Record(), false)
		if err == storage.ErrNotFound {
			heavy, err := h.JetCoordinator.Heavy(ctx, parcel.Pulse())
			if err != nil {
				return err
			}
			idx, err = h.saveIndexFromHeavy(ctx, jetID, msg.Parent, heavy)
			if err != nil {
				return errors.Wrap(err, "failed to fetch index from heavy")
			}
		} else if err != nil {
			return err
		}

		// Children exist and pointer does not match (preserving chain consistency).
		if idx.ChildPointer != nil && !childRec.PrevChild.Equal(idx.ChildPointer) {
			return errors.New("invalid child record")
		}

		child, err = tx.SetRecord(ctx, jetID, parcel.Pulse(), childRec)
		if err != nil {
			return err
		}
		idx.ChildPointer = child
		if msg.AsType != nil {
			idx.Delegates[*msg.AsType] = msg.Child
		}
		idx.LatestUpdate = parcel.Pulse()
		err = tx.SetObjectIndex(ctx, jetID, msg.Parent.Record(), idx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &reply.ID{ID: *child}, nil
}

func (h *MessageHandler) handleJetDrop(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	msg := parcel.Message().(*message.JetDrop)

	if !hack.SkipValidation(ctx) {
		for _, parcelBuff := range msg.Messages {
			jetDropMsg, err := message.Deserialize(bytes.NewBuffer(parcelBuff))
			if err != nil {
				return nil, err
			}
			handler, ok := h.replayHandlers[jetDropMsg.Type()]
			if !ok {
				return nil, errors.New("unknown message type")
			}

			_, err = handler(ctx, &message.Parcel{Msg: jetDropMsg})
			if err != nil {
				return nil, err
			}
		}
	}

	err := h.JetStorage.AddJets(ctx, msg.JetID)
	if err != nil {
		return nil, err
	}

	err = h.JetStorage.UpdateJetTree(
		ctx,
		parcel.Pulse(),
		true,
		msg.JetID,
	)
	if err != nil {
		return nil, err
	}

	return &reply.OK{}, nil
}

func (h *MessageHandler) handleValidateRecord(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	msg := parcel.Message().(*message.ValidateRecord)
	jetID := jetFromContext(ctx)

	err := h.DBContext.Update(ctx, func(tx *storage.TransactionManager) error {
		idx, err := tx.GetObjectIndex(ctx, jetID, msg.Object.Record(), true)
		if err == storage.ErrNotFound {
			heavy, err := h.JetCoordinator.Heavy(ctx, parcel.Pulse())
			if err != nil {
				return err
			}
			idx, err = h.saveIndexFromHeavy(ctx, jetID, msg.Object, heavy)
			if err != nil {
				return errors.Wrap(err, "failed to fetch index from heavy")
			}
		} else if err != nil {
			return err
		}

		// Find node that has this state.
		node, err := h.JetCoordinator.NodeForJet(ctx, jetID, parcel.Pulse(), msg.Object.Record().Pulse())
		if err != nil {
			return err
		}

		// Send checking message.
		genericReply, err := h.Bus.Send(ctx, &message.ValidationCheck{
			Object:              msg.Object,
			ValidatedState:      msg.State,
			LatestStateApproved: idx.LatestStateApproved,
		}, &core.MessageSendOptions{
			Receiver: node,
		})
		if err != nil {
			return err
		}
		switch genericReply.(type) {
		case *reply.OK:
			if msg.IsValid {
				idx.LatestStateApproved = &msg.State
			} else {
				idx.LatestState = idx.LatestStateApproved
			}
			idx.LatestUpdate = parcel.Pulse()
			err = tx.SetObjectIndex(ctx, jetID, msg.Object.Record(), idx)
			if err != nil {
				return errors.Wrap(err, "failed to save object index")
			}
		case *reply.NotOK:
			return errors.New("validation sequence integrity failure")
		default:
			return errors.New("handleValidateRecord: unexpected reply")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &reply.OK{}, nil
}

func (h *MessageHandler) handleGetObjectIndex(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	inslog := inslogger.FromContext(ctx)
	msg := parcel.Message().(*message.GetObjectIndex)
	jetID := jetFromContext(ctx)

	inslog.Debugf("handleGetObjectIndex: jetID: %v", jetID)
	inslog.Debug("handleGetObjectIndex: msg.Object.Record() ", msg.Object.Record())
	idx, err := h.ObjectStorage.GetObjectIndex(ctx, jetID, msg.Object.Record(), true)
	if err != nil {
		inslog.Debug("handleGetObjectIndex: failed to fetch object index, error - ", err)
		return nil, errors.Wrap(err, "failed to fetch object index")
	}

	buf, err := index.EncodeObjectLifeline(idx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to serialize index")
	}

	return &reply.ObjectIndex{Index: buf}, nil
}

func (h *MessageHandler) handleValidationCheck(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	msg := parcel.Message().(*message.ValidationCheck)
	jetID := jetFromContext(ctx)

	rec, err := h.ObjectStorage.GetRecord(ctx, jetID, &msg.ValidatedState)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch state record")
	}
	state, ok := rec.(record.ObjectState)
	if !ok {
		return nil, errors.New("failed to fetch state record")
	}
	approved := msg.LatestStateApproved
	validated := state.PrevStateID()
	if !approved.Equal(validated) && approved != nil && validated != nil {
		return &reply.NotOK{}, nil
	}

	return &reply.OK{}, nil
}

func (h *MessageHandler) getCode(ctx context.Context, id *core.RecordID) (*record.CodeRecord, error) {
	jetID := *jet.NewID(0, nil)

	rec, err := h.ObjectStorage.GetRecord(ctx, jetID, id)
	if err != nil {
		return nil, err
	}
	codeRec, ok := rec.(*record.CodeRecord)
	if !ok {
		return nil, errors.Wrap(ErrInvalidRef, "failed to retrieve code record")
	}

	return codeRec, nil
}

func validateState(old record.State, new record.State) error {
	if old == record.StateDeactivation {
		return ErrObjectDeactivated
	}
	if old == record.StateUndefined && new != record.StateActivation {
		return errors.New("object is not activated")
	}
	if old != record.StateUndefined && new == record.StateActivation {
		return errors.New("object is already activated")
	}
	return nil
}

func (h *MessageHandler) saveIndexFromHeavy(
	ctx context.Context, jetID core.RecordID, obj core.RecordRef, heavy *core.RecordRef,
) (*index.ObjectLifeline, error) {
	genericReply, err := h.Bus.Send(ctx, &message.GetObjectIndex{
		Object: obj,
	}, &core.MessageSendOptions{
		Receiver: heavy,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to send")
	}
	rep, ok := genericReply.(*reply.ObjectIndex)
	if !ok {
		return nil, fmt.Errorf("failed to fetch object index: unexpected reply type %T (reply=%+v)", genericReply, genericReply)
	}
	idx, err := index.DecodeObjectLifeline(rep.Index)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode")
	}

	err = h.ObjectStorage.SetObjectIndex(ctx, jetID, obj.Record(), idx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save")
	}
	return idx, nil
}

func (h *MessageHandler) handleHotRecords(ctx context.Context, parcel core.Parcel) (core.Reply, error) {
	logger := inslogger.FromContext(ctx)
	// if hack.SkipValidation(ctx) {
	// 	fmt.Println("handleHotRecords: SkipValidation")
	// 	return &reply.OK{}, nil
	// }

	msg := parcel.Message().(*message.HotData)
	// FIXME: check split signatures.
	jetID := *msg.Jet.Record()

	logger.Debugf("[jet]: %v got hot. Pulse: %v, DropPulse: %v, DropJet: %v\n", jetID.DebugString(), parcel.Pulse(), msg.Drop.Pulse, msg.DropJet.DebugString())

	err := h.DropStorage.SetDrop(ctx, msg.DropJet, &msg.Drop)
	if err == storage.ErrOverride {
		logger.Debugf("received drop duplicate for. jet: %v, pulse: %v", msg.DropJet.DebugString(), msg.Drop.Pulse)
		err = nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, "[jet]: drop error (pulse: %v)", msg.Drop.Pulse)
	}
	err = h.DropStorage.SetDropSizeHistory(ctx, msg.DropJet, msg.JetDropSizeHistory)
	if err != nil {
		return nil, errors.Wrap(err, "[ handleHotRecords ] Can't SetDropSizeHistory")
	}

	logger.WithFields(map[string]interface{}{
		"len": len(msg.RecentObjects),
		"jet": jetID.DebugString(),
	}).Debugf("received pending requests")
	recentStorage := h.RecentStorageProvider.GetStorage(ctx, jetID)
	for objID, requests := range msg.PendingRequests {
		for reqID := range requests {
			recentStorage.AddPendingRequest(ctx, objID, reqID)
		}
	}

	logger.WithFields(map[string]interface{}{
		"len": len(msg.RecentObjects),
		"jet": jetID.DebugString(),
	}).Debugf("received recent objects")
	for id, meta := range msg.RecentObjects {
		logger.Debugf("[got id] jet: %v, id: %v", jetID.DebugString(), id.DebugString())
		decodedIndex, err := index.DecodeObjectLifeline(meta.Index)
		if err != nil {
			fmt.Print("hot index write error")
			logger.Error(err)
			continue
		}

		err = h.ObjectStorage.SetObjectIndex(ctx, jetID, &id, decodedIndex)
		if err != nil {
			fmt.Print("hot index write error")
			logger.Error(err)
			continue
		}

		fmt.Println("[saved id] ", id.String())
		recentStorage.AddObjectWithTLL(ctx, id, meta.TTL)
	}

	err = h.JetStorage.UpdateJetTree(
		ctx,
		msg.PulseNumber,
		true,
		jetID,
	)
	if err != nil {
		fmt.Println("handleHotRecords: UpdateJetTree with err, ", err)
		return nil, err
	}
	err = h.JetStorage.AddJets(ctx, jetID)
	if err != nil {
		return nil, err
	}

	return &reply.OK{}, nil
}
