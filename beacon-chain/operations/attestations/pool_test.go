package attestations

import (
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/operations/attestations/kv"
)

var _ Pool = (*kv.AttCaches)(nil)
