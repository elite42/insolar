/*
 * The Clear BSD License
 *
 * Copyright (c) 2019 Insolar Technologies
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without modification, are permitted (subject to the limitations in the disclaimer below) provided that the following conditions are met:
 *
 *  Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 *  Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 *  Neither the name of Insolar Technologies nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
 *
 * NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package phases

import (
	"context"
	"fmt"

	"github.com/insolar/insolar/consensus/packets"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/log"
	"github.com/insolar/insolar/network"
	"github.com/insolar/insolar/network/merkle"
	"github.com/insolar/insolar/network/nodenetwork"
	"github.com/pkg/errors"
)

type SecondPhase interface {
	Execute(ctx context.Context, state *FirstPhaseState) (*SecondPhaseState, error)
	Execute21(ctx context.Context, state *SecondPhaseState) (*SecondPhaseState, error)
}

func NewSecondPhase() SecondPhase {
	return &secondPhase{}
}

type secondPhase struct {
	NodeKeeper   network.NodeKeeper       `inject:""`
	Calculator   merkle.Calculator        `inject:""`
	Communicator Communicator             `inject:""`
	Cryptography core.CryptographyService `inject:""`
}

func (sp *secondPhase) Execute(ctx context.Context, state *FirstPhaseState) (*SecondPhaseState, error) {
	prevCloudHash := sp.NodeKeeper.GetCloudHash()

	state.ValidProofs[sp.NodeKeeper.GetOrigin()] = state.PulseProof

	entry := &merkle.GlobuleEntry{
		PulseEntry:    state.PulseEntry,
		ProofSet:      state.ValidProofs,
		PulseHash:     state.PulseHash,
		PrevCloudHash: prevCloudHash,
		GlobuleID:     sp.NodeKeeper.GetOrigin().GetGlobuleID(),
	}
	globuleHash, globuleProof, err := sp.Calculator.GetGlobuleProof(entry)

	if err != nil {
		return nil, errors.Wrap(err, "[ SecondPhase ] Failed to calculate globule proof")
	}

	packet := packets.NewPhase2Packet()
	err = packet.SetGlobuleHashSignature(globuleProof.Signature.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "[ SecondPhase ] Failed to set globule proof in Phase2Packet")
	}
	bitset, err := sp.generatePhase2Bitset(state.UnsyncList, state.ValidProofs)
	if err != nil {
		return nil, errors.Wrap(err, "[ SecondPhase ] Failed to generate bitset for Phase2Packet")
	}
	packet.SetBitSet(bitset)
	activeNodes := state.UnsyncList.GetActiveNodes()
	packets, err := sp.Communicator.ExchangePhase2(ctx, state.UnsyncList, activeNodes, packet)
	if err != nil {
		return nil, errors.Wrap(err, "[ SecondPhase ] Failed to exchange packets on phase 2")
	}
	inslogger.FromContext(ctx).Infof("[ SecondPhase ] received responses: %d/%d", len(packets), len(activeNodes))

	origin := sp.NodeKeeper.GetOrigin().ID()
	stateMatrix := NewStateMatrix(state.UnsyncList)

	for ref, packet := range packets {
		err = nil
		if !ref.Equal(origin) {
			err = sp.checkPacketSignature(packet, ref, state.UnsyncList)
		}
		if err != nil {
			inslogger.FromContext(ctx).Warnf("Failed to check phase2 packet signature from %s: %s", ref, err.Error())
			continue
		}
		state.UnsyncList.SetGlobuleHashSignature(ref, packet.GetGlobuleHashSignature())
		err = stateMatrix.ApplyBitSet(ref, packet.GetBitSet())
		if err != nil {
			log.Warnf("[ SecondPhase ] Could not apply bitset from node %s", ref)
			continue
		}
	}

	matrixCalculation, err := stateMatrix.CalculatePhase2(origin)
	if err != nil {
		return nil, errors.Wrap(err, "[ SecondPhase ] Failed to calculate bitset matrix consensus result")
	}

	if len(matrixCalculation.TimedOut) > 0 {
		for _, nodeID := range matrixCalculation.TimedOut {
			state.UnsyncList.RemoveNode(nodeID)
		}

		type none struct{}
		newActive := make(map[core.RecordRef]none)
		for _, active := range matrixCalculation.Active {
			newActive[active] = none{}
		}

		newProofs := make(map[core.Node]*merkle.PulseProof)
		for node, proof := range state.ValidProofs {
			_, ok := newActive[node.ID()]
			if !ok {
				continue
			}
			newProofs[node] = proof
		}

		state.ValidProofs = newProofs

		entry := &merkle.GlobuleEntry{
			PulseEntry:    state.PulseEntry,
			ProofSet:      state.ValidProofs,
			PulseHash:     state.PulseHash,
			PrevCloudHash: prevCloudHash,
			GlobuleID:     sp.NodeKeeper.GetOrigin().GetGlobuleID(),
		}
		globuleHash, globuleProof, err = sp.Calculator.GetGlobuleProof(entry)

		if err != nil {
			return nil, errors.Wrap(err, "[ SecondPhase ] Failed to calculate globule proof")
		}
	}

	return &SecondPhaseState{
		FirstPhaseState: state,
		Matrix:          stateMatrix,
		MatrixState:     matrixCalculation,
		BitSet:          bitset,

		GlobuleHash:  globuleHash,
		GlobuleProof: globuleProof,
	}, nil
}

func (sp *secondPhase) Execute21(ctx context.Context, state *SecondPhaseState) (*SecondPhaseState, error) {
	additionalRequests := state.MatrixState.AdditionalRequestsPhase2

	count := len(additionalRequests)
	results := make(map[uint16]*packets.MissingNodeSupplementaryVote)
	claims := make(map[uint16]*packets.MissingNodeClaim)

	packet := packets.NewPhase2Packet()
	err := packet.SetGlobuleHashSignature(state.GlobuleProof.Signature.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "[ Phase 2.1 ] Failed to set pulse proof in Phase2Packet.")
	}
	packet.SetBitSet(state.BitSet)

	voteAnswers, err := sp.Communicator.ExchangePhase21(ctx, state.UnsyncList, packet, additionalRequests)
	if err != nil {
		return nil, errors.Wrap(err, "[ Phase 2.1 ] Failed to send additional requests on phase 2.1")
	}

	if len(additionalRequests) == 0 {
		return state, nil
	}

	for _, vote := range voteAnswers {
		switch v := vote.(type) {
		case *packets.MissingNodeSupplementaryVote:
			results[v.NodeIndex] = v
		case *packets.MissingNodeClaim:
			claims[v.NodeIndex] = v
		}
	}

	if len(results) != count {
		return nil, errors.New(fmt.Sprintf("[ Phase 2.1 ] Failed to receive enough MissingNodeSupplementaryVote responses,"+
			" received: %d/%d", len(results), count))
	}

	origin := sp.NodeKeeper.GetOrigin().ID()
	bitsetChanges := make([]packets.BitSetCell, 0)
	for index, result := range results {
		node, err := nodenetwork.ClaimToNode("", &result.NodeClaimUnsigned)
		if err != nil {
			return nil, errors.Wrapf(err, "[ Phase 2.1 ] Failed to convert claim to node, ref: %s",
				result.NodeClaimUnsigned.NodeRef)
		}

		merkleProof := &merkle.PulseProof{
			BaseProof: merkle.BaseProof{
				Signature: core.SignatureFromBytes(result.NodePulseProof.Signature()),
			},
			StateHash: result.NodePulseProof.StateHash(),
		}

		valid := validateProof(sp.Calculator, state.UnsyncList, state.PulseHash, node.ID(), merkleProof)
		if !valid {
			inslogger.FromContext(ctx).Warnf("Failed to validate proof from %s", node.ID())
			continue
		}

		err = state.Matrix.ReceivedProofFromNode(origin, node.ID())
		if err != nil {
			return nil, errors.Wrapf(err, "[ Phase 2.1 ] Failed to assign proof from node %s to state matrix",
				result.NodeClaimUnsigned.NodeRef)
		}
		state.UnsyncList.AddNode(node, index)
		state.UnsyncList.AddProof(node.ID(), &result.NodePulseProof)
		state.ValidProofs[node] = merkleProof
		bitsetChanges = append(bitsetChanges, packets.BitSetCell{NodeID: node.ID(), State: packets.Legit})
	}

	err = state.BitSet.ApplyChanges(bitsetChanges, state.UnsyncList)
	if err != nil {
		return nil, errors.Wrap(err, "[ Phase 2.1 ] Failed to apply changes to current bitset")
	}
	claimMap := make(map[core.RecordRef][]packets.ReferendumClaim)
	for index, claim := range claims {
		ref, err := state.UnsyncList.IndexToRef(int(index))
		if err != nil {
			return nil, errors.Wrapf(err, "[ Phase 2.1 ] Failed to map index %d to ref", index)
		}
		list := claimMap[ref]
		if list == nil {
			list = make([]packets.ReferendumClaim, 0)
		}
		list = append(list, claim.Claim)
		claimMap[ref] = list
	}
	if err = state.UnsyncList.AddClaims(claimMap); err != nil {
		return nil, errors.Wrapf(err, "[ Phase 2.1 ] Failed to add claims")
	}
	state.MatrixState, err = state.Matrix.CalculatePhase2(origin)
	if err != nil {
		return nil, errors.Wrap(err, "[ Phase 2.1 ] Failed to calculate matrix state")
	}
	addReqCount := len(state.MatrixState.AdditionalRequestsPhase2)
	if addReqCount != 0 {
		return nil, errors.New(fmt.Sprintf("[ Phase 2.1 ] Failed to get enough data during phase 2.1 "+
			"(still need additional %d requests)", addReqCount))
	}

	prevCloudHash := sp.NodeKeeper.GetCloudHash()
	entry := &merkle.GlobuleEntry{
		PulseEntry:    state.PulseEntry,
		ProofSet:      state.ValidProofs,
		PulseHash:     state.PulseHash,
		PrevCloudHash: prevCloudHash,
		GlobuleID:     sp.NodeKeeper.GetOrigin().GetGlobuleID(),
	}
	state.GlobuleHash, state.GlobuleProof, err = sp.Calculator.GetGlobuleProof(entry)
	if err != nil {
		return nil, errors.Wrap(err, "[ Phase 2.1 ] Failed to calculate globule proof")
	}
	var ghs packets.GlobuleHashSignature
	copy(ghs[:], state.GlobuleProof.Signature.Bytes()[:packets.SignatureLength])
	state.UnsyncList.SetGlobuleHashSignature(origin, ghs)

	return state, nil
}

func (sp *secondPhase) generatePhase2Bitset(list network.UnsyncList, proofs map[core.Node]*merkle.PulseProof) (packets.BitSet, error) {
	bitset, err := packets.NewBitSet(list.Length())
	if err != nil {
		return nil, err
	}
	cells := make([]packets.BitSetCell, 0)
	for node := range proofs {
		cells = append(cells, packets.BitSetCell{NodeID: node.ID(), State: packets.Legit})
	}
	cells = append(cells, packets.BitSetCell{NodeID: sp.NodeKeeper.GetOrigin().ID(), State: packets.Legit})
	err = bitset.ApplyChanges(cells, list)
	if err != nil {
		return nil, err
	}
	return bitset, nil
}

func (sp *secondPhase) checkPacketSignature(packet *packets.Phase2Packet, recordRef core.RecordRef, unsyncList network.UnsyncList) error {
	activeNode := unsyncList.GetActiveNode(recordRef)
	if activeNode == nil {
		return errors.New("failed to get active node")
	}
	key := activeNode.PublicKey()
	return packet.Verify(sp.Cryptography, key)
}
