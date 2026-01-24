package checkouts

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type Checkouts struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *Checkouts {
	return &Checkouts{http}
}
