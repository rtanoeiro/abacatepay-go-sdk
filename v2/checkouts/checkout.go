package checkouts

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type CheckoutManager struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *CheckoutManager {
	return &CheckoutManager{http}
}
