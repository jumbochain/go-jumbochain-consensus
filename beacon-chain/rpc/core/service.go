package core

import (
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/blockchain"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/cache"
	opfeed "github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/feed/operation"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/operations/synccommittee"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/p2p"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/state/stategen"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/sync"
)

type Service struct {
	HeadFetcher        blockchain.HeadFetcher
	GenesisTimeFetcher blockchain.TimeFetcher
	SyncChecker        sync.Checker
	Broadcaster        p2p.Broadcaster
	SyncCommitteePool  synccommittee.Pool
	OperationNotifier  opfeed.Notifier
	AttestationCache   *cache.AttestationCache
	StateGen           stategen.StateManager
	P2P                p2p.Broadcaster
}
