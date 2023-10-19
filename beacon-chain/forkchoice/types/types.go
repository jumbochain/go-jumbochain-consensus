package types

import (
	fieldparams "github.com/jumbochain/go-jumbochain-consensus/config/fieldparams"
	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/interfaces"
	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/primitives"
	ethpb "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
)

// Checkpoint is an array version of ethpb.Checkpoint. It is used internally in
// forkchoice, while the slice version is used in the interface to legacy code
// in other packages
type Checkpoint struct {
	Epoch primitives.Epoch
	Root  [fieldparams.RootLength]byte
}

// BlockAndCheckpoints to call the InsertOptimisticChain function
type BlockAndCheckpoints struct {
	Block               interfaces.ReadOnlyBeaconBlock
	JustifiedCheckpoint *ethpb.Checkpoint
	FinalizedCheckpoint *ethpb.Checkpoint
}
