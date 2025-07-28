package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"hardcover-cli/internal/client"
	"hardcover-cli/internal/testutil"
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
					"id": 123,
					"username": "testuser"
				}
			}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Setup config and command
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})

	// Set up context with config
	ctx := testutil.WithTestConfigAdapter(context.Background(), cfg)
	meCmd.SetContext(ctx)

	// Set up output capture
	var output bytes.Buffer
	meCmd.SetOut(&output)

	// Execute command
	err := meCmd.RunE(meCmd, []string{})
	require.NoError(t, err)

	// Verify output
	outputStr := output.String()
	assert.Contains(t, outputStr, "User Profile:")
	assert.Contains(t, outputStr, "ID: 123")
	assert.Contains(t, outputStr, "Username: testuser")
}

func TestMeCmd_MissingAPIKey(t *testing.T) {
	// Setup config with empty API key
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	})

	// Set up context with config
	ctx := testutil.WithTestConfigAdapter(context.Background(), cfg)
	meCmd.SetContext(ctx)

	// Execute command
	err := meCmd.RunE(meCmd, []string{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "API key is required")
}

func TestMeCmd_NoConfig(t *testing.T) {
	// Ensure globalConfig is nil for this test
	originalGlobalConfig := globalConfig
	globalConfig = nil
	defer func() { globalConfig = originalGlobalConfig }()

	// Create command without config
	meCmd.SetContext(context.Background())

	// Execute command
	err := meCmd.RunE(meCmd, []string{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get configuration")
}

func TestMeCmd_APIError(t *testing.T) {
	// Create test server that returns error
	errors := []testutil.GraphQLError{
		{Message: "User not found"},
	}
	server := testutil.CreateTestServer(t, testutil.ErrorResponse(errors))
	defer server.Close()

	// Setup config and command
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})

	// Set up context with config
	ctx := testutil.WithTestConfigAdapter(context.Background(), cfg)
	meCmd.SetContext(ctx)

	// Execute command
	err := meCmd.RunE(meCmd, []string{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get user profile")
}

func TestMeCmd_HTTPError(t *testing.T) {
	// Create test server that returns HTTP error
	server := testutil.CreateTestServer(t, testutil.HTTPErrorResponse(http.StatusUnauthorized, "Unauthorized"))
	defer server.Close()

	// Setup config and command
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})

	// Set up context with config
	ctx := testutil.WithTestConfigAdapter(context.Background(), cfg)
	meCmd.SetContext(ctx)

	// Execute command
	err := meCmd.RunE(meCmd, []string{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get user profile")
}

func TestMeCmd_PartialData(t *testing.T) {
	// Create test server with minimal user data
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{
				"me": {
					"id": 123,
					"username": "testuser"
				}
			}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Setup config and command
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})

	// Set up context with config
	ctx := testutil.WithTestConfigAdapter(context.Background(), cfg)
	meCmd.SetContext(ctx)

	// Set up output capture
	var output bytes.Buffer
	meCmd.SetOut(&output)

	// Execute command
	err := meCmd.RunE(meCmd, []string{})
	require.NoError(t, err)

	// Verify output shows partial user information
	outputStr := output.String()
	assert.Contains(t, outputStr, "ID: 123")
	assert.Contains(t, outputStr, "Username: testuser")
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
			"id": 123,
			"username": "testuser"
		}
	}`

	var response client.GetCurrentUserResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	require.NoError(t, err)

	// Verify the response contains the expected data
	require.NotNil(t, response.Me)
	assert.Equal(t, 123, response.Me.ID)
	assert.Equal(t, "testuser", response.Me.Username)
}
