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

package storage

import (
	"context"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/ledger/storage/object"
)

//go:generate minimock -i github.com/insolar/insolar/ledger/storage.ObjectStorage -o ./ -s _mock.go

// ObjectStorage returns objects and their meta
type ObjectStorage interface {
	IterateIndexIDs(
		ctx context.Context,
		jetID insolar.ID,
		handler func(id insolar.ID) error,
	) error

	GetObjectIndex(
		ctx context.Context,
		jetID insolar.ID,
		id *insolar.ID,
	) (*object.Lifeline, error)

	SetObjectIndex(
		ctx context.Context,
		jetID insolar.ID,
		id *insolar.ID,
		idx *object.Lifeline,
	) error
}

type objectStorage struct {
	DB                         DBContext                          `inject:""`
	PlatformCryptographyScheme insolar.PlatformCryptographyScheme `inject:""`
}

func NewObjectStorage() ObjectStorage {
	return new(objectStorage)
}

// IterateIndexIDs iterates over index IDs on provided Jet ID.
func (os *objectStorage) IterateIndexIDs(
	ctx context.Context,
	jetID insolar.ID,
	handler func(id insolar.ID) error,
) error {
	jetPrefix := insolar.JetID(jetID).Prefix()
	prefix := prefixkey(scopeIDLifeline, jetPrefix)

	return os.DB.iterate(ctx, prefix, func(k, v []byte) error {
		pn := pulseNumFromKey(0, k)
		id := insolar.NewID(pn, k[insolar.PulseNumberSize:])
		err := handler(*id)
		if err != nil {
			return err
		}
		return nil
	})
}

// GetObjectIndex wraps matching transaction manager method.
func (os *objectStorage) GetObjectIndex(
	ctx context.Context,
	jetID insolar.ID,
	id *insolar.ID,
) (*object.Lifeline, error) {
	prefix := insolar.JetID(jetID).Prefix()
	k := prefixkey(scopeIDLifeline, prefix, id[:])
	buf, err := os.DB.Get(ctx, k)
	if err != nil {
		return nil, err
	}
	res := object.DecodeIndex(buf)
	return &res, nil
}

// SetObjectIndex wraps matching transaction manager method.
func (os *objectStorage) SetObjectIndex(
	ctx context.Context,
	jetID insolar.ID,
	id *insolar.ID,
	idx *object.Lifeline,
) error {
	prefix := insolar.JetID(jetID).Prefix()
	k := prefixkey(scopeIDLifeline, prefix, id[:])
	if idx.Delegates == nil {
		idx.Delegates = map[insolar.Reference]insolar.Reference{}
	}
	encoded := object.EncodeIndex(*idx)
	return os.DB.Set(ctx, k, encoded)
}

func pulseNumFromKey(from int, key []byte) insolar.PulseNumber {
	return insolar.NewPulseNumber(key[from : from+insolar.PulseNumberSize])
}
