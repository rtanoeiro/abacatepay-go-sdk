package subscriptions

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type SubscriptionManager struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *SubscriptionManager {
	return &SubscriptionManager{http}
}
