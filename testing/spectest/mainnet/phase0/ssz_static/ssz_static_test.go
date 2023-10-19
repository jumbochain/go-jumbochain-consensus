package ssz_static

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/phase0/ssz_static"
)

func TestMainnet_Phase0_SSZStatic(t *testing.T) {
	ssz_static.RunSSZStaticTests(t, "mainnet")
}
