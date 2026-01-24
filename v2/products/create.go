package products

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (p *Products) Create(ctx context.Context, body CreateProductBody) (*CreateProductData, error) {
	var data CreateProductData

	if err := p.http.Post(ctx, v2.RouteCreateProduct, body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
