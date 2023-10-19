package db

import "github.com/jumbochain/go-jumbochain-consensus/beacon-chain/db/kv"

var _ Database = (*kv.Store)(nil)
