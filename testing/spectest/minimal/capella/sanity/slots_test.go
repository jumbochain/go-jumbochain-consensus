package sanity

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/capella/sanity"
)

func TestMinimal_Capella_Sanity_Slots(t *testing.T) {
	sanity.RunSlotProcessingTests(t, "minimal")
}
