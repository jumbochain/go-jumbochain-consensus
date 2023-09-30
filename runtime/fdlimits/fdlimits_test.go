package fdlimits_test

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v4/runtime/fdlimits"
	"github.com/prysmaticlabs/prysm/v4/testing/assert"
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
