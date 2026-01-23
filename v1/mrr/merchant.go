package mrr

import (
	"context"

	v1 "github.com/AbacatePay/go-types/v1"
)

func (m *MRRs) Merchant(ctx context.Context) (*Merchant, error) {
	var data Merchant

	if err := m.http.Get(ctx, v1.RouteGetMerchant, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
