package fetch_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"
)

type TestResponse struct {
	Message string `json:"message"`
}

func TestNewFetch(t *testing.T) {
	t.Run("valid config", func(t *testing.T) {
		client, err := fetch.New(fetch.Config{
			APIKey:  "test-key",
			BaseURL: "https://api.test.com",
			Version: "1.0.0",
			Timeout: 5 * time.Second,
		})

		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("error if api key is empty", func(t *testing.T) {
		_, err := fetch.New(fetch.Config{
			BaseURL: "https://api.test.com",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, fetch.ErrInvalidAPIKey)
	})

	t.Run("error if base url is empty", func(t *testing.T) {
		_, err := fetch.New(fetch.Config{
			APIKey: "test-key",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, fetch.ErrInvalidURL)
	})
}

func TestHTTPMethods(t *testing.T) {
	methods := []struct {
		name   string
		method string
		call   func(c *fetch.Fetch, ctx context.Context) (*http.Response, error)
	}{
		{
			name:   "GET",
			method: http.MethodGet,
			call: func(c *fetch.Fetch, ctx context.Context) (*http.Response, error) {
				return c.Get(ctx, "/test")
			},
		},
		{
			name:   "POST",
			method: http.MethodPost,
			call: func(c *fetch.Fetch, ctx context.Context) (*http.Response, error) {
				return c.Post(ctx, "/test", TestResponse{Message: "ok"})
			},
		},
		{
			name:   "PUT",
			method: http.MethodPut,
			call: func(c *fetch.Fetch, ctx context.Context) (*http.Response, error) {
				return c.Put(ctx, "/test", TestResponse{Message: "ok"})
			},
		},
		{
			name:   "DELETE",
			method: http.MethodDelete,
			call: func(c *fetch.Fetch, ctx context.Context) (*http.Response, error) {
				return c.Delete(ctx, "/test")
			},
		},
	}

	for _, tt := range methods {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tt.method, r.Method)
				assert.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

				json.NewEncoder(w).Encode(TestResponse{Message: "ok"})
			}))
			defer server.Close()

			client, err := fetch.New(fetch.Config{
				APIKey:  "test-key",
				BaseURL: server.URL,
				Version: "1.0.0",
			})
			assert.NoError(t, err)

			resp, err := tt.call(client, context.Background())
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}
}

func TestParseResponse(t *testing.T) {
	t.Run("success response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			json.NewEncoder(w).Encode(TestResponse{Message: "success"})
		}))
		defer server.Close()

		client, _ := fetch.New(fetch.Config{
			APIKey:  "test",
			BaseURL: server.URL,
			Version: "1.0.0",
		})

		resp, err := client.Get(context.Background(), "/")
		assert.NoError(t, err)

		out, err := fetch.ParseResponse[TestResponse](resp)
		assert.NoError(t, err)
		assert.Equal(t, "success", out.Message)
	})

	t.Run("error on non-2xx status", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"bad request"}`))
		}))
		defer server.Close()

		client, _ := fetch.New(fetch.Config{
			APIKey:  "test",
			BaseURL: server.URL,
			Version: "1.0.0",
		})

		resp, _ := client.Get(context.Background(), "/")
		_, err := fetch.ParseResponse[TestResponse](resp)

		assert.Error(t, err)
	})
}
