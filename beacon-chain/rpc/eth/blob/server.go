package blob

import (
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/blockchain"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/db"
)

type Server struct {
	ChainInfoFetcher blockchain.ChainInfoFetcher
	BeaconDB         db.ReadOnlyDatabase
}
