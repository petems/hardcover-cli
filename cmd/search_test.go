package cmd

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"hardcover-cli/internal/client"
	"hardcover-cli/internal/config"
)

// MockSearchClient for search tests
type MockSearchClient struct {
	GetCurrentUserFunc func(ctx context.Context) (*client.GetCurrentUserResponse, error)
	GetBookFunc        func(ctx context.Context, id string) (*client.GetBookResponse, error)
	SearchBooksFunc    func(ctx context.Context, query string) (*client.SearchBooksResponse, error)
}

func (m *MockSearchClient) GetCurrentUser(ctx context.Context) (*client.GetCurrentUserResponse, error) {
	if m.GetCurrentUserFunc != nil {
		return m.GetCurrentUserFunc(ctx)
	}
	return nil, fmt.Errorf("mock function not implemented")
}

func (m *MockSearchClient) GetBook(ctx context.Context, id string) (*client.GetBookResponse, error) {
	if m.GetBookFunc != nil {
		return m.GetBookFunc(ctx, id)
	}
	return nil, fmt.Errorf("mock function not implemented")
}

func (m *MockSearchClient) SearchBooks(ctx context.Context, query string) (*client.SearchBooksResponse, error) {
	if m.SearchBooksFunc != nil {
		return m.SearchBooksFunc(ctx, query)
	}
	return nil, fmt.Errorf("mock function not implemented")
}

// Store original NewClient function for search tests
var searchOriginalNewClient = client.NewClient

// Helper function to create a mock client factory for search tests
func withMockSearchClient(mockClient client.HardcoverClient) {
	client.NewClient = func(endpoint, apiKey string) client.HardcoverClient {
		return mockClient
	}
}

// Helper function to restore original client for search tests
func restoreOriginalSearchClient() {
	client.NewClient = searchOriginalNewClient
}

func TestSearchBooksCmd_MissingAPIKey(t *testing.T) {
	// Create command with empty API key
	cfg := &config.Config{
		APIKey:  "",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	ctx := withConfig(context.Background(), cfg)

	cmd := &cobra.Command{}
	cmd.SetContext(ctx)

	// Execute command
	err := searchBooksCmd.RunE(cmd, []string{"test"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "API key is required")
}

func TestSearchBooksCmd_NoConfig(t *testing.T) {
	// Create command without config
	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	// Execute command
	err := searchBooksCmd.RunE(cmd, []string{"test"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get configuration")
}

func TestSearchBooksCmd_APIError(t *testing.T) {
	// Create mock client that returns an error
	mockClient := &MockSearchClient{
		SearchBooksFunc: func(ctx context.Context, query string) (*client.SearchBooksResponse, error) {
			return nil, fmt.Errorf("GraphQL error: Search failed")
		},
	}

	// Create command with test context
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	ctx := withConfig(context.Background(), cfg)

	cmd := &cobra.Command{}
	cmd.SetContext(ctx)

	// Use mock client
	withMockSearchClient(mockClient)
	defer restoreOriginalSearchClient()

	// Execute command
	err := searchBooksCmd.RunE(cmd, []string{"test"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to search books")
	assert.Contains(t, err.Error(), "GraphQL error: Search failed")
}

func TestSearchCmd_CommandProperties(t *testing.T) {
	assert.Equal(t, "search", searchCmd.Use)
	assert.Equal(t, "Search for content on Hardcover.app", searchCmd.Short)
	assert.NotEmpty(t, searchCmd.Long)
}

func TestSearchBooksCmd_CommandProperties(t *testing.T) {
	assert.Equal(t, "books <query>", searchBooksCmd.Use)
	assert.Equal(t, "Search for books", searchBooksCmd.Short)
	assert.NotEmpty(t, searchBooksCmd.Long)
	assert.NotNil(t, searchBooksCmd.RunE)
}
