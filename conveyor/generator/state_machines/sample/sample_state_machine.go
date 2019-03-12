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

package sample

import (
	"github.com/insolar/insolar/conveyor/interfaces/adapter"
	"github.com/insolar/insolar/conveyor/interfaces/statemachine"
)

// custom types
type Event struct{}
type Payload struct{}
type TA1 string
type TAR string

// conveyor: state_machine
type TestStateMachine interface {
	GetTypeID() statemachine.ID

	initPresentHandler(input Event, payload interface{}) (*Payload, statemachine.ElementState, error)
	initFutureHandler(input Event, payload interface{}) (*Payload, statemachine.ElementState, error)
	initPastHandler(input Event, payload interface{}) (*Payload, statemachine.ElementState, error)

	errorPresentInit(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState)
	errorFutureInit(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState)
	errorPastInit(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState)

	// State Declaration
	stateFirst() statemachine.StateID

	// Migration
	migrateFromPresentFirst(input Event, payload *Payload) (*Payload, statemachine.ElementState, error)
	migrateFromFutureFirst(input Event, payload *Payload) (*Payload, statemachine.ElementState, error)

	// Transition
	transitPresentFirst(input Event, payload *Payload /* todo: , adapterHelper TA1*/) (*Payload, statemachine.ElementState, error)
	transitFutureFirst(input Event, payload *Payload /* todo: , adapterHelper TA1*/) (*Payload, statemachine.ElementState, error)
	transitPastFirst(input Event, payload *Payload) (*Payload, statemachine.ElementState, error)

	// TODO: Finalization
	// finalizePresentFirst(input Event, payload *Payload)
	// ...

	// Adapter Response
	responsePresentFirst(input Event, payload *Payload, respPayload TAR) (*Payload, statemachine.ElementState, error)
	responseFutureFirst(input Event, payload *Payload, respPayload TAR) (*Payload, statemachine.ElementState, error)
	responsePastFirst(input Event, payload *Payload, respPayload TAR) (*Payload, statemachine.ElementState, error)

	// State Error
	errorPresentFirst(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState)
	errorFutureFirst(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState)
	errorPastFirst(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState)

	// Adapter Response Error
	errorResponsePresentFirst(input interface{}, payload interface{}, ar adapter.IAdapterResponse, err error) (*Payload, statemachine.ElementState)
	errorResponseFutureFirst(input interface{}, payload interface{}, ar adapter.IAdapterResponse, err error) (*Payload, statemachine.ElementState)
	errorResponsePastFirst(input interface{}, payload interface{}, ar adapter.IAdapterResponse, err error) (*Payload, statemachine.ElementState)

	// State Declaration
	stateSecond() statemachine.StateID

	// Migration
	migrateFromPresentSecond(input Event, payload *Payload) (*Payload, statemachine.ElementState, error)
	migrateFromFutureSecond(input Event, payload *Payload) (*Payload, statemachine.ElementState, error)

	// Transition
	transitPresentFirstState(input Event, payload *Payload /* todo: , adapterHelper TA1*/) (*Payload, statemachine.ElementState, error)
	transitFutureFirstState(input Event, payload *Payload /* todo: , adapterHelper TA1*/) (*Payload, statemachine.ElementState, error)
	transitPastFirstState(input Event, payload *Payload) (*Payload, statemachine.ElementState, error)

	// TODO: Finalization
	// finalizePresentSecond(input Event, payload *Payload)
	// ...

	// Adapter Response
	responsePresentSecond(input Event, payload *Payload, respPayload TAR) (*Payload, statemachine.ElementState, error)
	responseFutureSecond(input Event, payload *Payload, respPayload TAR) (*Payload, statemachine.ElementState, error)
	responsePastSecond(input Event, payload *Payload, respPayload TAR) (*Payload, statemachine.ElementState, error)

	// State Error
	errorPresentSecond(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState)
	errorFutureSecond(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState)
	errorPastSecond(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState)

	// Adapter Response Error
	errorResponsePresentSecond(input interface{}, payload interface{}, ar adapter.IAdapterResponse, err error) (*Payload, statemachine.ElementState)
	errorResponseFutureSecond(input interface{}, payload interface{}, ar adapter.IAdapterResponse, err error) (*Payload, statemachine.ElementState)
	errorResponsePastSecond(input interface{}, payload interface{}, ar adapter.IAdapterResponse, err error) (*Payload, statemachine.ElementState)
}

type CleanTestStateMachine struct {
	BaseTestStateMachine
}

func (sm *CleanTestStateMachine) initPresentHandler(input Event, payload interface{}) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) initFutureHandler(input Event, payload interface{}) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) initPastHandler(input Event, payload interface{}) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorPresentInit(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorFutureInit(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorPastInit(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) migrateFromPresentFirst(input Event, payload *Payload) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) migrateFromFutureFirst(input Event, payload *Payload) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) transitPresentFirst(input Event, payload *Payload /* todo: , adapterHelper TA1*/) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) transitFutureFirst(input Event, payload *Payload /* todo: , adapterHelper TA1*/) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) transitPastFirst(input Event, payload *Payload) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) responsePresentFirst(input Event, payload *Payload, respPayload TAR) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) responseFutureFirst(input Event, payload *Payload, respPayload TAR) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) responsePastFirst(input Event, payload *Payload, respPayload TAR) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorPresentFirst(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorFutureFirst(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorPastFirst(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorResponsePresentFirst(input interface{}, payload interface{}, ar adapter.IAdapterResponse, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorResponseFutureFirst(input interface{}, payload interface{}, ar adapter.IAdapterResponse, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorResponsePastFirst(input interface{}, payload interface{}, ar adapter.IAdapterResponse, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) migrateFromPresentSecond(input Event, payload *Payload) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) migrateFromFutureSecond(input Event, payload *Payload) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) transitPresentFirstState(input Event, payload *Payload /* todo: , adapterHelper TA1*/) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) transitFutureFirstState(input Event, payload *Payload /* todo: , adapterHelper TA1*/) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) transitPastFirstState(input Event, payload *Payload) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) responsePresentSecond(input Event, payload *Payload, respPayload TAR) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) responseFutureSecond(input Event, payload *Payload, respPayload TAR) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) responsePastSecond(input Event, payload *Payload, respPayload TAR) (*Payload, statemachine.ElementState, error) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorPresentSecond(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorFutureSecond(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorPastSecond(input interface{}, payload interface{}, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorResponsePresentSecond(input interface{}, payload interface{}, ar adapter.IAdapterResponse, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorResponseFutureSecond(input interface{}, payload interface{}, ar adapter.IAdapterResponse, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}

func (sm *CleanTestStateMachine) errorResponsePastSecond(input interface{}, payload interface{}, ar adapter.IAdapterResponse, err error) (*Payload, statemachine.ElementState) {
	panic("implement me")
}