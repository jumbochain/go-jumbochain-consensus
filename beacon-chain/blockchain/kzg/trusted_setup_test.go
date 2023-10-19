package kzg

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
)

func TestStart(t *testing.T) {
	require.NoError(t, Start())
	require.NotNil(t, kzgContext)
}
