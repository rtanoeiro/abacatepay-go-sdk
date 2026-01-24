package payouts

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (p *Payouts) List(ctx context.Context, page, limit *int) (*ListPayoutsData, error) {
	var data ListPayoutsData

	if err := p.http.Get(ctx, v2.BuildListPayoutsURL(page, limit), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
