package shared

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/testing/assert"
	"github.com/pkg/errors"
)

func TestDecodeError(t *testing.T) {
	e := errors.New("not a number")
	de := NewDecodeError(e, "Z")
	de = NewDecodeError(de, "Y")
	de = NewDecodeError(de, "X")
	assert.Equal(t, "could not decode X.Y.Z: not a number", de.Error())
}
