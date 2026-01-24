package payouts

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (p *PayoutManager) Get(ctx context.Context, id string) (*GetPayoutData, error) {
	var data GetPayoutData

	if err := p.http.Get(ctx, v2.BuildGetPayoutURL(id), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
