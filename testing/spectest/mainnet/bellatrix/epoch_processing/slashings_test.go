package epoch_processing

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/bellatrix/epoch_processing"
)

func TestMainnet_Bellatrix_EpochProcessing_Slashings(t *testing.T) {
	epoch_processing.RunSlashingsTests(t, "mainnet")
}
