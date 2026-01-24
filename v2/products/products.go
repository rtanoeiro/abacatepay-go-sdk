package products

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type ProductManager struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *ProductManager {
	return &ProductManager{http}
}
