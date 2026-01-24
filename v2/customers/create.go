package customers

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (c *CustomerManager) Create(ctx context.Context, body CreateCustomerBody) (*CreateCustomerData, error) {
	var data CreateCustomerData

	if err := c.http.Post(ctx, v2.RouteCreateCustomer, body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
