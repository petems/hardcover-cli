package client

import (
	"context"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// HardcoverClient defines the interface for interacting with Hardcover API
type HardcoverClient interface {
	GetCurrentUser(ctx context.Context) (*GetCurrentUserResponse, error)
	GetBook(ctx context.Context, id string) (*GetBookResponse, error)
	SearchBooks(ctx context.Context, query string) (*SearchBooksResponse, error)
}

// Client represents a GraphQL client that uses genqlient
type Client struct {
	graphqlClient graphql.Client
}

// NewClient creates a new GraphQL client using genqlient
func NewClient(endpoint, apiKey string) HardcoverClient {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create a graphql client with auth headers
	graphqlClient := graphql.NewClient(endpoint, &authedTransport{
		wrapped: httpClient,
		apiKey:  apiKey,
	})

	return &Client{
		graphqlClient: graphqlClient,
	}
}

// authedTransport wraps an HTTP client to add authorization headers
type authedTransport struct {
	wrapped *http.Client
	apiKey  string
}

// Do implements the graphql.Doer interface
func (t *authedTransport) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "hardcover-cli/1.0.0")
	if t.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+t.apiKey)
	}
	return t.wrapped.Do(req)
}

// GetCurrentUser gets the current user profile using genqlient
func (c *Client) GetCurrentUser(ctx context.Context) (*GetCurrentUserResponse, error) {
	return GetCurrentUser(ctx, c.graphqlClient)
}

// GetBook gets a book by ID using genqlient
func (c *Client) GetBook(ctx context.Context, id string) (*GetBookResponse, error) {
	return GetBook(ctx, c.graphqlClient, id)
}

// SearchBooks searches for books using genqlient
func (c *Client) SearchBooks(ctx context.Context, query string) (*SearchBooksResponse, error) {
	return SearchBooks(ctx, c.graphqlClient, query)
}
