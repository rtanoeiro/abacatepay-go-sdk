package billing

import (
	"context"
	"fmt"

	"github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"
)

type Billing struct {
	HttpClient *fetch.Fetch
}

func New(httpClient *fetch.Fetch) *Billing {
	return &Billing{
		HttpClient: httpClient,
	}
}

func (b *Billing) Create(
	ctx context.Context,
	body *CreateBillingBody,
) (*CreateBillingResponse, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}

	if body.CustomerId == "" && (body.Customer == nil || body.Customer.Email == "") {
		return nil, fmt.Errorf("customerId or customer.email is required")
	}

	resp, err := b.HttpClient.Post(ctx, "/v1/billing/create", body)

	if err != nil {
		return nil, err
	}

	response, err := fetch.ParseResponse[CreateBillingResponse](resp)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (b *Billing) ListAll(ctx context.Context) (*ListBillingResponse, error) {
	resp, err := b.HttpClient.Get(ctx, "/v1/billing/list")

	if err != nil {
		return nil, err
	}

	response, err := fetch.ParseResponse[ListBillingResponse](resp)

	if err != nil {
		return nil, err
	}

	return response, nil
}
