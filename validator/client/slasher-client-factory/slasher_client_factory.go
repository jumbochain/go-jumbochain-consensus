package validator_client_factory

import (
	"github.com/jumbochain/go-jumbochain-consensus/config/features"
	beaconApi "github.com/jumbochain/go-jumbochain-consensus/validator/client/beacon-api"
	grpcApi "github.com/jumbochain/go-jumbochain-consensus/validator/client/grpc-api"
	"github.com/jumbochain/go-jumbochain-consensus/validator/client/iface"
	validatorHelpers "github.com/jumbochain/go-jumbochain-consensus/validator/helpers"
)

func NewSlasherClient(validatorConn validatorHelpers.NodeConnection) iface.SlasherClient {
	grpcClient := grpcApi.NewSlasherClient(validatorConn.GetGrpcClientConn())
	featureFlags := features.Get()

	if featureFlags.EnableBeaconRESTApi {
		return beaconApi.NewSlasherClientWithFallback(validatorConn.GetBeaconApiUrl(), validatorConn.GetBeaconApiTimeout(), grpcClient)
	} else {
		return grpcClient
	}
}
