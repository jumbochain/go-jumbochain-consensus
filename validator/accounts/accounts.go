package accounts

import (
	"github.com/jumbochain/go-jumbochain-consensus/validator/keymanager"
)

var (
	errKeymanagerNotSupported = "keymanager kind not supported: %s"
	// ErrCouldNotInitializeKeymanager informs about failed keymanager initialization
	ErrCouldNotInitializeKeymanager = "could not initialize keymanager"
)

// DeleteConfig specifies parameters for the accounts delete command.
type DeleteConfig struct {
	Keymanager       keymanager.IKeymanager
	DeletePublicKeys [][]byte
}
