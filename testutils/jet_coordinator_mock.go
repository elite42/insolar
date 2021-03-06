package testutils

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "JetCoordinator" can be found in github.com/insolar/insolar/insolar
*/
import (
	context "context"
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	insolar "github.com/insolar/insolar/insolar"

	testify_assert "github.com/stretchr/testify/assert"
)

//JetCoordinatorMock implements github.com/insolar/insolar/insolar.JetCoordinator
type JetCoordinatorMock struct {
	t minimock.Tester

	HeavyFunc       func(p context.Context, p1 insolar.PulseNumber) (r *insolar.Reference, r1 error)
	HeavyCounter    uint64
	HeavyPreCounter uint64
	HeavyMock       mJetCoordinatorMockHeavy

	IsAuthorizedFunc       func(p context.Context, p1 insolar.DynamicRole, p2 insolar.ID, p3 insolar.PulseNumber, p4 insolar.Reference) (r bool, r1 error)
	IsAuthorizedCounter    uint64
	IsAuthorizedPreCounter uint64
	IsAuthorizedMock       mJetCoordinatorMockIsAuthorized

	IsBeyondLimitFunc       func(p context.Context, p1 insolar.PulseNumber, p2 insolar.PulseNumber) (r bool, r1 error)
	IsBeyondLimitCounter    uint64
	IsBeyondLimitPreCounter uint64
	IsBeyondLimitMock       mJetCoordinatorMockIsBeyondLimit

	LightExecutorForJetFunc       func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r *insolar.Reference, r1 error)
	LightExecutorForJetCounter    uint64
	LightExecutorForJetPreCounter uint64
	LightExecutorForJetMock       mJetCoordinatorMockLightExecutorForJet

	LightExecutorForObjectFunc       func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r *insolar.Reference, r1 error)
	LightExecutorForObjectCounter    uint64
	LightExecutorForObjectPreCounter uint64
	LightExecutorForObjectMock       mJetCoordinatorMockLightExecutorForObject

	LightValidatorsForJetFunc       func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r []insolar.Reference, r1 error)
	LightValidatorsForJetCounter    uint64
	LightValidatorsForJetPreCounter uint64
	LightValidatorsForJetMock       mJetCoordinatorMockLightValidatorsForJet

	LightValidatorsForObjectFunc       func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r []insolar.Reference, r1 error)
	LightValidatorsForObjectCounter    uint64
	LightValidatorsForObjectPreCounter uint64
	LightValidatorsForObjectMock       mJetCoordinatorMockLightValidatorsForObject

	MeFunc       func() (r insolar.Reference)
	MeCounter    uint64
	MePreCounter uint64
	MeMock       mJetCoordinatorMockMe

	NodeForJetFunc       func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber, p3 insolar.PulseNumber) (r *insolar.Reference, r1 error)
	NodeForJetCounter    uint64
	NodeForJetPreCounter uint64
	NodeForJetMock       mJetCoordinatorMockNodeForJet

	NodeForObjectFunc       func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber, p3 insolar.PulseNumber) (r *insolar.Reference, r1 error)
	NodeForObjectCounter    uint64
	NodeForObjectPreCounter uint64
	NodeForObjectMock       mJetCoordinatorMockNodeForObject

	QueryRoleFunc       func(p context.Context, p1 insolar.DynamicRole, p2 insolar.ID, p3 insolar.PulseNumber) (r []insolar.Reference, r1 error)
	QueryRoleCounter    uint64
	QueryRolePreCounter uint64
	QueryRoleMock       mJetCoordinatorMockQueryRole

	VirtualExecutorForObjectFunc       func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r *insolar.Reference, r1 error)
	VirtualExecutorForObjectCounter    uint64
	VirtualExecutorForObjectPreCounter uint64
	VirtualExecutorForObjectMock       mJetCoordinatorMockVirtualExecutorForObject

	VirtualValidatorsForObjectFunc       func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r []insolar.Reference, r1 error)
	VirtualValidatorsForObjectCounter    uint64
	VirtualValidatorsForObjectPreCounter uint64
	VirtualValidatorsForObjectMock       mJetCoordinatorMockVirtualValidatorsForObject
}

//NewJetCoordinatorMock returns a mock for github.com/insolar/insolar/insolar.JetCoordinator
func NewJetCoordinatorMock(t minimock.Tester) *JetCoordinatorMock {
	m := &JetCoordinatorMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.HeavyMock = mJetCoordinatorMockHeavy{mock: m}
	m.IsAuthorizedMock = mJetCoordinatorMockIsAuthorized{mock: m}
	m.IsBeyondLimitMock = mJetCoordinatorMockIsBeyondLimit{mock: m}
	m.LightExecutorForJetMock = mJetCoordinatorMockLightExecutorForJet{mock: m}
	m.LightExecutorForObjectMock = mJetCoordinatorMockLightExecutorForObject{mock: m}
	m.LightValidatorsForJetMock = mJetCoordinatorMockLightValidatorsForJet{mock: m}
	m.LightValidatorsForObjectMock = mJetCoordinatorMockLightValidatorsForObject{mock: m}
	m.MeMock = mJetCoordinatorMockMe{mock: m}
	m.NodeForJetMock = mJetCoordinatorMockNodeForJet{mock: m}
	m.NodeForObjectMock = mJetCoordinatorMockNodeForObject{mock: m}
	m.QueryRoleMock = mJetCoordinatorMockQueryRole{mock: m}
	m.VirtualExecutorForObjectMock = mJetCoordinatorMockVirtualExecutorForObject{mock: m}
	m.VirtualValidatorsForObjectMock = mJetCoordinatorMockVirtualValidatorsForObject{mock: m}

	return m
}

type mJetCoordinatorMockHeavy struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockHeavyExpectation
	expectationSeries []*JetCoordinatorMockHeavyExpectation
}

type JetCoordinatorMockHeavyExpectation struct {
	input  *JetCoordinatorMockHeavyInput
	result *JetCoordinatorMockHeavyResult
}

type JetCoordinatorMockHeavyInput struct {
	p  context.Context
	p1 insolar.PulseNumber
}

type JetCoordinatorMockHeavyResult struct {
	r  *insolar.Reference
	r1 error
}

//Expect specifies that invocation of JetCoordinator.Heavy is expected from 1 to Infinity times
func (m *mJetCoordinatorMockHeavy) Expect(p context.Context, p1 insolar.PulseNumber) *mJetCoordinatorMockHeavy {
	m.mock.HeavyFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockHeavyExpectation{}
	}
	m.mainExpectation.input = &JetCoordinatorMockHeavyInput{p, p1}
	return m
}

//Return specifies results of invocation of JetCoordinator.Heavy
func (m *mJetCoordinatorMockHeavy) Return(r *insolar.Reference, r1 error) *JetCoordinatorMock {
	m.mock.HeavyFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockHeavyExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockHeavyResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.Heavy is expected once
func (m *mJetCoordinatorMockHeavy) ExpectOnce(p context.Context, p1 insolar.PulseNumber) *JetCoordinatorMockHeavyExpectation {
	m.mock.HeavyFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockHeavyExpectation{}
	expectation.input = &JetCoordinatorMockHeavyInput{p, p1}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockHeavyExpectation) Return(r *insolar.Reference, r1 error) {
	e.result = &JetCoordinatorMockHeavyResult{r, r1}
}

//Set uses given function f as a mock of JetCoordinator.Heavy method
func (m *mJetCoordinatorMockHeavy) Set(f func(p context.Context, p1 insolar.PulseNumber) (r *insolar.Reference, r1 error)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.HeavyFunc = f
	return m.mock
}

//Heavy implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) Heavy(p context.Context, p1 insolar.PulseNumber) (r *insolar.Reference, r1 error) {
	counter := atomic.AddUint64(&m.HeavyPreCounter, 1)
	defer atomic.AddUint64(&m.HeavyCounter, 1)

	if len(m.HeavyMock.expectationSeries) > 0 {
		if counter > uint64(len(m.HeavyMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.Heavy. %v %v", p, p1)
			return
		}

		input := m.HeavyMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetCoordinatorMockHeavyInput{p, p1}, "JetCoordinator.Heavy got unexpected parameters")

		result := m.HeavyMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.Heavy")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.HeavyMock.mainExpectation != nil {

		input := m.HeavyMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetCoordinatorMockHeavyInput{p, p1}, "JetCoordinator.Heavy got unexpected parameters")
		}

		result := m.HeavyMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.Heavy")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.HeavyFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.Heavy. %v %v", p, p1)
		return
	}

	return m.HeavyFunc(p, p1)
}

//HeavyMinimockCounter returns a count of JetCoordinatorMock.HeavyFunc invocations
func (m *JetCoordinatorMock) HeavyMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.HeavyCounter)
}

//HeavyMinimockPreCounter returns the value of JetCoordinatorMock.Heavy invocations
func (m *JetCoordinatorMock) HeavyMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.HeavyPreCounter)
}

//HeavyFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) HeavyFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.HeavyMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.HeavyCounter) == uint64(len(m.HeavyMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.HeavyMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.HeavyCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.HeavyFunc != nil {
		return atomic.LoadUint64(&m.HeavyCounter) > 0
	}

	return true
}

type mJetCoordinatorMockIsAuthorized struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockIsAuthorizedExpectation
	expectationSeries []*JetCoordinatorMockIsAuthorizedExpectation
}

type JetCoordinatorMockIsAuthorizedExpectation struct {
	input  *JetCoordinatorMockIsAuthorizedInput
	result *JetCoordinatorMockIsAuthorizedResult
}

type JetCoordinatorMockIsAuthorizedInput struct {
	p  context.Context
	p1 insolar.DynamicRole
	p2 insolar.ID
	p3 insolar.PulseNumber
	p4 insolar.Reference
}

type JetCoordinatorMockIsAuthorizedResult struct {
	r  bool
	r1 error
}

//Expect specifies that invocation of JetCoordinator.IsAuthorized is expected from 1 to Infinity times
func (m *mJetCoordinatorMockIsAuthorized) Expect(p context.Context, p1 insolar.DynamicRole, p2 insolar.ID, p3 insolar.PulseNumber, p4 insolar.Reference) *mJetCoordinatorMockIsAuthorized {
	m.mock.IsAuthorizedFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockIsAuthorizedExpectation{}
	}
	m.mainExpectation.input = &JetCoordinatorMockIsAuthorizedInput{p, p1, p2, p3, p4}
	return m
}

//Return specifies results of invocation of JetCoordinator.IsAuthorized
func (m *mJetCoordinatorMockIsAuthorized) Return(r bool, r1 error) *JetCoordinatorMock {
	m.mock.IsAuthorizedFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockIsAuthorizedExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockIsAuthorizedResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.IsAuthorized is expected once
func (m *mJetCoordinatorMockIsAuthorized) ExpectOnce(p context.Context, p1 insolar.DynamicRole, p2 insolar.ID, p3 insolar.PulseNumber, p4 insolar.Reference) *JetCoordinatorMockIsAuthorizedExpectation {
	m.mock.IsAuthorizedFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockIsAuthorizedExpectation{}
	expectation.input = &JetCoordinatorMockIsAuthorizedInput{p, p1, p2, p3, p4}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockIsAuthorizedExpectation) Return(r bool, r1 error) {
	e.result = &JetCoordinatorMockIsAuthorizedResult{r, r1}
}

//Set uses given function f as a mock of JetCoordinator.IsAuthorized method
func (m *mJetCoordinatorMockIsAuthorized) Set(f func(p context.Context, p1 insolar.DynamicRole, p2 insolar.ID, p3 insolar.PulseNumber, p4 insolar.Reference) (r bool, r1 error)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.IsAuthorizedFunc = f
	return m.mock
}

//IsAuthorized implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) IsAuthorized(p context.Context, p1 insolar.DynamicRole, p2 insolar.ID, p3 insolar.PulseNumber, p4 insolar.Reference) (r bool, r1 error) {
	counter := atomic.AddUint64(&m.IsAuthorizedPreCounter, 1)
	defer atomic.AddUint64(&m.IsAuthorizedCounter, 1)

	if len(m.IsAuthorizedMock.expectationSeries) > 0 {
		if counter > uint64(len(m.IsAuthorizedMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.IsAuthorized. %v %v %v %v %v", p, p1, p2, p3, p4)
			return
		}

		input := m.IsAuthorizedMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetCoordinatorMockIsAuthorizedInput{p, p1, p2, p3, p4}, "JetCoordinator.IsAuthorized got unexpected parameters")

		result := m.IsAuthorizedMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.IsAuthorized")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.IsAuthorizedMock.mainExpectation != nil {

		input := m.IsAuthorizedMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetCoordinatorMockIsAuthorizedInput{p, p1, p2, p3, p4}, "JetCoordinator.IsAuthorized got unexpected parameters")
		}

		result := m.IsAuthorizedMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.IsAuthorized")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.IsAuthorizedFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.IsAuthorized. %v %v %v %v %v", p, p1, p2, p3, p4)
		return
	}

	return m.IsAuthorizedFunc(p, p1, p2, p3, p4)
}

//IsAuthorizedMinimockCounter returns a count of JetCoordinatorMock.IsAuthorizedFunc invocations
func (m *JetCoordinatorMock) IsAuthorizedMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.IsAuthorizedCounter)
}

//IsAuthorizedMinimockPreCounter returns the value of JetCoordinatorMock.IsAuthorized invocations
func (m *JetCoordinatorMock) IsAuthorizedMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.IsAuthorizedPreCounter)
}

//IsAuthorizedFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) IsAuthorizedFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.IsAuthorizedMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.IsAuthorizedCounter) == uint64(len(m.IsAuthorizedMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.IsAuthorizedMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.IsAuthorizedCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.IsAuthorizedFunc != nil {
		return atomic.LoadUint64(&m.IsAuthorizedCounter) > 0
	}

	return true
}

type mJetCoordinatorMockIsBeyondLimit struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockIsBeyondLimitExpectation
	expectationSeries []*JetCoordinatorMockIsBeyondLimitExpectation
}

type JetCoordinatorMockIsBeyondLimitExpectation struct {
	input  *JetCoordinatorMockIsBeyondLimitInput
	result *JetCoordinatorMockIsBeyondLimitResult
}

type JetCoordinatorMockIsBeyondLimitInput struct {
	p  context.Context
	p1 insolar.PulseNumber
	p2 insolar.PulseNumber
}

type JetCoordinatorMockIsBeyondLimitResult struct {
	r  bool
	r1 error
}

//Expect specifies that invocation of JetCoordinator.IsBeyondLimit is expected from 1 to Infinity times
func (m *mJetCoordinatorMockIsBeyondLimit) Expect(p context.Context, p1 insolar.PulseNumber, p2 insolar.PulseNumber) *mJetCoordinatorMockIsBeyondLimit {
	m.mock.IsBeyondLimitFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockIsBeyondLimitExpectation{}
	}
	m.mainExpectation.input = &JetCoordinatorMockIsBeyondLimitInput{p, p1, p2}
	return m
}

//Return specifies results of invocation of JetCoordinator.IsBeyondLimit
func (m *mJetCoordinatorMockIsBeyondLimit) Return(r bool, r1 error) *JetCoordinatorMock {
	m.mock.IsBeyondLimitFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockIsBeyondLimitExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockIsBeyondLimitResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.IsBeyondLimit is expected once
func (m *mJetCoordinatorMockIsBeyondLimit) ExpectOnce(p context.Context, p1 insolar.PulseNumber, p2 insolar.PulseNumber) *JetCoordinatorMockIsBeyondLimitExpectation {
	m.mock.IsBeyondLimitFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockIsBeyondLimitExpectation{}
	expectation.input = &JetCoordinatorMockIsBeyondLimitInput{p, p1, p2}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockIsBeyondLimitExpectation) Return(r bool, r1 error) {
	e.result = &JetCoordinatorMockIsBeyondLimitResult{r, r1}
}

//Set uses given function f as a mock of JetCoordinator.IsBeyondLimit method
func (m *mJetCoordinatorMockIsBeyondLimit) Set(f func(p context.Context, p1 insolar.PulseNumber, p2 insolar.PulseNumber) (r bool, r1 error)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.IsBeyondLimitFunc = f
	return m.mock
}

//IsBeyondLimit implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) IsBeyondLimit(p context.Context, p1 insolar.PulseNumber, p2 insolar.PulseNumber) (r bool, r1 error) {
	counter := atomic.AddUint64(&m.IsBeyondLimitPreCounter, 1)
	defer atomic.AddUint64(&m.IsBeyondLimitCounter, 1)

	if len(m.IsBeyondLimitMock.expectationSeries) > 0 {
		if counter > uint64(len(m.IsBeyondLimitMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.IsBeyondLimit. %v %v %v", p, p1, p2)
			return
		}

		input := m.IsBeyondLimitMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetCoordinatorMockIsBeyondLimitInput{p, p1, p2}, "JetCoordinator.IsBeyondLimit got unexpected parameters")

		result := m.IsBeyondLimitMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.IsBeyondLimit")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.IsBeyondLimitMock.mainExpectation != nil {

		input := m.IsBeyondLimitMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetCoordinatorMockIsBeyondLimitInput{p, p1, p2}, "JetCoordinator.IsBeyondLimit got unexpected parameters")
		}

		result := m.IsBeyondLimitMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.IsBeyondLimit")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.IsBeyondLimitFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.IsBeyondLimit. %v %v %v", p, p1, p2)
		return
	}

	return m.IsBeyondLimitFunc(p, p1, p2)
}

//IsBeyondLimitMinimockCounter returns a count of JetCoordinatorMock.IsBeyondLimitFunc invocations
func (m *JetCoordinatorMock) IsBeyondLimitMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.IsBeyondLimitCounter)
}

//IsBeyondLimitMinimockPreCounter returns the value of JetCoordinatorMock.IsBeyondLimit invocations
func (m *JetCoordinatorMock) IsBeyondLimitMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.IsBeyondLimitPreCounter)
}

//IsBeyondLimitFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) IsBeyondLimitFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.IsBeyondLimitMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.IsBeyondLimitCounter) == uint64(len(m.IsBeyondLimitMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.IsBeyondLimitMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.IsBeyondLimitCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.IsBeyondLimitFunc != nil {
		return atomic.LoadUint64(&m.IsBeyondLimitCounter) > 0
	}

	return true
}

type mJetCoordinatorMockLightExecutorForJet struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockLightExecutorForJetExpectation
	expectationSeries []*JetCoordinatorMockLightExecutorForJetExpectation
}

type JetCoordinatorMockLightExecutorForJetExpectation struct {
	input  *JetCoordinatorMockLightExecutorForJetInput
	result *JetCoordinatorMockLightExecutorForJetResult
}

type JetCoordinatorMockLightExecutorForJetInput struct {
	p  context.Context
	p1 insolar.ID
	p2 insolar.PulseNumber
}

type JetCoordinatorMockLightExecutorForJetResult struct {
	r  *insolar.Reference
	r1 error
}

//Expect specifies that invocation of JetCoordinator.LightExecutorForJet is expected from 1 to Infinity times
func (m *mJetCoordinatorMockLightExecutorForJet) Expect(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) *mJetCoordinatorMockLightExecutorForJet {
	m.mock.LightExecutorForJetFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockLightExecutorForJetExpectation{}
	}
	m.mainExpectation.input = &JetCoordinatorMockLightExecutorForJetInput{p, p1, p2}
	return m
}

//Return specifies results of invocation of JetCoordinator.LightExecutorForJet
func (m *mJetCoordinatorMockLightExecutorForJet) Return(r *insolar.Reference, r1 error) *JetCoordinatorMock {
	m.mock.LightExecutorForJetFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockLightExecutorForJetExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockLightExecutorForJetResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.LightExecutorForJet is expected once
func (m *mJetCoordinatorMockLightExecutorForJet) ExpectOnce(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) *JetCoordinatorMockLightExecutorForJetExpectation {
	m.mock.LightExecutorForJetFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockLightExecutorForJetExpectation{}
	expectation.input = &JetCoordinatorMockLightExecutorForJetInput{p, p1, p2}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockLightExecutorForJetExpectation) Return(r *insolar.Reference, r1 error) {
	e.result = &JetCoordinatorMockLightExecutorForJetResult{r, r1}
}

//Set uses given function f as a mock of JetCoordinator.LightExecutorForJet method
func (m *mJetCoordinatorMockLightExecutorForJet) Set(f func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r *insolar.Reference, r1 error)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.LightExecutorForJetFunc = f
	return m.mock
}

//LightExecutorForJet implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) LightExecutorForJet(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r *insolar.Reference, r1 error) {
	counter := atomic.AddUint64(&m.LightExecutorForJetPreCounter, 1)
	defer atomic.AddUint64(&m.LightExecutorForJetCounter, 1)

	if len(m.LightExecutorForJetMock.expectationSeries) > 0 {
		if counter > uint64(len(m.LightExecutorForJetMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.LightExecutorForJet. %v %v %v", p, p1, p2)
			return
		}

		input := m.LightExecutorForJetMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetCoordinatorMockLightExecutorForJetInput{p, p1, p2}, "JetCoordinator.LightExecutorForJet got unexpected parameters")

		result := m.LightExecutorForJetMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.LightExecutorForJet")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.LightExecutorForJetMock.mainExpectation != nil {

		input := m.LightExecutorForJetMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetCoordinatorMockLightExecutorForJetInput{p, p1, p2}, "JetCoordinator.LightExecutorForJet got unexpected parameters")
		}

		result := m.LightExecutorForJetMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.LightExecutorForJet")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.LightExecutorForJetFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.LightExecutorForJet. %v %v %v", p, p1, p2)
		return
	}

	return m.LightExecutorForJetFunc(p, p1, p2)
}

//LightExecutorForJetMinimockCounter returns a count of JetCoordinatorMock.LightExecutorForJetFunc invocations
func (m *JetCoordinatorMock) LightExecutorForJetMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.LightExecutorForJetCounter)
}

//LightExecutorForJetMinimockPreCounter returns the value of JetCoordinatorMock.LightExecutorForJet invocations
func (m *JetCoordinatorMock) LightExecutorForJetMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.LightExecutorForJetPreCounter)
}

//LightExecutorForJetFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) LightExecutorForJetFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.LightExecutorForJetMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.LightExecutorForJetCounter) == uint64(len(m.LightExecutorForJetMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.LightExecutorForJetMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.LightExecutorForJetCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.LightExecutorForJetFunc != nil {
		return atomic.LoadUint64(&m.LightExecutorForJetCounter) > 0
	}

	return true
}

type mJetCoordinatorMockLightExecutorForObject struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockLightExecutorForObjectExpectation
	expectationSeries []*JetCoordinatorMockLightExecutorForObjectExpectation
}

type JetCoordinatorMockLightExecutorForObjectExpectation struct {
	input  *JetCoordinatorMockLightExecutorForObjectInput
	result *JetCoordinatorMockLightExecutorForObjectResult
}

type JetCoordinatorMockLightExecutorForObjectInput struct {
	p  context.Context
	p1 insolar.ID
	p2 insolar.PulseNumber
}

type JetCoordinatorMockLightExecutorForObjectResult struct {
	r  *insolar.Reference
	r1 error
}

//Expect specifies that invocation of JetCoordinator.LightExecutorForObject is expected from 1 to Infinity times
func (m *mJetCoordinatorMockLightExecutorForObject) Expect(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) *mJetCoordinatorMockLightExecutorForObject {
	m.mock.LightExecutorForObjectFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockLightExecutorForObjectExpectation{}
	}
	m.mainExpectation.input = &JetCoordinatorMockLightExecutorForObjectInput{p, p1, p2}
	return m
}

//Return specifies results of invocation of JetCoordinator.LightExecutorForObject
func (m *mJetCoordinatorMockLightExecutorForObject) Return(r *insolar.Reference, r1 error) *JetCoordinatorMock {
	m.mock.LightExecutorForObjectFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockLightExecutorForObjectExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockLightExecutorForObjectResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.LightExecutorForObject is expected once
func (m *mJetCoordinatorMockLightExecutorForObject) ExpectOnce(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) *JetCoordinatorMockLightExecutorForObjectExpectation {
	m.mock.LightExecutorForObjectFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockLightExecutorForObjectExpectation{}
	expectation.input = &JetCoordinatorMockLightExecutorForObjectInput{p, p1, p2}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockLightExecutorForObjectExpectation) Return(r *insolar.Reference, r1 error) {
	e.result = &JetCoordinatorMockLightExecutorForObjectResult{r, r1}
}

//Set uses given function f as a mock of JetCoordinator.LightExecutorForObject method
func (m *mJetCoordinatorMockLightExecutorForObject) Set(f func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r *insolar.Reference, r1 error)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.LightExecutorForObjectFunc = f
	return m.mock
}

//LightExecutorForObject implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) LightExecutorForObject(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r *insolar.Reference, r1 error) {
	counter := atomic.AddUint64(&m.LightExecutorForObjectPreCounter, 1)
	defer atomic.AddUint64(&m.LightExecutorForObjectCounter, 1)

	if len(m.LightExecutorForObjectMock.expectationSeries) > 0 {
		if counter > uint64(len(m.LightExecutorForObjectMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.LightExecutorForObject. %v %v %v", p, p1, p2)
			return
		}

		input := m.LightExecutorForObjectMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetCoordinatorMockLightExecutorForObjectInput{p, p1, p2}, "JetCoordinator.LightExecutorForObject got unexpected parameters")

		result := m.LightExecutorForObjectMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.LightExecutorForObject")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.LightExecutorForObjectMock.mainExpectation != nil {

		input := m.LightExecutorForObjectMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetCoordinatorMockLightExecutorForObjectInput{p, p1, p2}, "JetCoordinator.LightExecutorForObject got unexpected parameters")
		}

		result := m.LightExecutorForObjectMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.LightExecutorForObject")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.LightExecutorForObjectFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.LightExecutorForObject. %v %v %v", p, p1, p2)
		return
	}

	return m.LightExecutorForObjectFunc(p, p1, p2)
}

//LightExecutorForObjectMinimockCounter returns a count of JetCoordinatorMock.LightExecutorForObjectFunc invocations
func (m *JetCoordinatorMock) LightExecutorForObjectMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.LightExecutorForObjectCounter)
}

//LightExecutorForObjectMinimockPreCounter returns the value of JetCoordinatorMock.LightExecutorForObject invocations
func (m *JetCoordinatorMock) LightExecutorForObjectMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.LightExecutorForObjectPreCounter)
}

//LightExecutorForObjectFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) LightExecutorForObjectFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.LightExecutorForObjectMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.LightExecutorForObjectCounter) == uint64(len(m.LightExecutorForObjectMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.LightExecutorForObjectMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.LightExecutorForObjectCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.LightExecutorForObjectFunc != nil {
		return atomic.LoadUint64(&m.LightExecutorForObjectCounter) > 0
	}

	return true
}

type mJetCoordinatorMockLightValidatorsForJet struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockLightValidatorsForJetExpectation
	expectationSeries []*JetCoordinatorMockLightValidatorsForJetExpectation
}

type JetCoordinatorMockLightValidatorsForJetExpectation struct {
	input  *JetCoordinatorMockLightValidatorsForJetInput
	result *JetCoordinatorMockLightValidatorsForJetResult
}

type JetCoordinatorMockLightValidatorsForJetInput struct {
	p  context.Context
	p1 insolar.ID
	p2 insolar.PulseNumber
}

type JetCoordinatorMockLightValidatorsForJetResult struct {
	r  []insolar.Reference
	r1 error
}

//Expect specifies that invocation of JetCoordinator.LightValidatorsForJet is expected from 1 to Infinity times
func (m *mJetCoordinatorMockLightValidatorsForJet) Expect(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) *mJetCoordinatorMockLightValidatorsForJet {
	m.mock.LightValidatorsForJetFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockLightValidatorsForJetExpectation{}
	}
	m.mainExpectation.input = &JetCoordinatorMockLightValidatorsForJetInput{p, p1, p2}
	return m
}

//Return specifies results of invocation of JetCoordinator.LightValidatorsForJet
func (m *mJetCoordinatorMockLightValidatorsForJet) Return(r []insolar.Reference, r1 error) *JetCoordinatorMock {
	m.mock.LightValidatorsForJetFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockLightValidatorsForJetExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockLightValidatorsForJetResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.LightValidatorsForJet is expected once
func (m *mJetCoordinatorMockLightValidatorsForJet) ExpectOnce(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) *JetCoordinatorMockLightValidatorsForJetExpectation {
	m.mock.LightValidatorsForJetFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockLightValidatorsForJetExpectation{}
	expectation.input = &JetCoordinatorMockLightValidatorsForJetInput{p, p1, p2}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockLightValidatorsForJetExpectation) Return(r []insolar.Reference, r1 error) {
	e.result = &JetCoordinatorMockLightValidatorsForJetResult{r, r1}
}

//Set uses given function f as a mock of JetCoordinator.LightValidatorsForJet method
func (m *mJetCoordinatorMockLightValidatorsForJet) Set(f func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r []insolar.Reference, r1 error)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.LightValidatorsForJetFunc = f
	return m.mock
}

//LightValidatorsForJet implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) LightValidatorsForJet(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r []insolar.Reference, r1 error) {
	counter := atomic.AddUint64(&m.LightValidatorsForJetPreCounter, 1)
	defer atomic.AddUint64(&m.LightValidatorsForJetCounter, 1)

	if len(m.LightValidatorsForJetMock.expectationSeries) > 0 {
		if counter > uint64(len(m.LightValidatorsForJetMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.LightValidatorsForJet. %v %v %v", p, p1, p2)
			return
		}

		input := m.LightValidatorsForJetMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetCoordinatorMockLightValidatorsForJetInput{p, p1, p2}, "JetCoordinator.LightValidatorsForJet got unexpected parameters")

		result := m.LightValidatorsForJetMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.LightValidatorsForJet")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.LightValidatorsForJetMock.mainExpectation != nil {

		input := m.LightValidatorsForJetMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetCoordinatorMockLightValidatorsForJetInput{p, p1, p2}, "JetCoordinator.LightValidatorsForJet got unexpected parameters")
		}

		result := m.LightValidatorsForJetMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.LightValidatorsForJet")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.LightValidatorsForJetFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.LightValidatorsForJet. %v %v %v", p, p1, p2)
		return
	}

	return m.LightValidatorsForJetFunc(p, p1, p2)
}

//LightValidatorsForJetMinimockCounter returns a count of JetCoordinatorMock.LightValidatorsForJetFunc invocations
func (m *JetCoordinatorMock) LightValidatorsForJetMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.LightValidatorsForJetCounter)
}

//LightValidatorsForJetMinimockPreCounter returns the value of JetCoordinatorMock.LightValidatorsForJet invocations
func (m *JetCoordinatorMock) LightValidatorsForJetMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.LightValidatorsForJetPreCounter)
}

//LightValidatorsForJetFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) LightValidatorsForJetFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.LightValidatorsForJetMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.LightValidatorsForJetCounter) == uint64(len(m.LightValidatorsForJetMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.LightValidatorsForJetMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.LightValidatorsForJetCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.LightValidatorsForJetFunc != nil {
		return atomic.LoadUint64(&m.LightValidatorsForJetCounter) > 0
	}

	return true
}

type mJetCoordinatorMockLightValidatorsForObject struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockLightValidatorsForObjectExpectation
	expectationSeries []*JetCoordinatorMockLightValidatorsForObjectExpectation
}

type JetCoordinatorMockLightValidatorsForObjectExpectation struct {
	input  *JetCoordinatorMockLightValidatorsForObjectInput
	result *JetCoordinatorMockLightValidatorsForObjectResult
}

type JetCoordinatorMockLightValidatorsForObjectInput struct {
	p  context.Context
	p1 insolar.ID
	p2 insolar.PulseNumber
}

type JetCoordinatorMockLightValidatorsForObjectResult struct {
	r  []insolar.Reference
	r1 error
}

//Expect specifies that invocation of JetCoordinator.LightValidatorsForObject is expected from 1 to Infinity times
func (m *mJetCoordinatorMockLightValidatorsForObject) Expect(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) *mJetCoordinatorMockLightValidatorsForObject {
	m.mock.LightValidatorsForObjectFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockLightValidatorsForObjectExpectation{}
	}
	m.mainExpectation.input = &JetCoordinatorMockLightValidatorsForObjectInput{p, p1, p2}
	return m
}

//Return specifies results of invocation of JetCoordinator.LightValidatorsForObject
func (m *mJetCoordinatorMockLightValidatorsForObject) Return(r []insolar.Reference, r1 error) *JetCoordinatorMock {
	m.mock.LightValidatorsForObjectFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockLightValidatorsForObjectExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockLightValidatorsForObjectResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.LightValidatorsForObject is expected once
func (m *mJetCoordinatorMockLightValidatorsForObject) ExpectOnce(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) *JetCoordinatorMockLightValidatorsForObjectExpectation {
	m.mock.LightValidatorsForObjectFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockLightValidatorsForObjectExpectation{}
	expectation.input = &JetCoordinatorMockLightValidatorsForObjectInput{p, p1, p2}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockLightValidatorsForObjectExpectation) Return(r []insolar.Reference, r1 error) {
	e.result = &JetCoordinatorMockLightValidatorsForObjectResult{r, r1}
}

//Set uses given function f as a mock of JetCoordinator.LightValidatorsForObject method
func (m *mJetCoordinatorMockLightValidatorsForObject) Set(f func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r []insolar.Reference, r1 error)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.LightValidatorsForObjectFunc = f
	return m.mock
}

//LightValidatorsForObject implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) LightValidatorsForObject(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r []insolar.Reference, r1 error) {
	counter := atomic.AddUint64(&m.LightValidatorsForObjectPreCounter, 1)
	defer atomic.AddUint64(&m.LightValidatorsForObjectCounter, 1)

	if len(m.LightValidatorsForObjectMock.expectationSeries) > 0 {
		if counter > uint64(len(m.LightValidatorsForObjectMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.LightValidatorsForObject. %v %v %v", p, p1, p2)
			return
		}

		input := m.LightValidatorsForObjectMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetCoordinatorMockLightValidatorsForObjectInput{p, p1, p2}, "JetCoordinator.LightValidatorsForObject got unexpected parameters")

		result := m.LightValidatorsForObjectMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.LightValidatorsForObject")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.LightValidatorsForObjectMock.mainExpectation != nil {

		input := m.LightValidatorsForObjectMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetCoordinatorMockLightValidatorsForObjectInput{p, p1, p2}, "JetCoordinator.LightValidatorsForObject got unexpected parameters")
		}

		result := m.LightValidatorsForObjectMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.LightValidatorsForObject")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.LightValidatorsForObjectFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.LightValidatorsForObject. %v %v %v", p, p1, p2)
		return
	}

	return m.LightValidatorsForObjectFunc(p, p1, p2)
}

//LightValidatorsForObjectMinimockCounter returns a count of JetCoordinatorMock.LightValidatorsForObjectFunc invocations
func (m *JetCoordinatorMock) LightValidatorsForObjectMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.LightValidatorsForObjectCounter)
}

//LightValidatorsForObjectMinimockPreCounter returns the value of JetCoordinatorMock.LightValidatorsForObject invocations
func (m *JetCoordinatorMock) LightValidatorsForObjectMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.LightValidatorsForObjectPreCounter)
}

//LightValidatorsForObjectFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) LightValidatorsForObjectFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.LightValidatorsForObjectMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.LightValidatorsForObjectCounter) == uint64(len(m.LightValidatorsForObjectMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.LightValidatorsForObjectMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.LightValidatorsForObjectCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.LightValidatorsForObjectFunc != nil {
		return atomic.LoadUint64(&m.LightValidatorsForObjectCounter) > 0
	}

	return true
}

type mJetCoordinatorMockMe struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockMeExpectation
	expectationSeries []*JetCoordinatorMockMeExpectation
}

type JetCoordinatorMockMeExpectation struct {
	result *JetCoordinatorMockMeResult
}

type JetCoordinatorMockMeResult struct {
	r insolar.Reference
}

//Expect specifies that invocation of JetCoordinator.Me is expected from 1 to Infinity times
func (m *mJetCoordinatorMockMe) Expect() *mJetCoordinatorMockMe {
	m.mock.MeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockMeExpectation{}
	}

	return m
}

//Return specifies results of invocation of JetCoordinator.Me
func (m *mJetCoordinatorMockMe) Return(r insolar.Reference) *JetCoordinatorMock {
	m.mock.MeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockMeExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockMeResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.Me is expected once
func (m *mJetCoordinatorMockMe) ExpectOnce() *JetCoordinatorMockMeExpectation {
	m.mock.MeFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockMeExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockMeExpectation) Return(r insolar.Reference) {
	e.result = &JetCoordinatorMockMeResult{r}
}

//Set uses given function f as a mock of JetCoordinator.Me method
func (m *mJetCoordinatorMockMe) Set(f func() (r insolar.Reference)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.MeFunc = f
	return m.mock
}

//Me implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) Me() (r insolar.Reference) {
	counter := atomic.AddUint64(&m.MePreCounter, 1)
	defer atomic.AddUint64(&m.MeCounter, 1)

	if len(m.MeMock.expectationSeries) > 0 {
		if counter > uint64(len(m.MeMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.Me.")
			return
		}

		result := m.MeMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.Me")
			return
		}

		r = result.r

		return
	}

	if m.MeMock.mainExpectation != nil {

		result := m.MeMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.Me")
		}

		r = result.r

		return
	}

	if m.MeFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.Me.")
		return
	}

	return m.MeFunc()
}

//MeMinimockCounter returns a count of JetCoordinatorMock.MeFunc invocations
func (m *JetCoordinatorMock) MeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.MeCounter)
}

//MeMinimockPreCounter returns the value of JetCoordinatorMock.Me invocations
func (m *JetCoordinatorMock) MeMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.MePreCounter)
}

//MeFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) MeFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.MeMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.MeCounter) == uint64(len(m.MeMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.MeMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.MeCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.MeFunc != nil {
		return atomic.LoadUint64(&m.MeCounter) > 0
	}

	return true
}

type mJetCoordinatorMockNodeForJet struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockNodeForJetExpectation
	expectationSeries []*JetCoordinatorMockNodeForJetExpectation
}

type JetCoordinatorMockNodeForJetExpectation struct {
	input  *JetCoordinatorMockNodeForJetInput
	result *JetCoordinatorMockNodeForJetResult
}

type JetCoordinatorMockNodeForJetInput struct {
	p  context.Context
	p1 insolar.ID
	p2 insolar.PulseNumber
	p3 insolar.PulseNumber
}

type JetCoordinatorMockNodeForJetResult struct {
	r  *insolar.Reference
	r1 error
}

//Expect specifies that invocation of JetCoordinator.NodeForJet is expected from 1 to Infinity times
func (m *mJetCoordinatorMockNodeForJet) Expect(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber, p3 insolar.PulseNumber) *mJetCoordinatorMockNodeForJet {
	m.mock.NodeForJetFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockNodeForJetExpectation{}
	}
	m.mainExpectation.input = &JetCoordinatorMockNodeForJetInput{p, p1, p2, p3}
	return m
}

//Return specifies results of invocation of JetCoordinator.NodeForJet
func (m *mJetCoordinatorMockNodeForJet) Return(r *insolar.Reference, r1 error) *JetCoordinatorMock {
	m.mock.NodeForJetFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockNodeForJetExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockNodeForJetResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.NodeForJet is expected once
func (m *mJetCoordinatorMockNodeForJet) ExpectOnce(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber, p3 insolar.PulseNumber) *JetCoordinatorMockNodeForJetExpectation {
	m.mock.NodeForJetFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockNodeForJetExpectation{}
	expectation.input = &JetCoordinatorMockNodeForJetInput{p, p1, p2, p3}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockNodeForJetExpectation) Return(r *insolar.Reference, r1 error) {
	e.result = &JetCoordinatorMockNodeForJetResult{r, r1}
}

//Set uses given function f as a mock of JetCoordinator.NodeForJet method
func (m *mJetCoordinatorMockNodeForJet) Set(f func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber, p3 insolar.PulseNumber) (r *insolar.Reference, r1 error)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.NodeForJetFunc = f
	return m.mock
}

//NodeForJet implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) NodeForJet(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber, p3 insolar.PulseNumber) (r *insolar.Reference, r1 error) {
	counter := atomic.AddUint64(&m.NodeForJetPreCounter, 1)
	defer atomic.AddUint64(&m.NodeForJetCounter, 1)

	if len(m.NodeForJetMock.expectationSeries) > 0 {
		if counter > uint64(len(m.NodeForJetMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.NodeForJet. %v %v %v %v", p, p1, p2, p3)
			return
		}

		input := m.NodeForJetMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetCoordinatorMockNodeForJetInput{p, p1, p2, p3}, "JetCoordinator.NodeForJet got unexpected parameters")

		result := m.NodeForJetMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.NodeForJet")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.NodeForJetMock.mainExpectation != nil {

		input := m.NodeForJetMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetCoordinatorMockNodeForJetInput{p, p1, p2, p3}, "JetCoordinator.NodeForJet got unexpected parameters")
		}

		result := m.NodeForJetMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.NodeForJet")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.NodeForJetFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.NodeForJet. %v %v %v %v", p, p1, p2, p3)
		return
	}

	return m.NodeForJetFunc(p, p1, p2, p3)
}

//NodeForJetMinimockCounter returns a count of JetCoordinatorMock.NodeForJetFunc invocations
func (m *JetCoordinatorMock) NodeForJetMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.NodeForJetCounter)
}

//NodeForJetMinimockPreCounter returns the value of JetCoordinatorMock.NodeForJet invocations
func (m *JetCoordinatorMock) NodeForJetMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.NodeForJetPreCounter)
}

//NodeForJetFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) NodeForJetFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.NodeForJetMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.NodeForJetCounter) == uint64(len(m.NodeForJetMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.NodeForJetMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.NodeForJetCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.NodeForJetFunc != nil {
		return atomic.LoadUint64(&m.NodeForJetCounter) > 0
	}

	return true
}

type mJetCoordinatorMockNodeForObject struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockNodeForObjectExpectation
	expectationSeries []*JetCoordinatorMockNodeForObjectExpectation
}

type JetCoordinatorMockNodeForObjectExpectation struct {
	input  *JetCoordinatorMockNodeForObjectInput
	result *JetCoordinatorMockNodeForObjectResult
}

type JetCoordinatorMockNodeForObjectInput struct {
	p  context.Context
	p1 insolar.ID
	p2 insolar.PulseNumber
	p3 insolar.PulseNumber
}

type JetCoordinatorMockNodeForObjectResult struct {
	r  *insolar.Reference
	r1 error
}

//Expect specifies that invocation of JetCoordinator.NodeForObject is expected from 1 to Infinity times
func (m *mJetCoordinatorMockNodeForObject) Expect(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber, p3 insolar.PulseNumber) *mJetCoordinatorMockNodeForObject {
	m.mock.NodeForObjectFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockNodeForObjectExpectation{}
	}
	m.mainExpectation.input = &JetCoordinatorMockNodeForObjectInput{p, p1, p2, p3}
	return m
}

//Return specifies results of invocation of JetCoordinator.NodeForObject
func (m *mJetCoordinatorMockNodeForObject) Return(r *insolar.Reference, r1 error) *JetCoordinatorMock {
	m.mock.NodeForObjectFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockNodeForObjectExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockNodeForObjectResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.NodeForObject is expected once
func (m *mJetCoordinatorMockNodeForObject) ExpectOnce(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber, p3 insolar.PulseNumber) *JetCoordinatorMockNodeForObjectExpectation {
	m.mock.NodeForObjectFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockNodeForObjectExpectation{}
	expectation.input = &JetCoordinatorMockNodeForObjectInput{p, p1, p2, p3}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockNodeForObjectExpectation) Return(r *insolar.Reference, r1 error) {
	e.result = &JetCoordinatorMockNodeForObjectResult{r, r1}
}

//Set uses given function f as a mock of JetCoordinator.NodeForObject method
func (m *mJetCoordinatorMockNodeForObject) Set(f func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber, p3 insolar.PulseNumber) (r *insolar.Reference, r1 error)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.NodeForObjectFunc = f
	return m.mock
}

//NodeForObject implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) NodeForObject(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber, p3 insolar.PulseNumber) (r *insolar.Reference, r1 error) {
	counter := atomic.AddUint64(&m.NodeForObjectPreCounter, 1)
	defer atomic.AddUint64(&m.NodeForObjectCounter, 1)

	if len(m.NodeForObjectMock.expectationSeries) > 0 {
		if counter > uint64(len(m.NodeForObjectMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.NodeForObject. %v %v %v %v", p, p1, p2, p3)
			return
		}

		input := m.NodeForObjectMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetCoordinatorMockNodeForObjectInput{p, p1, p2, p3}, "JetCoordinator.NodeForObject got unexpected parameters")

		result := m.NodeForObjectMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.NodeForObject")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.NodeForObjectMock.mainExpectation != nil {

		input := m.NodeForObjectMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetCoordinatorMockNodeForObjectInput{p, p1, p2, p3}, "JetCoordinator.NodeForObject got unexpected parameters")
		}

		result := m.NodeForObjectMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.NodeForObject")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.NodeForObjectFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.NodeForObject. %v %v %v %v", p, p1, p2, p3)
		return
	}

	return m.NodeForObjectFunc(p, p1, p2, p3)
}

//NodeForObjectMinimockCounter returns a count of JetCoordinatorMock.NodeForObjectFunc invocations
func (m *JetCoordinatorMock) NodeForObjectMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.NodeForObjectCounter)
}

//NodeForObjectMinimockPreCounter returns the value of JetCoordinatorMock.NodeForObject invocations
func (m *JetCoordinatorMock) NodeForObjectMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.NodeForObjectPreCounter)
}

//NodeForObjectFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) NodeForObjectFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.NodeForObjectMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.NodeForObjectCounter) == uint64(len(m.NodeForObjectMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.NodeForObjectMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.NodeForObjectCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.NodeForObjectFunc != nil {
		return atomic.LoadUint64(&m.NodeForObjectCounter) > 0
	}

	return true
}

type mJetCoordinatorMockQueryRole struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockQueryRoleExpectation
	expectationSeries []*JetCoordinatorMockQueryRoleExpectation
}

type JetCoordinatorMockQueryRoleExpectation struct {
	input  *JetCoordinatorMockQueryRoleInput
	result *JetCoordinatorMockQueryRoleResult
}

type JetCoordinatorMockQueryRoleInput struct {
	p  context.Context
	p1 insolar.DynamicRole
	p2 insolar.ID
	p3 insolar.PulseNumber
}

type JetCoordinatorMockQueryRoleResult struct {
	r  []insolar.Reference
	r1 error
}

//Expect specifies that invocation of JetCoordinator.QueryRole is expected from 1 to Infinity times
func (m *mJetCoordinatorMockQueryRole) Expect(p context.Context, p1 insolar.DynamicRole, p2 insolar.ID, p3 insolar.PulseNumber) *mJetCoordinatorMockQueryRole {
	m.mock.QueryRoleFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockQueryRoleExpectation{}
	}
	m.mainExpectation.input = &JetCoordinatorMockQueryRoleInput{p, p1, p2, p3}
	return m
}

//Return specifies results of invocation of JetCoordinator.QueryRole
func (m *mJetCoordinatorMockQueryRole) Return(r []insolar.Reference, r1 error) *JetCoordinatorMock {
	m.mock.QueryRoleFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockQueryRoleExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockQueryRoleResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.QueryRole is expected once
func (m *mJetCoordinatorMockQueryRole) ExpectOnce(p context.Context, p1 insolar.DynamicRole, p2 insolar.ID, p3 insolar.PulseNumber) *JetCoordinatorMockQueryRoleExpectation {
	m.mock.QueryRoleFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockQueryRoleExpectation{}
	expectation.input = &JetCoordinatorMockQueryRoleInput{p, p1, p2, p3}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockQueryRoleExpectation) Return(r []insolar.Reference, r1 error) {
	e.result = &JetCoordinatorMockQueryRoleResult{r, r1}
}

//Set uses given function f as a mock of JetCoordinator.QueryRole method
func (m *mJetCoordinatorMockQueryRole) Set(f func(p context.Context, p1 insolar.DynamicRole, p2 insolar.ID, p3 insolar.PulseNumber) (r []insolar.Reference, r1 error)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.QueryRoleFunc = f
	return m.mock
}

//QueryRole implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) QueryRole(p context.Context, p1 insolar.DynamicRole, p2 insolar.ID, p3 insolar.PulseNumber) (r []insolar.Reference, r1 error) {
	counter := atomic.AddUint64(&m.QueryRolePreCounter, 1)
	defer atomic.AddUint64(&m.QueryRoleCounter, 1)

	if len(m.QueryRoleMock.expectationSeries) > 0 {
		if counter > uint64(len(m.QueryRoleMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.QueryRole. %v %v %v %v", p, p1, p2, p3)
			return
		}

		input := m.QueryRoleMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetCoordinatorMockQueryRoleInput{p, p1, p2, p3}, "JetCoordinator.QueryRole got unexpected parameters")

		result := m.QueryRoleMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.QueryRole")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.QueryRoleMock.mainExpectation != nil {

		input := m.QueryRoleMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetCoordinatorMockQueryRoleInput{p, p1, p2, p3}, "JetCoordinator.QueryRole got unexpected parameters")
		}

		result := m.QueryRoleMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.QueryRole")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.QueryRoleFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.QueryRole. %v %v %v %v", p, p1, p2, p3)
		return
	}

	return m.QueryRoleFunc(p, p1, p2, p3)
}

//QueryRoleMinimockCounter returns a count of JetCoordinatorMock.QueryRoleFunc invocations
func (m *JetCoordinatorMock) QueryRoleMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryRoleCounter)
}

//QueryRoleMinimockPreCounter returns the value of JetCoordinatorMock.QueryRole invocations
func (m *JetCoordinatorMock) QueryRoleMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.QueryRolePreCounter)
}

//QueryRoleFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) QueryRoleFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.QueryRoleMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.QueryRoleCounter) == uint64(len(m.QueryRoleMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.QueryRoleMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.QueryRoleCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.QueryRoleFunc != nil {
		return atomic.LoadUint64(&m.QueryRoleCounter) > 0
	}

	return true
}

type mJetCoordinatorMockVirtualExecutorForObject struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockVirtualExecutorForObjectExpectation
	expectationSeries []*JetCoordinatorMockVirtualExecutorForObjectExpectation
}

type JetCoordinatorMockVirtualExecutorForObjectExpectation struct {
	input  *JetCoordinatorMockVirtualExecutorForObjectInput
	result *JetCoordinatorMockVirtualExecutorForObjectResult
}

type JetCoordinatorMockVirtualExecutorForObjectInput struct {
	p  context.Context
	p1 insolar.ID
	p2 insolar.PulseNumber
}

type JetCoordinatorMockVirtualExecutorForObjectResult struct {
	r  *insolar.Reference
	r1 error
}

//Expect specifies that invocation of JetCoordinator.VirtualExecutorForObject is expected from 1 to Infinity times
func (m *mJetCoordinatorMockVirtualExecutorForObject) Expect(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) *mJetCoordinatorMockVirtualExecutorForObject {
	m.mock.VirtualExecutorForObjectFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockVirtualExecutorForObjectExpectation{}
	}
	m.mainExpectation.input = &JetCoordinatorMockVirtualExecutorForObjectInput{p, p1, p2}
	return m
}

//Return specifies results of invocation of JetCoordinator.VirtualExecutorForObject
func (m *mJetCoordinatorMockVirtualExecutorForObject) Return(r *insolar.Reference, r1 error) *JetCoordinatorMock {
	m.mock.VirtualExecutorForObjectFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockVirtualExecutorForObjectExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockVirtualExecutorForObjectResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.VirtualExecutorForObject is expected once
func (m *mJetCoordinatorMockVirtualExecutorForObject) ExpectOnce(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) *JetCoordinatorMockVirtualExecutorForObjectExpectation {
	m.mock.VirtualExecutorForObjectFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockVirtualExecutorForObjectExpectation{}
	expectation.input = &JetCoordinatorMockVirtualExecutorForObjectInput{p, p1, p2}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockVirtualExecutorForObjectExpectation) Return(r *insolar.Reference, r1 error) {
	e.result = &JetCoordinatorMockVirtualExecutorForObjectResult{r, r1}
}

//Set uses given function f as a mock of JetCoordinator.VirtualExecutorForObject method
func (m *mJetCoordinatorMockVirtualExecutorForObject) Set(f func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r *insolar.Reference, r1 error)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.VirtualExecutorForObjectFunc = f
	return m.mock
}

//VirtualExecutorForObject implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) VirtualExecutorForObject(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r *insolar.Reference, r1 error) {
	counter := atomic.AddUint64(&m.VirtualExecutorForObjectPreCounter, 1)
	defer atomic.AddUint64(&m.VirtualExecutorForObjectCounter, 1)

	if len(m.VirtualExecutorForObjectMock.expectationSeries) > 0 {
		if counter > uint64(len(m.VirtualExecutorForObjectMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.VirtualExecutorForObject. %v %v %v", p, p1, p2)
			return
		}

		input := m.VirtualExecutorForObjectMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetCoordinatorMockVirtualExecutorForObjectInput{p, p1, p2}, "JetCoordinator.VirtualExecutorForObject got unexpected parameters")

		result := m.VirtualExecutorForObjectMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.VirtualExecutorForObject")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.VirtualExecutorForObjectMock.mainExpectation != nil {

		input := m.VirtualExecutorForObjectMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetCoordinatorMockVirtualExecutorForObjectInput{p, p1, p2}, "JetCoordinator.VirtualExecutorForObject got unexpected parameters")
		}

		result := m.VirtualExecutorForObjectMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.VirtualExecutorForObject")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.VirtualExecutorForObjectFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.VirtualExecutorForObject. %v %v %v", p, p1, p2)
		return
	}

	return m.VirtualExecutorForObjectFunc(p, p1, p2)
}

//VirtualExecutorForObjectMinimockCounter returns a count of JetCoordinatorMock.VirtualExecutorForObjectFunc invocations
func (m *JetCoordinatorMock) VirtualExecutorForObjectMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.VirtualExecutorForObjectCounter)
}

//VirtualExecutorForObjectMinimockPreCounter returns the value of JetCoordinatorMock.VirtualExecutorForObject invocations
func (m *JetCoordinatorMock) VirtualExecutorForObjectMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.VirtualExecutorForObjectPreCounter)
}

//VirtualExecutorForObjectFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) VirtualExecutorForObjectFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.VirtualExecutorForObjectMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.VirtualExecutorForObjectCounter) == uint64(len(m.VirtualExecutorForObjectMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.VirtualExecutorForObjectMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.VirtualExecutorForObjectCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.VirtualExecutorForObjectFunc != nil {
		return atomic.LoadUint64(&m.VirtualExecutorForObjectCounter) > 0
	}

	return true
}

type mJetCoordinatorMockVirtualValidatorsForObject struct {
	mock              *JetCoordinatorMock
	mainExpectation   *JetCoordinatorMockVirtualValidatorsForObjectExpectation
	expectationSeries []*JetCoordinatorMockVirtualValidatorsForObjectExpectation
}

type JetCoordinatorMockVirtualValidatorsForObjectExpectation struct {
	input  *JetCoordinatorMockVirtualValidatorsForObjectInput
	result *JetCoordinatorMockVirtualValidatorsForObjectResult
}

type JetCoordinatorMockVirtualValidatorsForObjectInput struct {
	p  context.Context
	p1 insolar.ID
	p2 insolar.PulseNumber
}

type JetCoordinatorMockVirtualValidatorsForObjectResult struct {
	r  []insolar.Reference
	r1 error
}

//Expect specifies that invocation of JetCoordinator.VirtualValidatorsForObject is expected from 1 to Infinity times
func (m *mJetCoordinatorMockVirtualValidatorsForObject) Expect(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) *mJetCoordinatorMockVirtualValidatorsForObject {
	m.mock.VirtualValidatorsForObjectFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockVirtualValidatorsForObjectExpectation{}
	}
	m.mainExpectation.input = &JetCoordinatorMockVirtualValidatorsForObjectInput{p, p1, p2}
	return m
}

//Return specifies results of invocation of JetCoordinator.VirtualValidatorsForObject
func (m *mJetCoordinatorMockVirtualValidatorsForObject) Return(r []insolar.Reference, r1 error) *JetCoordinatorMock {
	m.mock.VirtualValidatorsForObjectFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetCoordinatorMockVirtualValidatorsForObjectExpectation{}
	}
	m.mainExpectation.result = &JetCoordinatorMockVirtualValidatorsForObjectResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetCoordinator.VirtualValidatorsForObject is expected once
func (m *mJetCoordinatorMockVirtualValidatorsForObject) ExpectOnce(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) *JetCoordinatorMockVirtualValidatorsForObjectExpectation {
	m.mock.VirtualValidatorsForObjectFunc = nil
	m.mainExpectation = nil

	expectation := &JetCoordinatorMockVirtualValidatorsForObjectExpectation{}
	expectation.input = &JetCoordinatorMockVirtualValidatorsForObjectInput{p, p1, p2}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetCoordinatorMockVirtualValidatorsForObjectExpectation) Return(r []insolar.Reference, r1 error) {
	e.result = &JetCoordinatorMockVirtualValidatorsForObjectResult{r, r1}
}

//Set uses given function f as a mock of JetCoordinator.VirtualValidatorsForObject method
func (m *mJetCoordinatorMockVirtualValidatorsForObject) Set(f func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r []insolar.Reference, r1 error)) *JetCoordinatorMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.VirtualValidatorsForObjectFunc = f
	return m.mock
}

//VirtualValidatorsForObject implements github.com/insolar/insolar/insolar.JetCoordinator interface
func (m *JetCoordinatorMock) VirtualValidatorsForObject(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber) (r []insolar.Reference, r1 error) {
	counter := atomic.AddUint64(&m.VirtualValidatorsForObjectPreCounter, 1)
	defer atomic.AddUint64(&m.VirtualValidatorsForObjectCounter, 1)

	if len(m.VirtualValidatorsForObjectMock.expectationSeries) > 0 {
		if counter > uint64(len(m.VirtualValidatorsForObjectMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetCoordinatorMock.VirtualValidatorsForObject. %v %v %v", p, p1, p2)
			return
		}

		input := m.VirtualValidatorsForObjectMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetCoordinatorMockVirtualValidatorsForObjectInput{p, p1, p2}, "JetCoordinator.VirtualValidatorsForObject got unexpected parameters")

		result := m.VirtualValidatorsForObjectMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.VirtualValidatorsForObject")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.VirtualValidatorsForObjectMock.mainExpectation != nil {

		input := m.VirtualValidatorsForObjectMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetCoordinatorMockVirtualValidatorsForObjectInput{p, p1, p2}, "JetCoordinator.VirtualValidatorsForObject got unexpected parameters")
		}

		result := m.VirtualValidatorsForObjectMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetCoordinatorMock.VirtualValidatorsForObject")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.VirtualValidatorsForObjectFunc == nil {
		m.t.Fatalf("Unexpected call to JetCoordinatorMock.VirtualValidatorsForObject. %v %v %v", p, p1, p2)
		return
	}

	return m.VirtualValidatorsForObjectFunc(p, p1, p2)
}

//VirtualValidatorsForObjectMinimockCounter returns a count of JetCoordinatorMock.VirtualValidatorsForObjectFunc invocations
func (m *JetCoordinatorMock) VirtualValidatorsForObjectMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.VirtualValidatorsForObjectCounter)
}

//VirtualValidatorsForObjectMinimockPreCounter returns the value of JetCoordinatorMock.VirtualValidatorsForObject invocations
func (m *JetCoordinatorMock) VirtualValidatorsForObjectMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.VirtualValidatorsForObjectPreCounter)
}

//VirtualValidatorsForObjectFinished returns true if mock invocations count is ok
func (m *JetCoordinatorMock) VirtualValidatorsForObjectFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.VirtualValidatorsForObjectMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.VirtualValidatorsForObjectCounter) == uint64(len(m.VirtualValidatorsForObjectMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.VirtualValidatorsForObjectMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.VirtualValidatorsForObjectCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.VirtualValidatorsForObjectFunc != nil {
		return atomic.LoadUint64(&m.VirtualValidatorsForObjectCounter) > 0
	}

	return true
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *JetCoordinatorMock) ValidateCallCounters() {

	if !m.HeavyFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.Heavy")
	}

	if !m.IsAuthorizedFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.IsAuthorized")
	}

	if !m.IsBeyondLimitFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.IsBeyondLimit")
	}

	if !m.LightExecutorForJetFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.LightExecutorForJet")
	}

	if !m.LightExecutorForObjectFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.LightExecutorForObject")
	}

	if !m.LightValidatorsForJetFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.LightValidatorsForJet")
	}

	if !m.LightValidatorsForObjectFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.LightValidatorsForObject")
	}

	if !m.MeFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.Me")
	}

	if !m.NodeForJetFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.NodeForJet")
	}

	if !m.NodeForObjectFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.NodeForObject")
	}

	if !m.QueryRoleFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.QueryRole")
	}

	if !m.VirtualExecutorForObjectFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.VirtualExecutorForObject")
	}

	if !m.VirtualValidatorsForObjectFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.VirtualValidatorsForObject")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *JetCoordinatorMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *JetCoordinatorMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *JetCoordinatorMock) MinimockFinish() {

	if !m.HeavyFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.Heavy")
	}

	if !m.IsAuthorizedFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.IsAuthorized")
	}

	if !m.IsBeyondLimitFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.IsBeyondLimit")
	}

	if !m.LightExecutorForJetFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.LightExecutorForJet")
	}

	if !m.LightExecutorForObjectFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.LightExecutorForObject")
	}

	if !m.LightValidatorsForJetFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.LightValidatorsForJet")
	}

	if !m.LightValidatorsForObjectFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.LightValidatorsForObject")
	}

	if !m.MeFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.Me")
	}

	if !m.NodeForJetFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.NodeForJet")
	}

	if !m.NodeForObjectFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.NodeForObject")
	}

	if !m.QueryRoleFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.QueryRole")
	}

	if !m.VirtualExecutorForObjectFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.VirtualExecutorForObject")
	}

	if !m.VirtualValidatorsForObjectFinished() {
		m.t.Fatal("Expected call to JetCoordinatorMock.VirtualValidatorsForObject")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *JetCoordinatorMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *JetCoordinatorMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && m.HeavyFinished()
		ok = ok && m.IsAuthorizedFinished()
		ok = ok && m.IsBeyondLimitFinished()
		ok = ok && m.LightExecutorForJetFinished()
		ok = ok && m.LightExecutorForObjectFinished()
		ok = ok && m.LightValidatorsForJetFinished()
		ok = ok && m.LightValidatorsForObjectFinished()
		ok = ok && m.MeFinished()
		ok = ok && m.NodeForJetFinished()
		ok = ok && m.NodeForObjectFinished()
		ok = ok && m.QueryRoleFinished()
		ok = ok && m.VirtualExecutorForObjectFinished()
		ok = ok && m.VirtualValidatorsForObjectFinished()

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if !m.HeavyFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.Heavy")
			}

			if !m.IsAuthorizedFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.IsAuthorized")
			}

			if !m.IsBeyondLimitFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.IsBeyondLimit")
			}

			if !m.LightExecutorForJetFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.LightExecutorForJet")
			}

			if !m.LightExecutorForObjectFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.LightExecutorForObject")
			}

			if !m.LightValidatorsForJetFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.LightValidatorsForJet")
			}

			if !m.LightValidatorsForObjectFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.LightValidatorsForObject")
			}

			if !m.MeFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.Me")
			}

			if !m.NodeForJetFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.NodeForJet")
			}

			if !m.NodeForObjectFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.NodeForObject")
			}

			if !m.QueryRoleFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.QueryRole")
			}

			if !m.VirtualExecutorForObjectFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.VirtualExecutorForObject")
			}

			if !m.VirtualValidatorsForObjectFinished() {
				m.t.Error("Expected call to JetCoordinatorMock.VirtualValidatorsForObject")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *JetCoordinatorMock) AllMocksCalled() bool {

	if !m.HeavyFinished() {
		return false
	}

	if !m.IsAuthorizedFinished() {
		return false
	}

	if !m.IsBeyondLimitFinished() {
		return false
	}

	if !m.LightExecutorForJetFinished() {
		return false
	}

	if !m.LightExecutorForObjectFinished() {
		return false
	}

	if !m.LightValidatorsForJetFinished() {
		return false
	}

	if !m.LightValidatorsForObjectFinished() {
		return false
	}

	if !m.MeFinished() {
		return false
	}

	if !m.NodeForJetFinished() {
		return false
	}

	if !m.NodeForObjectFinished() {
		return false
	}

	if !m.QueryRoleFinished() {
		return false
	}

	if !m.VirtualExecutorForObjectFinished() {
		return false
	}

	if !m.VirtualValidatorsForObjectFinished() {
		return false
	}

	return true
}
