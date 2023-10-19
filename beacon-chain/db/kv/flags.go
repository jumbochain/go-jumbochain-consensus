package kv

import (
	"fmt"

	"github.com/jumbochain/go-jumbochain-consensus/cmd/beacon-chain/flags"
	"github.com/jumbochain/go-jumbochain-consensus/config/params"
	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/primitives"
	"github.com/urfave/cli/v2"
)

var maxEpochsToPersistBlobs = params.BeaconNetworkConfig().MinEpochsForBlobsSidecarsRequest

// ConfigureBlobRetentionEpoch sets the for blob retention based on command-line context. It sets the local config `maxEpochsToPersistBlobs`.
// If the flag is not set, the spec default `MinEpochsForBlobsSidecarsRequest` is used.
// An error if the input epoch is smaller than the spec default value.
func ConfigureBlobRetentionEpoch(cliCtx *cli.Context) error {
	// Check if the blob retention epoch flag is set.
	if cliCtx.IsSet(flags.BlobRetentionEpoch.Name) {
		// Retrieve and cast the epoch value.
		epochValue := cliCtx.Uint64(flags.BlobRetentionEpoch.Name)
		e := primitives.Epoch(epochValue)

		// Validate the epoch value against the spec default.
		if e < params.BeaconNetworkConfig().MinEpochsForBlobsSidecarsRequest {
			return fmt.Errorf("%s smaller than spec default, %d < %d", flags.BlobRetentionEpoch.Name, e, params.BeaconNetworkConfig().MinEpochsForBlobsSidecarsRequest)
		}

		maxEpochsToPersistBlobs = e
	}

	return nil
}
