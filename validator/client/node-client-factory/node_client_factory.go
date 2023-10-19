package validator_client_factory

import (
	"github.com/jumbochain/go-jumbochain-consensus/config/features"
	beaconApi "github.com/jumbochain/go-jumbochain-consensus/validator/client/beacon-api"
	grpcApi "github.com/jumbochain/go-jumbochain-consensus/validator/client/grpc-api"
	"github.com/jumbochain/go-jumbochain-consensus/validator/client/iface"
	validatorHelpers "github.com/jumbochain/go-jumbochain-consensus/validator/helpers"
)

func NewNodeClient(validatorConn validatorHelpers.NodeConnection) iface.NodeClient {
	grpcClient := grpcApi.NewNodeClient(validatorConn.GetGrpcClientConn())
	featureFlags := features.Get()

	if featureFlags.EnableBeaconRESTApi {
		return beaconApi.NewNodeClientWithFallback(validatorConn.GetBeaconApiUrl(), validatorConn.GetBeaconApiTimeout(), grpcClient)
	} else {
		return grpcClient
	}
}
