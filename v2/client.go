package abacatepay

import (
	"errors"
	"os"
	"time"

	"github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/store"
)

const Version = "2"

const DefaultTimeout = 5 * time.Second

var ErrInvalidAPIKey = errors.New("abacatepay: api key is required")

type ClientConfig struct {
	BaseURL string
	APIKey  string
	Timeout time.Duration
}

type Client struct {
	http *fetch.Fetch

	Store *store.Stores
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
		BaseURL: baseURL,
		APIKey:  cfg.APIKey,
		Version: Version,
		Timeout: timeout,
	})

	if err != nil {
		return nil, err
	}

	return &Client{
		http: httpClient,
		Store: store.New(httpClient),
	}, nil
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
