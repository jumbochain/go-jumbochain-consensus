package endtoend

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/runtime/version"
	"github.com/jumbochain/go-jumbochain-consensus/testing/endtoend/types"
)

func TestEndToEnd_MinimalConfig_WithBuilder(t *testing.T) {
	r := e2eMinimal(t, version.Phase0, types.WithCheckpointSync(), types.WithBuilder())
	r.run()
}
