package customers

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (c *Customers) Delete(ctx context.Context, id string) (*DeleteCustomerData, error) {
	var data DeleteCustomerData

	if err := c.http.Delete(ctx, v2.RouteDeleteCustomer, v2.RESTDeleteCustomerBody{ID: id}, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
