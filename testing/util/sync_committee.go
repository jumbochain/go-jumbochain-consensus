package util

import (
	fieldparams "github.com/jumbochain/go-jumbochain-consensus/config/fieldparams"
	"github.com/jumbochain/go-jumbochain-consensus/config/params"
	"github.com/jumbochain/go-jumbochain-consensus/encoding/bytesutil"
	ethpb "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
)

// HydrateSyncCommittee hydrates the provided sync committee message.
func HydrateSyncCommittee(s *ethpb.SyncCommitteeMessage) *ethpb.SyncCommitteeMessage {
	if s.Signature == nil {
		s.Signature = make([]byte, 96)
	}
	if s.BlockRoot == nil {
		s.BlockRoot = make([]byte, fieldparams.RootLength)
	}
	return s
}

// ConvertToCommittee takes a list of pubkeys and returns a SyncCommittee with
// these keys as members. Some keys may appear repeated
func ConvertToCommittee(inputKeys [][]byte) *ethpb.SyncCommittee {
	var pubKeys [][]byte
	for i := uint64(0); i < params.BeaconConfig().SyncCommitteeSize; i++ {
		if i < uint64(len(inputKeys)) {
			pubKeys = append(pubKeys, bytesutil.PadTo(inputKeys[i], params.BeaconConfig().BLSPubkeyLength))
		} else {
			pubKeys = append(pubKeys, bytesutil.PadTo([]byte{}, params.BeaconConfig().BLSPubkeyLength))
		}
	}

	return &ethpb.SyncCommittee{
		Pubkeys:         pubKeys,
		AggregatePubkey: bytesutil.PadTo([]byte{}, params.BeaconConfig().BLSPubkeyLength),
	}
}
