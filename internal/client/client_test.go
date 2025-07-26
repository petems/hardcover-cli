package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	// Test that NewClient creates a client successfully
	client := NewClient("https://example.com/graphql", "test-api-key")
	assert.NotNil(t, client)

	// Test that the client implements the HardcoverClient interface
	var _ HardcoverClient = client
}

func TestClient_Interface(t *testing.T) {
	// Test that the concrete Client type implements all required methods
	client := &Client{}

	// These should compile without errors
	var _ HardcoverClient = client

	// Test method signatures exist
	assert.NotNil(t, client.GetCurrentUser)
	assert.NotNil(t, client.GetBook)
	assert.NotNil(t, client.SearchBooks)
}

func TestAuthedTransport_Do(t *testing.T) {
	// Test that authedTransport adds the correct headers
	transport := &authedTransport{
		wrapped: &http.Client{},
		apiKey:  "test-api-key",
	}

	req, err := http.NewRequest("POST", "https://example.com/graphql", nil)
	require.NoError(t, err)

	// Mock the underlying client to verify headers are set
	originalClient := transport.wrapped
	transport.wrapped = &http.Client{
		Transport: &mockTransport{
			roundTripFunc: func(req *http.Request) (*http.Response, error) {
				// Verify headers are set correctly
				assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
				assert.Equal(t, "hardcover-cli/1.0.0", req.Header.Get("User-Agent"))
				assert.Equal(t, "Bearer test-api-key", req.Header.Get("Authorization"))

				return &http.Response{
					StatusCode: 200,
					Body:       http.NoBody,
				}, nil
			},
		},
	}

	_, err = transport.Do(req)
	require.NoError(t, err)

	// Restore original client
	transport.wrapped = originalClient
}

func TestAuthedTransport_DoWithoutAPIKey(t *testing.T) {
	// Test that authedTransport works without API key
	transport := &authedTransport{
		wrapped: &http.Client{},
		apiKey:  "",
	}

	req, err := http.NewRequest("POST", "https://example.com/graphql", nil)
	require.NoError(t, err)

	// Mock the underlying client to verify headers are set correctly
	transport.wrapped = &http.Client{
		Transport: &mockTransport{
			roundTripFunc: func(req *http.Request) (*http.Response, error) {
				// Verify headers are set correctly
				assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
				assert.Equal(t, "hardcover-cli/1.0.0", req.Header.Get("User-Agent"))
				assert.Equal(t, "", req.Header.Get("Authorization")) // No auth header when no API key

				return &http.Response{
					StatusCode: 200,
					Body:       http.NoBody,
				}, nil
			},
		},
	}

	_, err = transport.Do(req)
	require.NoError(t, err)
}

// mockTransport is a mock HTTP transport for testing
type mockTransport struct {
	roundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.roundTripFunc(req)
}

func TestClient_Methods(t *testing.T) {
	// Test that all client methods exist and have the correct signatures
	client := &Client{}

	// Test that methods exist (without calling them to avoid nil pointer issues)
	assert.NotNil(t, client.GetCurrentUser)
	assert.NotNil(t, client.GetBook)
	assert.NotNil(t, client.SearchBooks)

	// Test that the methods have the correct signatures by checking function types
	// GetCurrentUser should take (context.Context) and return (*GetCurrentUserResponse, error)
	var getCurrentUserFunc func(context.Context) (*GetCurrentUserResponse, error)
	getCurrentUserFunc = client.GetCurrentUser
	assert.NotNil(t, getCurrentUserFunc)

	// GetBook should take (context.Context, string) and return (*GetBookResponse, error)
	var getBookFunc func(context.Context, string) (*GetBookResponse, error)
	getBookFunc = client.GetBook
	assert.NotNil(t, getBookFunc)

	// SearchBooks should take (context.Context, string) and return (*SearchBooksResponse, error)
	var searchBooksFunc func(context.Context, string) (*SearchBooksResponse, error)
	searchBooksFunc = client.SearchBooks
	assert.NotNil(t, searchBooksFunc)
}
