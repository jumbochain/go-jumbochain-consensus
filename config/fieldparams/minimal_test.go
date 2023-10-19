//go:build minimal

package field_params_test

import (
	"testing"

	fieldparams "github.com/jumbochain/go-jumbochain-consensus/config/fieldparams"
	"github.com/jumbochain/go-jumbochain-consensus/config/params"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
)

func TestFieldParametersValues(t *testing.T) {
	params.SetupTestConfigCleanup(t)
	min := params.MinimalSpecConfig().Copy()
	params.OverrideBeaconConfig(min)
	require.Equal(t, "minimal", fieldparams.Preset)
	testFieldParametersMatchConfig(t)
}
