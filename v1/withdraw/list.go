package withdraw

import (
	"context"

	v1 "github.com/AbacatePay/go-types/v1"
)

func (w *Withdraws) List(ctx context.Context) (*ListWithdrawsData, error) {
	var data ListWithdrawsData

	if err := w.http.Get(ctx, v1.RouteListWithdraws, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
