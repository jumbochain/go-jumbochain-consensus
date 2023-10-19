package sanity

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/deneb/sanity"
)

func TestMinimal_Deneb_Sanity_Slots(t *testing.T) {
	sanity.RunSlotProcessingTests(t, "minimal")
}
