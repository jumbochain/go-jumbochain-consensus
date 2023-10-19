package validator_client_factory

import (
	"github.com/jumbochain/go-jumbochain-consensus/config/features"
	beaconApi "github.com/jumbochain/go-jumbochain-consensus/validator/client/beacon-api"
	grpcApi "github.com/jumbochain/go-jumbochain-consensus/validator/client/grpc-api"
	"github.com/jumbochain/go-jumbochain-consensus/validator/client/iface"
	validatorHelpers "github.com/jumbochain/go-jumbochain-consensus/validator/helpers"
)

func NewValidatorClient(validatorConn validatorHelpers.NodeConnection) iface.ValidatorClient {
	featureFlags := features.Get()

	if featureFlags.EnableBeaconRESTApi {
		return beaconApi.NewBeaconApiValidatorClient(validatorConn.GetBeaconApiUrl(), validatorConn.GetBeaconApiTimeout())
	} else {
		return grpcApi.NewGrpcValidatorClient(validatorConn.GetGrpcClientConn())
	}
}
