package transparents

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (t *TransparentManager) CheckStatus(ctx context.Context, id string) (*CheckStatusData, error) {
	var data CheckStatusData

	if err := t.http.Get(ctx, v2.BuildCheckQRCodePIXStatusURL(id), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
