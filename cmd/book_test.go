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

// MockHardcoverClient for book tests - same as in me_test.go but duplicated to avoid import issues
type MockBookClient struct {
	GetCurrentUserFunc func(ctx context.Context) (*client.GetCurrentUserResponse, error)
	GetBookFunc        func(ctx context.Context, id string) (*client.GetBookResponse, error)
	SearchBooksFunc    func(ctx context.Context, query string) (*client.SearchBooksResponse, error)
}

func (m *MockBookClient) GetCurrentUser(ctx context.Context) (*client.GetCurrentUserResponse, error) {
	if m.GetCurrentUserFunc != nil {
		return m.GetCurrentUserFunc(ctx)
	}
	return nil, fmt.Errorf("mock function not implemented")
}

func (m *MockBookClient) GetBook(ctx context.Context, id string) (*client.GetBookResponse, error) {
	if m.GetBookFunc != nil {
		return m.GetBookFunc(ctx, id)
	}
	return nil, fmt.Errorf("mock function not implemented")
}

func (m *MockBookClient) SearchBooks(ctx context.Context, query string) (*client.SearchBooksResponse, error) {
	if m.SearchBooksFunc != nil {
		return m.SearchBooksFunc(ctx, query)
	}
	return nil, fmt.Errorf("mock function not implemented")
}

// Store original NewClient function for book tests
var bookOriginalNewClient = client.NewClient

// Helper function to create a mock client factory for book tests
func withMockBookClient(mockClient client.HardcoverClient) {
	client.NewClient = func(endpoint, apiKey string) client.HardcoverClient {
		return mockClient
	}
}

// Helper function to restore original client for book tests
func restoreOriginalBookClient() {
	client.NewClient = bookOriginalNewClient
}

func TestBookGetCmd_Success(t *testing.T) {
	// Create mock client
	mockClient := &MockBookClient{
		GetBookFunc: func(ctx context.Context, id string) (*client.GetBookResponse, error) {
			assert.Equal(t, "book123", id)
			return &client.GetBookResponse{
				Book: client.GetBookBook{
					Id:              "book123",
					Title:           "Test Book",
					Description:     "A test book description",
					Slug:            "test-book",
					Isbn:            "978-0123456789",
					PublicationYear: 2023,
					PageCount:       350,
					Cached_contributors: []client.GetBookBookCached_contributorsContributor{
						{Name: "John Doe", Role: "author"},
						{Name: "Jane Smith", Role: "editor"},
					},
					Cached_genres: []client.GetBookBookCached_genresGenre{
						{Name: "Fiction"},
						{Name: "Science Fiction"},
					},
					Image:         "https://example.com/cover.jpg",
					AverageRating: 4.5,
					RatingsCount:  100,
					CreatedAt:     "2023-01-01T00:00:00Z",
					UpdatedAt:     "2023-01-02T00:00:00Z",
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
	withMockBookClient(mockClient)
	defer restoreOriginalBookClient()

	// Execute command
	err := bookGetCmd.RunE(cmd, []string{"book123"})
	require.NoError(t, err)

	// Verify output
	outputStr := output.String()
	assert.Contains(t, outputStr, "Book Details:")
	assert.Contains(t, outputStr, "Title: Test Book")
	assert.Contains(t, outputStr, "ID: book123")
	assert.Contains(t, outputStr, "Description: A test book description")
	assert.Contains(t, outputStr, "ISBN: 978-0123456789")
	assert.Contains(t, outputStr, "Publication Year: 2023")
	assert.Contains(t, outputStr, "Page Count: 350")
	assert.Contains(t, outputStr, "John Doe (author)")
	assert.Contains(t, outputStr, "Jane Smith (editor)")
	assert.Contains(t, outputStr, "Genres: Fiction, Science Fiction")
	assert.Contains(t, outputStr, "Average Rating: 4.50 (100 ratings)")
	assert.Contains(t, outputStr, "Image: https://example.com/cover.jpg")
	assert.Contains(t, outputStr, "Hardcover URL: https://hardcover.app/books/test-book")
}

func TestBookGetCmd_MissingAPIKey(t *testing.T) {
	// Create command with empty API key
	cfg := &config.Config{
		APIKey:  "",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	ctx := withConfig(context.Background(), cfg)

	cmd := &cobra.Command{}
	cmd.SetContext(ctx)

	// Execute command
	err := bookGetCmd.RunE(cmd, []string{"book123"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "API key is required")
}

func TestBookGetCmd_NoConfig(t *testing.T) {
	// Create command without config
	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	// Execute command
	err := bookGetCmd.RunE(cmd, []string{"book123"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get configuration")
}

func TestBookGetCmd_APIError(t *testing.T) {
	// Create mock client that returns an error
	mockClient := &MockBookClient{
		GetBookFunc: func(ctx context.Context, id string) (*client.GetBookResponse, error) {
			return nil, fmt.Errorf("GraphQL error: Book not found")
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
	withMockBookClient(mockClient)
	defer restoreOriginalBookClient()

	// Execute command
	err := bookGetCmd.RunE(cmd, []string{"book123"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get book")
	assert.Contains(t, err.Error(), "GraphQL error: Book not found")
}

func TestBookGetCmd_PartialData(t *testing.T) {
	// Create mock client with minimal book data
	mockClient := &MockBookClient{
		GetBookFunc: func(ctx context.Context, id string) (*client.GetBookResponse, error) {
			return &client.GetBookResponse{
				Book: client.GetBookBook{
					Id:    "book456",
					Title: "Minimal Book",
					Slug:  "minimal-book",
					// Other fields are empty/zero values
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
	withMockBookClient(mockClient)
	defer restoreOriginalBookClient()

	// Execute command
	err := bookGetCmd.RunE(cmd, []string{"book456"})
	require.NoError(t, err)

	// Verify output contains required fields but not optional ones
	outputStr := output.String()
	assert.Contains(t, outputStr, "Book Details:")
	assert.Contains(t, outputStr, "Title: Minimal Book")
	assert.Contains(t, outputStr, "ID: book456")
	assert.Contains(t, outputStr, "Hardcover URL: https://hardcover.app/books/minimal-book")
	
	// Should not contain empty optional fields
	assert.NotContains(t, outputStr, "Description:")
	assert.NotContains(t, outputStr, "ISBN:")
	assert.NotContains(t, outputStr, "Publication Year:")
}

func TestBookCmd_CommandProperties(t *testing.T) {
	assert.Equal(t, "book", bookCmd.Use)
	assert.Equal(t, "Manage and retrieve book information", bookCmd.Short)
	assert.NotEmpty(t, bookCmd.Long)
}

func TestBookGetCmd_CommandProperties(t *testing.T) {
	assert.Equal(t, "get <book_id>", bookGetCmd.Use)
	assert.Equal(t, "Get detailed information about a specific book", bookGetCmd.Short)
	assert.NotEmpty(t, bookGetCmd.Long)
	assert.NotNil(t, bookGetCmd.RunE)
}