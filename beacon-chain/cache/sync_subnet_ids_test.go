package cache

import (
	"testing"

	fieldparams "github.com/jumbochain/go-jumbochain-consensus/config/fieldparams"
	"github.com/jumbochain/go-jumbochain-consensus/config/params"
	"github.com/jumbochain/go-jumbochain-consensus/testing/assert"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
)

func TestSyncSubnetIDsCache_Roundtrip(t *testing.T) {
	c := newSyncSubnetIDs()

	for i := 0; i < 20; i++ {
		pubkey := [fieldparams.BLSPubkeyLength]byte{byte(i)}
		c.AddSyncCommitteeSubnets(pubkey[:], 100, []uint64{uint64(i)}, 0)
	}

	for i := uint64(0); i < 20; i++ {
		pubkey := [fieldparams.BLSPubkeyLength]byte{byte(i)}

		idxs, _, ok, _ := c.GetSyncCommitteeSubnets(pubkey[:], 100)
		if !ok {
			t.Errorf("Couldn't find entry in cache for pubkey %#x", pubkey)
			continue
		}
		require.Equal(t, i, idxs[0])
	}
	coms := c.GetAllSubnets(100)
	assert.Equal(t, 20, len(coms))
}

func TestSyncSubnetIDsCache_ValidateCurrentEpoch(t *testing.T) {
	c := newSyncSubnetIDs()

	for i := 0; i < 20; i++ {
		pubkey := [fieldparams.BLSPubkeyLength]byte{byte(i)}
		c.AddSyncCommitteeSubnets(pubkey[:], 100, []uint64{uint64(i)}, 0)
	}

	coms := c.GetAllSubnets(50)
	assert.Equal(t, 0, len(coms))

	for i := uint64(0); i < 20; i++ {
		pubkey := [fieldparams.BLSPubkeyLength]byte{byte(i)}

		_, jEpoch, ok, _ := c.GetSyncCommitteeSubnets(pubkey[:], 100)
		if !ok {
			t.Errorf("Couldn't find entry in cache for pubkey %#x", pubkey)
			continue
		}
		require.Equal(t, true, uint64(jEpoch) >= 100-params.BeaconConfig().SyncCommitteeSubnetCount)
	}

	coms = c.GetAllSubnets(99)
	assert.Equal(t, 20, len(coms))
}
