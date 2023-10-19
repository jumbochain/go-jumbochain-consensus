package logs

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
)

func TestStreamServer_BackfillsMessages(t *testing.T) {
	ss := NewStreamServer()
	msgs := [][]byte{
		[]byte("foo"),
		[]byte("bar"),
		[]byte("buzz"),
	}
	for _, msg := range msgs {
		_, err := ss.Write(msg)
		require.NoError(t, err)
	}

	recentMessages := ss.GetLastFewLogs()
	require.DeepEqual(t, msgs, recentMessages)
}
