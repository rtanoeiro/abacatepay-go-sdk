package products

import (
	"context"
)

func (p *Products) Delete(ctx context.Context, id string) (*DeleteProductData, error) {
	var data DeleteProductData

	// TODO: Add this in go-types
	if err := p.http.Delete(ctx, "/products/delete", DeleteProductBody{ID: id}, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
