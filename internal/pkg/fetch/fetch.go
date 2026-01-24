package fetch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	ErrInvalidAPIKey = errors.New("fetch: api key is required")
	ErrInvalidURL    = errors.New("fetch: base url is required")
)

func New(cfg Config) (*Fetch, error) {
	if cfg.APIKey == "" {
		return nil, ErrInvalidAPIKey
	}

	if cfg.BaseURL == "" {
		return nil, ErrInvalidURL
	}

	timeout := cfg.Timeout

	if timeout <= 0 {
		timeout = 5 * time.Second
	}

	return &Fetch{
		apiKey:  cfg.APIKey,
		baseURL: cfg.BaseURL,
		version: cfg.Version,
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func WithTimeout(d time.Duration) RequestOption {
	return func(c *requestConfig) {
		if d > 0 {
			c.timeout = d
		}
	}
}

func WithHeader(k, v string) RequestOption {
	return func(c *requestConfig) {
		c.headers[k] = v
	}
}

func (f *Fetch) Request(
	ctx context.Context,
	method string,
	endpoint string,
	body any,
	opts ...RequestOption,
) (*http.Response, error) {
	url := f.baseURL + "/v" + f.version + endpoint

	var reader io.Reader

	if body != nil {
		b, err := json.Marshal(body)

		if err != nil {
			return nil, fmt.Errorf("fetch: serialize body: %w", err)
		}

		reader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reader)

	if err != nil {
		return nil, fmt.Errorf("fetch: create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+f.apiKey)
	req.Header.Set("User-Agent", "AbacatePay-Go/"+f.version)

	cfg := requestConfig{
		headers: make(map[string]string),
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	for k, v := range cfg.headers {
		req.Header.Set(k, v)
	}

	return f.client.Do(req)
}

func decode(resp *http.Response, out any) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("fetch: read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Body:       body,
		}
	}

	if out == nil || len(body) == 0 {
		return nil
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("fetch: deserialize response: %w", err)
	}

	return nil
}

func (f *Fetch) Get(
	ctx context.Context,
	endpoint string,
	out any,
	opts ...RequestOption,
) error {
	resp, err := f.Request(ctx, http.MethodGet, endpoint, nil, opts...)

	if err != nil {
		return err
	}

	return decode(resp, out)
}

func (f *Fetch) Post(
	ctx context.Context,
	endpoint string,
	body any,
	out any,
	opts ...RequestOption,
) error {
	resp, err := f.Request(ctx, http.MethodPost, endpoint, body, opts...)

	if err != nil {
		return err
	}

	return decode(resp, out)
}

func (f *Fetch) Put(
	ctx context.Context,
	endpoint string,
	body any,
	out any,
	opts ...RequestOption,
) error {
	resp, err := f.Request(ctx, http.MethodPut, endpoint, body, opts...)

	if err != nil {
		return err
	}

	return decode(resp, out)
}

func (f *Fetch) Patch(
	ctx context.Context,
	endpoint string,
	body any,
	out any,
	opts ...RequestOption,
) error {
	resp, err := f.Request(ctx, http.MethodPatch, endpoint, body, opts...)

	if err != nil {
		return err
	}

	return decode(resp, out)
}

func (f *Fetch) Delete(
	ctx context.Context,
	endpoint string,
	body any,
	out any,
	opts ...RequestOption,
) error {
	resp, err := f.Request(ctx, http.MethodDelete, endpoint, body, opts...)

	if err != nil {
		return err
	}

	return decode(resp, out)
}
