package transparents

import (
	"context"

	v2 "github.com/AbacatePay/go-types/v2"
)

func (t *Transparents) Simulate(ctx context.Context, id string, metadata map[string]any) (*SimulatePaymentData, error) {
	var data SimulatePaymentData

	if err := t.http.Post(ctx, v2.BuildSimulatePaymentURL(id), v2.RESTPostSimulateQRCodePixPaymentBody{Metadata: metadata}, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
