package checkouts

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (c *Checkouts) Get(ctx context.Context, id string) (*GetCheckoutData, error) {
	var data GetCheckoutData

	if err := c.http.Get(ctx, v2.BuildGetCheckoutURL(id), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
