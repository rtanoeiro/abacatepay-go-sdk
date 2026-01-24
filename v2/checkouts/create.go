package checkouts

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (c *Checkouts) Create(ctx context.Context, body CreateCheckoutBody) (*CreateCheckoutData, error) {
	var data CreateCheckoutData

	if err := c.http.Post(ctx, v2.RouteCreateCheckout, body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
