package forkchoice

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/runtime/version"
	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/common/forkchoice"
)

func TestMainnet_Capella_Forkchoice(t *testing.T) {
	forkchoice.Run(t, "mainnet", version.Capella)
}
