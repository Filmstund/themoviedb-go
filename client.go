package themoviedb

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const theMovieDBAPIBaseURL = "https://api.themoviedb.org/3"

type Option func(*Client) *Client

type Client struct {
	baseURL   string
	http      *http.Client
	apiKey    string
	language  string
	apiConfig *APIConfiguration
}

// NewClient constructs a new The Movie DB client.
func NewClient(ctx context.Context, apiKey string, opts ...Option) (*Client, error) {
	c := &Client{
		apiKey: apiKey,
	}
	for _, opt := range opts {
		c = opt(c)
	}
	if c.baseURL == "" {
		c.baseURL = theMovieDBAPIBaseURL
	}
	if c.http == nil {
		c.http = &http.Client{}
	}
	if c.apiConfig == nil {
		apiConfig, err := c.APIConfiguration(ctx)
		if err != nil {
			return nil, fmt.Errorf("newClient: %w", err)
		}
		c.apiConfig = apiConfig
	}
	return c, nil
}

// WithLanguage sets the language parameter on all requests that supports it.
// Example: 'en-us'.
func WithLanguage(lang string) Option {
	return func(client *Client) *Client {
		client.language = lang
		return client
	}
}

// WithTimeout sets the HTTP request timeout for all requests. A zero timeout means no timeout.
func WithTimeout(timeout time.Duration) Option {
	return func(client *Client) *Client {
		client.http.Timeout = timeout
		return client
	}
}

// WithAPIConfiguration manually sets the API config, instead of fetching it live.
func WithAPIConfiguration(apiConfig APIConfiguration) Option {
	return func(client *Client) *Client {
		client.apiConfig = &apiConfig
		return client
	}
}

// WithClient sets a custom http client.
func WithClient(httpClient *http.Client) Option {
	return func(client *Client) *Client {
		client.http = httpClient
		return client
	}
}

// WithBaseURL sets the base url for all API calls.
func WithBaseURL(baseURL string) Option {
	return func(client *Client) *Client {
		client.baseURL = baseURL
		return client
	}
}