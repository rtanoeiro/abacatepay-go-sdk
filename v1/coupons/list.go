package coupons

import (
	"context"

	v1 "github.com/AbacatePay/go-types/v1"
)

func (c *Coupons) List(ctx context.Context) (*ListCouponsData, error) {
	var coupons ListCouponsData

	if err := c.http.Get(ctx, v1.RouteListCoupons, &coupons); err != nil {
		return nil, err
	}

	return &coupons, nil
}
