// Package testutil provides testing utilities and helpers for the hardcover-cli application.
package testutil_test

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"hardcover-cli/internal/testutil"
)

func TestDefaultTestConfig(t *testing.T) {
	cfg := testutil.DefaultTestConfig()
	assert.Equal(t, "test-api-key", cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

func TestSetupTestConfig(t *testing.T) {
	testCfg := &testutil.TestConfig{
		APIKey:  "custom-key",
		BaseURL: "custom-url",
	}
	cfg := testutil.SetupTestConfig(testCfg)
	assert.Equal(t, "custom-key", cfg.APIKey)
	assert.Equal(t, "custom-url", cfg.BaseURL)
}

func TestConfig_Methods(t *testing.T) {
	cfg := &testutil.Config{
		APIKey:  "test-key",
		BaseURL: "test-url",
	}
	assert.Equal(t, "test-key", cfg.APIKey)
	assert.Equal(t, "test-url", cfg.BaseURL)
}

func TestCreateTestServer_Success(t *testing.T) {
	data := map[string]interface{}{
		"test": "value",
	}
	response := testutil.SuccessResponse(data)

	server := testutil.CreateTestServer(t, response)
	defer server.Close()

	// Make a request to the test server
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, http.NoBody)
	require.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			t.Errorf("Failed to close response body: %v", closeErr)
		}
	}()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	var graphqlResp testutil.GraphQLResponse
	err = json.NewDecoder(resp.Body).Decode(&graphqlResp)
	require.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(graphqlResp.Data, &result)
	require.NoError(t, err)
	assert.Equal(t, "value", result["test"])
}

func TestCreateTestServer_Error(t *testing.T) {
	errors := []testutil.GraphQLError{
		{Message: "Test error"},
	}
	response := testutil.ErrorResponse(errors)

	server := testutil.CreateTestServer(t, response)
	defer server.Close()

	// Make a request to the test server
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, http.NoBody)
	require.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			t.Errorf("Failed to close response body: %v", closeErr)
		}
	}()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var graphqlResp testutil.GraphQLResponse
	err = json.NewDecoder(resp.Body).Decode(&graphqlResp)
	require.NoError(t, err)

	assert.Len(t, graphqlResp.Errors, 1)
	assert.Equal(t, "Test error", graphqlResp.Errors[0].Message)
}

func TestCreateTestServer_HTTPError(t *testing.T) {
	response := testutil.HTTPErrorResponse(http.StatusBadRequest, "Bad Request")

	server := testutil.CreateTestServer(t, response)
	defer server.Close()

	// Make a request to the test server
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, http.NoBody)
	require.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			t.Errorf("Failed to close response body: %v", closeErr)
		}
	}()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestEnvironmentManager(t *testing.T) {
	envMgr := testutil.NewEnvironmentManager(t)

	// Test setting a new environment variable
	envMgr.SetEnv("TEST_VAR", "test_value")
	assert.Equal(t, "test_value", os.Getenv("TEST_VAR"))

	// Test unsetting an environment variable
	envMgr.UnsetEnv("TEST_VAR")
	assert.Empty(t, os.Getenv("TEST_VAR"))

	// Test modifying existing environment variable
	if err := os.Setenv("EXISTING_VAR", "original"); err != nil {
		t.Errorf("Failed to set EXISTING_VAR: %v", err)
	}
	envMgr.SetEnv("EXISTING_VAR", "modified")
	assert.Equal(t, "modified", os.Getenv("EXISTING_VAR"))

	// Test cleanup restores original values
	envMgr.Cleanup()
	assert.Equal(t, "original", os.Getenv("EXISTING_VAR"))

	// Clean up
	if err := os.Unsetenv("EXISTING_VAR"); err != nil {
		t.Errorf("Failed to unset EXISTING_VAR: %v", err)
	}
}

func TestTempDirManager(t *testing.T) {
	tempDirMgr := testutil.NewTempDirManager(t)
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
	cfg := &testutil.Config{
		APIKey:  "test-key",
		BaseURL: "test-url",
	}
	tempDirMgr.CreateConfig(t, cfg)

	// Verify config file exists
	_, err := os.Stat(configPath)
	assert.NoError(t, err)
}

func TestSetupTestCommand(t *testing.T) {
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-key",
		BaseURL: "test-url",
	})

	// Mock withConfig function
	type testConfigKey struct{}
	mockWithConfig := func(ctx context.Context, cfg interface{}) context.Context {
		return context.WithValue(ctx, testConfigKey{}, cfg)
	}

	cmd, output := testutil.SetupTestCommand(t, cfg, mockWithConfig)

	assert.NotNil(t, cmd)
	assert.NotNil(t, output)

	// Verify context was set
	ctx := cmd.Context()
	assert.NotNil(t, ctx.Value(testConfigKey{}))
}

func TestResponseBuilders(t *testing.T) {
	// Test SuccessResponse
	data := map[string]interface{}{"key": "value"}
	successResp := testutil.SuccessResponse(data)
	assert.Equal(t, data, successResp.Data)
	assert.Equal(t, http.StatusOK, successResp.StatusCode)
	assert.Equal(t, "application/json", successResp.Headers["Content-Type"])

	// Test ErrorResponse
	errors := []testutil.GraphQLError{{Message: "error"}}
	errorResp := testutil.ErrorResponse(errors)
	assert.Nil(t, errorResp.Data)
	assert.Equal(t, errors, errorResp.Errors)
	assert.Equal(t, http.StatusOK, errorResp.StatusCode)

	// Test HTTPErrorResponse
	httpErrorResp := testutil.HTTPErrorResponse(http.StatusNotFound, "Not Found")
	assert.Equal(t, "Not Found", httpErrorResp.Data)
	assert.Equal(t, http.StatusNotFound, httpErrorResp.StatusCode)
}

func TestGraphQLError_Error(t *testing.T) {
	err := testutil.GraphQLError{
		Message: "Test error message",
		Locations: []testutil.GraphQLErrorLocation{
			{Line: 1, Column: 5},
		},
	}

	assert.Equal(t, "Test error message", err.Error())
}
