package store

import (
	"context"

	v1 "github.com/AbacatePay/go-types/v1"
)

func (s *Stores) Get(ctx context.Context) (*GetStoreData, error) {
	var store GetStoreData

	if err := s.http.Get(ctx, v1.RouteStoreDetails, &store); err != nil {
		return nil, err
	}

	return &store, nil
}
