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

package ledgertestutils

import (
	"context"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/ledger"
	"github.com/insolar/insolar/ledger/artifactmanager"
	"github.com/insolar/insolar/ledger/localstorage"
	"github.com/insolar/insolar/ledger/pulsemanager"
	"github.com/insolar/insolar/ledger/recentstorage"
	"github.com/insolar/insolar/ledger/storage"
	"github.com/insolar/insolar/ledger/storage/jet"
	"github.com/insolar/insolar/ledger/storage/storagetest"
	"github.com/insolar/insolar/log"
	"github.com/insolar/insolar/messagebus"
	"github.com/insolar/insolar/network/nodenetwork"
	"github.com/insolar/insolar/platformpolicy"
	"github.com/insolar/insolar/testutils"
	"github.com/insolar/insolar/testutils/testmessagebus"
	"github.com/stretchr/testify/require"
)

// TmpLedger crteates ledger on top of temporary database.
// Returns *ledger.Ledger and cleanup function.
// FIXME: THIS METHOD IS DEPRECATED. USE MOCKS.
func TmpLedger(t *testing.T, dir string, handlersRole core.StaticRole, c core.Components, closeJets bool) (*ledger.Ledger, *storage.DB, func()) {
	log.Warn("TmpLedger is deprecated. Use mocks.")

	pcs := platformpolicy.NewPlatformCryptographyScheme()
	mc := minimock.NewController(t)

	// Init subcomponents.
	ctx := inslogger.TestContext(t)
	conf := configuration.NewLedger()
	db, dbcancel := storagetest.TmpDB(ctx, t, storagetest.Dir(dir))
	pulseStorage := storage.NewPulseStorage()
	pulseStorage.PulseTracker = db

	pulse, err := db.GetLatestPulse(ctx)
	require.NoError(t, err)
	pulseStorage.Set(&pulse.Pulse)

	am := artifactmanager.NewArtifactManger(db)
	am.PlatformCryptographyScheme = pcs
	conf.PulseManager.HeavySyncEnabled = false
	pm := pulsemanager.NewPulseManager(conf)
	ls := localstorage.NewLocalStorage(db)
	jc := testutils.NewJetCoordinatorMock(mc)
	jc.IsAuthorizedMock.Return(true, nil)
	jc.LightExecutorForJetMock.Return(&core.RecordRef{}, nil)
	jc.HeavyMock.Return(&core.RecordRef{}, nil)
	jc.MeMock.Return(core.RecordRef{})

	// Init components.
	if c.MessageBus == nil {
		mb := testmessagebus.NewTestMessageBus(t)
		mb.PulseStorage = pulseStorage
		c.MessageBus = mb
	} else {
		switch mb := c.MessageBus.(type) {
		case *messagebus.MessageBus:
			mb.PulseStorage = pulseStorage
		case *testmessagebus.TestMessageBus:
			mb.PulseStorage = pulseStorage
		default:
			panic("unknown message bus")
		}
	}
	if c.NodeNetwork == nil {
		c.NodeNetwork = nodenetwork.NewNodeKeeper(nodenetwork.NewNode(core.RecordRef{}, core.StaticRoleLightMaterial, nil, "127.0.0.1:5432", ""))
	}

	certificate := testutils.NewCertificateMock(t)
	certificate.GetRoleMock.Return(handlersRole)

	handler := artifactmanager.NewMessageHandler(&conf, certificate)
	handler.PulseTracker = db
	handler.JetStorage = db
	handler.ActiveNodesStorage = db
	handler.DBContext = db
	handler.PulseTracker = db
	handler.ObjectStorage = db

	handler.PlatformCryptographyScheme = pcs
	handler.JetCoordinator = jc

	gilMock := testutils.NewGlobalInsolarLockMock(t)
	gilMock.AcquireFunc = func(context.Context) {}
	gilMock.ReleaseFunc = func(context.Context) {}

	alsMock := testutils.NewActiveListSwapperMock(t)
	alsMock.MoveSyncToActiveFunc = func() error { return nil }

	handler.Bus = c.MessageBus
	am.DefaultBus = c.MessageBus
	pm.NodeNet = c.NodeNetwork
	pm.GIL = gilMock
	pm.Bus = c.MessageBus
	pm.LR = c.LogicRunner
	pm.ActiveListSwapper = alsMock
	pm.PulseStorage = pulseStorage

	hdw := artifactmanager.NewHotDataWaiterConcrete()

	pm.HotDataWaiter = hdw
	handler.HotDataWaiter = hdw

	recentStorageMock := recentstorage.NewRecentStorageMock(t)
	recentStorageMock.AddPendingRequestMock.Return()
	recentStorageMock.AddObjectMock.Return()
	recentStorageMock.RemovePendingRequestMock.Return()
	recentStorageMock.GetRequestsForObjectMock.Return(nil)

	provideMock := recentstorage.NewProviderMock(t)
	provideMock.GetStorageFunc = func(ctx context.Context, p core.RecordID) (r recentstorage.RecentStorage) {
		return recentStorageMock
	}

	handler.RecentStorageProvider = provideMock

	err = handler.Init(ctx)
	if err != nil {
		panic(err)
	}

	if closeJets {
		pm.HotDataWaiter.Unlock(ctx, *jet.NewID(0, nil))
	}

	// Create ledger.
	l := ledger.NewTestLedger(db, am, pm, jc, ls)

	return l, db, dbcancel
}
