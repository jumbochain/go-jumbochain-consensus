package epoch_processing

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/phase0/epoch_processing"
)

func TestMainnet_Phase0_EpochProcessing_Eth1DataReset(t *testing.T) {
	epoch_processing.RunEth1DataResetTests(t, "mainnet")
}
