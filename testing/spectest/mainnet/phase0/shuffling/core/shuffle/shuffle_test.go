package shuffle

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/phase0/shuffling/core/shuffle"
)

func TestMainnet_Phase0_Shuffling_Core_Shuffle(t *testing.T) {
	shuffle.RunShuffleTests(t, "mainnet")
}
