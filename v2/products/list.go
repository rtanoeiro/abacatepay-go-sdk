package products

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (p *ProductManager) List(ctx context.Context, page, limit *int) (*ListProductsData, error) {
	var data ListProductsData

	if err := p.http.Get(ctx, v2.BuildListProductsURL(page, limit), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
