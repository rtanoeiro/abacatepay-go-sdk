package billing_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/billing"
)

func TestNewBilling(t *testing.T) {
	client := billing.New(nil)
	assert.NotNil(t, client)
}

func TestCreateBilling(t *testing.T) {
	t.Run("invalid body validation", func(t *testing.T) {
		b := billing.New(nil)

		resp, err := b.Create(context.Background(), &billing.CreateBillingBody{})
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("create billing successfully", func(t *testing.T) {
		body := &billing.CreateBillingBody{
			Frequency:     billing.OneTime,
			Methods:       []billing.Method{billing.PIX},
			CompletionUrl: "https://example.com/completion",
			ReturnUrl:     "https://example.com/return",
			Products: []*billing.BillingProduct{
				{
					ExternalId: "pix-1",
					Name:       "PIX",
					Quantity:   1,
					Price:      100,
				},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/v1/billing/create", r.URL.Path)
			assert.Equal(t, http.MethodPost, r.Method)

			var received billing.CreateBillingBody
			json.NewDecoder(r.Body).Decode(&received)

			assert.Equal(t, *body, received)

			json.NewEncoder(w).Encode(billing.CreateBillingResponse{
				Data: billing.CreateBillingResponseItem{
					PublicID: "pix-1",
				},
			})
		}))
		defer server.Close()

		client, _ := fetch.New(fetch.Config{
			APIKey:  "test-key",
			BaseURL: server.URL,
			Version: "1.0.0",
		})

		b := billing.New(client)

		resp, err := b.Create(context.Background(), body)
		assert.NoError(t, err)
		assert.Equal(t, "pix-1", resp.Data.PublicID)
	})
}

func TestListAllBillings(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/billing/list", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)

		json.NewEncoder(w).Encode(billing.ListBillingResponse{
			Data: []billing.BillingListItem{
				{PublicID: "pix-1"},
			},
		})
	}))
	defer server.Close()

	client, _ := fetch.New(fetch.Config{
		APIKey:  "test-key",
		BaseURL: server.URL,
		Version: "1.0.0",
	})

	b := billing.New(client)

	resp, err := b.ListAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, resp.Data, 1)
}
