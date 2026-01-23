package customers

import (
	"context"

	types "github.com/AbacatePay/go-types/v1"
)

func (c *Customers) List(ctx context.Context) (*ListCustomersData, error) {
	var data ListCustomersData

	if err := c.http.Get(ctx, types.RouteListCustomers, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
