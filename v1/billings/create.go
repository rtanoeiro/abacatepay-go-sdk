package billings

import (
	"context"

	v1 "github.com/AbacatePay/go-types/v1"
)

func (b *Billings) Create(ctx context.Context, body CreateBillingBody) (*CreateBillingData, error) {
	var data CreateBillingData

	if err := b.http.Post(ctx, v1.RouteCreateCharge, body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
