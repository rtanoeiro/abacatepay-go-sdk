package coupons

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (c *CouponManager) ToggleStatus(ctx context.Context, id string) (*ToggleStatusData, error) {
	var data ToggleStatusData

	if err := c.http.Patch(ctx, v2.RouteToggleStatusCoupon, v2.RESTPatchToggleCouponStatusBody{ID: id}, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
