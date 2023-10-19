package types

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/config/params"
	"github.com/jumbochain/go-jumbochain-consensus/encoding/bytesutil"
	"github.com/jumbochain/go-jumbochain-consensus/testing/assert"
)

func TestInitializeDataMaps(t *testing.T) {
	params.SetupTestConfigCleanup(t)
	tests := []struct {
		name   string
		action func()
		exists bool
	}{
		{
			name: "no change",
			action: func() {
			},
			exists: true,
		},
		{
			name: "fork version changes",
			action: func() {
				cfg := params.BeaconConfig().Copy()
				cfg.GenesisForkVersion = []byte{0x01, 0x02, 0x00, 0x00}
				params.OverrideBeaconConfig(cfg)
			},
			exists: false,
		},
		{
			name: "fork version changes with reset",
			action: func() {
				cfg := params.BeaconConfig().Copy()
				cfg.GenesisForkVersion = []byte{0x01, 0x02, 0x00, 0x00}
				params.OverrideBeaconConfig(cfg)
				InitializeDataMaps()
			},
			exists: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.action()
			_, ok := BlockMap[bytesutil.ToBytes4(params.BeaconConfig().GenesisForkVersion)]
			assert.Equal(t, tt.exists, ok)
		})
	}
}
