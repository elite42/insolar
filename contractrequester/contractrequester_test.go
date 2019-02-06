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

package contractrequester

import (
	"context"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/gojuno/minimock"
	"github.com/stretchr/testify/assert"

	"github.com/insolar/insolar/component"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/core/message"
	"github.com/insolar/insolar/core/reply"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
)

func mockMessageBus(t *testing.T, result core.Reply) *testutils.MessageBusMock {
	mbMock := testutils.NewMessageBusMock(t)
	mbMock.SendFunc = func(c context.Context, m core.Message, o *core.MessageSendOptions) (r core.Reply, r1 error) {
		return result, nil
	}
	return mbMock
}

func TestNew(t *testing.T) {
	ps := testutils.NewPulseStorageMock(t)
	messageBus := mockMessageBus(t, nil)

	contractRequester, err := New()

	cm := &component.Manager{}
	cm.Inject(ps, messageBus, contractRequester)

	require.NoError(t, err)
	require.Equal(t, messageBus, contractRequester.MessageBus)
	require.Equal(t, ps, contractRequester.PulseStorage)
}

func TestContractRequester_SendRequest(t *testing.T) {
	ctx := inslogger.TestContext(t)
	ref := testutils.RandomRef()

	pm := testutils.NewPulseStorageMock(t)
	pm.CurrentMock.Return(core.GenesisPulse, nil)

	mbm := mockMessageBus(t, &reply.RegisterRequest{})
	cReq, err := New()
	assert.NoError(t, err)
	cReq.MessageBus = mbm
	cReq.PulseStorage = pm

	mbm.MustRegisterMock.Return()
	cReq.Start(ctx)

	go func() {
		resLen := 0
		for resLen == 0 {
			cReq.ResultMutex.Lock()
			resLen = len(cReq.ResultMap)
			cReq.ResultMutex.Unlock()
			runtime.Gosched()
		}

		cReq.ResultMutex.Lock()
		for k, v := range cReq.ResultMap {
			v <- &message.ReturnResults{
				Sequence: k,
				Reply:    &reply.CallMethod{},
			}
		}
		cReq.ResultMutex.Unlock()
	}()
	result, err := cReq.SendRequest(ctx, &ref, "TestMethod", []interface{}{})

	require.NoError(t, err)
	require.Equal(t, &reply.CallMethod{}, result)
}

func TestContractRequester_SendRequest_RouteError(t *testing.T) {
	ctx := inslogger.TestContext(t)
	ref := testutils.RandomRef()

	pm := testutils.NewPulseStorageMock(t)
	pm.CurrentMock.Return(core.GenesisPulse, nil)

	mbm := mockMessageBus(t, &reply.CallMethod{})
	cReq, err := New()
	assert.NoError(t, err)
	cReq.MessageBus = mbm
	cReq.PulseStorage = pm

	mbm.MustRegisterMock.Return()
	err = cReq.Start(ctx)
	require.NoError(t, err)

	ifResultMapEmpty := func() bool {
		cReq.ResultMutex.Lock()
		defer cReq.ResultMutex.Unlock()
		return len(cReq.ResultMap) == 0
	}

	go func() {
		for ifResultMapEmpty() {
			runtime.Gosched()
		}
		cReq.ResultMutex.Lock()
		for k, v := range cReq.ResultMap {
			v <- &message.ReturnResults{
				Sequence: k,
				Reply:    nil,
			}
		}
		cReq.ResultMutex.Unlock()
	}()

	result, err := cReq.SendRequest(ctx, &ref, "TestMethod", []interface{}{})
	require.Nil(t, result)
}

func TestCallMethodCanceled(t *testing.T) {
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second)
	defer cancelFunc()

	cr, err := New()
	require.NoError(t, err)

	mc := minimock.NewController(t)
	defer mc.Finish()

	mb := testutils.NewMessageBusMock(mc)
	cr.MessageBus = mb

	msg := &message.BaseLogicMessage{
		Nonce: randomUint64(),
	}
	ref := testutils.RandomRef()
	prototypeRef := testutils.RandomRef()
	method := testutils.RandomString()

	mb.SendFunc = func(p context.Context, p1 core.Message, p2 *core.MessageSendOptions) (r core.Reply, r1 error) {
		return &reply.RegisterRequest{}, nil
	}

	_, err = cr.CallMethod(ctx, msg, false, &ref, method, core.Arguments{}, &prototypeRef)
	require.Error(t, err)
	assert.Contains(t, "canceled", err.Error())

	_, ok := cr.ResultMap[msg.Sequence]
	assert.Equal(t, false, ok)
}

func TestCallMethodWaitResults(t *testing.T) {
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*10)
	defer cancelFunc()

	cr, err := New()
	require.NoError(t, err)

	mc := minimock.NewController(t)
	defer mc.Finish()

	mb := testutils.NewMessageBusMock(mc)
	cr.MessageBus = mb

	msg := &message.BaseLogicMessage{
		Nonce: randomUint64(),
	}
	ref := testutils.RandomRef()
	prototypeRef := testutils.RandomRef()
	method := testutils.RandomString()

	mb.SendFunc = func(p context.Context, p1 core.Message, p2 *core.MessageSendOptions) (r core.Reply, r1 error) {
		go func() {
			r, ok := p1.(*message.CallMethod)
			require.Equal(t, ok, true)

			cr.ResultMutex.Lock()
			defer cr.ResultMutex.Unlock()
			resChan, ok := cr.ResultMap[r.Sequence]
			resChan <- &message.ReturnResults{
				Reply: &reply.CallMethod{},
			}
		}()
		return &reply.RegisterRequest{}, nil
	}
	_, err = cr.CallMethod(ctx, msg, false, &ref, method, core.Arguments{}, &prototypeRef)
	require.NoError(t, err)
}

type callTestSuite struct {
	suite.Suite

	mc         *minimock.Controller
	ctx        context.Context
	cancelFunc context.CancelFunc
	cr         *ContractRequester
	mb         *testutils.MessageBusMock

	msg          *message.BaseLogicMessage
	ref          core.RecordRef
	prototypeRef core.RecordRef
	method       string
}

func (suite callTestSuite) BeforeTest(suiteName string, testName string) {
	suite.mc = minimock.NewController(suite.T())

	suite.ctx = inslogger.TestContext(suite.T())
	suite.ctx, suite.cancelFunc = context.WithTimeout(suite.ctx, time.Second)

	suite.mb = testutils.NewMessageBusMock(suite.mc)

	suite.msg = &message.BaseLogicMessage{
		Nonce: randomUint64(),
	}

	suite.cr, _ = New() // TODO why can't handle error here?
	suite.cr.MessageBus = suite.mb

	suite.ref = testutils.RandomRef()
	suite.prototypeRef = testutils.RandomRef()
	suite.method = testutils.RandomString()
}

func (suite *callTestSuite) AfterTest(suiteName, testName string) {
	suite.mc.Finish()
	suite.cancelFunc()
}

func (suite *callTestSuite) TestCallMethodWaitResults() {
	//mb := testutils.NewMessageBusMock(suite.mc)
	suite.mb.SendMock.Set(
		func(p context.Context, p1 core.Message, p2 *core.MessageSendOptions) (r core.Reply, r1 error) {
			go func() {
				r, ok := p1.(*message.CallMethod)
				suite.Equal(ok, true)

				suite.cr.ResultMutex.Lock()
				defer suite.cr.ResultMutex.Unlock()
				resChan, ok := suite.cr.ResultMap[r.Sequence]
				resChan <- &message.ReturnResults{
					Reply: &reply.CallMethod{},
				}
			}()
			return &reply.RegisterRequest{}, nil
		})

	//suite.cr.MessageBus = mb
	_, err := suite.cr.CallMethod(suite.ctx, suite.msg, false, &suite.ref, suite.method, core.Arguments{}, &suite.prototypeRef)
	suite.NoError(err)
}

func TestContractRequesterCalls(t *testing.T) {
	suite.Run(t, new(callTestSuite))
}
