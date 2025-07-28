package testutil

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultTestConfig(t *testing.T) {
	cfg := DefaultTestConfig()
	assert.NotNil(t, cfg)
	assert.Equal(t, "test-api-key", cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

func TestSetupTestConfig(t *testing.T) {
	// Test with nil input
	cfg := SetupTestConfig(nil)
	assert.NotNil(t, cfg)
	assert.Equal(t, "test-api-key", cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)

	// Test with custom input
	customCfg := &TestConfig{
		APIKey:  "custom-key",
		BaseURL: "https://custom.url",
	}
	cfg = SetupTestConfig(customCfg)
	assert.Equal(t, "custom-key", cfg.APIKey)
	assert.Equal(t, "https://custom.url", cfg.BaseURL)
}

func TestConfig_Methods(t *testing.T) {
	cfg := &Config{
		APIKey:  "test-key",
		BaseURL: "test-url",
	}

	assert.Equal(t, "test-key", cfg.GetAPIKey())
	assert.Equal(t, "test-url", cfg.GetBaseURL())
}

func TestCreateTestServer_Success(t *testing.T) {
	data := map[string]interface{}{
		"test": "value",
	}
	response := SuccessResponse(data)

	server := CreateTestServer(t, response)
	defer server.Close()

	// Make a request to the test server
	resp, err := http.Get(server.URL)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	var graphqlResp GraphQLResponse
	err = json.NewDecoder(resp.Body).Decode(&graphqlResp)
	require.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(graphqlResp.Data, &result)
	require.NoError(t, err)
	assert.Equal(t, "value", result["test"])
}

func TestCreateTestServer_Error(t *testing.T) {
	errors := []GraphQLError{
		{Message: "Test error"},
	}
	response := ErrorResponse(errors)

	server := CreateTestServer(t, response)
	defer server.Close()

	// Make a request to the test server
	resp, err := http.Get(server.URL)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var graphqlResp GraphQLResponse
	err = json.NewDecoder(resp.Body).Decode(&graphqlResp)
	require.NoError(t, err)

	assert.Len(t, graphqlResp.Errors, 1)
	assert.Equal(t, "Test error", graphqlResp.Errors[0].Message)
}

func TestCreateTestServer_HTTPError(t *testing.T) {
	response := HTTPErrorResponse(http.StatusBadRequest, "Bad Request")

	server := CreateTestServer(t, response)
	defer server.Close()

	// Make a request to the test server
	resp, err := http.Get(server.URL)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestEnvironmentManager(t *testing.T) {
	envMgr := NewEnvironmentManager(t)

	// Test setting a new environment variable
	envMgr.SetEnv("TEST_VAR", "test_value")
	assert.Equal(t, "test_value", os.Getenv("TEST_VAR"))

	// Test unsetting an environment variable
	envMgr.UnsetEnv("TEST_VAR")
	assert.Empty(t, os.Getenv("TEST_VAR"))

	// Test modifying existing environment variable
	os.Setenv("EXISTING_VAR", "original")
	envMgr.SetEnv("EXISTING_VAR", "modified")
	assert.Equal(t, "modified", os.Getenv("EXISTING_VAR"))

	// Test cleanup restores original values
	envMgr.Cleanup()
	assert.Equal(t, "original", os.Getenv("EXISTING_VAR"))

	// Clean up
	os.Unsetenv("EXISTING_VAR")
}

func TestTempDirManager(t *testing.T) {
	tempDirMgr := NewTempDirManager(t)
	defer tempDirMgr.Cleanup()

	// Test temp directory creation
	tempDir := tempDirMgr.GetTempDir()
	assert.NotEmpty(t, tempDir)

	// Verify HOME is set to temp directory
	assert.Equal(t, tempDir, os.Getenv("HOME"))

	// Test config path
	configPath := tempDirMgr.GetConfigPath()
	expectedPath := tempDir + "/.hardcover/config.yaml"
	assert.Equal(t, expectedPath, configPath)

	// Test config creation
	cfg := &Config{
		APIKey:  "test-key",
		BaseURL: "test-url",
	}
	tempDirMgr.CreateConfig(t, cfg)

	// Verify config file exists
	_, err := os.Stat(configPath)
	assert.NoError(t, err)
}

func TestSetupTestCommand(t *testing.T) {
	cfg := SetupTestConfig(&TestConfig{
		APIKey:  "test-key",
		BaseURL: "test-url",
	})

	// Mock withConfig function
	mockWithConfig := func(ctx context.Context, cfg interface{}) context.Context {
		return context.WithValue(ctx, "test-config", cfg)
	}

	cmd, output := SetupTestCommand(t, cfg, mockWithConfig)

	assert.NotNil(t, cmd)
	assert.NotNil(t, output)

	// Verify context was set
	ctx := cmd.Context()
	assert.NotNil(t, ctx.Value("test-config"))
}

func TestResponseBuilders(t *testing.T) {
	// Test SuccessResponse
	data := map[string]interface{}{"key": "value"}
	successResp := SuccessResponse(data)
	assert.Equal(t, data, successResp.Data)
	assert.Equal(t, http.StatusOK, successResp.StatusCode)
	assert.Equal(t, "application/json", successResp.Headers["Content-Type"])

	// Test ErrorResponse
	errors := []GraphQLError{{Message: "error"}}
	errorResp := ErrorResponse(errors)
	assert.Nil(t, errorResp.Data)
	assert.Equal(t, errors, errorResp.Errors)
	assert.Equal(t, http.StatusOK, errorResp.StatusCode)

	// Test HTTPErrorResponse
	httpErrorResp := HTTPErrorResponse(http.StatusNotFound, "Not Found")
	assert.Equal(t, "Not Found", httpErrorResp.Data)
	assert.Equal(t, http.StatusNotFound, httpErrorResp.StatusCode)
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
