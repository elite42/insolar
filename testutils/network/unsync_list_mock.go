package network

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "UnsyncList" can be found in github.com/insolar/insolar/network
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	packets "github.com/insolar/insolar/consensus/packets"
	core "github.com/insolar/insolar/core"

	testify_assert "github.com/stretchr/testify/assert"
)

//UnsyncListMock implements github.com/insolar/insolar/network.UnsyncList
type UnsyncListMock struct {
	t minimock.Tester

	AddNodeFunc       func(p core.Node, p1 uint16)
	AddNodeCounter    uint64
	AddNodePreCounter uint64
	AddNodeMock       mUnsyncListMockAddNode

	AddProofFunc       func(p core.RecordRef, p1 *packets.NodePulseProof)
	AddProofCounter    uint64
	AddProofPreCounter uint64
	AddProofMock       mUnsyncListMockAddProof

	GetActiveNodeFunc       func(p core.RecordRef) (r core.Node)
	GetActiveNodeCounter    uint64
	GetActiveNodePreCounter uint64
	GetActiveNodeMock       mUnsyncListMockGetActiveNode

	GetActiveNodesFunc       func() (r []core.Node)
	GetActiveNodesCounter    uint64
	GetActiveNodesPreCounter uint64
	GetActiveNodesMock       mUnsyncListMockGetActiveNodes

	GetGlobuleHashSignatureFunc       func(p core.RecordRef) (r packets.GlobuleHashSignature, r1 bool)
	GetGlobuleHashSignatureCounter    uint64
	GetGlobuleHashSignaturePreCounter uint64
	GetGlobuleHashSignatureMock       mUnsyncListMockGetGlobuleHashSignature

	GetOriginFunc       func() (r core.Node)
	GetOriginCounter    uint64
	GetOriginPreCounter uint64
	GetOriginMock       mUnsyncListMockGetOrigin

	GetProofFunc       func(p core.RecordRef) (r *packets.NodePulseProof)
	GetProofCounter    uint64
	GetProofPreCounter uint64
	GetProofMock       mUnsyncListMockGetProof

	IndexToRefFunc       func(p int) (r core.RecordRef, r1 error)
	IndexToRefCounter    uint64
	IndexToRefPreCounter uint64
	IndexToRefMock       mUnsyncListMockIndexToRef

	LengthFunc       func() (r int)
	LengthCounter    uint64
	LengthPreCounter uint64
	LengthMock       mUnsyncListMockLength

	RefToIndexFunc       func(p core.RecordRef) (r int, r1 error)
	RefToIndexCounter    uint64
	RefToIndexPreCounter uint64
	RefToIndexMock       mUnsyncListMockRefToIndex

	RemoveNodeFunc       func(p core.RecordRef)
	RemoveNodeCounter    uint64
	RemoveNodePreCounter uint64
	RemoveNodeMock       mUnsyncListMockRemoveNode

	SetGlobuleHashSignatureFunc       func(p core.RecordRef, p1 packets.GlobuleHashSignature)
	SetGlobuleHashSignatureCounter    uint64
	SetGlobuleHashSignaturePreCounter uint64
	SetGlobuleHashSignatureMock       mUnsyncListMockSetGlobuleHashSignature
}

//NewUnsyncListMock returns a mock for github.com/insolar/insolar/network.UnsyncList
func NewUnsyncListMock(t minimock.Tester) *UnsyncListMock {
	m := &UnsyncListMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddNodeMock = mUnsyncListMockAddNode{mock: m}
	m.AddProofMock = mUnsyncListMockAddProof{mock: m}
	m.GetActiveNodeMock = mUnsyncListMockGetActiveNode{mock: m}
	m.GetActiveNodesMock = mUnsyncListMockGetActiveNodes{mock: m}
	m.GetGlobuleHashSignatureMock = mUnsyncListMockGetGlobuleHashSignature{mock: m}
	m.GetOriginMock = mUnsyncListMockGetOrigin{mock: m}
	m.GetProofMock = mUnsyncListMockGetProof{mock: m}
	m.IndexToRefMock = mUnsyncListMockIndexToRef{mock: m}
	m.LengthMock = mUnsyncListMockLength{mock: m}
	m.RefToIndexMock = mUnsyncListMockRefToIndex{mock: m}
	m.RemoveNodeMock = mUnsyncListMockRemoveNode{mock: m}
	m.SetGlobuleHashSignatureMock = mUnsyncListMockSetGlobuleHashSignature{mock: m}

	return m
}

type mUnsyncListMockAddNode struct {
	mock              *UnsyncListMock
	mainExpectation   *UnsyncListMockAddNodeExpectation
	expectationSeries []*UnsyncListMockAddNodeExpectation
}

type UnsyncListMockAddNodeExpectation struct {
	input *UnsyncListMockAddNodeInput
}

type UnsyncListMockAddNodeInput struct {
	p  core.Node
	p1 uint16
}

//Expect specifies that invocation of UnsyncList.AddNode is expected from 1 to Infinity times
func (m *mUnsyncListMockAddNode) Expect(p core.Node, p1 uint16) *mUnsyncListMockAddNode {
	m.mock.AddNodeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockAddNodeExpectation{}
	}
	m.mainExpectation.input = &UnsyncListMockAddNodeInput{p, p1}
	return m
}

//Return specifies results of invocation of UnsyncList.AddNode
func (m *mUnsyncListMockAddNode) Return() *UnsyncListMock {
	m.mock.AddNodeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockAddNodeExpectation{}
	}

	return m.mock
}

//ExpectOnce specifies that invocation of UnsyncList.AddNode is expected once
func (m *mUnsyncListMockAddNode) ExpectOnce(p core.Node, p1 uint16) *UnsyncListMockAddNodeExpectation {
	m.mock.AddNodeFunc = nil
	m.mainExpectation = nil

	expectation := &UnsyncListMockAddNodeExpectation{}
	expectation.input = &UnsyncListMockAddNodeInput{p, p1}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

//Set uses given function f as a mock of UnsyncList.AddNode method
func (m *mUnsyncListMockAddNode) Set(f func(p core.Node, p1 uint16)) *UnsyncListMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.AddNodeFunc = f
	return m.mock
}

//AddNode implements github.com/insolar/insolar/network.UnsyncList interface
func (m *UnsyncListMock) AddNode(p core.Node, p1 uint16) {
	counter := atomic.AddUint64(&m.AddNodePreCounter, 1)
	defer atomic.AddUint64(&m.AddNodeCounter, 1)

	if len(m.AddNodeMock.expectationSeries) > 0 {
		if counter > uint64(len(m.AddNodeMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to UnsyncListMock.AddNode. %v %v", p, p1)
			return
		}

		input := m.AddNodeMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, UnsyncListMockAddNodeInput{p, p1}, "UnsyncList.AddNode got unexpected parameters")

		return
	}

	if m.AddNodeMock.mainExpectation != nil {

		input := m.AddNodeMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, UnsyncListMockAddNodeInput{p, p1}, "UnsyncList.AddNode got unexpected parameters")
		}

		return
	}

	if m.AddNodeFunc == nil {
		m.t.Fatalf("Unexpected call to UnsyncListMock.AddNode. %v %v", p, p1)
		return
	}

	m.AddNodeFunc(p, p1)
}

//AddNodeMinimockCounter returns a count of UnsyncListMock.AddNodeFunc invocations
func (m *UnsyncListMock) AddNodeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.AddNodeCounter)
}

//AddNodeMinimockPreCounter returns the value of UnsyncListMock.AddNode invocations
func (m *UnsyncListMock) AddNodeMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.AddNodePreCounter)
}

//AddNodeFinished returns true if mock invocations count is ok
func (m *UnsyncListMock) AddNodeFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.AddNodeMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.AddNodeCounter) == uint64(len(m.AddNodeMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.AddNodeMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.AddNodeCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.AddNodeFunc != nil {
		return atomic.LoadUint64(&m.AddNodeCounter) > 0
	}

	return true
}

type mUnsyncListMockAddProof struct {
	mock              *UnsyncListMock
	mainExpectation   *UnsyncListMockAddProofExpectation
	expectationSeries []*UnsyncListMockAddProofExpectation
}

type UnsyncListMockAddProofExpectation struct {
	input *UnsyncListMockAddProofInput
}

type UnsyncListMockAddProofInput struct {
	p  core.RecordRef
	p1 *packets.NodePulseProof
}

//Expect specifies that invocation of UnsyncList.AddProof is expected from 1 to Infinity times
func (m *mUnsyncListMockAddProof) Expect(p core.RecordRef, p1 *packets.NodePulseProof) *mUnsyncListMockAddProof {
	m.mock.AddProofFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockAddProofExpectation{}
	}
	m.mainExpectation.input = &UnsyncListMockAddProofInput{p, p1}
	return m
}

//Return specifies results of invocation of UnsyncList.AddProof
func (m *mUnsyncListMockAddProof) Return() *UnsyncListMock {
	m.mock.AddProofFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockAddProofExpectation{}
	}

	return m.mock
}

//ExpectOnce specifies that invocation of UnsyncList.AddProof is expected once
func (m *mUnsyncListMockAddProof) ExpectOnce(p core.RecordRef, p1 *packets.NodePulseProof) *UnsyncListMockAddProofExpectation {
	m.mock.AddProofFunc = nil
	m.mainExpectation = nil

	expectation := &UnsyncListMockAddProofExpectation{}
	expectation.input = &UnsyncListMockAddProofInput{p, p1}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

//Set uses given function f as a mock of UnsyncList.AddProof method
func (m *mUnsyncListMockAddProof) Set(f func(p core.RecordRef, p1 *packets.NodePulseProof)) *UnsyncListMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.AddProofFunc = f
	return m.mock
}

//AddProof implements github.com/insolar/insolar/network.UnsyncList interface
func (m *UnsyncListMock) AddProof(p core.RecordRef, p1 *packets.NodePulseProof) {
	counter := atomic.AddUint64(&m.AddProofPreCounter, 1)
	defer atomic.AddUint64(&m.AddProofCounter, 1)

	if len(m.AddProofMock.expectationSeries) > 0 {
		if counter > uint64(len(m.AddProofMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to UnsyncListMock.AddProof. %v %v", p, p1)
			return
		}

		input := m.AddProofMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, UnsyncListMockAddProofInput{p, p1}, "UnsyncList.AddProof got unexpected parameters")

		return
	}

	if m.AddProofMock.mainExpectation != nil {

		input := m.AddProofMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, UnsyncListMockAddProofInput{p, p1}, "UnsyncList.AddProof got unexpected parameters")
		}

		return
	}

	if m.AddProofFunc == nil {
		m.t.Fatalf("Unexpected call to UnsyncListMock.AddProof. %v %v", p, p1)
		return
	}

	m.AddProofFunc(p, p1)
}

//AddProofMinimockCounter returns a count of UnsyncListMock.AddProofFunc invocations
func (m *UnsyncListMock) AddProofMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.AddProofCounter)
}

//AddProofMinimockPreCounter returns the value of UnsyncListMock.AddProof invocations
func (m *UnsyncListMock) AddProofMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.AddProofPreCounter)
}

//AddProofFinished returns true if mock invocations count is ok
func (m *UnsyncListMock) AddProofFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.AddProofMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.AddProofCounter) == uint64(len(m.AddProofMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.AddProofMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.AddProofCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.AddProofFunc != nil {
		return atomic.LoadUint64(&m.AddProofCounter) > 0
	}

	return true
}

type mUnsyncListMockGetActiveNode struct {
	mock              *UnsyncListMock
	mainExpectation   *UnsyncListMockGetActiveNodeExpectation
	expectationSeries []*UnsyncListMockGetActiveNodeExpectation
}

type UnsyncListMockGetActiveNodeExpectation struct {
	input  *UnsyncListMockGetActiveNodeInput
	result *UnsyncListMockGetActiveNodeResult
}

type UnsyncListMockGetActiveNodeInput struct {
	p core.RecordRef
}

type UnsyncListMockGetActiveNodeResult struct {
	r core.Node
}

//Expect specifies that invocation of UnsyncList.GetActiveNode is expected from 1 to Infinity times
func (m *mUnsyncListMockGetActiveNode) Expect(p core.RecordRef) *mUnsyncListMockGetActiveNode {
	m.mock.GetActiveNodeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockGetActiveNodeExpectation{}
	}
	m.mainExpectation.input = &UnsyncListMockGetActiveNodeInput{p}
	return m
}

//Return specifies results of invocation of UnsyncList.GetActiveNode
func (m *mUnsyncListMockGetActiveNode) Return(r core.Node) *UnsyncListMock {
	m.mock.GetActiveNodeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockGetActiveNodeExpectation{}
	}
	m.mainExpectation.result = &UnsyncListMockGetActiveNodeResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of UnsyncList.GetActiveNode is expected once
func (m *mUnsyncListMockGetActiveNode) ExpectOnce(p core.RecordRef) *UnsyncListMockGetActiveNodeExpectation {
	m.mock.GetActiveNodeFunc = nil
	m.mainExpectation = nil

	expectation := &UnsyncListMockGetActiveNodeExpectation{}
	expectation.input = &UnsyncListMockGetActiveNodeInput{p}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *UnsyncListMockGetActiveNodeExpectation) Return(r core.Node) {
	e.result = &UnsyncListMockGetActiveNodeResult{r}
}

//Set uses given function f as a mock of UnsyncList.GetActiveNode method
func (m *mUnsyncListMockGetActiveNode) Set(f func(p core.RecordRef) (r core.Node)) *UnsyncListMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetActiveNodeFunc = f
	return m.mock
}

//GetActiveNode implements github.com/insolar/insolar/network.UnsyncList interface
func (m *UnsyncListMock) GetActiveNode(p core.RecordRef) (r core.Node) {
	counter := atomic.AddUint64(&m.GetActiveNodePreCounter, 1)
	defer atomic.AddUint64(&m.GetActiveNodeCounter, 1)

	if len(m.GetActiveNodeMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetActiveNodeMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to UnsyncListMock.GetActiveNode. %v", p)
			return
		}

		input := m.GetActiveNodeMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, UnsyncListMockGetActiveNodeInput{p}, "UnsyncList.GetActiveNode got unexpected parameters")

		result := m.GetActiveNodeMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.GetActiveNode")
			return
		}

		r = result.r

		return
	}

	if m.GetActiveNodeMock.mainExpectation != nil {

		input := m.GetActiveNodeMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, UnsyncListMockGetActiveNodeInput{p}, "UnsyncList.GetActiveNode got unexpected parameters")
		}

		result := m.GetActiveNodeMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.GetActiveNode")
		}

		r = result.r

		return
	}

	if m.GetActiveNodeFunc == nil {
		m.t.Fatalf("Unexpected call to UnsyncListMock.GetActiveNode. %v", p)
		return
	}

	return m.GetActiveNodeFunc(p)
}

//GetActiveNodeMinimockCounter returns a count of UnsyncListMock.GetActiveNodeFunc invocations
func (m *UnsyncListMock) GetActiveNodeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetActiveNodeCounter)
}

//GetActiveNodeMinimockPreCounter returns the value of UnsyncListMock.GetActiveNode invocations
func (m *UnsyncListMock) GetActiveNodeMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetActiveNodePreCounter)
}

//GetActiveNodeFinished returns true if mock invocations count is ok
func (m *UnsyncListMock) GetActiveNodeFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetActiveNodeMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetActiveNodeCounter) == uint64(len(m.GetActiveNodeMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetActiveNodeMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetActiveNodeCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetActiveNodeFunc != nil {
		return atomic.LoadUint64(&m.GetActiveNodeCounter) > 0
	}

	return true
}

type mUnsyncListMockGetActiveNodes struct {
	mock              *UnsyncListMock
	mainExpectation   *UnsyncListMockGetActiveNodesExpectation
	expectationSeries []*UnsyncListMockGetActiveNodesExpectation
}

type UnsyncListMockGetActiveNodesExpectation struct {
	result *UnsyncListMockGetActiveNodesResult
}

type UnsyncListMockGetActiveNodesResult struct {
	r []core.Node
}

//Expect specifies that invocation of UnsyncList.GetActiveNodes is expected from 1 to Infinity times
func (m *mUnsyncListMockGetActiveNodes) Expect() *mUnsyncListMockGetActiveNodes {
	m.mock.GetActiveNodesFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockGetActiveNodesExpectation{}
	}

	return m
}

//Return specifies results of invocation of UnsyncList.GetActiveNodes
func (m *mUnsyncListMockGetActiveNodes) Return(r []core.Node) *UnsyncListMock {
	m.mock.GetActiveNodesFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockGetActiveNodesExpectation{}
	}
	m.mainExpectation.result = &UnsyncListMockGetActiveNodesResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of UnsyncList.GetActiveNodes is expected once
func (m *mUnsyncListMockGetActiveNodes) ExpectOnce() *UnsyncListMockGetActiveNodesExpectation {
	m.mock.GetActiveNodesFunc = nil
	m.mainExpectation = nil

	expectation := &UnsyncListMockGetActiveNodesExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *UnsyncListMockGetActiveNodesExpectation) Return(r []core.Node) {
	e.result = &UnsyncListMockGetActiveNodesResult{r}
}

//Set uses given function f as a mock of UnsyncList.GetActiveNodes method
func (m *mUnsyncListMockGetActiveNodes) Set(f func() (r []core.Node)) *UnsyncListMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetActiveNodesFunc = f
	return m.mock
}

//GetActiveNodes implements github.com/insolar/insolar/network.UnsyncList interface
func (m *UnsyncListMock) GetActiveNodes() (r []core.Node) {
	counter := atomic.AddUint64(&m.GetActiveNodesPreCounter, 1)
	defer atomic.AddUint64(&m.GetActiveNodesCounter, 1)

	if len(m.GetActiveNodesMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetActiveNodesMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to UnsyncListMock.GetActiveNodes.")
			return
		}

		result := m.GetActiveNodesMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.GetActiveNodes")
			return
		}

		r = result.r

		return
	}

	if m.GetActiveNodesMock.mainExpectation != nil {

		result := m.GetActiveNodesMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.GetActiveNodes")
		}

		r = result.r

		return
	}

	if m.GetActiveNodesFunc == nil {
		m.t.Fatalf("Unexpected call to UnsyncListMock.GetActiveNodes.")
		return
	}

	return m.GetActiveNodesFunc()
}

//GetActiveNodesMinimockCounter returns a count of UnsyncListMock.GetActiveNodesFunc invocations
func (m *UnsyncListMock) GetActiveNodesMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetActiveNodesCounter)
}

//GetActiveNodesMinimockPreCounter returns the value of UnsyncListMock.GetActiveNodes invocations
func (m *UnsyncListMock) GetActiveNodesMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetActiveNodesPreCounter)
}

//GetActiveNodesFinished returns true if mock invocations count is ok
func (m *UnsyncListMock) GetActiveNodesFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetActiveNodesMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetActiveNodesCounter) == uint64(len(m.GetActiveNodesMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetActiveNodesMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetActiveNodesCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetActiveNodesFunc != nil {
		return atomic.LoadUint64(&m.GetActiveNodesCounter) > 0
	}

	return true
}

type mUnsyncListMockGetGlobuleHashSignature struct {
	mock              *UnsyncListMock
	mainExpectation   *UnsyncListMockGetGlobuleHashSignatureExpectation
	expectationSeries []*UnsyncListMockGetGlobuleHashSignatureExpectation
}

type UnsyncListMockGetGlobuleHashSignatureExpectation struct {
	input  *UnsyncListMockGetGlobuleHashSignatureInput
	result *UnsyncListMockGetGlobuleHashSignatureResult
}

type UnsyncListMockGetGlobuleHashSignatureInput struct {
	p core.RecordRef
}

type UnsyncListMockGetGlobuleHashSignatureResult struct {
	r  packets.GlobuleHashSignature
	r1 bool
}

//Expect specifies that invocation of UnsyncList.GetGlobuleHashSignature is expected from 1 to Infinity times
func (m *mUnsyncListMockGetGlobuleHashSignature) Expect(p core.RecordRef) *mUnsyncListMockGetGlobuleHashSignature {
	m.mock.GetGlobuleHashSignatureFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockGetGlobuleHashSignatureExpectation{}
	}
	m.mainExpectation.input = &UnsyncListMockGetGlobuleHashSignatureInput{p}
	return m
}

//Return specifies results of invocation of UnsyncList.GetGlobuleHashSignature
func (m *mUnsyncListMockGetGlobuleHashSignature) Return(r packets.GlobuleHashSignature, r1 bool) *UnsyncListMock {
	m.mock.GetGlobuleHashSignatureFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockGetGlobuleHashSignatureExpectation{}
	}
	m.mainExpectation.result = &UnsyncListMockGetGlobuleHashSignatureResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of UnsyncList.GetGlobuleHashSignature is expected once
func (m *mUnsyncListMockGetGlobuleHashSignature) ExpectOnce(p core.RecordRef) *UnsyncListMockGetGlobuleHashSignatureExpectation {
	m.mock.GetGlobuleHashSignatureFunc = nil
	m.mainExpectation = nil

	expectation := &UnsyncListMockGetGlobuleHashSignatureExpectation{}
	expectation.input = &UnsyncListMockGetGlobuleHashSignatureInput{p}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *UnsyncListMockGetGlobuleHashSignatureExpectation) Return(r packets.GlobuleHashSignature, r1 bool) {
	e.result = &UnsyncListMockGetGlobuleHashSignatureResult{r, r1}
}

//Set uses given function f as a mock of UnsyncList.GetGlobuleHashSignature method
func (m *mUnsyncListMockGetGlobuleHashSignature) Set(f func(p core.RecordRef) (r packets.GlobuleHashSignature, r1 bool)) *UnsyncListMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetGlobuleHashSignatureFunc = f
	return m.mock
}

//GetGlobuleHashSignature implements github.com/insolar/insolar/network.UnsyncList interface
func (m *UnsyncListMock) GetGlobuleHashSignature(p core.RecordRef) (r packets.GlobuleHashSignature, r1 bool) {
	counter := atomic.AddUint64(&m.GetGlobuleHashSignaturePreCounter, 1)
	defer atomic.AddUint64(&m.GetGlobuleHashSignatureCounter, 1)

	if len(m.GetGlobuleHashSignatureMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetGlobuleHashSignatureMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to UnsyncListMock.GetGlobuleHashSignature. %v", p)
			return
		}

		input := m.GetGlobuleHashSignatureMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, UnsyncListMockGetGlobuleHashSignatureInput{p}, "UnsyncList.GetGlobuleHashSignature got unexpected parameters")

		result := m.GetGlobuleHashSignatureMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.GetGlobuleHashSignature")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.GetGlobuleHashSignatureMock.mainExpectation != nil {

		input := m.GetGlobuleHashSignatureMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, UnsyncListMockGetGlobuleHashSignatureInput{p}, "UnsyncList.GetGlobuleHashSignature got unexpected parameters")
		}

		result := m.GetGlobuleHashSignatureMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.GetGlobuleHashSignature")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.GetGlobuleHashSignatureFunc == nil {
		m.t.Fatalf("Unexpected call to UnsyncListMock.GetGlobuleHashSignature. %v", p)
		return
	}

	return m.GetGlobuleHashSignatureFunc(p)
}

//GetGlobuleHashSignatureMinimockCounter returns a count of UnsyncListMock.GetGlobuleHashSignatureFunc invocations
func (m *UnsyncListMock) GetGlobuleHashSignatureMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetGlobuleHashSignatureCounter)
}

//GetGlobuleHashSignatureMinimockPreCounter returns the value of UnsyncListMock.GetGlobuleHashSignature invocations
func (m *UnsyncListMock) GetGlobuleHashSignatureMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetGlobuleHashSignaturePreCounter)
}

//GetGlobuleHashSignatureFinished returns true if mock invocations count is ok
func (m *UnsyncListMock) GetGlobuleHashSignatureFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetGlobuleHashSignatureMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetGlobuleHashSignatureCounter) == uint64(len(m.GetGlobuleHashSignatureMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetGlobuleHashSignatureMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetGlobuleHashSignatureCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetGlobuleHashSignatureFunc != nil {
		return atomic.LoadUint64(&m.GetGlobuleHashSignatureCounter) > 0
	}

	return true
}

type mUnsyncListMockGetOrigin struct {
	mock              *UnsyncListMock
	mainExpectation   *UnsyncListMockGetOriginExpectation
	expectationSeries []*UnsyncListMockGetOriginExpectation
}

type UnsyncListMockGetOriginExpectation struct {
	result *UnsyncListMockGetOriginResult
}

type UnsyncListMockGetOriginResult struct {
	r core.Node
}

//Expect specifies that invocation of UnsyncList.GetOrigin is expected from 1 to Infinity times
func (m *mUnsyncListMockGetOrigin) Expect() *mUnsyncListMockGetOrigin {
	m.mock.GetOriginFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockGetOriginExpectation{}
	}

	return m
}

//Return specifies results of invocation of UnsyncList.GetOrigin
func (m *mUnsyncListMockGetOrigin) Return(r core.Node) *UnsyncListMock {
	m.mock.GetOriginFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockGetOriginExpectation{}
	}
	m.mainExpectation.result = &UnsyncListMockGetOriginResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of UnsyncList.GetOrigin is expected once
func (m *mUnsyncListMockGetOrigin) ExpectOnce() *UnsyncListMockGetOriginExpectation {
	m.mock.GetOriginFunc = nil
	m.mainExpectation = nil

	expectation := &UnsyncListMockGetOriginExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *UnsyncListMockGetOriginExpectation) Return(r core.Node) {
	e.result = &UnsyncListMockGetOriginResult{r}
}

//Set uses given function f as a mock of UnsyncList.GetOrigin method
func (m *mUnsyncListMockGetOrigin) Set(f func() (r core.Node)) *UnsyncListMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetOriginFunc = f
	return m.mock
}

//GetOrigin implements github.com/insolar/insolar/network.UnsyncList interface
func (m *UnsyncListMock) GetOrigin() (r core.Node) {
	counter := atomic.AddUint64(&m.GetOriginPreCounter, 1)
	defer atomic.AddUint64(&m.GetOriginCounter, 1)

	if len(m.GetOriginMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetOriginMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to UnsyncListMock.GetOrigin.")
			return
		}

		result := m.GetOriginMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.GetOrigin")
			return
		}

		r = result.r

		return
	}

	if m.GetOriginMock.mainExpectation != nil {

		result := m.GetOriginMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.GetOrigin")
		}

		r = result.r

		return
	}

	if m.GetOriginFunc == nil {
		m.t.Fatalf("Unexpected call to UnsyncListMock.GetOrigin.")
		return
	}

	return m.GetOriginFunc()
}

//GetOriginMinimockCounter returns a count of UnsyncListMock.GetOriginFunc invocations
func (m *UnsyncListMock) GetOriginMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetOriginCounter)
}

//GetOriginMinimockPreCounter returns the value of UnsyncListMock.GetOrigin invocations
func (m *UnsyncListMock) GetOriginMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetOriginPreCounter)
}

//GetOriginFinished returns true if mock invocations count is ok
func (m *UnsyncListMock) GetOriginFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetOriginMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetOriginCounter) == uint64(len(m.GetOriginMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetOriginMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetOriginCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetOriginFunc != nil {
		return atomic.LoadUint64(&m.GetOriginCounter) > 0
	}

	return true
}

type mUnsyncListMockGetProof struct {
	mock              *UnsyncListMock
	mainExpectation   *UnsyncListMockGetProofExpectation
	expectationSeries []*UnsyncListMockGetProofExpectation
}

type UnsyncListMockGetProofExpectation struct {
	input  *UnsyncListMockGetProofInput
	result *UnsyncListMockGetProofResult
}

type UnsyncListMockGetProofInput struct {
	p core.RecordRef
}

type UnsyncListMockGetProofResult struct {
	r *packets.NodePulseProof
}

//Expect specifies that invocation of UnsyncList.GetProof is expected from 1 to Infinity times
func (m *mUnsyncListMockGetProof) Expect(p core.RecordRef) *mUnsyncListMockGetProof {
	m.mock.GetProofFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockGetProofExpectation{}
	}
	m.mainExpectation.input = &UnsyncListMockGetProofInput{p}
	return m
}

//Return specifies results of invocation of UnsyncList.GetProof
func (m *mUnsyncListMockGetProof) Return(r *packets.NodePulseProof) *UnsyncListMock {
	m.mock.GetProofFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockGetProofExpectation{}
	}
	m.mainExpectation.result = &UnsyncListMockGetProofResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of UnsyncList.GetProof is expected once
func (m *mUnsyncListMockGetProof) ExpectOnce(p core.RecordRef) *UnsyncListMockGetProofExpectation {
	m.mock.GetProofFunc = nil
	m.mainExpectation = nil

	expectation := &UnsyncListMockGetProofExpectation{}
	expectation.input = &UnsyncListMockGetProofInput{p}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *UnsyncListMockGetProofExpectation) Return(r *packets.NodePulseProof) {
	e.result = &UnsyncListMockGetProofResult{r}
}

//Set uses given function f as a mock of UnsyncList.GetProof method
func (m *mUnsyncListMockGetProof) Set(f func(p core.RecordRef) (r *packets.NodePulseProof)) *UnsyncListMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetProofFunc = f
	return m.mock
}

//GetProof implements github.com/insolar/insolar/network.UnsyncList interface
func (m *UnsyncListMock) GetProof(p core.RecordRef) (r *packets.NodePulseProof) {
	counter := atomic.AddUint64(&m.GetProofPreCounter, 1)
	defer atomic.AddUint64(&m.GetProofCounter, 1)

	if len(m.GetProofMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetProofMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to UnsyncListMock.GetProof. %v", p)
			return
		}

		input := m.GetProofMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, UnsyncListMockGetProofInput{p}, "UnsyncList.GetProof got unexpected parameters")

		result := m.GetProofMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.GetProof")
			return
		}

		r = result.r

		return
	}

	if m.GetProofMock.mainExpectation != nil {

		input := m.GetProofMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, UnsyncListMockGetProofInput{p}, "UnsyncList.GetProof got unexpected parameters")
		}

		result := m.GetProofMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.GetProof")
		}

		r = result.r

		return
	}

	if m.GetProofFunc == nil {
		m.t.Fatalf("Unexpected call to UnsyncListMock.GetProof. %v", p)
		return
	}

	return m.GetProofFunc(p)
}

//GetProofMinimockCounter returns a count of UnsyncListMock.GetProofFunc invocations
func (m *UnsyncListMock) GetProofMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetProofCounter)
}

//GetProofMinimockPreCounter returns the value of UnsyncListMock.GetProof invocations
func (m *UnsyncListMock) GetProofMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetProofPreCounter)
}

//GetProofFinished returns true if mock invocations count is ok
func (m *UnsyncListMock) GetProofFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetProofMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetProofCounter) == uint64(len(m.GetProofMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetProofMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetProofCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetProofFunc != nil {
		return atomic.LoadUint64(&m.GetProofCounter) > 0
	}

	return true
}

type mUnsyncListMockIndexToRef struct {
	mock              *UnsyncListMock
	mainExpectation   *UnsyncListMockIndexToRefExpectation
	expectationSeries []*UnsyncListMockIndexToRefExpectation
}

type UnsyncListMockIndexToRefExpectation struct {
	input  *UnsyncListMockIndexToRefInput
	result *UnsyncListMockIndexToRefResult
}

type UnsyncListMockIndexToRefInput struct {
	p int
}

type UnsyncListMockIndexToRefResult struct {
	r  core.RecordRef
	r1 error
}

//Expect specifies that invocation of UnsyncList.IndexToRef is expected from 1 to Infinity times
func (m *mUnsyncListMockIndexToRef) Expect(p int) *mUnsyncListMockIndexToRef {
	m.mock.IndexToRefFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockIndexToRefExpectation{}
	}
	m.mainExpectation.input = &UnsyncListMockIndexToRefInput{p}
	return m
}

//Return specifies results of invocation of UnsyncList.IndexToRef
func (m *mUnsyncListMockIndexToRef) Return(r core.RecordRef, r1 error) *UnsyncListMock {
	m.mock.IndexToRefFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockIndexToRefExpectation{}
	}
	m.mainExpectation.result = &UnsyncListMockIndexToRefResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of UnsyncList.IndexToRef is expected once
func (m *mUnsyncListMockIndexToRef) ExpectOnce(p int) *UnsyncListMockIndexToRefExpectation {
	m.mock.IndexToRefFunc = nil
	m.mainExpectation = nil

	expectation := &UnsyncListMockIndexToRefExpectation{}
	expectation.input = &UnsyncListMockIndexToRefInput{p}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *UnsyncListMockIndexToRefExpectation) Return(r core.RecordRef, r1 error) {
	e.result = &UnsyncListMockIndexToRefResult{r, r1}
}

//Set uses given function f as a mock of UnsyncList.IndexToRef method
func (m *mUnsyncListMockIndexToRef) Set(f func(p int) (r core.RecordRef, r1 error)) *UnsyncListMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.IndexToRefFunc = f
	return m.mock
}

//IndexToRef implements github.com/insolar/insolar/network.UnsyncList interface
func (m *UnsyncListMock) IndexToRef(p int) (r core.RecordRef, r1 error) {
	counter := atomic.AddUint64(&m.IndexToRefPreCounter, 1)
	defer atomic.AddUint64(&m.IndexToRefCounter, 1)

	if len(m.IndexToRefMock.expectationSeries) > 0 {
		if counter > uint64(len(m.IndexToRefMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to UnsyncListMock.IndexToRef. %v", p)
			return
		}

		input := m.IndexToRefMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, UnsyncListMockIndexToRefInput{p}, "UnsyncList.IndexToRef got unexpected parameters")

		result := m.IndexToRefMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.IndexToRef")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.IndexToRefMock.mainExpectation != nil {

		input := m.IndexToRefMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, UnsyncListMockIndexToRefInput{p}, "UnsyncList.IndexToRef got unexpected parameters")
		}

		result := m.IndexToRefMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.IndexToRef")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.IndexToRefFunc == nil {
		m.t.Fatalf("Unexpected call to UnsyncListMock.IndexToRef. %v", p)
		return
	}

	return m.IndexToRefFunc(p)
}

//IndexToRefMinimockCounter returns a count of UnsyncListMock.IndexToRefFunc invocations
func (m *UnsyncListMock) IndexToRefMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.IndexToRefCounter)
}

//IndexToRefMinimockPreCounter returns the value of UnsyncListMock.IndexToRef invocations
func (m *UnsyncListMock) IndexToRefMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.IndexToRefPreCounter)
}

//IndexToRefFinished returns true if mock invocations count is ok
func (m *UnsyncListMock) IndexToRefFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.IndexToRefMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.IndexToRefCounter) == uint64(len(m.IndexToRefMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.IndexToRefMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.IndexToRefCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.IndexToRefFunc != nil {
		return atomic.LoadUint64(&m.IndexToRefCounter) > 0
	}

	return true
}

type mUnsyncListMockLength struct {
	mock              *UnsyncListMock
	mainExpectation   *UnsyncListMockLengthExpectation
	expectationSeries []*UnsyncListMockLengthExpectation
}

type UnsyncListMockLengthExpectation struct {
	result *UnsyncListMockLengthResult
}

type UnsyncListMockLengthResult struct {
	r int
}

//Expect specifies that invocation of UnsyncList.Length is expected from 1 to Infinity times
func (m *mUnsyncListMockLength) Expect() *mUnsyncListMockLength {
	m.mock.LengthFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockLengthExpectation{}
	}

	return m
}

//Return specifies results of invocation of UnsyncList.Length
func (m *mUnsyncListMockLength) Return(r int) *UnsyncListMock {
	m.mock.LengthFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockLengthExpectation{}
	}
	m.mainExpectation.result = &UnsyncListMockLengthResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of UnsyncList.Length is expected once
func (m *mUnsyncListMockLength) ExpectOnce() *UnsyncListMockLengthExpectation {
	m.mock.LengthFunc = nil
	m.mainExpectation = nil

	expectation := &UnsyncListMockLengthExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *UnsyncListMockLengthExpectation) Return(r int) {
	e.result = &UnsyncListMockLengthResult{r}
}

//Set uses given function f as a mock of UnsyncList.Length method
func (m *mUnsyncListMockLength) Set(f func() (r int)) *UnsyncListMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.LengthFunc = f
	return m.mock
}

//Length implements github.com/insolar/insolar/network.UnsyncList interface
func (m *UnsyncListMock) Length() (r int) {
	counter := atomic.AddUint64(&m.LengthPreCounter, 1)
	defer atomic.AddUint64(&m.LengthCounter, 1)

	if len(m.LengthMock.expectationSeries) > 0 {
		if counter > uint64(len(m.LengthMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to UnsyncListMock.Length.")
			return
		}

		result := m.LengthMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.Length")
			return
		}

		r = result.r

		return
	}

	if m.LengthMock.mainExpectation != nil {

		result := m.LengthMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.Length")
		}

		r = result.r

		return
	}

	if m.LengthFunc == nil {
		m.t.Fatalf("Unexpected call to UnsyncListMock.Length.")
		return
	}

	return m.LengthFunc()
}

//LengthMinimockCounter returns a count of UnsyncListMock.LengthFunc invocations
func (m *UnsyncListMock) LengthMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.LengthCounter)
}

//LengthMinimockPreCounter returns the value of UnsyncListMock.Length invocations
func (m *UnsyncListMock) LengthMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.LengthPreCounter)
}

//LengthFinished returns true if mock invocations count is ok
func (m *UnsyncListMock) LengthFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.LengthMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.LengthCounter) == uint64(len(m.LengthMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.LengthMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.LengthCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.LengthFunc != nil {
		return atomic.LoadUint64(&m.LengthCounter) > 0
	}

	return true
}

type mUnsyncListMockRefToIndex struct {
	mock              *UnsyncListMock
	mainExpectation   *UnsyncListMockRefToIndexExpectation
	expectationSeries []*UnsyncListMockRefToIndexExpectation
}

type UnsyncListMockRefToIndexExpectation struct {
	input  *UnsyncListMockRefToIndexInput
	result *UnsyncListMockRefToIndexResult
}

type UnsyncListMockRefToIndexInput struct {
	p core.RecordRef
}

type UnsyncListMockRefToIndexResult struct {
	r  int
	r1 error
}

//Expect specifies that invocation of UnsyncList.RefToIndex is expected from 1 to Infinity times
func (m *mUnsyncListMockRefToIndex) Expect(p core.RecordRef) *mUnsyncListMockRefToIndex {
	m.mock.RefToIndexFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockRefToIndexExpectation{}
	}
	m.mainExpectation.input = &UnsyncListMockRefToIndexInput{p}
	return m
}

//Return specifies results of invocation of UnsyncList.RefToIndex
func (m *mUnsyncListMockRefToIndex) Return(r int, r1 error) *UnsyncListMock {
	m.mock.RefToIndexFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockRefToIndexExpectation{}
	}
	m.mainExpectation.result = &UnsyncListMockRefToIndexResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of UnsyncList.RefToIndex is expected once
func (m *mUnsyncListMockRefToIndex) ExpectOnce(p core.RecordRef) *UnsyncListMockRefToIndexExpectation {
	m.mock.RefToIndexFunc = nil
	m.mainExpectation = nil

	expectation := &UnsyncListMockRefToIndexExpectation{}
	expectation.input = &UnsyncListMockRefToIndexInput{p}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *UnsyncListMockRefToIndexExpectation) Return(r int, r1 error) {
	e.result = &UnsyncListMockRefToIndexResult{r, r1}
}

//Set uses given function f as a mock of UnsyncList.RefToIndex method
func (m *mUnsyncListMockRefToIndex) Set(f func(p core.RecordRef) (r int, r1 error)) *UnsyncListMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.RefToIndexFunc = f
	return m.mock
}

//RefToIndex implements github.com/insolar/insolar/network.UnsyncList interface
func (m *UnsyncListMock) RefToIndex(p core.RecordRef) (r int, r1 error) {
	counter := atomic.AddUint64(&m.RefToIndexPreCounter, 1)
	defer atomic.AddUint64(&m.RefToIndexCounter, 1)

	if len(m.RefToIndexMock.expectationSeries) > 0 {
		if counter > uint64(len(m.RefToIndexMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to UnsyncListMock.RefToIndex. %v", p)
			return
		}

		input := m.RefToIndexMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, UnsyncListMockRefToIndexInput{p}, "UnsyncList.RefToIndex got unexpected parameters")

		result := m.RefToIndexMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.RefToIndex")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.RefToIndexMock.mainExpectation != nil {

		input := m.RefToIndexMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, UnsyncListMockRefToIndexInput{p}, "UnsyncList.RefToIndex got unexpected parameters")
		}

		result := m.RefToIndexMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the UnsyncListMock.RefToIndex")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.RefToIndexFunc == nil {
		m.t.Fatalf("Unexpected call to UnsyncListMock.RefToIndex. %v", p)
		return
	}

	return m.RefToIndexFunc(p)
}

//RefToIndexMinimockCounter returns a count of UnsyncListMock.RefToIndexFunc invocations
func (m *UnsyncListMock) RefToIndexMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RefToIndexCounter)
}

//RefToIndexMinimockPreCounter returns the value of UnsyncListMock.RefToIndex invocations
func (m *UnsyncListMock) RefToIndexMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.RefToIndexPreCounter)
}

//RefToIndexFinished returns true if mock invocations count is ok
func (m *UnsyncListMock) RefToIndexFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.RefToIndexMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.RefToIndexCounter) == uint64(len(m.RefToIndexMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.RefToIndexMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.RefToIndexCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.RefToIndexFunc != nil {
		return atomic.LoadUint64(&m.RefToIndexCounter) > 0
	}

	return true
}

type mUnsyncListMockRemoveNode struct {
	mock              *UnsyncListMock
	mainExpectation   *UnsyncListMockRemoveNodeExpectation
	expectationSeries []*UnsyncListMockRemoveNodeExpectation
}

type UnsyncListMockRemoveNodeExpectation struct {
	input *UnsyncListMockRemoveNodeInput
}

type UnsyncListMockRemoveNodeInput struct {
	p core.RecordRef
}

//Expect specifies that invocation of UnsyncList.RemoveNode is expected from 1 to Infinity times
func (m *mUnsyncListMockRemoveNode) Expect(p core.RecordRef) *mUnsyncListMockRemoveNode {
	m.mock.RemoveNodeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockRemoveNodeExpectation{}
	}
	m.mainExpectation.input = &UnsyncListMockRemoveNodeInput{p}
	return m
}

//Return specifies results of invocation of UnsyncList.RemoveNode
func (m *mUnsyncListMockRemoveNode) Return() *UnsyncListMock {
	m.mock.RemoveNodeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockRemoveNodeExpectation{}
	}

	return m.mock
}

//ExpectOnce specifies that invocation of UnsyncList.RemoveNode is expected once
func (m *mUnsyncListMockRemoveNode) ExpectOnce(p core.RecordRef) *UnsyncListMockRemoveNodeExpectation {
	m.mock.RemoveNodeFunc = nil
	m.mainExpectation = nil

	expectation := &UnsyncListMockRemoveNodeExpectation{}
	expectation.input = &UnsyncListMockRemoveNodeInput{p}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

//Set uses given function f as a mock of UnsyncList.RemoveNode method
func (m *mUnsyncListMockRemoveNode) Set(f func(p core.RecordRef)) *UnsyncListMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.RemoveNodeFunc = f
	return m.mock
}

//RemoveNode implements github.com/insolar/insolar/network.UnsyncList interface
func (m *UnsyncListMock) RemoveNode(p core.RecordRef) {
	counter := atomic.AddUint64(&m.RemoveNodePreCounter, 1)
	defer atomic.AddUint64(&m.RemoveNodeCounter, 1)

	if len(m.RemoveNodeMock.expectationSeries) > 0 {
		if counter > uint64(len(m.RemoveNodeMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to UnsyncListMock.RemoveNode. %v", p)
			return
		}

		input := m.RemoveNodeMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, UnsyncListMockRemoveNodeInput{p}, "UnsyncList.RemoveNode got unexpected parameters")

		return
	}

	if m.RemoveNodeMock.mainExpectation != nil {

		input := m.RemoveNodeMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, UnsyncListMockRemoveNodeInput{p}, "UnsyncList.RemoveNode got unexpected parameters")
		}

		return
	}

	if m.RemoveNodeFunc == nil {
		m.t.Fatalf("Unexpected call to UnsyncListMock.RemoveNode. %v", p)
		return
	}

	m.RemoveNodeFunc(p)
}

//RemoveNodeMinimockCounter returns a count of UnsyncListMock.RemoveNodeFunc invocations
func (m *UnsyncListMock) RemoveNodeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RemoveNodeCounter)
}

//RemoveNodeMinimockPreCounter returns the value of UnsyncListMock.RemoveNode invocations
func (m *UnsyncListMock) RemoveNodeMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.RemoveNodePreCounter)
}

//RemoveNodeFinished returns true if mock invocations count is ok
func (m *UnsyncListMock) RemoveNodeFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.RemoveNodeMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.RemoveNodeCounter) == uint64(len(m.RemoveNodeMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.RemoveNodeMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.RemoveNodeCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.RemoveNodeFunc != nil {
		return atomic.LoadUint64(&m.RemoveNodeCounter) > 0
	}

	return true
}

type mUnsyncListMockSetGlobuleHashSignature struct {
	mock              *UnsyncListMock
	mainExpectation   *UnsyncListMockSetGlobuleHashSignatureExpectation
	expectationSeries []*UnsyncListMockSetGlobuleHashSignatureExpectation
}

type UnsyncListMockSetGlobuleHashSignatureExpectation struct {
	input *UnsyncListMockSetGlobuleHashSignatureInput
}

type UnsyncListMockSetGlobuleHashSignatureInput struct {
	p  core.RecordRef
	p1 packets.GlobuleHashSignature
}

//Expect specifies that invocation of UnsyncList.SetGlobuleHashSignature is expected from 1 to Infinity times
func (m *mUnsyncListMockSetGlobuleHashSignature) Expect(p core.RecordRef, p1 packets.GlobuleHashSignature) *mUnsyncListMockSetGlobuleHashSignature {
	m.mock.SetGlobuleHashSignatureFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockSetGlobuleHashSignatureExpectation{}
	}
	m.mainExpectation.input = &UnsyncListMockSetGlobuleHashSignatureInput{p, p1}
	return m
}

//Return specifies results of invocation of UnsyncList.SetGlobuleHashSignature
func (m *mUnsyncListMockSetGlobuleHashSignature) Return() *UnsyncListMock {
	m.mock.SetGlobuleHashSignatureFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &UnsyncListMockSetGlobuleHashSignatureExpectation{}
	}

	return m.mock
}

//ExpectOnce specifies that invocation of UnsyncList.SetGlobuleHashSignature is expected once
func (m *mUnsyncListMockSetGlobuleHashSignature) ExpectOnce(p core.RecordRef, p1 packets.GlobuleHashSignature) *UnsyncListMockSetGlobuleHashSignatureExpectation {
	m.mock.SetGlobuleHashSignatureFunc = nil
	m.mainExpectation = nil

	expectation := &UnsyncListMockSetGlobuleHashSignatureExpectation{}
	expectation.input = &UnsyncListMockSetGlobuleHashSignatureInput{p, p1}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

//Set uses given function f as a mock of UnsyncList.SetGlobuleHashSignature method
func (m *mUnsyncListMockSetGlobuleHashSignature) Set(f func(p core.RecordRef, p1 packets.GlobuleHashSignature)) *UnsyncListMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.SetGlobuleHashSignatureFunc = f
	return m.mock
}

//SetGlobuleHashSignature implements github.com/insolar/insolar/network.UnsyncList interface
func (m *UnsyncListMock) SetGlobuleHashSignature(p core.RecordRef, p1 packets.GlobuleHashSignature) {
	counter := atomic.AddUint64(&m.SetGlobuleHashSignaturePreCounter, 1)
	defer atomic.AddUint64(&m.SetGlobuleHashSignatureCounter, 1)

	if len(m.SetGlobuleHashSignatureMock.expectationSeries) > 0 {
		if counter > uint64(len(m.SetGlobuleHashSignatureMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to UnsyncListMock.SetGlobuleHashSignature. %v %v", p, p1)
			return
		}

		input := m.SetGlobuleHashSignatureMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, UnsyncListMockSetGlobuleHashSignatureInput{p, p1}, "UnsyncList.SetGlobuleHashSignature got unexpected parameters")

		return
	}

	if m.SetGlobuleHashSignatureMock.mainExpectation != nil {

		input := m.SetGlobuleHashSignatureMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, UnsyncListMockSetGlobuleHashSignatureInput{p, p1}, "UnsyncList.SetGlobuleHashSignature got unexpected parameters")
		}

		return
	}

	if m.SetGlobuleHashSignatureFunc == nil {
		m.t.Fatalf("Unexpected call to UnsyncListMock.SetGlobuleHashSignature. %v %v", p, p1)
		return
	}

	m.SetGlobuleHashSignatureFunc(p, p1)
}

//SetGlobuleHashSignatureMinimockCounter returns a count of UnsyncListMock.SetGlobuleHashSignatureFunc invocations
func (m *UnsyncListMock) SetGlobuleHashSignatureMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SetGlobuleHashSignatureCounter)
}

//SetGlobuleHashSignatureMinimockPreCounter returns the value of UnsyncListMock.SetGlobuleHashSignature invocations
func (m *UnsyncListMock) SetGlobuleHashSignatureMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.SetGlobuleHashSignaturePreCounter)
}

//SetGlobuleHashSignatureFinished returns true if mock invocations count is ok
func (m *UnsyncListMock) SetGlobuleHashSignatureFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.SetGlobuleHashSignatureMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.SetGlobuleHashSignatureCounter) == uint64(len(m.SetGlobuleHashSignatureMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.SetGlobuleHashSignatureMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.SetGlobuleHashSignatureCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.SetGlobuleHashSignatureFunc != nil {
		return atomic.LoadUint64(&m.SetGlobuleHashSignatureCounter) > 0
	}

	return true
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *UnsyncListMock) ValidateCallCounters() {

	if !m.AddNodeFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.AddNode")
	}

	if !m.AddProofFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.AddProof")
	}

	if !m.GetActiveNodeFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.GetActiveNode")
	}

	if !m.GetActiveNodesFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.GetActiveNodes")
	}

	if !m.GetGlobuleHashSignatureFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.GetGlobuleHashSignature")
	}

	if !m.GetOriginFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.GetOrigin")
	}

	if !m.GetProofFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.GetProof")
	}

	if !m.IndexToRefFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.IndexToRef")
	}

	if !m.LengthFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.Length")
	}

	if !m.RefToIndexFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.RefToIndex")
	}

	if !m.RemoveNodeFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.RemoveNode")
	}

	if !m.SetGlobuleHashSignatureFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.SetGlobuleHashSignature")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *UnsyncListMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *UnsyncListMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *UnsyncListMock) MinimockFinish() {

	if !m.AddNodeFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.AddNode")
	}

	if !m.AddProofFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.AddProof")
	}

	if !m.GetActiveNodeFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.GetActiveNode")
	}

	if !m.GetActiveNodesFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.GetActiveNodes")
	}

	if !m.GetGlobuleHashSignatureFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.GetGlobuleHashSignature")
	}

	if !m.GetOriginFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.GetOrigin")
	}

	if !m.GetProofFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.GetProof")
	}

	if !m.IndexToRefFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.IndexToRef")
	}

	if !m.LengthFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.Length")
	}

	if !m.RefToIndexFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.RefToIndex")
	}

	if !m.RemoveNodeFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.RemoveNode")
	}

	if !m.SetGlobuleHashSignatureFinished() {
		m.t.Fatal("Expected call to UnsyncListMock.SetGlobuleHashSignature")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *UnsyncListMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *UnsyncListMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && m.AddNodeFinished()
		ok = ok && m.AddProofFinished()
		ok = ok && m.GetActiveNodeFinished()
		ok = ok && m.GetActiveNodesFinished()
		ok = ok && m.GetGlobuleHashSignatureFinished()
		ok = ok && m.GetOriginFinished()
		ok = ok && m.GetProofFinished()
		ok = ok && m.IndexToRefFinished()
		ok = ok && m.LengthFinished()
		ok = ok && m.RefToIndexFinished()
		ok = ok && m.RemoveNodeFinished()
		ok = ok && m.SetGlobuleHashSignatureFinished()

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if !m.AddNodeFinished() {
				m.t.Error("Expected call to UnsyncListMock.AddNode")
			}

			if !m.AddProofFinished() {
				m.t.Error("Expected call to UnsyncListMock.AddProof")
			}

			if !m.GetActiveNodeFinished() {
				m.t.Error("Expected call to UnsyncListMock.GetActiveNode")
			}

			if !m.GetActiveNodesFinished() {
				m.t.Error("Expected call to UnsyncListMock.GetActiveNodes")
			}

			if !m.GetGlobuleHashSignatureFinished() {
				m.t.Error("Expected call to UnsyncListMock.GetGlobuleHashSignature")
			}

			if !m.GetOriginFinished() {
				m.t.Error("Expected call to UnsyncListMock.GetOrigin")
			}

			if !m.GetProofFinished() {
				m.t.Error("Expected call to UnsyncListMock.GetProof")
			}

			if !m.IndexToRefFinished() {
				m.t.Error("Expected call to UnsyncListMock.IndexToRef")
			}

			if !m.LengthFinished() {
				m.t.Error("Expected call to UnsyncListMock.Length")
			}

			if !m.RefToIndexFinished() {
				m.t.Error("Expected call to UnsyncListMock.RefToIndex")
			}

			if !m.RemoveNodeFinished() {
				m.t.Error("Expected call to UnsyncListMock.RemoveNode")
			}

			if !m.SetGlobuleHashSignatureFinished() {
				m.t.Error("Expected call to UnsyncListMock.SetGlobuleHashSignature")
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
func (m *UnsyncListMock) AllMocksCalled() bool {

	if !m.AddNodeFinished() {
		return false
	}

	if !m.AddProofFinished() {
		return false
	}

	if !m.GetActiveNodeFinished() {
		return false
	}

	if !m.GetActiveNodesFinished() {
		return false
	}

	if !m.GetGlobuleHashSignatureFinished() {
		return false
	}

	if !m.GetOriginFinished() {
		return false
	}

	if !m.GetProofFinished() {
		return false
	}

	if !m.IndexToRefFinished() {
		return false
	}

	if !m.LengthFinished() {
		return false
	}

	if !m.RefToIndexFinished() {
		return false
	}

	if !m.RemoveNodeFinished() {
		return false
	}

	if !m.SetGlobuleHashSignatureFinished() {
		return false
	}

	return true
}
