package payouts

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type PayoutManager struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *PayoutManager {
	return &PayoutManager{http}
}
