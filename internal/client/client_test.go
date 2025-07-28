package client_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"hardcover-cli/internal/client"
	"hardcover-cli/internal/testutil"
)

func TestNewClient(t *testing.T) {
	endpoint := "https://api.hardcover.app/v1/graphql"
	apiKey := "test-api-key"

	c := client.NewClient(endpoint, apiKey)

	assert.NotNil(t, c)
	// Test that the client can be used (behavioral test)
	// We can't test unexported fields directly, but we can test the public API
}

func TestClient_Execute_Success(t *testing.T) {
	// Create test server with custom handler for validation
	server := testutil.CreateTestServerWithHandler(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and headers
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "hardcover-cli/1.0.0", r.Header.Get("User-Agent"))
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))

		// Verify request body
		var req client.GraphQLRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			t.Errorf("Failed to decode request body: %v", err)
			return
		}
		assert.Equal(t, "query { test }", req.Query)
		assert.Equal(t, "test-value", req.Variables["test-var"])

		// Send response
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{"test": "success"}`),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("Failed to encode response: %v", err)
		}
	})
	defer server.Close()

	c := client.NewClient(server.URL, "test-api-key")

	variables := map[string]interface{}{
		"test-var": "test-value",
	}

	var result map[string]interface{}
	err := c.Execute(context.Background(), "query { test }", variables, &result)

	require.NoError(t, err)
	assert.Equal(t, "success", result["test"])
}

func TestClient_Execute_GraphQLError(t *testing.T) {
	// Setup GraphQL errors
	errors := []testutil.GraphQLError{
		{
			Message: "Field 'test' doesn't exist",
			Locations: []testutil.GraphQLErrorLocation{
				{Line: 1, Column: 9},
			},
		},
	}

	// Create test server that returns GraphQL errors
	server := testutil.CreateTestServer(t, testutil.ErrorResponse(errors))
	defer server.Close()

	c := client.NewClient(server.URL, "test-api-key")

	var result map[string]interface{}
	err := c.Execute(context.Background(), "query { test }", nil, &result)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "GraphQL errors")
	assert.Contains(t, err.Error(), "Field 'test' doesn't exist")
}

func TestClient_Execute_HTTPError(t *testing.T) {
	// Create test server that returns HTTP error
	server := testutil.CreateTestServer(t, testutil.HTTPErrorResponse(http.StatusUnauthorized, "Unauthorized"))
	defer server.Close()

	c := client.NewClient(server.URL, "test-api-key")

	var result map[string]interface{}
	err := c.Execute(context.Background(), "query { test }", nil, &result)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "HTTP error 401")
	assert.Contains(t, err.Error(), "Unauthorized")
}

func TestClient_Execute_NetworkError(t *testing.T) {
	// Use invalid endpoint to trigger network error
	c := client.NewClient("http://invalid-endpoint:99999", "test-api-key")

	var result map[string]interface{}
	err := c.Execute(context.Background(), "query { test }", nil, &result)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to execute request")
}

func TestClient_Execute_InvalidJSON(t *testing.T) {
	// Create test server that returns invalid JSON
	server := testutil.CreateTestServerWithHandler(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte("invalid json")); err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	})
	defer server.Close()

	c := client.NewClient(server.URL, "test-api-key")

	var result map[string]interface{}
	err := c.Execute(context.Background(), "query { test }", nil, &result)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal response")
}

func TestClient_Execute_WithoutAPIKey(t *testing.T) {
	// Create test server with custom handler to verify no auth header
	server := testutil.CreateTestServerWithHandler(func(w http.ResponseWriter, r *http.Request) {
		// Verify no Authorization header is set
		assert.Empty(t, r.Header.Get("Authorization"))

		// Send response
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{"test": "success"}`),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("Failed to encode response: %v", err)
		}
	})
	defer server.Close()

	c := client.NewClient(server.URL, "")

	var result map[string]interface{}
	err := c.Execute(context.Background(), "query { test }", nil, &result)

	require.NoError(t, err)
	assert.Equal(t, "success", result["test"])
}

func TestClient_Execute_WithContext(t *testing.T) {
	// Create test server with delay
	server := testutil.CreateTestServerWithHandler(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(100 * time.Millisecond)

		response := client.GraphQLResponse{
			Data: json.RawMessage(`{"test": "success"}`),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("Failed to encode response: %v", err)
		}
	})
	defer server.Close()

	c := client.NewClient(server.URL, "test-api-key")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	var result map[string]interface{}
	err := c.Execute(ctx, "query { test }", nil, &result)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}

func TestClient_Execute_NilResult(t *testing.T) {
	// Setup test data
	responseData := map[string]interface{}{
		"test": "success",
	}

	// Create test server
	server := testutil.CreateTestServer(t, testutil.SuccessResponse(responseData))
	defer server.Close()

	c := client.NewClient(server.URL, "test-api-key")

	// Pass nil result to test that it doesn't crash
	err := c.Execute(context.Background(), "query { test }", nil, nil)

	require.NoError(t, err)
}

func TestGraphQLError_Error(t *testing.T) {
	err := client.GraphQLError{
		Message: "Test error message",
		Locations: []client.GraphQLErrorLocation{
			{Line: 1, Column: 5},
		},
	}

	assert.Equal(t, "Test error message", err.Error())
}
