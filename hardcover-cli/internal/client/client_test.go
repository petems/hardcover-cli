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
	assert.Equal(t, endpoint, client.endpoint)
	assert.Equal(t, apiKey, client.apiKey)
	assert.NotNil(t, client.httpClient)
	assert.Equal(t, 30*time.Second, client.httpClient.Timeout)
}

func TestClient_Execute_Success(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and headers
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "hardcover-cli/1.0.0", r.Header.Get("User-Agent"))
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		
		// Verify request body
		var req GraphQLRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)
		assert.Equal(t, "query { test }", req.Query)
		assert.Equal(t, "test-value", req.Variables["test-var"])
		
		// Send response
		response := GraphQLResponse{
			Data: json.RawMessage(`{"test": "success"}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()
	
	client := NewClient(server.URL, "test-api-key")
	
	variables := map[string]interface{}{
		"test-var": "test-value",
	}
	
	var result map[string]interface{}
	err := client.Execute(context.Background(), "query { test }", variables, &result)
	
	require.NoError(t, err)
	assert.Equal(t, "success", result["test"])
}

func TestClient_Execute_GraphQLError(t *testing.T) {
	// Create test server that returns GraphQL errors
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := GraphQLResponse{
			Data: json.RawMessage(`null`),
			Errors: []GraphQLError{
				{
					Message: "Field 'test' doesn't exist",
					Locations: []GraphQLErrorLocation{
						{Line: 1, Column: 9},
					},
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()
	
	client := NewClient(server.URL, "test-api-key")
	
	var result map[string]interface{}
	err := client.Execute(context.Background(), "query { test }", nil, &result)
	
	require.Error(t, err)
	assert.Contains(t, err.Error(), "GraphQL errors")
	assert.Contains(t, err.Error(), "Field 'test' doesn't exist")
}

func TestClient_Execute_HTTPError(t *testing.T) {
	// Create test server that returns HTTP error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
	}))
	defer server.Close()
	
	client := NewClient(server.URL, "test-api-key")
	
	var result map[string]interface{}
	err := client.Execute(context.Background(), "query { test }", nil, &result)
	
	require.Error(t, err)
	assert.Contains(t, err.Error(), "HTTP error 401")
	assert.Contains(t, err.Error(), "Unauthorized")
}

func TestClient_Execute_NetworkError(t *testing.T) {
	// Use invalid endpoint to trigger network error
	client := NewClient("http://invalid-endpoint:99999", "test-api-key")
	
	var result map[string]interface{}
	err := client.Execute(context.Background(), "query { test }", nil, &result)
	
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to execute request")
}

func TestClient_Execute_InvalidJSON(t *testing.T) {
	// Create test server that returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()
	
	client := NewClient(server.URL, "test-api-key")
	
	var result map[string]interface{}
	err := client.Execute(context.Background(), "query { test }", nil, &result)
	
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal response")
}

func TestClient_Execute_WithoutAPIKey(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify no Authorization header is set
		assert.Empty(t, r.Header.Get("Authorization"))
		
		// Send response
		response := GraphQLResponse{
			Data: json.RawMessage(`{"test": "success"}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()
	
	client := NewClient(server.URL, "")
	
	var result map[string]interface{}
	err := client.Execute(context.Background(), "query { test }", nil, &result)
	
	require.NoError(t, err)
	assert.Equal(t, "success", result["test"])
}

func TestClient_Execute_WithContext(t *testing.T) {
	// Create test server with delay
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		
		response := GraphQLResponse{
			Data: json.RawMessage(`{"test": "success"}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()
	
	client := NewClient(server.URL, "test-api-key")
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	
	var result map[string]interface{}
	err := client.Execute(ctx, "query { test }", nil, &result)
	
	require.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}

func TestClient_Execute_NilResult(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := GraphQLResponse{
			Data: json.RawMessage(`{"test": "success"}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()
	
	client := NewClient(server.URL, "test-api-key")
	
	// Pass nil result to test that it doesn't crash
	err := client.Execute(context.Background(), "query { test }", nil, nil)
	
	require.NoError(t, err)
}

func TestGraphQLError_Error(t *testing.T) {
	err := GraphQLError{
		Message: "Test error message",
		Locations: []GraphQLErrorLocation{
			{Line: 1, Column: 5},
		},
	}
	
	assert.Equal(t, "Test error message", err.Error())
}