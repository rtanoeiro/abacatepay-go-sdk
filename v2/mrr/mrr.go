package mrr

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type MRRManager struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *MRRManager {
	return &MRRManager{http}
}
