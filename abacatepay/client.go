package abacatepay

import (
	"errors"
	"os"
	"time"

	"github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/billing"
)

var Version = "dev"

const DefaultTimeout = 5 * time.Second

var ErrInvalidAPIKey = errors.New("abacatepay: api key is required")

type Client struct {
	http *fetch.Fetch

	Billing *billing.Billing
}

type ClientConfig struct {
	BaseURL string
	APIKey  string
	Timeout time.Duration
}

func New(cfg ClientConfig) (*Client, error) {
	if cfg.APIKey == "" {
		return nil, ErrInvalidAPIKey
	}

	baseURL := cfg.BaseURL

	if baseURL == "" {
		baseURL = getenv("ABACATEPAY_API_URL", "https://api.abacatepay.com")
	}

	timeout := cfg.Timeout

	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	httpClient, err := fetch.New(fetch.Config{
		APIKey:  cfg.APIKey,
		BaseURL: baseURL,
		Version: Version,
		Timeout: timeout,
	})

	if err != nil {
		return nil, err
	}

	return &Client{
		http:    httpClient,
		Billing: billing.New(httpClient),
	}, nil
}

func (c *Client) HTTP() *fetch.Fetch {
	return c.http
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	
	return fallback
}
