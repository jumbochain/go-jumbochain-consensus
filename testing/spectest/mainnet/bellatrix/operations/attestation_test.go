package operations

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/bellatrix/operations"
)

func TestMainnet_Bellatrix_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}
