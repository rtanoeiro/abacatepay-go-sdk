package mrr

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (m *MRRs) Merchant(ctx context.Context) (*GetMerchantData, error) {
	var data GetMerchantData

	if err := m.http.Get(ctx, v2.RouteGetMerchant, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
