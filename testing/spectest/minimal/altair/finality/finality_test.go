package finality

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/altair/finality"
)

func TestMinimal_Altair_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
