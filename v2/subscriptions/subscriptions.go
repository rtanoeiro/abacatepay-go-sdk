package subscriptions

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type Subscriptions struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *Subscriptions {
	return &Subscriptions{http}
}
