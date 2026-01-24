package transparents

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type TransparentManager struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *TransparentManager {
	return &TransparentManager{http}
}
