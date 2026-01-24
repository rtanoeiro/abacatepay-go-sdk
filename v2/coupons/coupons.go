package coupons

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type Coupons struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *Coupons {
	return &Coupons{http}
}
