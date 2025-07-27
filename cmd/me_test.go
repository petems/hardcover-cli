package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"hardcover-cli/internal/client"
	"hardcover-cli/internal/config"
)

func TestMeCmd_Success(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))

		// Verify GraphQL query
		var req client.GraphQLRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)
		assert.Contains(t, req.Query, "query GetCurrentUser")
		assert.Contains(t, req.Query, "me")

		// Send response
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{
				"me": {
					"id": "user123",
					"username": "testuser"
				}
			}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create command with test context
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}
	ctx := withConfig(context.Background(), cfg)

	cmd := &cobra.Command{}
	cmd.SetContext(ctx)

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err := meCmd.RunE(cmd, []string{})
	require.NoError(t, err)

	// Verify output
	outputStr := output.String()
	assert.Contains(t, outputStr, "User Profile:")
	assert.Contains(t, outputStr, "ID: user123")
	assert.Contains(t, outputStr, "Username: testuser")
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
	assert.Contains(t, err.Error(), "API key is required")
}

func TestMeCmd_APIError(t *testing.T) {
	// Create test server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`null`),
			Errors: []client.GraphQLError{
				{
					Message: "User not found",
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create command with test context
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}
	ctx := withConfig(context.Background(), cfg)

	cmd := &cobra.Command{}
	cmd.SetContext(ctx)

	// Execute command
	err := meCmd.RunE(cmd, []string{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get user profile")
}

func TestMeCmd_HTTPError(t *testing.T) {
	// Create test server that returns HTTP error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
	}))
	defer server.Close()

	// Create command with test context
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}
	ctx := withConfig(context.Background(), cfg)

	cmd := &cobra.Command{}
	cmd.SetContext(ctx)

	// Execute command
	err := meCmd.RunE(cmd, []string{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get user profile")
}

func TestMeCmd_PartialData(t *testing.T) {
	// Create test server with minimal user data
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{
				"me": {
					"id": "user123",
					"username": "testuser"
				}
			}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create command with test context
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}
	ctx := withConfig(context.Background(), cfg)

	cmd := &cobra.Command{}
	cmd.SetContext(ctx)

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err := meCmd.RunE(cmd, []string{})
	require.NoError(t, err)

	// Verify output contains required fields but not optional ones
	outputStr := output.String()
	assert.Contains(t, outputStr, "ID: user123")
	assert.Contains(t, outputStr, "Username: testuser")
	assert.NotContains(t, outputStr, "Email:")
	assert.NotContains(t, outputStr, "Created:")
	assert.NotContains(t, outputStr, "Updated:")
}

func TestMeCmd_CommandProperties(t *testing.T) {
	// Test command properties
	assert.Equal(t, "me", meCmd.Use)
	assert.Equal(t, "Get the current user's profile information based on the API key", meCmd.Short)
	assert.NotEmpty(t, meCmd.Long)
	assert.Contains(t, meCmd.Long, "User ID")
	assert.Contains(t, meCmd.Long, "Username")
	assert.Contains(t, meCmd.Long, "hardcover-cli me")
}

func TestMeCmd_Integration(t *testing.T) {
	// Setup commands for testing
	setupMeCommands()

	// Test the command is properly registered
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "me" {
			found = true
			break
		}
	}
	assert.True(t, found, "me command should be registered with root command")
}

func TestGetCurrentUserResponse_JSONUnmarshal(t *testing.T) {
	// Test JSON unmarshaling
	jsonData := `{
		"me": {
			"id": "user123",
			"username": "testuser"
		}
	}`

	var response GetCurrentUserResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	require.NoError(t, err)

	// Verify the response contains the expected data
	if userData, ok := response.Me.(map[string]interface{}); ok {
		assert.Equal(t, "user123", userData["id"])
		assert.Equal(t, "testuser", userData["username"])
	}
}
