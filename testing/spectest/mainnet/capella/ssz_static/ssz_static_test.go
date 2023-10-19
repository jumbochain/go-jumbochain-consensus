package ssz_static

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/capella/ssz_static"
)

func TestMainnet_Capella_SSZStatic(t *testing.T) {
	ssz_static.RunSSZStaticTests(t, "mainnet")
}
