package testutil

import (
	"context"
	"strconv"

	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/interfaces"
	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/primitives"
	"github.com/jumbochain/go-jumbochain-consensus/encoding/bytesutil"
)

// MockBlocker is a fake implementation of lookup.Blocker.
type MockBlocker struct {
	BlockToReturn interfaces.ReadOnlySignedBeaconBlock
	ErrorToReturn error
	SlotBlockMap  map[primitives.Slot]interfaces.ReadOnlySignedBeaconBlock
	RootBlockMap  map[[32]byte]interfaces.ReadOnlySignedBeaconBlock
}

// Block --
func (m *MockBlocker) Block(_ context.Context, b []byte) (interfaces.ReadOnlySignedBeaconBlock, error) {
	if m.ErrorToReturn != nil {
		return nil, m.ErrorToReturn
	}
	if m.BlockToReturn != nil {
		return m.BlockToReturn, nil
	}
	slotNumber, parseErr := strconv.ParseUint(string(b), 10, 64)
	if parseErr != nil {
		//nolint:nilerr
		return m.RootBlockMap[bytesutil.ToBytes32(b)], nil
	}
	return m.SlotBlockMap[primitives.Slot(slotNumber)], nil
}
