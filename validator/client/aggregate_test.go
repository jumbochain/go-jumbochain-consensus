package client

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	fieldparams "github.com/jumbochain/go-jumbochain-consensus/config/fieldparams"
	"github.com/jumbochain/go-jumbochain-consensus/config/params"
	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/primitives"
	"github.com/jumbochain/go-jumbochain-consensus/crypto/bls"
	ethpb "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
	"github.com/jumbochain/go-jumbochain-consensus/testing/assert"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
	"github.com/jumbochain/go-jumbochain-consensus/testing/util"
	"github.com/jumbochain/go-jumbochain-consensus/time"
	"github.com/jumbochain/go-jumbochain-consensus/time/slots"
	"github.com/prysmaticlabs/go-bitfield"
	logTest "github.com/sirupsen/logrus/hooks/test"
)

func TestSubmitAggregateAndProof_GetDutiesRequestFailure(t *testing.T) {
	hook := logTest.NewGlobal()
	validator, _, validatorKey, finish := setup(t)
	validator.duties = &ethpb.DutiesResponse{Duties: []*ethpb.DutiesResponse_Duty{}}
	defer finish()

	var pubKey [fieldparams.BLSPubkeyLength]byte
	copy(pubKey[:], validatorKey.PublicKey().Marshal())
	validator.SubmitAggregateAndProof(context.Background(), 0, pubKey)

	require.LogsContain(t, hook, "Could not fetch validator assignment")
}

func TestSubmitAggregateAndProof_SignFails(t *testing.T) {
	validator, m, validatorKey, finish := setup(t)
	defer finish()
	var pubKey [fieldparams.BLSPubkeyLength]byte
	copy(pubKey[:], validatorKey.PublicKey().Marshal())
	validator.duties = &ethpb.DutiesResponse{
		Duties: []*ethpb.DutiesResponse_Duty{
			{
				PublicKey: validatorKey.PublicKey().Marshal(),
			},
		},
	}

	m.validatorClient.EXPECT().DomainData(
		gomock.Any(), // ctx
		gomock.Any(), // epoch
	).Return(&ethpb.DomainResponse{SignatureDomain: make([]byte, 32)}, nil /*err*/)

	m.validatorClient.EXPECT().SubmitAggregateSelectionProof(
		gomock.Any(), // ctx
		gomock.AssignableToTypeOf(&ethpb.AggregateSelectionRequest{}),
	).Return(&ethpb.AggregateSelectionResponse{
		AggregateAndProof: &ethpb.AggregateAttestationAndProof{
			AggregatorIndex: 0,
			Aggregate: util.HydrateAttestation(&ethpb.Attestation{
				AggregationBits: make([]byte, 1),
			}),
			SelectionProof: make([]byte, 96),
		},
	}, nil)

	m.validatorClient.EXPECT().DomainData(
		gomock.Any(), // ctx
		gomock.Any(), // epoch
	).Return(&ethpb.DomainResponse{SignatureDomain: nil}, errors.New("bad domain root"))

	validator.SubmitAggregateAndProof(context.Background(), 0, pubKey)
}

func TestSubmitAggregateAndProof_Ok(t *testing.T) {
	validator, m, validatorKey, finish := setup(t)
	defer finish()
	var pubKey [fieldparams.BLSPubkeyLength]byte
	copy(pubKey[:], validatorKey.PublicKey().Marshal())
	validator.duties = &ethpb.DutiesResponse{
		Duties: []*ethpb.DutiesResponse_Duty{
			{
				PublicKey: validatorKey.PublicKey().Marshal(),
			},
		},
	}

	m.validatorClient.EXPECT().DomainData(
		gomock.Any(), // ctx
		gomock.Any(), // epoch
	).Return(&ethpb.DomainResponse{SignatureDomain: make([]byte, 32)}, nil /*err*/)

	m.validatorClient.EXPECT().SubmitAggregateSelectionProof(
		gomock.Any(), // ctx
		gomock.AssignableToTypeOf(&ethpb.AggregateSelectionRequest{}),
	).Return(&ethpb.AggregateSelectionResponse{
		AggregateAndProof: &ethpb.AggregateAttestationAndProof{
			AggregatorIndex: 0,
			Aggregate: util.HydrateAttestation(&ethpb.Attestation{
				AggregationBits: make([]byte, 1),
			}),
			SelectionProof: make([]byte, 96),
		},
	}, nil)

	m.validatorClient.EXPECT().DomainData(
		gomock.Any(), // ctx
		gomock.Any(), // epoch
	).Return(&ethpb.DomainResponse{SignatureDomain: make([]byte, 32)}, nil /*err*/)

	m.validatorClient.EXPECT().SubmitSignedAggregateSelectionProof(
		gomock.Any(), // ctx
		gomock.AssignableToTypeOf(&ethpb.SignedAggregateSubmitRequest{}),
	).Return(&ethpb.SignedAggregateSubmitResponse{AttestationDataRoot: make([]byte, 32)}, nil)

	validator.SubmitAggregateAndProof(context.Background(), 0, pubKey)
}

func TestWaitForSlotTwoThird_WaitCorrectly(t *testing.T) {
	validator, _, _, finish := setup(t)
	defer finish()
	currentTime := time.Now()
	numOfSlots := primitives.Slot(4)
	validator.genesisTime = uint64(currentTime.Unix()) - uint64(numOfSlots.Mul(params.BeaconConfig().SecondsPerSlot))
	oneThird := slots.DivideSlotBy(3 /* one third of slot duration */)
	timeToSleep := oneThird + oneThird

	twoThirdTime := currentTime.Add(timeToSleep)
	validator.waitToSlotTwoThirds(context.Background(), numOfSlots)
	currentTime = time.Now()
	assert.Equal(t, twoThirdTime.Unix(), currentTime.Unix())
}

func TestWaitForSlotTwoThird_DoneContext_ReturnsImmediately(t *testing.T) {
	validator, _, _, finish := setup(t)
	defer finish()
	currentTime := time.Now()
	numOfSlots := primitives.Slot(4)
	validator.genesisTime = uint64(currentTime.Unix()) - uint64(numOfSlots.Mul(params.BeaconConfig().SecondsPerSlot))

	expectedTime := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	validator.waitToSlotTwoThirds(ctx, numOfSlots)
	currentTime = time.Now()
	assert.Equal(t, expectedTime.Unix(), currentTime.Unix())
}

func TestAggregateAndProofSignature_CanSignValidSignature(t *testing.T) {
	validator, m, validatorKey, finish := setup(t)
	defer finish()

	var pubKey [fieldparams.BLSPubkeyLength]byte
	copy(pubKey[:], validatorKey.PublicKey().Marshal())
	m.validatorClient.EXPECT().DomainData(
		gomock.Any(), // ctx
		&ethpb.DomainRequest{Epoch: 0, Domain: params.BeaconConfig().DomainAggregateAndProof[:]},
	).Return(&ethpb.DomainResponse{SignatureDomain: make([]byte, 32)}, nil /*err*/)

	agg := &ethpb.AggregateAttestationAndProof{
		AggregatorIndex: 0,
		Aggregate: util.HydrateAttestation(&ethpb.Attestation{
			AggregationBits: bitfield.NewBitlist(1),
		}),
		SelectionProof: make([]byte, 96),
	}
	sig, err := validator.aggregateAndProofSig(context.Background(), pubKey, agg, 0 /* slot */)
	require.NoError(t, err)
	_, err = bls.SignatureFromBytes(sig)
	require.NoError(t, err)
}
