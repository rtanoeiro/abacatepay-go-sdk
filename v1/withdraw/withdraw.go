package withdraw

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type Withdraws struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *Withdraws {
	return &Withdraws{http}
}
