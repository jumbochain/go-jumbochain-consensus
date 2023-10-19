package util

import (
	"bytes"
	"context"
	"encoding/binary"
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/helpers"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/signing"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/time"
	fieldparams "github.com/jumbochain/go-jumbochain-consensus/config/fieldparams"
	"github.com/jumbochain/go-jumbochain-consensus/config/params"
	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/primitives"
	"github.com/jumbochain/go-jumbochain-consensus/testing/assert"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
	"github.com/jumbochain/go-jumbochain-consensus/time/slots"
)

func TestBlockSignature(t *testing.T) {
	beaconState, privKeys := DeterministicGenesisState(t, 100)
	block, err := GenerateFullBlock(beaconState, privKeys, nil, 0)
	require.NoError(t, err)

	require.NoError(t, beaconState.SetSlot(beaconState.Slot()+1))
	proposerIdx, err := helpers.BeaconProposerIndex(context.Background(), beaconState)
	assert.NoError(t, err)

	assert.NoError(t, beaconState.SetSlot(beaconState.Slot()-1))
	epoch := slots.ToEpoch(block.Block.Slot)
	blockSig, err := signing.ComputeDomainAndSign(beaconState, epoch, block.Block, params.BeaconConfig().DomainBeaconProposer, privKeys[proposerIdx])
	require.NoError(t, err)

	signature, err := BlockSignature(beaconState, block.Block, privKeys)
	assert.NoError(t, err)

	if !bytes.Equal(blockSig, signature.Marshal()) {
		t.Errorf("Expected block signatures to be equal, received %#x != %#x", blockSig, signature.Marshal())
	}
}

func TestRandaoReveal(t *testing.T) {
	beaconState, privKeys := DeterministicGenesisState(t, 100)

	epoch := time.CurrentEpoch(beaconState)
	randaoReveal, err := RandaoReveal(beaconState, epoch, privKeys)
	assert.NoError(t, err)

	proposerIdx, err := helpers.BeaconProposerIndex(context.Background(), beaconState)
	assert.NoError(t, err)
	buf := make([]byte, fieldparams.RootLength)
	binary.LittleEndian.PutUint64(buf, uint64(epoch))
	// We make the previous validator's index sign the message instead of the proposer.
	sszUint := primitives.SSZUint64(epoch)
	epochSignature, err := signing.ComputeDomainAndSign(beaconState, epoch, &sszUint, params.BeaconConfig().DomainRandao, privKeys[proposerIdx])
	require.NoError(t, err)

	if !bytes.Equal(randaoReveal, epochSignature) {
		t.Errorf("Expected randao reveals to be equal, received %#x != %#x", randaoReveal, epochSignature)
	}
}
