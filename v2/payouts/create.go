package payouts

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (p *Payouts) Create(ctx context.Context, body CreatePayoutBody) (*CreatePayoutData, error) {
	var data CreatePayoutData

	if err := p.http.Post(ctx, v2.RouteCreatePayout, body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
