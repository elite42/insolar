//
// Modified BSD 3-Clause Clear License
//
// Copyright (c) 2019 Insolar Technologies GmbH
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted (subject to the limitations in the disclaimer below) provided that
// the following conditions are met:
//  * Redistributions of source code must retain the above copyright notice, this list
//    of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright notice, this list
//    of conditions and the following disclaimer in the documentation and/or other materials
//    provided with the distribution.
//  * Neither the name of Insolar Technologies GmbH nor the names of its contributors
//    may be used to endorse or promote products derived from this software without
//    specific prior written permission.
//
// NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED
// BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS
// AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
// BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
// OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Notwithstanding any other provisions of this license, it is prohibited to:
//    (a) use this software,
//
//    (b) prepare modifications and derivative works of this software,
//
//    (c) distribute this software (including without limitation in source code, binary or
//        object code form), and
//
//    (d) reproduce copies of this software
//
//    for any commercial purposes, and/or
//
//    for the purposes of making available this software to third parties as a service,
//    including, without limitation, any software-as-a-service, platform-as-a-service,
//    infrastructure-as-a-service or other similar online service, irrespective of
//    whether it competes with the products or services of Insolar Technologies GmbH.
//

package storage

import (
	"context"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/network/node"
	"github.com/pkg/errors"
	"sync"
)

//go:generate minimock -i github.com/insolar/insolar/network/storage.SnapshotAccessor -o ../../testutils/network -s _mock.go

// SnapshotAccessor provides methods for accessing Snapshot.
type SnapshotAccessor interface {
	ForPulseNumber(context.Context, insolar.PulseNumber) (*node.Snapshot, error)
	Latest(ctx context.Context) (node.Snapshot, error)
}

//go:generate minimock -i github.com/insolar/insolar/network/storage.SnapshotAppender -o ../../testutils/network -s _mock.go

// SnapshotAppender provides method for appending Snapshot to storage.
type SnapshotAppender interface {
	Append(ctx context.Context, pulse insolar.PulseNumber, snapshot *node.Snapshot) error
}

// NewSnapshotStorage constructor creates PulseStorage
func NewSnapshotStorage() *SnapshotStorage {
	return &SnapshotStorage{}
}

type SnapshotStorage struct {
	DB   DB `inject:""`
	lock sync.RWMutex
}

func (s *SnapshotStorage) Append(ctx context.Context, pulse insolar.PulseNumber, snapshot *node.Snapshot) error {
	s.lock.RLock()
	defer s.lock.RUnlock()

	buff, err := snapshot.Encode()
	if err != nil {
		return errors.Wrap(err, "[SnapshotStorage] Failed to append snapshot")
	}
	return s.DB.Set(pulseKey(pulse), buff)
}

func (s *SnapshotStorage) ForPulseNumber(ctx context.Context, pulse insolar.PulseNumber) (*node.Snapshot, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	buf, err := s.DB.Get(pulseKey(pulse))
	if err != nil {
		return nil, errors.Wrap(err, "[SnapshotStorage] Failed to get snapshot from DB")
	}
	result := &node.Snapshot{}
	err = result.Decode(buf)
	if err != nil {
		return nil, errors.Wrap(err, "[SnapshotStorage] Failed to decode snapshot")
	}
	return result, nil
}

func (s *SnapshotStorage) Latest(ctx context.Context) (*node.Snapshot, error) {
	panic("implement me")
}
