package beacon_api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/rpc/eth/shared"
	ethpb "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
	"github.com/pkg/errors"
)

func (c *beaconApiValidatorClient) submitValidatorRegistrations(ctx context.Context, registrations []*ethpb.SignedValidatorRegistrationV1) error {
	const endpoint = "/eth/v1/validator/register_validator"

	jsonRegistration := make([]*shared.SignedValidatorRegistration, len(registrations))

	for index, registration := range registrations {
		outMessage, err := shared.SignedValidatorRegistrationFromConsensus(registration)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to encode to SignedValidatorRegistration at index %d", index))
		}
		jsonRegistration[index] = outMessage
	}

	marshalledJsonRegistration, err := json.Marshal(jsonRegistration)
	if err != nil {
		return errors.Wrap(err, "failed to marshal registration")
	}

	if _, err := c.jsonRestHandler.PostRestJson(ctx, endpoint, nil, bytes.NewBuffer(marshalledJsonRegistration), nil); err != nil {
		return errors.Wrapf(err, "failed to send POST data to `%s` REST endpoint", endpoint)
	}

	return nil
}
