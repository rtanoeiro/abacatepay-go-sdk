package pix

import (
	"context"

	v1 "github.com/AbacatePay/go-types/v1"
)

func (p *Pix) Status(ctx context.Context, id string) (*CheckPIXStatusData, error) {
	var data CheckPIXStatusData

	// TODO: Fix this later
	if err := p.http.Get(ctx, v1.RouteCheckQRCodePIX+"?"+"id"+id, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
