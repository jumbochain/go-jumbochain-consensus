package validator

import (
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/blockchain"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/builder"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/cache"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/feed/operation"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/db"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/operations/attestations"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/operations/synccommittee"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/p2p"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/rpc/core"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/rpc/lookup"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/sync"
	eth "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
)

// Server defines a server implementation of the gRPC Validator service,
// providing RPC endpoints intended for validator clients.
type Server struct {
	HeadFetcher            blockchain.HeadFetcher
	TimeFetcher            blockchain.TimeFetcher
	SyncChecker            sync.Checker
	AttestationsPool       attestations.Pool
	PeerManager            p2p.PeerManager
	Broadcaster            p2p.Broadcaster
	Stater                 lookup.Stater
	OptimisticModeFetcher  blockchain.OptimisticModeFetcher
	SyncCommitteePool      synccommittee.Pool
	V1Alpha1Server         eth.BeaconNodeValidatorServer
	ProposerSlotIndexCache *cache.ProposerPayloadIDsCache
	ChainInfoFetcher       blockchain.ChainInfoFetcher
	BeaconDB               db.HeadAccessDatabase
	BlockBuilder           builder.BlockBuilder
	OperationNotifier      operation.Notifier
	CoreService            *core.Service
}
