package transparents

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (t *TransparentManager) Create(ctx context.Context, body CreateQRCodePIXBody) (*CreateQRCodePIXData, error) {
	var data CreateQRCodePIXData

	if err := t.http.Post(ctx, v2.RouteCreatePixQRCode, body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
