package transparents

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type Transparents struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *Transparents {
	return &Transparents{http}
}
