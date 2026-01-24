package customers

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type Customers struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *Customers {
	return &Customers{http}
}
