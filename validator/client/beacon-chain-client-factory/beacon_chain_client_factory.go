package validator_client_factory

import (
	"github.com/jumbochain/go-jumbochain-consensus/config/features"
	beaconApi "github.com/jumbochain/go-jumbochain-consensus/validator/client/beacon-api"
	grpcApi "github.com/jumbochain/go-jumbochain-consensus/validator/client/grpc-api"
	"github.com/jumbochain/go-jumbochain-consensus/validator/client/iface"
	validatorHelpers "github.com/jumbochain/go-jumbochain-consensus/validator/helpers"
)

func NewBeaconChainClient(validatorConn validatorHelpers.NodeConnection) iface.BeaconChainClient {
	grpcClient := grpcApi.NewGrpcBeaconChainClient(validatorConn.GetGrpcClientConn())
	featureFlags := features.Get()

	if featureFlags.EnableBeaconRESTApi {
		return beaconApi.NewBeaconApiBeaconChainClientWithFallback(validatorConn.GetBeaconApiUrl(), validatorConn.GetBeaconApiTimeout(), grpcClient)
	} else {
		return grpcClient
	}
}
