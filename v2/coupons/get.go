package coupons

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (c *CouponManager) Get(ctx context.Context, id string) (*GetCouponData, error) {
	var data GetCouponData

	if err := c.http.Get(ctx, v2.BuildGetCouponURL(id), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
