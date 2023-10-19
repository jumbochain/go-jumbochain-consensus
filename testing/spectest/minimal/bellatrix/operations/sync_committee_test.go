package operations

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/spectest/shared/bellatrix/operations"
)

func TestMinimal_Bellatrix_Operations_SyncCommittee(t *testing.T) {
	operations.RunProposerSlashingTest(t, "minimal")
}
