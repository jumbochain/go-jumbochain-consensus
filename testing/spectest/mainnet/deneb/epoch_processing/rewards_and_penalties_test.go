package epoch_processing

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/deneb/epoch_processing"
)

func TestMainnet_Deneb_EpochProcessing_RewardsAndPenalties(t *testing.T) {
	epoch_processing.RunRewardsAndPenaltiesTests(t, "mainnet")
}
