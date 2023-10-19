package testing

import (
	"github.com/jumbochain/go-jumbochain-consensus/time/slots"
)

var _ slots.Ticker = (*MockTicker)(nil)
