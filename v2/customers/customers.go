package customers

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type CustomerManager struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *CustomerManager {
	return &CustomerManager{http}
}
