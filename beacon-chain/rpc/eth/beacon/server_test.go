package beacon

import ethpbservice "github.com/jumbochain/go-jumbochain-consensus/proto/eth/service"

var _ ethpbservice.BeaconChainServer = (*Server)(nil)
