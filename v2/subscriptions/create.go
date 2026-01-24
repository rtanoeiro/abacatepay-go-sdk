package subscriptions

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (s *Subscriptions) Create(ctx context.Context, body CreateSubscriptionBody) (*CreateSubscriptionData, error) {
	var data CreateSubscriptionData

	if err := s.http.Post(ctx, v2.RouteCreateSubscription, body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
