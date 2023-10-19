package epoch_processing

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/capella/epoch_processing"
)

func TestMainnet_Capella_EpochProcessing_Eth1DataReset(t *testing.T) {
	epoch_processing.RunEth1DataResetTests(t, "mainnet")
}
