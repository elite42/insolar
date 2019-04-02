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
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/andreyromancev/belt"
	"github.com/insolar/insolar/ledger/artifactmanager/conver"
	"github.com/insolar/insolar/ledger/storage/blob"
	"github.com/pkg/errors"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"

	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/delegationtoken"
	"github.com/insolar/insolar/insolar/jet"
	"github.com/insolar/insolar/insolar/message"
	"github.com/insolar/insolar/insolar/reply"
	"github.com/insolar/insolar/instrumentation/hack"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/instrumentation/insmetrics"
	"github.com/insolar/insolar/ledger/recentstorage"
	"github.com/insolar/insolar/ledger/storage"
	"github.com/insolar/insolar/ledger/storage/drop"
	"github.com/insolar/insolar/ledger/storage/node"
	"github.com/insolar/insolar/ledger/storage/object"
)

// MessageHandler processes messages for local storage interaction.
type MessageHandler struct {
	RecentStorageProvider      recentstorage.Provider             `inject:""`
	Bus                        insolar.MessageBus                 `inject:""`
	PlatformCryptographyScheme insolar.PlatformCryptographyScheme `inject:""`
	JetCoordinator             insolar.JetCoordinator             `inject:""`
	CryptographyService        insolar.CryptographyService        `inject:""`
	DelegationTokenFactory     insolar.DelegationTokenFactory     `inject:""`
	PulseStorage               insolar.PulseStorage               `inject:""`
	JetStorage                 jet.Storage                        `inject:""`

	DropModifier drop.Modifier `inject:""`

	BlobModifier blob.Modifier `inject:""`
	BlobAccessor blob.Accessor `inject:""`

	IDLocker storage.IDLocker `inject:""`

	ObjectStorage storage.ObjectStorage `inject:""`
	Nodes         node.Accessor         `inject:""`
	PulseTracker  storage.PulseTracker  `inject:""`
	DBContext     storage.DBContext     `inject:""`
	HotDataWaiter HotDataWaiter         `inject:""`

	replayHandlers map[insolar.MessageType]insolar.MessageHandler
	conf           *configuration.Ledger
	middleware     *middleware
	jetTreeUpdater *jetTreeUpdater

	// Belt.
	beltHandlers map[insolar.MessageType]insolar.MessageHandler
	events       chan belt.Event
	Sorter       *conver.Sorter
}

// NewMessageHandler creates new handler.
func NewMessageHandler(conf *configuration.Ledger) *MessageHandler {
	return &MessageHandler{
		replayHandlers: map[insolar.MessageType]insolar.MessageHandler{},
		conf:           conf,
	}
}

func instrumentHandler(name string) Handler {
	return func(handler insolar.MessageHandler) insolar.MessageHandler {
		return func(ctx context.Context, p insolar.Parcel) (insolar.Reply, error) {
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
	h.initBelt(ctx)

	m := newMiddleware(h)
	h.middleware = m

	h.jetTreeUpdater = newJetTreeUpdater(h.Nodes, h.JetStorage, h.Bus, h.JetCoordinator)
	h.setHandlersForLight(m)
	h.setReplayHandlers(m)

	return nil
}

func (h *MessageHandler) setHandlersForLight(m *middleware) {
	// Generic.
	h.Bus.MustRegister(insolar.TypeGetCode, BuildMiddleware(h.handleGetCode,
		instrumentHandler("handleGetCode"),
		m.addFieldsToLogger,
		m.checkJet,
	))

	// h.beltHandlers[insolar.TypeGetObject] = BuildMiddleware(h.handleGetObject)
	h.Bus.MustRegister(insolar.TypeGetObject, h.WrapMessageBus)
	//h.Bus.MustRegister(insolar.TypeGetObject, h.beltHandlers[insolar.TypeGetObject])

	h.Bus.MustRegister(insolar.TypeGetDelegate,
		BuildMiddleware(h.handleGetDelegate,
			instrumentHandler("handleGetDelegate"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(insolar.TypeGetChildren,
		BuildMiddleware(h.handleGetChildren,
			instrumentHandler("handleGetChildren"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(insolar.TypeSetRecord,
		BuildMiddleware(h.handleSetRecord,
			instrumentHandler("handleSetRecord"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(insolar.TypeUpdateObject,
		BuildMiddleware(h.handleUpdateObject,
			instrumentHandler("handleUpdateObject"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(insolar.TypeRegisterChild,
		BuildMiddleware(h.handleRegisterChild,
			instrumentHandler("handleRegisterChild"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(insolar.TypeSetBlob,
		BuildMiddleware(h.handleSetBlob,
			instrumentHandler("handleSetBlob"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(insolar.TypeGetObjectIndex,
		BuildMiddleware(h.handleGetObjectIndex,
			instrumentHandler("handleGetObjectIndex"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(insolar.TypeGetPendingRequests,
		BuildMiddleware(h.handleHasPendingRequests,
			instrumentHandler("handleHasPendingRequests"),
			m.addFieldsToLogger,
			m.checkJet,
			m.waitForHotData))

	h.Bus.MustRegister(insolar.TypeGetJet,
		BuildMiddleware(h.handleGetJet,
			instrumentHandler("handleGetJet")))

	h.Bus.MustRegister(insolar.TypeHotRecords,
		BuildMiddleware(h.handleHotRecords,
			instrumentHandler("handleHotRecords"),
			m.releaseHotDataWaiters))

	h.Bus.MustRegister(
		insolar.TypeGetRequest,
		BuildMiddleware(
			h.handleGetRequest,
			instrumentHandler("handleGetRequest"),
			m.checkJet,
		),
	)

	h.Bus.MustRegister(
		insolar.TypeGetPendingRequestID,
		BuildMiddleware(
			h.handleGetPendingRequestID,
			instrumentHandler("handleGetPendingRequestID"),
			m.checkJet,
		),
	)

	// Validation.
	h.Bus.MustRegister(insolar.TypeValidateRecord,
		BuildMiddleware(h.handleValidateRecord,
			m.addFieldsToLogger,
			m.checkJet))

	h.Bus.MustRegister(insolar.TypeValidationCheck,
		BuildMiddleware(h.handleValidationCheck,
			m.addFieldsToLogger,
			m.checkJet))

	h.Bus.MustRegister(insolar.TypeJetDrop,
		BuildMiddleware(h.handleJetDrop,
			m.addFieldsToLogger,
			m.checkJet))
}

func (h *MessageHandler) setReplayHandlers(m *middleware) {
	// Generic.
	h.replayHandlers[insolar.TypeGetCode] = BuildMiddleware(h.handleGetCode, m.addFieldsToLogger)
	// h.replayHandlers[insolar.TypeGetObject] = BuildMiddleware(h.handleGetObject, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[insolar.TypeGetDelegate] = BuildMiddleware(h.handleGetDelegate, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[insolar.TypeGetChildren] = BuildMiddleware(h.handleGetChildren, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[insolar.TypeSetRecord] = BuildMiddleware(h.handleSetRecord, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[insolar.TypeUpdateObject] = BuildMiddleware(h.handleUpdateObject, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[insolar.TypeRegisterChild] = BuildMiddleware(h.handleRegisterChild, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[insolar.TypeSetBlob] = BuildMiddleware(h.handleSetBlob, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[insolar.TypeGetObjectIndex] = BuildMiddleware(h.handleGetObjectIndex, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[insolar.TypeGetPendingRequests] = BuildMiddleware(h.handleHasPendingRequests, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[insolar.TypeGetJet] = BuildMiddleware(h.handleGetJet)

	// Validation.
	h.replayHandlers[insolar.TypeValidateRecord] = BuildMiddleware(h.handleValidateRecord, m.addFieldsToLogger, m.checkJet)
	h.replayHandlers[insolar.TypeValidationCheck] = BuildMiddleware(h.handleValidationCheck, m.addFieldsToLogger, m.checkJet)
}

func (h *MessageHandler) handleSetRecord(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	msg := parcel.Message().(*message.SetRecord)
	rec := object.DeserializeRecord(msg.Record)
	jetID := jetFromContext(ctx)

	calculatedID := object.NewRecordIDFromRecord(h.PlatformCryptographyScheme, parcel.Pulse(), rec)

	switch r := rec.(type) {
	case object.Request:
		if h.RecentStorageProvider.Count() > h.conf.PendingRequestsLimit {
			return &reply.Error{ErrType: reply.ErrTooManyPendingRequests}, nil
		}
		recentStorage := h.RecentStorageProvider.GetPendingStorage(ctx, jetID)
		recentStorage.AddPendingRequest(ctx, r.GetObject(), *calculatedID)
	case *object.ResultRecord:
		recentStorage := h.RecentStorageProvider.GetPendingStorage(ctx, jetID)
		recentStorage.RemovePendingRequest(ctx, r.Object, *r.Request.Record())
	}

	id, err := h.ObjectStorage.SetRecord(ctx, jetID, parcel.Pulse(), rec)
	if err == storage.ErrOverride {
		inslogger.FromContext(ctx).WithField("type", fmt.Sprintf("%T", rec)).Warn("set record override")
		id = calculatedID
	} else if err != nil {
		return nil, err
	}

	return &reply.ID{ID: *id}, nil
}

func (h *MessageHandler) handleSetBlob(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	msg := parcel.Message().(*message.SetBlob)
	jetID := jetFromContext(ctx)
	calculatedID := object.CalculateIDForBlob(h.PlatformCryptographyScheme, parcel.Pulse(), msg.Memory)

	_, err := h.BlobAccessor.ForID(ctx, *calculatedID)
	if err == nil {
		return &reply.ID{ID: *calculatedID}, nil
	}
	if err != nil && err != blob.ErrNotFound {
		return nil, err
	}

	err = h.BlobModifier.Set(ctx, *calculatedID, blob.Blob{Value: msg.Memory, JetID: insolar.JetID(jetID)})
	if err == nil {
		return &reply.ID{ID: *calculatedID}, nil
	}
	if err == blob.ErrOverride {
		return &reply.ID{ID: *calculatedID}, nil
	}
	return nil, err
}

func (h *MessageHandler) handleGetCode(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	msg := parcel.Message().(*message.GetCode)
	jetID := jetFromContext(ctx)

	codeRec, err := h.getCode(ctx, msg.Code.Record())
	if err == insolar.ErrNotFound {
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
	code, err := h.BlobAccessor.ForID(ctx, *codeRec.Code)
	if err == blob.ErrNotFound {
		hNode, err := h.JetCoordinator.Heavy(ctx, parcel.Pulse())
		if err != nil {
			return nil, err
		}
		return h.saveCodeFromHeavy(ctx, insolar.JetID(jetID), msg.Code, *codeRec.Code, hNode)
	}

	rep := reply.Code{
		Code:        code.Value,
		MachineType: codeRec.MachineType,
	}

	return &rep, nil
}

func (h *MessageHandler) handleHasPendingRequests(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	msg := parcel.Message().(*message.GetPendingRequests)
	jetID := jetFromContext(ctx)

	for _, reqID := range h.RecentStorageProvider.GetPendingStorage(ctx, jetID).GetRequestsForObject(*msg.Object.Record()) {
		if reqID.Pulse() < parcel.Pulse() {
			return &reply.HasPendingRequests{Has: true}, nil
		}
	}

	return &reply.HasPendingRequests{Has: false}, nil
}

func (h *MessageHandler) handleGetJet(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	msg := parcel.Message().(*message.GetJet)

	jetID, actual := h.JetStorage.ForID(ctx, msg.Pulse, msg.Object)

	return &reply.Jet{ID: insolar.ID(jetID), Actual: actual}, nil
}

func (h *MessageHandler) handleGetDelegate(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	msg := parcel.Message().(*message.GetDelegate)
	jetID := jetFromContext(ctx)

	h.RecentStorageProvider.GetIndexStorage(ctx, jetID).AddObject(ctx, *msg.Head.Record())

	h.IDLocker.Lock(msg.Head.Record())
	defer h.IDLocker.Unlock(msg.Head.Record())

	idx, err := h.ObjectStorage.GetObjectIndex(ctx, jetID, msg.Head.Record())
	if err == insolar.ErrNotFound {
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
	ctx context.Context, parcel insolar.Parcel,
) (insolar.Reply, error) {
	msg := parcel.Message().(*message.GetChildren)
	jetID := jetFromContext(ctx)

	h.RecentStorageProvider.GetIndexStorage(ctx, jetID).AddObject(ctx, *msg.Parent.Record())

	h.IDLocker.Lock(msg.Parent.Record())
	defer h.IDLocker.Unlock(msg.Parent.Record())

	idx, err := h.ObjectStorage.GetObjectIndex(ctx, jetID, msg.Parent.Record())
	if err == insolar.ErrNotFound {
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
		refs         []insolar.Reference
		currentChild *insolar.ID
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

	var childJet *insolar.ID
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

	childJetID, actual := h.JetStorage.ForID(ctx, currentChild.Pulse(), *msg.Parent.Record())
	childJet = (*insolar.ID)(&childJetID)

	if !actual {
		actualJet, err := h.jetTreeUpdater.fetchJet(ctx, *msg.Parent.Record(), currentChild.Pulse())
		if err != nil {
			return nil, err
		}
		childJet = actualJet
	}

	// Try to fetch the first child.
	_, err = h.ObjectStorage.GetRecord(ctx, *childJet, currentChild)
	if err == insolar.ErrNotFound {
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
		if err == insolar.ErrNotFound {
			return &reply.Children{Refs: refs, NextFrom: currentChild}, nil
		}
		if err != nil {
			return nil, errors.New("failed to retrieve children")
		}

		childRec, ok := rec.(*object.ChildRecord)
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

func (h *MessageHandler) handleGetRequest(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	jetID := jetFromContext(ctx)
	msg := parcel.Message().(*message.GetRequest)

	rec, err := h.ObjectStorage.GetRecord(ctx, jetID, &msg.Request)
	if err != nil {
		return nil, errors.New("failed to fetch request")
	}

	req, ok := rec.(*object.RequestRecord)
	if !ok {
		return nil, errors.New("failed to decode request")
	}

	rep := reply.Request{
		ID:     msg.Request,
		Record: object.SerializeRecord(req),
	}

	return &rep, nil
}

func (h *MessageHandler) handleGetPendingRequestID(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	jetID := jetFromContext(ctx)
	msg := parcel.Message().(*message.GetPendingRequestID)

	requests := h.RecentStorageProvider.GetPendingStorage(ctx, jetID).GetRequestsForObject(msg.ObjectID)
	if len(requests) == 0 {
		return &reply.Error{ErrType: reply.ErrNoPendingRequests}, nil
	}

	rep := reply.ID{
		ID: requests[0],
	}

	return &rep, nil
}

func (h *MessageHandler) handleUpdateObject(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	msg := parcel.Message().(*message.UpdateObject)
	jetID := jetFromContext(ctx)
	logger := inslogger.FromContext(ctx).WithFields(map[string]interface{}{
		"object": msg.Object.Record().DebugString(),
		"pulse":  parcel.Pulse(),
	})

	rec := object.DeserializeRecord(msg.Record)
	state, ok := rec.(object.State)
	if !ok {
		return nil, errors.New("wrong object state record")
	}

	h.RecentStorageProvider.GetIndexStorage(ctx, jetID).AddObject(ctx, *msg.Object.Record())

	calculatedID := object.CalculateIDForBlob(h.PlatformCryptographyScheme, parcel.Pulse(), msg.Memory)
	// FIXME: temporary fix. If we calculate blob id on the client, pulse can change before message sending and this
	//  id will not match the one calculated on the server.
	err := h.BlobModifier.Set(ctx, *calculatedID, blob.Blob{JetID: insolar.JetID(jetID), Value: msg.Memory})
	if err != nil && err != blob.ErrOverride {
		return nil, errors.Wrap(err, "failed to set blob")
	}

	switch s := state.(type) {
	case *object.ActivateRecord:
		s.Memory = calculatedID
	case *object.AmendRecord:
		s.Memory = calculatedID
	}

	h.IDLocker.Lock(msg.Object.Record())
	defer h.IDLocker.Unlock(msg.Object.Record())

	idx, err := h.ObjectStorage.GetObjectIndex(ctx, jetID, msg.Object.Record())
	// No index on our node.
	if err == insolar.ErrNotFound {
		if state.ID() == object.StateActivation {
			// We are activating the object. There is no index for it anywhere.
			idx = &object.Lifeline{State: object.StateUndefined}
		} else {
			logger.Debug("failed to fetch index (fetching from heavy)")
			// We are updating object. Index should be on the heavy executor.
			heavy, err := h.JetCoordinator.Heavy(ctx, parcel.Pulse())
			if err != nil {
				return nil, err
			}
			idx, err = h.saveIndexFromHeavy(ctx, jetID, msg.Object, heavy)
			if err != nil {
				return nil, errors.Wrap(err, "failed to fetch index from heavy")
			}
		}
	} else if err != nil {
		return nil, err
	}

	if err = validateState(idx.State, state.ID()); err != nil {
		return &reply.Error{ErrType: reply.ErrDeactivated}, nil
	}

	recID := object.NewRecordIDFromRecord(h.PlatformCryptographyScheme, parcel.Pulse(), rec)

	// Index exists and latest record id does not match (preserving chain consistency).
	// For the case when vm can't save or send result to another vm and it tries to update the same record again
	if idx.LatestState != nil && !state.PrevStateID().Equal(*idx.LatestState) && idx.LatestState != recID {
		return nil, errors.New("invalid state record")
	}

	id, err := h.ObjectStorage.SetRecord(ctx, jetID, parcel.Pulse(), rec)
	if err == storage.ErrOverride {
		logger.WithField("type", fmt.Sprintf("%T", rec)).Warn("set record override (#1)")
		id = recID
	} else if err != nil {
		return nil, err
	}
	idx.LatestState = id
	idx.State = state.ID()
	if state.ID() == object.StateActivation {
		idx.Parent = state.(*object.ActivateRecord).Parent
	}

	idx.LatestUpdate = parcel.Pulse()
	err = h.ObjectStorage.SetObjectIndex(ctx, jetID, msg.Object.Record(), idx)
	if err != nil {
		return nil, err
	}

	logger.WithField("state", idx.LatestState.DebugString()).Debug("saved object")

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

func (h *MessageHandler) handleRegisterChild(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	logger := inslogger.FromContext(ctx)

	msg := parcel.Message().(*message.RegisterChild)
	jetID := jetFromContext(ctx)
	rec := object.DeserializeRecord(msg.Record)
	childRec, ok := rec.(*object.ChildRecord)
	if !ok {
		return nil, errors.New("wrong child record")
	}

	h.RecentStorageProvider.GetIndexStorage(ctx, jetID).AddObject(ctx, *msg.Parent.Record())

	h.IDLocker.Lock(msg.Parent.Record())
	defer h.IDLocker.Unlock(msg.Parent.Record())

	var child *insolar.ID
	idx, err := h.ObjectStorage.GetObjectIndex(ctx, jetID, msg.Parent.Record())
	if err == insolar.ErrNotFound {
		heavy, err := h.JetCoordinator.Heavy(ctx, parcel.Pulse())
		if err != nil {
			return nil, err
		}
		idx, err = h.saveIndexFromHeavy(ctx, jetID, msg.Parent, heavy)
		if err != nil {
			return nil, errors.Wrap(err, "failed to fetch index from heavy")
		}
	} else if err != nil {
		return nil, err
	}

	recID := object.NewRecordIDFromRecord(h.PlatformCryptographyScheme, parcel.Pulse(), childRec)

	// Children exist and pointer does not match (preserving chain consistency).
	// For the case when vm can't save or send result to another vm and it tries to update the same record again
	if idx.ChildPointer != nil && !childRec.PrevChild.Equal(*idx.ChildPointer) && idx.ChildPointer != recID {
		return nil, errors.New("invalid child record")
	}

	child, err = h.ObjectStorage.SetRecord(ctx, jetID, parcel.Pulse(), childRec)
	if err == storage.ErrOverride {
		logger.WithField("type", fmt.Sprintf("%T", rec)).Warn("set record override (#2)")
		child = recID
	} else if err != nil {
		return nil, err
	}

	idx.ChildPointer = child
	if msg.AsType != nil {
		idx.Delegates[*msg.AsType] = msg.Child
	}
	idx.LatestUpdate = parcel.Pulse()
	err = h.ObjectStorage.SetObjectIndex(ctx, jetID, msg.Parent.Record(), idx)
	if err != nil {
		return nil, err
	}

	return &reply.ID{ID: *child}, nil
}

func (h *MessageHandler) handleJetDrop(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
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

	h.JetStorage.Update(
		ctx, parcel.Pulse(), true, insolar.JetID(msg.JetID),
	)

	return &reply.OK{}, nil
}

func (h *MessageHandler) handleValidateRecord(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	msg := parcel.Message().(*message.ValidateRecord)
	jetID := jetFromContext(ctx)

	h.IDLocker.Lock(msg.Object.Record())
	defer h.IDLocker.Unlock(msg.Object.Record())

	idx, err := h.ObjectStorage.GetObjectIndex(ctx, jetID, msg.Object.Record())
	if err == insolar.ErrNotFound {
		heavy, err := h.JetCoordinator.Heavy(ctx, parcel.Pulse())
		if err != nil {
			return nil, err
		}
		idx, err = h.saveIndexFromHeavy(ctx, jetID, msg.Object, heavy)
		if err != nil {
			return nil, errors.Wrap(err, "failed to fetch index from heavy")
		}
	} else if err != nil {
		return nil, err
	}

	// Find node that has this state.
	node, err := h.JetCoordinator.NodeForJet(ctx, jetID, parcel.Pulse(), msg.Object.Record().Pulse())
	if err != nil {
		return nil, err
	}

	// Send checking message.
	genericReply, err := h.Bus.Send(ctx, &message.ValidationCheck{
		Object:              msg.Object,
		ValidatedState:      msg.State,
		LatestStateApproved: idx.LatestStateApproved,
	}, &insolar.MessageSendOptions{
		Receiver: node,
	})
	if err != nil {
		return nil, err
	}
	switch genericReply.(type) {
	case *reply.OK:
		if msg.IsValid {
			idx.LatestStateApproved = &msg.State
		} else {
			idx.LatestState = idx.LatestStateApproved
		}
		idx.LatestUpdate = parcel.Pulse()
		err = h.ObjectStorage.SetObjectIndex(ctx, jetID, msg.Object.Record(), idx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to save object index")
		}
	case *reply.NotOK:
		return nil, errors.New("validation sequence integrity failure")
	default:
		return nil, errors.New("handleValidateRecord: unexpected reply")
	}

	return &reply.OK{}, nil
}

func (h *MessageHandler) handleGetObjectIndex(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	msg := parcel.Message().(*message.GetObjectIndex)
	jetID := jetFromContext(ctx)

	h.IDLocker.Lock(msg.Object.Record())
	defer h.IDLocker.Unlock(msg.Object.Record())

	idx, err := h.ObjectStorage.GetObjectIndex(ctx, jetID, msg.Object.Record())
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch object index")
	}

	buf := object.EncodeIndex(*idx)

	return &reply.ObjectIndex{Index: buf}, nil
}

func (h *MessageHandler) handleValidationCheck(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	msg := parcel.Message().(*message.ValidationCheck)
	jetID := jetFromContext(ctx)

	rec, err := h.ObjectStorage.GetRecord(ctx, jetID, &msg.ValidatedState)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch state record")
	}
	state, ok := rec.(object.State)
	if !ok {
		return nil, errors.New("failed to fetch state record")
	}
	approved := msg.LatestStateApproved
	validated := state.PrevStateID()
	if validated != nil && approved != nil && !approved.Equal(*validated) {
		return &reply.NotOK{}, nil
	}

	return &reply.OK{}, nil
}

func (h *MessageHandler) getCode(ctx context.Context, id *insolar.ID) (*object.CodeRecord, error) {
	jetID := jetFromContext(ctx)

	rec, err := h.ObjectStorage.GetRecord(ctx, insolar.ID(jetID), id)
	if err != nil {
		return nil, err
	}
	codeRec, ok := rec.(*object.CodeRecord)
	if !ok {
		return nil, errors.Wrap(ErrInvalidRef, "failed to retrieve code record")
	}

	return codeRec, nil
}

func validateState(old object.StateID, new object.StateID) error {
	if old == object.StateDeactivation {
		return ErrObjectDeactivated
	}
	if old == object.StateUndefined && new != object.StateActivation {
		return errors.New("object is not activated")
	}
	if old != object.StateUndefined && new == object.StateActivation {
		return errors.New("object is already activated")
	}
	return nil
}

func (h *MessageHandler) saveIndexFromHeavy(
	ctx context.Context, jetID insolar.ID, obj insolar.Reference, heavy *insolar.Reference,
) (*object.Lifeline, error) {
	genericReply, err := h.Bus.Send(ctx, &message.GetObjectIndex{
		Object: obj,
	}, &insolar.MessageSendOptions{
		Receiver: heavy,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to send")
	}
	rep, ok := genericReply.(*reply.ObjectIndex)
	if !ok {
		return nil, fmt.Errorf("failed to fetch object index: unexpected reply type %T (reply=%+v)", genericReply, genericReply)
	}
	idx := object.DecodeIndex(rep.Index)

	err = h.ObjectStorage.SetObjectIndex(ctx, jetID, obj.Record(), &idx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save")
	}
	return &idx, nil
}

func (h *MessageHandler) saveCodeFromHeavy(
	ctx context.Context, jetID insolar.JetID, code insolar.Reference, blobID insolar.ID, heavy *insolar.Reference,
) (*reply.Code, error) {
	genericReply, err := h.Bus.Send(ctx, &message.GetCode{
		Code: code,
	}, &insolar.MessageSendOptions{
		Receiver: heavy,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to send")
	}
	rep, ok := genericReply.(*reply.Code)
	if !ok {
		return nil, fmt.Errorf("failed to fetch code: unexpected reply type %T (reply=%+v)", genericReply, genericReply)
	}

	err = h.BlobModifier.Set(ctx, blobID, blob.Blob{JetID: jetID, Value: rep.Code})
	if err != nil {
		return nil, errors.Wrap(err, "failed to save")
	}
	return rep, nil
}

func (h *MessageHandler) fetchObject(
	ctx context.Context, obj insolar.Reference, node insolar.Reference, stateID *insolar.ID, pulse insolar.PulseNumber,
) (*reply.Object, error) {
	sender := BuildSender(
		h.Bus.Send,
		followRedirectSender(h.Bus),
		retryJetSender(pulse, h.JetStorage),
	)
	genericReply, err := sender(
		ctx,
		&message.GetObject{
			Head:     obj,
			Approved: false,
			State:    stateID,
		},
		&insolar.MessageSendOptions{
			Receiver: &node,
			Token:    &delegationtoken.GetObjectRedirectToken{},
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch object state")
	}
	if rep, ok := genericReply.(*reply.Error); ok {
		return nil, rep.Error()
	}

	rep, ok := genericReply.(*reply.Object)
	if !ok {
		return nil, fmt.Errorf("failed to fetch object state: unexpected reply type %T (reply=%+v)", genericReply, genericReply)
	}
	return rep, nil
}

func (h *MessageHandler) handleHotRecords(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	logger := inslogger.FromContext(ctx)

	msg := parcel.Message().(*message.HotData)
	jetID := *msg.Jet.Record()

	logger.WithFields(map[string]interface{}{
		"jet": jetID.DebugString(),
	}).Info("received hot data")

	err := h.DropModifier.Set(ctx, msg.Drop)
	if err == storage.ErrOverride {
		err = nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, "[jet]: drop error (pulse: %v)", msg.Drop.Pulse)
	}

	pendingStorage := h.RecentStorageProvider.GetPendingStorage(ctx, jetID)
	logger.Debugf("received %d pending requests", len(msg.PendingRequests))

	var notificationList []insolar.ID
	for objID, objContext := range msg.PendingRequests {
		if !objContext.Active {
			notificationList = append(notificationList, objID)
		}

		objContext.Active = false
		pendingStorage.SetContextToObject(ctx, objID, objContext)
	}

	go func() {
		for _, objID := range notificationList {
			go func(objID insolar.ID) {
				rep, err := h.Bus.Send(ctx, &message.AbandonedRequestsNotification{
					Object: objID,
				}, nil)

				if err != nil {
					logger.Error("failed to notify about pending requests")
					return
				}
				if _, ok := rep.(*reply.OK); !ok {
					logger.Error("received unexpected reply on pending notification")
				}
			}(objID)
		}
	}()

	indexStorage := h.RecentStorageProvider.GetIndexStorage(ctx, jetID)
	for id, meta := range msg.RecentObjects {
		decodedIndex := object.DecodeIndex(meta.Index)

		err = h.ObjectStorage.SetObjectIndex(ctx, jetID, &id, &decodedIndex)
		if err != nil {
			logger.Error(err)
			continue
		}

		indexStorage.AddObjectWithTLL(ctx, id, meta.TTL)
	}

	h.JetStorage.Update(
		ctx, msg.PulseNumber, true, insolar.JetID(jetID),
	)

	h.jetTreeUpdater.releaseJet(ctx, jetID, msg.PulseNumber)

	return &reply.OK{}, nil
}
