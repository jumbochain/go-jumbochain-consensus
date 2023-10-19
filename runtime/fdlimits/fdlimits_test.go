package fdlimits_test

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/runtime/fdlimits"
	"github.com/jumbochain/go-jumbochain-consensus/testing/assert"
	gethLimit "jumbochain.org/common/fdlimit"
)

func TestSetMaxFdLimits(t *testing.T) {
	assert.NoError(t, fdlimits.SetMaxFdLimits())

	curr, err := gethLimit.Current()
	assert.NoError(t, err)

	max, err := gethLimit.Maximum()
	assert.NoError(t, err)

	assert.Equal(t, max, curr, "current and maximum file descriptor limits do not match up.")

}
