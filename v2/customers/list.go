package customers

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (c *Customers) List(ctx context.Context, page, limit *int) (*ListCustomersData, error) {
	var data ListCustomersData

	if err := c.http.Get(ctx, v2.BuildListCustomersURL(page, limit), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
