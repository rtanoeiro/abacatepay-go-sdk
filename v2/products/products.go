package products

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type Products struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *Products {
	return &Products{http}
}
