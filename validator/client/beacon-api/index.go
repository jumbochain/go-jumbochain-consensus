package beacon_api

import (
	"context"
	"strconv"

	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/primitives"
	ethpb "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
	"github.com/jumbochain/jumbochain-parlia-go/common/hexutil"
	"github.com/pkg/errors"
)

func (c beaconApiValidatorClient) validatorIndex(ctx context.Context, in *ethpb.ValidatorIndexRequest) (*ethpb.ValidatorIndexResponse, error) {
	stringPubKey := hexutil.Encode(in.PublicKey)

	stateValidator, err := c.stateValidatorsProvider.GetStateValidators(ctx, []string{stringPubKey}, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get state validator")
	}

	if len(stateValidator.Data) == 0 {
		return nil, errors.Errorf("could not find validator index for public key `%s`", stringPubKey)
	}

	stringValidatorIndex := stateValidator.Data[0].Index

	index, err := strconv.ParseUint(stringValidatorIndex, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse validator index")
	}

	return &ethpb.ValidatorIndexResponse{Index: primitives.ValidatorIndex(index)}, nil
}
