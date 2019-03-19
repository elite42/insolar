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
)

type dropKey struct {
	pulse core.PulseNumber
	jetID core.JetID
}

type dropStorageMemory struct {
	lock  sync.RWMutex
	drops map[dropKey]Drop
}

// NewStorageMemory creates a new storage, that holds data in a memory.
func NewStorageMemory() *dropStorageMemory { // nolint: golint
	return &dropStorageMemory{
		drops: map[dropKey]Drop{},
	}
}

// ForPulse returns a Drop for a provided pulse, that is stored in a memory
func (m *dropStorageMemory) ForPulse(ctx context.Context, jetID core.JetID, pulse core.PulseNumber) (Drop, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	key := dropKey{jetID: jetID, pulse: pulse}
	d, ok := m.drops[key]
	if !ok {
		return Drop{}, ErrNotFound
	}

	return d, nil
}

// Set saves a provided Drop to a memory
func (m *dropStorageMemory) Set(ctx context.Context, drop Drop) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	key := dropKey{jetID: drop.JetID, pulse: drop.Pulse}
	_, ok := m.drops[key]
	if ok {
		return ErrOverride
	}
	m.drops[key] = drop

	return nil
}

// Delete methods removes a drop from a memory storage.
func (m *dropStorageMemory) Delete(pulse core.PulseNumber) {
	m.lock.Lock()
	for key := range m.drops {
		if key.pulse == pulse {
			delete(m.drops, key)
		}
	}
	m.lock.Unlock()
}