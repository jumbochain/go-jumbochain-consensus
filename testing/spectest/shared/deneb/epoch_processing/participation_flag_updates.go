package epoch_processing

import (
	"path"
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/altair"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/state"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/utils"
)

// RunParticipationFlagUpdatesTests executes "epoch_processing/participation_flag_updates" tests.
func RunParticipationFlagUpdatesTests(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))

	testFolders, testsFolderPath := utils.TestFolders(t, config, "deneb", "epoch_processing/participation_flag_updates/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			RunEpochOperationTest(t, folderPath, processParticipationFlagUpdatesWrapper)
		})
	}
}

func processParticipationFlagUpdatesWrapper(t *testing.T, st state.BeaconState) (state.BeaconState, error) {
	st, err := altair.ProcessParticipationFlagUpdates(st)
	require.NoError(t, err, "Could not process participation flag update")
	return st, nil
}
