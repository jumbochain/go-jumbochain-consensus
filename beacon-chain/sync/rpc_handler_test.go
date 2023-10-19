package sync

import (
	"context"
	"testing"
	"time"

	p2ptest "github.com/jumbochain/go-jumbochain-consensus/beacon-chain/p2p/testing"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
	"github.com/jumbochain/go-jumbochain-consensus/testing/util"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type rpcHandlerTest struct {
	t       *testing.T
	topic   protocol.ID
	timeout time.Duration
	err     error
	s       *Service
}

func (rt *rpcHandlerTest) testHandler(nh network.StreamHandler, rh rpcHandler, rhi interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), rt.timeout)
	defer func() {
		cancel()
	}()

	w := util.NewWaiter()
	server := p2ptest.NewTestP2P(rt.t)

	client, ok := rt.s.cfg.p2p.(*p2ptest.TestP2P)
	require.Equal(rt.t, true, ok)

	client.Connect(server)
	defer func() {
		require.NoError(rt.t, client.Disconnect(server.PeerID()))
	}()
	require.Equal(rt.t, 1, len(client.BHost.Network().Peers()), "Expected peers to be connected")
	h := func(stream network.Stream) {
		defer w.Done()
		nh(stream)
	}
	server.BHost.SetStreamHandler(protocol.ID(rt.topic), h)
	stream, err := client.BHost.NewStream(ctx, server.BHost.ID(), protocol.ID(rt.topic))
	require.NoError(rt.t, err)

	err = rh(ctx, rhi, stream)
	if rt.err == nil {
		require.NoError(rt.t, err)
	} else {
		require.ErrorIs(rt.t, err, rt.err)
	}

	w.RequireDoneBeforeCancel(ctx, rt.t)
}
