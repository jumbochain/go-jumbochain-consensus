package operations

import (
	"context"
	"path"
	"testing"

	"github.com/golang/snappy"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/blocks"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/validators"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/state"
	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/interfaces"
	ethpb "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/utils"
	"github.com/jumbochain/go-jumbochain-consensus/testing/util"
)

func RunProposerSlashingTest(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))
	testFolders, testsFolderPath := utils.TestFolders(t, config, "bellatrix", "operations/proposer_slashing/pyspec_tests")
	if len(testFolders) == 0 {
		t.Fatalf("No test folders found for %s/%s/%s", config, "bellatrix", "operations/proposer_slashing/pyspec_tests")
	}
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			proposerSlashingFile, err := util.BazelFileBytes(folderPath, "proposer_slashing.ssz_snappy")
			require.NoError(t, err)
			proposerSlashingSSZ, err := snappy.Decode(nil /* dst */, proposerSlashingFile)
			require.NoError(t, err, "Failed to decompress")
			proposerSlashing := &ethpb.ProposerSlashing{}
			require.NoError(t, proposerSlashing.UnmarshalSSZ(proposerSlashingSSZ), "Failed to unmarshal")

			body := &ethpb.BeaconBlockBodyBellatrix{ProposerSlashings: []*ethpb.ProposerSlashing{proposerSlashing}}
			RunBlockOperationTest(t, folderPath, body, func(ctx context.Context, s state.BeaconState, b interfaces.ReadOnlySignedBeaconBlock) (state.BeaconState, error) {
				return blocks.ProcessProposerSlashings(ctx, s, b.Block().Body().ProposerSlashings(), validators.SlashValidator)
			})
		})
	}
}
