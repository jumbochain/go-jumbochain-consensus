package state

import "github.com/jumbochain/go-jumbochain-consensus/async/event"

// Notifier interface defines the methods of the service that provides state updates to consumers.
type Notifier interface {
	StateFeed() *event.Feed
}
