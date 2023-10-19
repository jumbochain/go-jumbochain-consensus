package rpc

import (
	pb "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1/validator-client"
)

var _ pb.AuthServer = (*Server)(nil)
