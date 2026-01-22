package mrr

import (
	"context"

	v1 "github.com/almeidazs/go-abacate-types/v1"
)

func (m *MRRs) Get(ctx context.Context) (*MRR, error) {
	var data MRR

	if err := m.http.Get(ctx, v1.RouteGetMRR, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
