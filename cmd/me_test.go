package cmd

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"hardcover-cli/internal/client"
	"hardcover-cli/internal/config"
)

// MockHardcoverClient is a mock implementation of client.HardcoverClient for testing
type MockHardcoverClient struct {
	GetCurrentUserFunc func(ctx context.Context) (*client.GetCurrentUserResponse, error)
	GetBookFunc        func(ctx context.Context, id string) (*client.GetBookResponse, error)
	SearchBooksFunc    func(ctx context.Context, query string) (*client.SearchBooksResponse, error)
}

func (m *MockHardcoverClient) GetCurrentUser(ctx context.Context) (*client.GetCurrentUserResponse, error) {
	if m.GetCurrentUserFunc != nil {
		return m.GetCurrentUserFunc(ctx)
	}
	return nil, fmt.Errorf("mock function not implemented")
}

func (m *MockHardcoverClient) GetBook(ctx context.Context, id string) (*client.GetBookResponse, error) {
	if m.GetBookFunc != nil {
		return m.GetBookFunc(ctx, id)
	}
	return nil, fmt.Errorf("mock function not implemented")
}

func (m *MockHardcoverClient) SearchBooks(ctx context.Context, query string) (*client.SearchBooksResponse, error) {
	if m.SearchBooksFunc != nil {
		return m.SearchBooksFunc(ctx, query)
	}
	return nil, fmt.Errorf("mock function not implemented")
}

// Store original NewClient function
var originalNewClient = client.NewClient

// Helper function to create a mock client factory
func withMockClient(mockClient client.HardcoverClient) {
	client.NewClient = func(endpoint, apiKey string) client.HardcoverClient {
		return mockClient
	}
}

// Helper function to restore original client
func restoreOriginalClient() {
	client.NewClient = originalNewClient
}

func TestMeCmd_Success(t *testing.T) {
	// Create mock client
	mockClient := &MockHardcoverClient{
		GetCurrentUserFunc: func(ctx context.Context) (*client.GetCurrentUserResponse, error) {
			return &client.GetCurrentUserResponse{
				Me: client.GetCurrentUserMeUser{
					Id:        "user123",
					Username:  "testuser",
					Email:     "test@example.com",
					CreatedAt: "2023-01-01T00:00:00Z",
					UpdatedAt: "2023-01-02T00:00:00Z",
				},
			}, nil
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
	
	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)
	
	// Use mock client
	withMockClient(mockClient)
	defer restoreOriginalClient()
	
	// Execute command
	err := meCmd.RunE(cmd, []string{})
	require.NoError(t, err)
	
	// Verify output
	outputStr := output.String()
	assert.Contains(t, outputStr, "User Profile:")
	assert.Contains(t, outputStr, "ID: user123")
	assert.Contains(t, outputStr, "Username: testuser")
	assert.Contains(t, outputStr, "Email: test@example.com")
	assert.Contains(t, outputStr, "Created: 2023-01-01T00:00:00Z")
	assert.Contains(t, outputStr, "Updated: 2023-01-02T00:00:00Z")
}

func TestMeCmd_MissingAPIKey(t *testing.T) {
	// Create command with empty API key
	cfg := &config.Config{
		APIKey:  "",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	ctx := withConfig(context.Background(), cfg)
	
	cmd := &cobra.Command{}
	cmd.SetContext(ctx)
	
	// Execute command
	err := meCmd.RunE(cmd, []string{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "API key is required")
}

func TestMeCmd_NoConfig(t *testing.T) {
	// Create command without config
	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())
	
	// Execute command
	err := meCmd.RunE(cmd, []string{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get configuration")
}

func TestMeCmd_APIError(t *testing.T) {
	// Create mock client that returns an error
	mockClient := &MockHardcoverClient{
		GetCurrentUserFunc: func(ctx context.Context) (*client.GetCurrentUserResponse, error) {
			return nil, fmt.Errorf("GraphQL error: Invalid API key")
		},
	}

	// Create command with test context
	cfg := &config.Config{
		APIKey:  "invalid-key",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	ctx := withConfig(context.Background(), cfg)
	
	cmd := &cobra.Command{}
	cmd.SetContext(ctx)
	
	// Use mock client
	withMockClient(mockClient)
	defer restoreOriginalClient()
	
	// Execute command
	err := meCmd.RunE(cmd, []string{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get user profile")
	assert.Contains(t, err.Error(), "GraphQL error: Invalid API key")
}

func TestMeCmd_PartialData(t *testing.T) {
	// Create mock client with partial data
	mockClient := &MockHardcoverClient{
		GetCurrentUserFunc: func(ctx context.Context) (*client.GetCurrentUserResponse, error) {
			return &client.GetCurrentUserResponse{
				Me: client.GetCurrentUserMeUser{
					Id:       "user456",
					Username: "testuser2",
					// Email, CreatedAt, UpdatedAt are empty
				},
			}, nil
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
	
	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)
	
	// Use mock client
	withMockClient(mockClient)
	defer restoreOriginalClient()
	
	// Execute command
	err := meCmd.RunE(cmd, []string{})
	require.NoError(t, err)
	
	// Verify output contains required fields but not optional ones
	outputStr := output.String()
	assert.Contains(t, outputStr, "User Profile:")
	assert.Contains(t, outputStr, "ID: user456")
	assert.Contains(t, outputStr, "Username: testuser2")
	assert.NotContains(t, outputStr, "Email:")
	assert.NotContains(t, outputStr, "Created:")
	assert.NotContains(t, outputStr, "Updated:")
}

func TestMeCmd_CommandProperties(t *testing.T) {
	assert.Equal(t, "me", meCmd.Use)
	assert.Equal(t, "Get your user profile information", meCmd.Short)
	assert.NotEmpty(t, meCmd.Long)
	assert.NotNil(t, meCmd.RunE)
}