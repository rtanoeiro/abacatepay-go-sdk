package customers

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (c *Customers) Get(ctx context.Context, id string) (*GetCustomerData, error) {
	var data GetCustomerData

	if err := c.http.Get(ctx, v2.BuildGetCustomerURL(id), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
