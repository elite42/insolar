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
	"bytes"
	"context"
	"errors"

	"github.com/dgraph-io/badger"
	"github.com/insolar/insolar/insolar"
)

// iterstate stores iterator state
type iterstate struct {
	prefix []byte
	start  []byte
	end    []byte
}

// ReplicaIter provides partial iterator over BadgerDB key/value pairs
// required for replication to Heavy Material node in provided pulses range.
//
// "Required KV pairs" are all keys with record's namespaces
// in provided pulses range and all indexes from zero pulse to the end of provided range.
//
// "Partial" means it fetches data in chunks of the specified size.
// After a chunk has been fetched, an iterator saves current position.
//
// NOTE: This is not an "honest" alogrithm, because the last record size can exceed the limit.
// Better implementation is for the future work.
type ReplicaIter struct {
	ctx        context.Context
	dbContext  DBContext
	limitBytes int
	istates    []*iterstate
	lastpulse  insolar.PulseNumber
}

// NewReplicaIter creates ReplicaIter what iterates over records on jet,
// required for heavy material replication.
//
// Params 'start' and 'end' defines pulses from which scan should happen,
// and on which it should be stopped, but indexes scan are always started
// from insolar.FirstPulseNumber.
//
// Param 'limit' sets per message limit.
func NewReplicaIter(
	ctx context.Context,
	dbContext DBContext,
	jetID insolar.ID,
	start insolar.PulseNumber,
	end insolar.PulseNumber,
	limit int,
) *ReplicaIter {
	// fmt.Printf("CALL NewReplicaIter [%v:%v] (jet=%v)\n", start, end, jetID)
	newit := func(prefixbyte byte, jetID insolar.ID, start, end insolar.PulseNumber) *iterstate {
		prefix := []byte{prefixbyte}
		jetPrefix := insolar.JetID(jetID).Prefix()
		iter := &iterstate{prefix: prefix}
		iter.start = bytes.Join([][]byte{prefix, jetPrefix[:], start.Bytes()}, nil)
		iter.end = bytes.Join([][]byte{prefix, jetPrefix[:], end.Bytes()}, nil)
		return iter
	}

	return &ReplicaIter{
		ctx:        ctx,
		dbContext:  dbContext,
		limitBytes: limit,
		// record iterators (order matters for heavy node consistency)
		istates: []*iterstate{
			newit(scopeIDLifeline, jetID, insolar.FirstPulseNumber, end),
		},
	}
}

// NextRecords fetches next part of key value pairs.
func (r *ReplicaIter) NextRecords() ([]insolar.KV, error) {
	if r.isDone() {
		return nil, ErrReplicatorDone
	}
	fc := &fetchchunk{
		db:    r.dbContext.GetBadgerDB(),
		limit: r.limitBytes,
	}
	for _, is := range r.istates {
		if is.start == nil {
			continue
		}
		var fetcherr error
		var lastpulse insolar.PulseNumber
		is.start, lastpulse, fetcherr = fc.fetch(r.ctx, is.prefix, is.start, is.end)
		if fetcherr != nil {
			return nil, fetcherr
		}
		if lastpulse > r.lastpulse {
			r.lastpulse = lastpulse
		}
	}
	return fc.records, nil
}

// LastPulse returns maximum pulse number of returned keys after each fetch.
func (r *ReplicaIter) LastSeenPulse() insolar.PulseNumber {
	return r.lastpulse
}

// ErrReplicatorDone is returned by an Replicator NextRecords method when the iteration is complete.
var ErrReplicatorDone = errors.New("no more items in iterator")

func (r *ReplicaIter) isDone() bool {
	for _, is := range r.istates {
		if is.start != nil {
			return false
		}
	}
	return true
}

type fetchchunk struct {
	db      *badger.DB
	records []insolar.KV
	size    int
	limit   int
}

func (fc *fetchchunk) fetch(
	ctx context.Context,
	prefix []byte,
	start []byte,
	end []byte,
) ([]byte, insolar.PulseNumber, error) {
	if fc.size > fc.limit {
		return start, 0, nil
	}

	var nextstart []byte
	var lastpulse insolar.PulseNumber
	err := fc.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek(start); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			if item == nil {
				break
			}
			// key prefix < end
			if bytes.Compare(item.Key()[:len(end)], end) != -1 {
				break
			}

			key := item.KeyCopy(nil)
			if fc.size > fc.limit {
				nextstart = key
				// inslogger.FromContext(ctx).Warnf("size > r.limit: %v > %v (nextstart=%v)",
				// 	fc.size, fc.limit, hex.EncodeToString(key))
				return nil
			}

			lastpulse = pulseFromKey(key)
			// fmt.Printf("Replica> key: %v (pulse=%v)\n", hex.EncodeToString(key), lastpulse)

			value, err := it.Item().ValueCopy(nil)
			if err != nil {
				return err
			}

			NullifyJetInKey(key)
			fc.records = append(fc.records, insolar.KV{K: key, V: value})
			fc.size += len(key) + len(value)
		}
		nextstart = nil
		return nil
	})
	return nextstart, lastpulse, err
}

// NullifyJetInKey nullify jet part in record.
func NullifyJetInKey(key []byte) {
	for i := 1; i < insolar.RecordHashSize; i++ {
		key[i] = 0
	}
}
