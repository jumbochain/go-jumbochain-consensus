package stategen

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/state"
	state_native "github.com/jumbochain/go-jumbochain-consensus/beacon-chain/state/state-native"
	ethpb "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
	"github.com/jumbochain/go-jumbochain-consensus/testing/assert"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
)

func TestHotStateCache_RoundTrip(t *testing.T) {
	c := newHotStateCache()
	root := [32]byte{'A'}
	s := c.get(root)
	assert.Equal(t, state.BeaconState(nil), s)
	assert.Equal(t, false, c.has(root), "Empty cache has an object")

	s, err := state_native.InitializeFromProtoPhase0(&ethpb.BeaconState{
		Slot: 10,
	})
	require.NoError(t, err)

	c.put(root, s)
	assert.Equal(t, true, c.has(root), "Empty cache does not have an object")

	res := c.get(root)
	assert.NotNil(t, s)
	assert.DeepEqual(t, res.ToProto(), s.ToProto(), "Expected equal protos to return from cache")

	c.delete(root)
	assert.Equal(t, false, c.has(root), "Cache not supposed to have the object")
}
