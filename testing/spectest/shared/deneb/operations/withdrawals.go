package operations

import (
	"context"
	"path"
	"testing"

	"github.com/golang/snappy"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/blocks"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/state"
	consensusblocks "github.com/jumbochain/go-jumbochain-consensus/consensus-types/blocks"
	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/interfaces"
	enginev1 "github.com/jumbochain/go-jumbochain-consensus/proto/engine/v1"
	ethpb "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/utils"
	"github.com/jumbochain/go-jumbochain-consensus/testing/util"
)

func RunWithdrawalsTest(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))
	testFolders, testsFolderPath := utils.TestFolders(t, config, "deneb", "operations/withdrawals/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			payloadFile, err := util.BazelFileBytes(folderPath, "execution_payload.ssz_snappy")
			require.NoError(t, err)
			payloadSSZ, err := snappy.Decode(nil /* dst */, payloadFile)
			require.NoError(t, err, "Failed to decompress")
			payload := &enginev1.ExecutionPayloadDeneb{}
			require.NoError(t, payload.UnmarshalSSZ(payloadSSZ), "Failed to unmarshal")

			body := &ethpb.BeaconBlockBodyDeneb{ExecutionPayload: payload}
			RunBlockOperationTest(t, folderPath, body, func(_ context.Context, s state.BeaconState, b interfaces.SignedBeaconBlock) (state.BeaconState, error) {
				payload, err := b.Block().Body().Execution()
				if err != nil {
					return nil, err
				}
				withdrawals, err := payload.Withdrawals()
				if err != nil {
					return nil, err
				}
				p, err := consensusblocks.WrappedExecutionPayloadDeneb(&enginev1.ExecutionPayloadDeneb{Withdrawals: withdrawals}, 0)
				require.NoError(t, err)
				return blocks.ProcessWithdrawals(s, p)
			})
		})
	}
}
