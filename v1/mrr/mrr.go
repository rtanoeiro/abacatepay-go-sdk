package mrr

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type MRRs struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *MRRs {
	return &MRRs{http}
}
