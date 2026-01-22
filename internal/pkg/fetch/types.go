package fetch

import (
	"net/http"
	"time"
)

type Config struct {
	APIKey  string
	BaseURL string
	Version string
	Timeout time.Duration
}

type Fetch struct {
	apiKey  string
	baseURL string
	version string
	client  *http.Client
}

type RequestOption func(*requestConfig)

type requestConfig struct {
	timeout time.Duration
	headers map[string]string
}
