package coupons

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (c *CouponManager) Create(ctx context.Context, body CreateCouponBody) (*CreateCouponData, error) {
	var data CreateCouponData

	if err := c.http.Post(ctx, v2.RouteCreateCoupon, body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
