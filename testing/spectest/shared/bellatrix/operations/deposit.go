package operations

import (
	"context"
	"path"
	"testing"

	"github.com/golang/snappy"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/altair"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/state"
	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/interfaces"
	ethpb "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/utils"
	"github.com/jumbochain/go-jumbochain-consensus/testing/util"
)

func RunDepositTest(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))
	testFolders, testsFolderPath := utils.TestFolders(t, config, "bellatrix", "operations/deposit/pyspec_tests")
	if len(testFolders) == 0 {
		t.Fatalf("No test folders found for %s/%s/%s", config, "bellatrix", "operations/deposit/pyspec_tests")
	}
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			depositFile, err := util.BazelFileBytes(folderPath, "deposit.ssz_snappy")
			require.NoError(t, err)
			depositSSZ, err := snappy.Decode(nil /* dst */, depositFile)
			require.NoError(t, err, "Failed to decompress")
			deposit := &ethpb.Deposit{}
			require.NoError(t, deposit.UnmarshalSSZ(depositSSZ), "Failed to unmarshal")

			body := &ethpb.BeaconBlockBodyBellatrix{Deposits: []*ethpb.Deposit{deposit}}
			processDepositsFunc := func(ctx context.Context, s state.BeaconState, b interfaces.ReadOnlySignedBeaconBlock) (state.BeaconState, error) {
				return altair.ProcessDeposits(ctx, s, b.Block().Body().Deposits())
			}
			RunBlockOperationTest(t, folderPath, body, processDepositsFunc)
		})
	}
}
