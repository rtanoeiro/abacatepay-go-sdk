package coupons

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (c *CouponManager) List(ctx context.Context, page, limit *int) (*ListCouponsData, error) {
	var data ListCouponsData

	if err := c.http.Get(ctx, v2.BuildListCouponsURL(page, limit), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
