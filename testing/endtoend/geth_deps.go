package endtoend

// This file contains the dependencies required for github.com/jumbochain/jumbochain-parlia-go/cmd/geth.
// Having these dependencies listed here helps go mod understand that these dependencies are
// necessary for end to end tests since we build go-ethereum binary for this test.
import (
	_ "github.com/jumbochain/jumbochain-parlia-go/accounts"          // Required for go-ethereum e2e.
	_ "github.com/jumbochain/jumbochain-parlia-go/accounts/keystore" // Required for go-ethereum e2e.
	_ "github.com/jumbochain/jumbochain-parlia-go/cmd/utils"         // Required for go-ethereum e2e.
	_ "github.com/jumbochain/jumbochain-parlia-go/common"            // Required for go-ethereum e2e.
	_ "github.com/jumbochain/jumbochain-parlia-go/console"           // Required for go-ethereum e2e.
	_ "github.com/jumbochain/jumbochain-parlia-go/eth"               // Required for go-ethereum e2e.
	_ "github.com/jumbochain/jumbochain-parlia-go/eth/downloader"    // Required for go-ethereum e2e.
	_ "github.com/jumbochain/jumbochain-parlia-go/ethclient"         // Required for go-ethereum e2e.
	_ "github.com/jumbochain/jumbochain-parlia-go/les"               // Required for go-ethereum e2e.
	_ "github.com/jumbochain/jumbochain-parlia-go/log"               // Required for go-ethereum e2e.
	_ "github.com/jumbochain/jumbochain-parlia-go/metrics"           // Required for go-ethereum e2e.
	_ "github.com/jumbochain/jumbochain-parlia-go/node"              // Required for go-ethereum e2e.
)
