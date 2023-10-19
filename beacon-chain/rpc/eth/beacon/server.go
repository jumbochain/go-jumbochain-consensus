// Package beacon defines a gRPC beacon service implementation,
// following the official API standards https://ethereum.github.io/beacon-apis/#/.
// This package includes the beacon and config endpoints.
package beacon

import (
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/blockchain"
	blockfeed "github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/feed/block"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/feed/operation"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/db"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/execution"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/operations/attestations"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/operations/blstoexec"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/operations/slashings"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/operations/voluntaryexits"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/p2p"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/rpc/core"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/rpc/lookup"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/state/stategen"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/sync"
	eth "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
)

// Server defines a server implementation of the gRPC Beacon Chain service,
// providing RPC endpoints to access data relevant to the Ethereum Beacon Chain.
type Server struct {
	BeaconDB                      db.ReadOnlyDatabase
	ChainInfoFetcher              blockchain.ChainInfoFetcher
	GenesisTimeFetcher            blockchain.TimeFetcher
	BlockReceiver                 blockchain.BlockReceiver
	BlockNotifier                 blockfeed.Notifier
	OperationNotifier             operation.Notifier
	Broadcaster                   p2p.Broadcaster
	AttestationsPool              attestations.Pool
	SlashingsPool                 slashings.PoolManager
	VoluntaryExitsPool            voluntaryexits.PoolManager
	StateGenService               stategen.StateManager
	Stater                        lookup.Stater
	Blocker                       lookup.Blocker
	HeadFetcher                   blockchain.HeadFetcher
	TimeFetcher                   blockchain.TimeFetcher
	OptimisticModeFetcher         blockchain.OptimisticModeFetcher
	V1Alpha1ValidatorServer       eth.BeaconNodeValidatorServer
	SyncChecker                   sync.Checker
	CanonicalHistory              *stategen.CanonicalHistory
	ExecutionPayloadReconstructor execution.ExecutionPayloadReconstructor
	FinalizationFetcher           blockchain.FinalizationFetcher
	BLSChangesPool                blstoexec.PoolManager
	ForkchoiceFetcher             blockchain.ForkchoiceFetcher
	CoreService                   *core.Service
}
