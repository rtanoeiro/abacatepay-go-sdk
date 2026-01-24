package store

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (s *StoreManager) Get(ctx context.Context) (*GetStoreData, error) {
	var data GetStoreData

	if err := s.http.Get(ctx, v2.RouteStoreDetails, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
