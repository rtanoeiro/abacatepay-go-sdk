package withdraw

import (
	"context"

	v1 "github.com/almeidazs/go-abacate-types/v1"
)

func (w *Withdraws) Create(ctx context.Context, body CreateWithdrawBody) (*CreateWithdrawResponse, error) {
	var data CreateWithdrawResponse

	if err := w.http.Post(ctx, v1.RouteCreateWithdraw, body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
