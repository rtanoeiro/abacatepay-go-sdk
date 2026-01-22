package customers

import (
	"context"

	types "github.com/almeidazs/go-abacate-types/v1"
)

func (c *Customers) Create(ctx context.Context, body CreateCustomerBody) (*CreateCustomerData, error) {
	var customer CreateCustomerData

	if err := c.http.Post(ctx, types.RouteCreateCustomer, body, &customer); err != nil {
		return nil, err
	}

	return &customer, nil
}
