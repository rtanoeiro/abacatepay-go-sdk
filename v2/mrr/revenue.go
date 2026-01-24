package mrr

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (m *MRRs) Revenue(ctx context.Context, start, end string) (*GetRevenueData, error) {
	var data GetRevenueData

	if err := m.http.Get(ctx, v2.BuildGetRevenueURL(start, end), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
