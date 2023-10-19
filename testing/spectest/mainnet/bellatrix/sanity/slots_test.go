package sanity

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/bellatrix/sanity"
)

func TestMainnet_Bellatrix_Sanity_Slots(t *testing.T) {
	sanity.RunSlotProcessingTests(t, "mainnet")
}
