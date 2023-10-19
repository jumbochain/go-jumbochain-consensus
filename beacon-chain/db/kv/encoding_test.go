package kv

import (
	"context"
	"testing"

	testpb "github.com/jumbochain/go-jumbochain-consensus/proto/testing"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
)

func Test_encode_handlesNilFromFunction(t *testing.T) {
	foo := func() *testpb.Puzzle {
		return nil
	}
	_, err := encode(context.Background(), foo())
	require.ErrorContains(t, "cannot encode nil message", err)
}
