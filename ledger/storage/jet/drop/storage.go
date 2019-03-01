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

package drop

import (
	"context"
	"sync"

	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/ledger/storage"
	"github.com/insolar/insolar/ledger/storage/jet"
)

// Modifier provides interface for modifying jetdrops
//go:generate minimock -i github.com/insolar/insolar/ledger/storage/jet/drop.Modifier -o ./ -s _mock.go
type Modifier interface {
	Set(ctx context.Context, jetID core.JetID, drop jet.Drop) error
}

// Accessor provides interface for accessing jetdrops
//go:generate minimock -i github.com/insolar/insolar/ledger/storage/jet/drop.Accessor -o ./ -s _mock.go
type Accessor interface {
	ForPulse(ctx context.Context, jetID core.JetID, pulse core.PulseNumber) (jet.Drop, error)
}

type dropForPulseManager struct {
	lock  sync.RWMutex
	drops map[core.PulseNumber]jet.Drop
}

func (m *dropForPulseManager) set(drop jet.Drop, pulse core.PulseNumber) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, ok := m.drops[pulse]; ok {
		return storage.ErrOverride
	}

	m.drops[pulse] = drop

	return nil
}

func (m *dropForPulseManager) forPulse(jetID core.JetID, pulse core.PulseNumber) (jet.Drop, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	drop, ok := m.drops[pulse]
	if !ok {
		return jet.Drop{}, core.ErrNotFound
	}

	return drop, nil
}

type dropStorageMemory struct {
	lock sync.Mutex
	jets map[core.JetID]*dropForPulseManager
}

// NewStorageMemory creates new storage, that holds data in-memory
func NewStorageMemory() *dropStorageMemory {
	return &dropStorageMemory{
		jets: map[core.JetID]*dropForPulseManager{},
	}
}

func (m *dropStorageMemory) fetchStorage(jetID core.JetID) (ds *dropForPulseManager) {
	m.lock.Lock()
	defer m.lock.Unlock()

	ds, ok := m.jets[jetID]
	if !ok {
		m.jets[jetID] = new(dropForPulseManager)
		ds = m.jets[jetID]
	}
	return
}

// ForPulse returns a jet.Drop for a provided pulse, that is stored in memory
func (m *dropStorageMemory) ForPulse(ctx context.Context, jetID core.JetID, pulse core.PulseNumber) (jet.Drop, error) {
	ds := m.fetchStorage(jetID)
	return ds.forPulse(jetID, pulse)
}

// Set saves a provided jet.Drop to memory
func (m *dropStorageMemory) Set(ctx context.Context, jetID core.JetID, drop jet.Drop) error {
	ds := m.fetchStorage(jetID)
	return ds.set(drop, drop.Pulse)
}

type dropStorageDB struct {
	DB storage.DBContext `inject:""`
}

// NewStorageDB creates new storage, that holds data in db
func NewStorageDB() *dropStorageDB {
	return &dropStorageDB{}
}

// ForPulse returns a jet.Drop for a provided pulse, that is stored in db
func (ds *dropStorageDB) ForPulse(ctx context.Context, jetID core.JetID, pulse core.PulseNumber) (jet.Drop, error) {
	_, prefix := jetID.Jet()
	k := storage.JetDropPrefixKey(prefix, pulse)

	// buf, err := db.get(ctx, k)
	buf, err := ds.DB.Get(ctx, k)
	if err != nil {
		return jet.Drop{}, err
	}
	drop, err := jet.Decode(buf)
	if err != nil {
		return jet.Drop{}, err
	}
	return *drop, nil
}

// Set saves a provided jet.Drop to db
func (ds *dropStorageDB) Set(ctx context.Context, jetID core.JetID, drop jet.Drop) error {
	_, prefix := jetID.Jet()
	k := storage.JetDropPrefixKey(prefix, drop.Pulse)
	_, err := ds.DB.Get(ctx, k)
	if err == nil {
		return storage.ErrOverride
	}

	encoded, err := jet.Encode(&drop)
	if err != nil {
		return err
	}
	return ds.DB.Set(ctx, k, encoded)
}
