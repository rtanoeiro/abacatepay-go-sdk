package payouts

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type Payouts struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *Payouts {
	return &Payouts{http}
}
