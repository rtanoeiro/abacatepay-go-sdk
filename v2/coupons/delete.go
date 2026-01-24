package coupons

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (c *Coupons) Delete(ctx context.Context, id string) (*DeleteCouponData, error) {
	var data DeleteCouponData

	if err := c.http.Delete(ctx, v2.RouteDeleteCoupon, v2.RESTDeleteCouponBody{ID: id}, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
