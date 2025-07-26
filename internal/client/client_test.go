package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	endpoint := "https://api.example.com/graphql"
	apiKey := "test-api-key"

	client := NewClient(endpoint, apiKey)

	assert.NotNil(t, client)
	// Note: We can't access internal fields anymore since we're using an interface
	// The client should be functional though
}

func TestClient_GetCurrentUser_Success(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and headers
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))

		// Verify request body
		var req map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)
		assert.Contains(t, req["query"], "query GetCurrentUser")
		assert.Contains(t, req["query"], "me")

		// Send response
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"me": map[string]interface{}{
					"id":        "user123",
					"username":  "testuser",
					"email":     "test@example.com",
					"createdAt": "2023-01-01T00:00:00Z",
					"updatedAt": "2023-01-02T00:00:00Z",
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-api-key")

	response, err := client.GetCurrentUser(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, response)
	user := response.GetMe()
	assert.Equal(t, "user123", user.GetId())
	assert.Equal(t, "testuser", user.GetUsername())
	assert.Equal(t, "test@example.com", user.GetEmail())
}

func TestClient_GetCurrentUser_GraphQLError(t *testing.T) {
	// Create test server that returns GraphQL errors
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"data": nil,
			"errors": []map[string]interface{}{
				{
					"message": "User not found",
					"locations": []map[string]interface{}{
						{"line": 1, "column": 9},
					},
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-api-key")

	_, err := client.GetCurrentUser(context.Background())

	require.Error(t, err)
	assert.Contains(t, err.Error(), "User not found")
}

func TestClient_GetCurrentUser_HTTPError(t *testing.T) {
	// Create test server that returns HTTP error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-api-key")

	_, err := client.GetCurrentUser(context.Background())

	require.Error(t, err)
	assert.Contains(t, err.Error(), "returned error 401")
}

func TestClient_GetCurrentUser_NetworkError(t *testing.T) {
	// Use invalid endpoint to trigger network error
	client := NewClient("http://invalid-endpoint:99999", "test-api-key")

	_, err := client.GetCurrentUser(context.Background())

	require.Error(t, err)
	assert.Contains(t, err.Error(), "dial tcp: address 99999: invalid port")
}

func TestClient_GetCurrentUser_WithoutAPIKey(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify no Authorization header is set
		assert.Empty(t, r.Header.Get("Authorization"))

		// Send response
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"me": map[string]interface{}{
					"id":       "user123",
					"username": "testuser",
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, "")

	response, err := client.GetCurrentUser(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, response)
	user := response.GetMe()
	assert.Equal(t, "user123", user.GetId())
}

func TestClient_GetCurrentUser_WithContext(t *testing.T) {
	// Create test server with delay
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)

		response := map[string]interface{}{
			"data": map[string]interface{}{
				"me": map[string]interface{}{
					"id":       "user123",
					"username": "testuser",
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-api-key")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	_, err := client.GetCurrentUser(ctx)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}

func TestClient_GetBook_Success(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))

		// Send response
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"book": map[string]interface{}{
					"id":    "book123",
					"title": "Test Book",
					"slug":  "test-book",
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-api-key")

	response, err := client.GetBook(context.Background(), "book123")

	require.NoError(t, err)
	assert.NotNil(t, response)
	book := response.GetBook()
	assert.Equal(t, "book123", book.GetId())
	assert.Equal(t, "Test Book", book.GetTitle())
}

func TestClient_SearchBooks_Success(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))

		// Send response
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"search": map[string]interface{}{
					"__typename": "BookSearchResults",
					"totalCount": 1,
					"results": []map[string]interface{}{
						{
							"id":    "book123",
							"title": "Test Book",
							"slug":  "test-book",
						},
					},
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-api-key")

	response, err := client.SearchBooks(context.Background(), "test")

	require.NoError(t, err)
	assert.NotNil(t, response)
	searchResults := response.GetSearch()
	assert.Equal(t, "BookSearchResults", searchResults.GetTypename())
}
