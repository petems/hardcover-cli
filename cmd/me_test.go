package cmd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"hardcover-cli/internal/testutil"
)

func TestMeCmd_Success(t *testing.T) {
	// Setup test data
	userData := map[string]interface{}{
		"me": map[string]interface{}{
			"id":       "user123",
			"username": "testuser",
		},
	}

	// Create test server
	server := testutil.CreateTestServer(t, testutil.SuccessResponse(userData))
	defer server.Close()

	// Setup config and command
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})
	cmd, output := testutil.SetupTestCommand(t, cfg, testutil.WithTestConfigAdapter)

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
	// Setup config with empty API key
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	})
	cmd, _ := testutil.SetupTestCommand(t, cfg, testutil.WithTestConfigAdapter)

	// Execute command
	err := meCmd.RunE(cmd, []string{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "API key is required")
}

func TestMeCmd_NoConfig(t *testing.T) {
	// Create command without config
	cmd, _ := testutil.SetupTestCommand(t, nil, testutil.WithTestConfigAdapter)

	// Execute command
	err := meCmd.RunE(cmd, []string{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "API key is required")
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
	cmd, _ := testutil.SetupTestCommand(t, cfg, testutil.WithTestConfigAdapter)

	// Execute command
	err := meCmd.RunE(cmd, []string{})
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
	cmd, _ := testutil.SetupTestCommand(t, cfg, testutil.WithTestConfigAdapter)

	// Execute command
	err := meCmd.RunE(cmd, []string{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get user profile")
}

func TestMeCmd_PartialData(t *testing.T) {
	// Setup test data with minimal user data
	userData := map[string]interface{}{
		"me": map[string]interface{}{
			"id":       "user123",
			"username": "testuser",
		},
	}

	// Create test server
	server := testutil.CreateTestServer(t, testutil.SuccessResponse(userData))
	defer server.Close()

	// Setup config and command
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})
	cmd, output := testutil.SetupTestCommand(t, cfg, testutil.WithTestConfigAdapter)

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
