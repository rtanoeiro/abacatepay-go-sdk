package mrr

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (m *MRRManager) Get(ctx context.Context) (*GetMRRData, error) {
	var data GetMRRData

	if err := m.http.Get(ctx, v2.RouteGetMRR, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
