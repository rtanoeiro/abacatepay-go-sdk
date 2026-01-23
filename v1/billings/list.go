package billings

import (
	"context"

	v1 "github.com/AbacatePay/go-types/v1"
)

func (b *Billings) List(ctx context.Context) (*ListBillingsData, error) {
	var data ListBillingsData

	if err := b.http.Get(ctx, v1.RouteListCharges, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
