package coupons

import (
	"context"

	v1 "github.com/almeidazs/go-abacate-types/v1"
)

func (c *Coupons) Create(ctx context.Context, body CreateCouponBody) (*CreateCouponData, error) {
	var coupon CreateCouponData

	if err := c.http.Post(ctx, v1.RouteCreateCoupon, body, &coupon); err != nil {
		return nil, err
	}

	return &coupon, nil
}
