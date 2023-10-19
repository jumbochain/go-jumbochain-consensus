package ssz_static

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/deneb/ssz_static"
)

func TestMainnet_Deneb_SSZStatic(t *testing.T) {
	ssz_static.RunSSZStaticTests(t, "mainnet")
}
