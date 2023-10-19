package sync

import (
	"context"
	"testing"
	"time"

	mockChain "github.com/jumbochain/go-jumbochain-consensus/beacon-chain/blockchain/testing"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/signing"
	testingdb "github.com/jumbochain/go-jumbochain-consensus/beacon-chain/db/testing"
	doublylinkedtree "github.com/jumbochain/go-jumbochain-consensus/beacon-chain/forkchoice/doubly-linked-tree"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/operations/blstoexec"
	mockp2p "github.com/jumbochain/go-jumbochain-consensus/beacon-chain/p2p/testing"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/state/stategen"
	mockSync "github.com/jumbochain/go-jumbochain-consensus/beacon-chain/sync/initial-sync/testing"
	"github.com/jumbochain/go-jumbochain-consensus/config/params"
	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/primitives"
	"github.com/jumbochain/go-jumbochain-consensus/encoding/bytesutil"
	ethpb "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
	"github.com/jumbochain/go-jumbochain-consensus/testing/assert"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
	"github.com/jumbochain/go-jumbochain-consensus/testing/util"
	"github.com/jumbochain/go-jumbochain-consensus/time/slots"
	logTest "github.com/sirupsen/logrus/hooks/test"
)

func TestBroadcastBLSChanges(t *testing.T) {
	params.SetupTestConfigCleanup(t)
	c := params.BeaconConfig()
	c.CapellaForkEpoch = c.BellatrixForkEpoch.Add(2)
	params.OverrideBeaconConfig(c)
	chainService := &mockChain.ChainService{
		Genesis:        time.Now(),
		ValidatorsRoot: [32]byte{'A'},
	}
	s := NewService(context.Background(),
		WithP2P(mockp2p.NewTestP2P(t)),
		WithInitialSync(&mockSync.Sync{IsSyncing: false}),
		WithChainService(chainService),
		WithOperationNotifier(chainService.OperationNotifier()),
		WithBlsToExecPool(blstoexec.NewPool()),
	)
	var emptySig [96]byte
	s.cfg.blsToExecPool.InsertBLSToExecChange(&ethpb.SignedBLSToExecutionChange{
		Message: &ethpb.BLSToExecutionChange{
			ValidatorIndex:     10,
			FromBlsPubkey:      make([]byte, 48),
			ToExecutionAddress: make([]byte, 20),
		},
		Signature: emptySig[:],
	})

	capellaStart, err := slots.EpochStart(params.BeaconConfig().CapellaForkEpoch)
	require.NoError(t, err)
	s.broadcastBLSChanges(capellaStart + 1)
}

func TestRateBLSChanges(t *testing.T) {
	logHook := logTest.NewGlobal()
	params.SetupTestConfigCleanup(t)
	c := params.BeaconConfig()
	c.CapellaForkEpoch = c.BellatrixForkEpoch.Add(2)
	params.OverrideBeaconConfig(c)
	chainService := &mockChain.ChainService{
		Genesis:        time.Now(),
		ValidatorsRoot: [32]byte{'A'},
	}
	p1 := mockp2p.NewTestP2P(t)
	s := NewService(context.Background(),
		WithP2P(p1),
		WithInitialSync(&mockSync.Sync{IsSyncing: false}),
		WithChainService(chainService),
		WithOperationNotifier(chainService.OperationNotifier()),
		WithBlsToExecPool(blstoexec.NewPool()),
	)
	beaconDB := testingdb.SetupDB(t)
	s.cfg.stateGen = stategen.New(beaconDB, doublylinkedtree.New())
	s.cfg.beaconDB = beaconDB
	s.initCaches()
	st, keys := util.DeterministicGenesisStateCapella(t, 256)
	s.cfg.chain = &mockChain.ChainService{
		ValidatorsRoot: [32]byte{'A'},
		Genesis:        time.Now().Add(-time.Second * time.Duration(params.BeaconConfig().SecondsPerSlot) * time.Duration(10)),
		State:          st,
	}

	for i := 0; i < 200; i++ {
		message := &ethpb.BLSToExecutionChange{
			ValidatorIndex:     primitives.ValidatorIndex(i),
			FromBlsPubkey:      keys[i+1].PublicKey().Marshal(),
			ToExecutionAddress: bytesutil.PadTo([]byte("address"), 20),
		}
		epoch := params.BeaconConfig().CapellaForkEpoch + 1
		domain, err := signing.Domain(st.Fork(), epoch, params.BeaconConfig().DomainBLSToExecutionChange, st.GenesisValidatorsRoot())
		assert.NoError(t, err)
		htr, err := signing.SigningData(message.HashTreeRoot, domain)
		assert.NoError(t, err)
		signed := &ethpb.SignedBLSToExecutionChange{
			Message:   message,
			Signature: keys[i+1].Sign(htr[:]).Marshal(),
		}

		s.cfg.blsToExecPool.InsertBLSToExecChange(signed)
	}

	require.Equal(t, false, p1.BroadcastCalled)
	slot, err := slots.EpochStart(params.BeaconConfig().CapellaForkEpoch)
	require.NoError(t, err)
	s.broadcastBLSChanges(slot)
	time.Sleep(100 * time.Millisecond) // Need a sleep for the go routine to be ready
	require.Equal(t, true, p1.BroadcastCalled)
	require.LogsDoNotContain(t, logHook, "could not")

	p1.BroadcastCalled = false
	time.Sleep(500 * time.Millisecond) // Need a sleep for the second batch to be broadcast
	require.Equal(t, true, p1.BroadcastCalled)
	require.LogsDoNotContain(t, logHook, "could not")
}

func TestBroadcastBLSBatch_changes_slice(t *testing.T) {
	message := &ethpb.BLSToExecutionChange{
		FromBlsPubkey:      make([]byte, 48),
		ToExecutionAddress: make([]byte, 20),
	}
	signed := &ethpb.SignedBLSToExecutionChange{
		Message:   message,
		Signature: make([]byte, 96),
	}
	changes := make([]*ethpb.SignedBLSToExecutionChange, 200)
	for i := 0; i < len(changes); i++ {
		changes[i] = signed
	}
	p1 := mockp2p.NewTestP2P(t)
	chainService := &mockChain.ChainService{
		Genesis:        time.Now(),
		ValidatorsRoot: [32]byte{'A'},
	}
	s := NewService(context.Background(),
		WithP2P(p1),
		WithInitialSync(&mockSync.Sync{IsSyncing: false}),
		WithChainService(chainService),
		WithOperationNotifier(chainService.OperationNotifier()),
		WithBlsToExecPool(blstoexec.NewPool()),
	)
	beaconDB := testingdb.SetupDB(t)
	s.cfg.stateGen = stategen.New(beaconDB, doublylinkedtree.New())
	s.cfg.beaconDB = beaconDB
	s.initCaches()
	st, _ := util.DeterministicGenesisStateCapella(t, 32)
	s.cfg.chain = &mockChain.ChainService{
		ValidatorsRoot: [32]byte{'A'},
		Genesis:        time.Now().Add(-time.Second * time.Duration(params.BeaconConfig().SecondsPerSlot) * time.Duration(10)),
		State:          st,
	}

	s.broadcastBLSBatch(s.ctx, &changes)
	require.Equal(t, 200-128, len(changes))
}
