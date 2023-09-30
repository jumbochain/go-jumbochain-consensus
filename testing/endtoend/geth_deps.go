package endtoend

// This file contains the dependencies required for jumbochain.org/cmd/geth.
// Having these dependencies listed here helps go mod understand that these dependencies are
// necessary for end to end tests since we build go-ethereum binary for this test.
import (
	_ "jumbochain.org/accounts"          // Required for go-ethereum e2e.
	_ "jumbochain.org/accounts/keystore" // Required for go-ethereum e2e.
	_ "jumbochain.org/cmd/utils"         // Required for go-ethereum e2e.
	_ "jumbochain.org/common"            // Required for go-ethereum e2e.
	_ "jumbochain.org/console"           // Required for go-ethereum e2e.
	_ "jumbochain.org/eth"               // Required for go-ethereum e2e.
	_ "jumbochain.org/eth/downloader"    // Required for go-ethereum e2e.
	_ "jumbochain.org/ethclient"         // Required for go-ethereum e2e.
	_ "jumbochain.org/les"               // Required for go-ethereum e2e.
	_ "jumbochain.org/log"               // Required for go-ethereum e2e.
	_ "jumbochain.org/metrics"           // Required for go-ethereum e2e.
	_ "jumbochain.org/node"              // Required for go-ethereum e2e.
)
