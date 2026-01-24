package checkouts

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (c *Checkouts) List(ctx context.Context) (*ListCheckoutsData, error) {
	var data ListCheckoutsData

	if err := c.http.Get(ctx, v2.RouteListCheckouts, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
