package withdraw

import (
	"context"

	v1 "github.com/AbacatePay/go-types/v1"
)

func (w *Withdraws) Get(ctx context.Context, query GetWithdrawQuery) (*GetWithdrawResponse, error) {
	var data GetWithdrawResponse

	if err := w.http.Get(ctx, v1.BuildGetWithdrawURL(query), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
