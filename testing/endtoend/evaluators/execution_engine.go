package evaluators

import (
	"context"
	"strconv"

	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/primitives"
	mathutil "github.com/jumbochain/go-jumbochain-consensus/math"
	"github.com/jumbochain/go-jumbochain-consensus/proto/eth/service"
	v1 "github.com/jumbochain/go-jumbochain-consensus/proto/eth/v1"
	v2 "github.com/jumbochain/go-jumbochain-consensus/proto/eth/v2"
	"github.com/jumbochain/go-jumbochain-consensus/testing/endtoend/policies"
	"github.com/jumbochain/go-jumbochain-consensus/testing/endtoend/types"
	"github.com/jumbochain/go-jumbochain-consensus/time/slots"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// OptimisticSyncEnabled checks that the node is in an optimistic state.
var OptimisticSyncEnabled = types.Evaluator{
	Name:       "optimistic_sync_at_epoch_%d",
	Policy:     policies.AllEpochs,
	Evaluation: optimisticSyncEnabled,
}

func optimisticSyncEnabled(_ *types.EvaluationContext, conns ...*grpc.ClientConn) error {
	for _, conn := range conns {
		client := service.NewBeaconChainClient(conn)
		head, err := client.GetBlindedBlock(context.Background(), &v1.BlockRequest{BlockId: []byte("head")})
		if err != nil {
			return err
		}
		headSlot := uint64(0)
		switch hb := head.Data.Message.(type) {
		case *v2.SignedBlindedBeaconBlockContainer_Phase0Block:
			headSlot = uint64(hb.Phase0Block.Slot)
		case *v2.SignedBlindedBeaconBlockContainer_AltairBlock:
			headSlot = uint64(hb.AltairBlock.Slot)
		case *v2.SignedBlindedBeaconBlockContainer_BellatrixBlock:
			headSlot = uint64(hb.BellatrixBlock.Slot)
		case *v2.SignedBlindedBeaconBlockContainer_CapellaBlock:
			headSlot = uint64(hb.CapellaBlock.Slot)
		default:
			return errors.New("no valid block type retrieved")
		}
		currEpoch := slots.ToEpoch(primitives.Slot(headSlot))
		startSlot, err := slots.EpochStart(currEpoch)
		if err != nil {
			return err
		}
		isOptimistic := false
		for i := startSlot; i <= primitives.Slot(headSlot); i++ {
			castI, err := mathutil.Int(uint64(i))
			if err != nil {
				return err
			}
			block, err := client.GetBlindedBlock(context.Background(), &v1.BlockRequest{BlockId: []byte(strconv.Itoa(castI))})
			if err != nil {
				// Continue in the event of non-existent blocks.
				continue
			}
			if !block.ExecutionOptimistic {
				return errors.New("expected block to be optimistic, but it is not")
			}
			isOptimistic = true
		}
		if !isOptimistic {
			return errors.New("expected block to be optimistic, but it is not")
		}
	}
	return nil
}
