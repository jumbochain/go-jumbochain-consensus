package node

import (
	ethpbservice "github.com/jumbochain/go-jumbochain-consensus/proto/eth/service"
)

var _ ethpbservice.BeaconNodeServer = (*Server)(nil)
