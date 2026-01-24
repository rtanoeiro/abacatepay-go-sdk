package coupons

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type CouponManager struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *CouponManager {
	return &CouponManager{http}
}
