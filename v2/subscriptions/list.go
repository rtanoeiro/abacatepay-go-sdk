package subscriptions

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (s *Subscriptions) List(ctx context.Context, limit *int, cursor *string) (*ListSubscriptionsData, error) {
	var data ListSubscriptionsData

	if err := s.http.Get(ctx, v2.BuildListSubscriptionsURL(limit, cursor), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
