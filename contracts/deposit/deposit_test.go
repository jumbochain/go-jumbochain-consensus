package deposit_test

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/core/signing"
	"github.com/jumbochain/go-jumbochain-consensus/config/params"
	"github.com/jumbochain/go-jumbochain-consensus/contracts/deposit"
	"github.com/jumbochain/go-jumbochain-consensus/crypto/bls"
	ethpb "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
	"github.com/jumbochain/go-jumbochain-consensus/testing/assert"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
	"github.com/jumbochain/go-jumbochain-consensus/testing/util"
)

func TestDepositInput_GeneratesPb(t *testing.T) {
	k1, err := bls.RandKey()
	require.NoError(t, err)
	k2, err := bls.RandKey()
	require.NoError(t, err)

	result, _, err := deposit.DepositInput(k1, k2, 0)
	require.NoError(t, err)
	assert.DeepEqual(t, k1.PublicKey().Marshal(), result.PublicKey)

	sig, err := bls.SignatureFromBytes(result.Signature)
	require.NoError(t, err)
	testData := &ethpb.DepositMessage{
		PublicKey:             result.PublicKey,
		WithdrawalCredentials: result.WithdrawalCredentials,
		Amount:                result.Amount,
	}
	sr, err := testData.HashTreeRoot()
	require.NoError(t, err)
	domain, err := signing.ComputeDomain(
		params.BeaconConfig().DomainDeposit,
		nil, /*forkVersion*/
		nil, /*genesisValidatorsRoot*/
	)
	require.NoError(t, err)
	root, err := (&ethpb.SigningData{ObjectRoot: sr[:], Domain: domain}).HashTreeRoot()
	require.NoError(t, err)
	assert.Equal(t, true, sig.Verify(k1.PublicKey(), root[:]))
}

func TestVerifyDepositSignature_ValidSig(t *testing.T) {
	deposits, _, err := util.DeterministicDepositsAndKeys(1)
	require.NoError(t, err)
	dep := deposits[0]
	domain, err := signing.ComputeDomain(
		params.BeaconConfig().DomainDeposit,
		params.BeaconConfig().GenesisForkVersion,
		params.BeaconConfig().ZeroHash[:],
	)
	require.NoError(t, err)
	err = deposit.VerifyDepositSignature(dep.Data, domain)
	require.NoError(t, err)
}

func TestVerifyDepositSignature_InvalidSig(t *testing.T) {
	deposits, _, err := util.DeterministicDepositsAndKeys(1)
	require.NoError(t, err)
	dep := deposits[0]
	domain, err := signing.ComputeDomain(
		params.BeaconConfig().DomainDeposit,
		params.BeaconConfig().GenesisForkVersion,
		params.BeaconConfig().ZeroHash[:],
	)
	require.NoError(t, err)
	dep.Data.Signature = dep.Data.Signature[1:]
	err = deposit.VerifyDepositSignature(dep.Data, domain)
	if err == nil {
		t.Fatal("Deposit Verification succeeds with a invalid signature")
	}
}
