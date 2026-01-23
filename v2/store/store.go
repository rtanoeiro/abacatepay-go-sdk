package store

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type Stores struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *Stores {
	return &Stores{http}
}
