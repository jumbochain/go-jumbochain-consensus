package forkchoice

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/runtime/version"
	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/common/forkchoice"
)

func TestMainnet_Deneb_Forkchoice(t *testing.T) {
	forkchoice.Run(t, "mainnet", version.Deneb)
}
