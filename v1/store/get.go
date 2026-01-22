package store

import (
	"context"

	v1 "github.com/almeidazs/go-abacate-types/v1"
)

func (s *Stores) Get(ctx context.Context) (*Store, error) {
	var store Store

	if err := s.http.Get(ctx, v1.RouteStoreDetails, &store); err != nil {
		return nil, err
	}

	return &store, nil
}
