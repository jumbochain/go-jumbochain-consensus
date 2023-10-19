package rewards

import (
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/blockchain"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/rpc/lookup"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/state/stategen"
)

type Server struct {
	Blocker               lookup.Blocker
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
	FinalizationFetcher   blockchain.FinalizationFetcher
	ReplayerBuilder       stategen.ReplayerBuilder
	TimeFetcher           blockchain.TimeFetcher
	Stater                lookup.Stater
	HeadFetcher           blockchain.HeadFetcher
}
