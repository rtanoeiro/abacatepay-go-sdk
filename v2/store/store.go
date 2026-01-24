package store

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type StoreManager struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *StoreManager {
	return &StoreManager{http}
}
