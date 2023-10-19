//go:build !noMainnetGenesis
// +build !noMainnetGenesis

package genesis

import (
	_ "embed"

	"github.com/jumbochain/go-jumbochain-consensus/config/params"
)

var (
	//go:embed mainnet.ssz.snappy
	mainnetRawSSZCompressed []byte // 1.8Mb
)

func init() {
	embeddedStates[params.MainnetName] = &mainnetRawSSZCompressed
}
