package kv

import (
	"flag"
	"strconv"
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/cmd/beacon-chain/flags"
	"github.com/jumbochain/go-jumbochain-consensus/config/params"
	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/primitives"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
	"github.com/urfave/cli/v2"
)

func TestConfigureBlobRetentionEpoch(t *testing.T) {
	maxEpochsToPersistBlobs = params.BeaconNetworkConfig().MinEpochsForBlobsSidecarsRequest
	params.SetupTestConfigCleanup(t)
	app := cli.App{}
	set := flag.NewFlagSet("test", 0)

	// Test case: Spec default.
	require.NoError(t, ConfigureBlobRetentionEpoch(cli.NewContext(&app, set, nil)))
	require.Equal(t, params.BeaconNetworkConfig().MinEpochsForBlobsSidecarsRequest, maxEpochsToPersistBlobs)

	set.Uint64(flags.BlobRetentionEpoch.Name, 0, "")
	minEpochsForSidecarRequest := uint64(params.BeaconNetworkConfig().MinEpochsForBlobsSidecarsRequest)
	require.NoError(t, set.Set(flags.BlobRetentionEpoch.Name, strconv.FormatUint(2*minEpochsForSidecarRequest, 10)))
	cliCtx := cli.NewContext(&app, set, nil)

	// Test case: Input epoch is greater than or equal to spec value.
	require.NoError(t, ConfigureBlobRetentionEpoch(cliCtx))
	require.Equal(t, primitives.Epoch(2*minEpochsForSidecarRequest), maxEpochsToPersistBlobs)

	// Test case: Input epoch is less than spec value.
	require.NoError(t, set.Set(flags.BlobRetentionEpoch.Name, strconv.FormatUint(minEpochsForSidecarRequest-1, 10)))
	cliCtx = cli.NewContext(&app, set, nil)
	err := ConfigureBlobRetentionEpoch(cliCtx)
	require.ErrorContains(t, "extend-blob-retention-epoch smaller than spec default", err)
}
