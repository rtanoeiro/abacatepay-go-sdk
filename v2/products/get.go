package products

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (p *ProductManager) Get(ctx context.Context, id, externalId *string) (*GetProductData, error) {
	var data GetProductData

	if err := p.http.Get(ctx, v2.BuildGetProductURL(id, externalId), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
