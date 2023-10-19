package finality

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/capella/finality"
)

func TestMinimal_Capella_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
