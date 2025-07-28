package testutil

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// TestConfig represents common test configuration options
type TestConfig struct {
	APIKey  string
	BaseURL string
}

// DefaultTestConfig returns a standard test configuration
func DefaultTestConfig() *TestConfig {
	return &TestConfig{
		APIKey:  "test-api-key",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
}

// GraphQLError represents a GraphQL error (copied to avoid import cycle)
type GraphQLError struct {
	Message   string                   `json:"message"`
	Locations []GraphQLErrorLocation   `json:"locations,omitempty"`
}

// Error implements the error interface
func (e GraphQLError) Error() string {
	return e.Message
}

// GraphQLErrorLocation represents the location of a GraphQL error
type GraphQLErrorLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// GraphQLResponse represents a GraphQL response (copied to avoid import cycle)
type GraphQLResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []GraphQLError  `json:"errors,omitempty"`
}

// TestServerResponse represents a configurable test server response
type TestServerResponse struct {
	Data       interface{}
	Errors     []GraphQLError
	StatusCode int
	Headers    map[string]string
}

// TestServerHandler represents a function that handles test server requests
type TestServerHandler func(w http.ResponseWriter, r *http.Request)

// CreateTestServer creates an HTTP test server with default GraphQL response behavior
func CreateTestServer(t *testing.T, response *TestServerResponse) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set default status code
		statusCode := response.StatusCode
		if statusCode == 0 {
			statusCode = http.StatusOK
		}

		// Set headers
		for key, value := range response.Headers {
			w.Header().Set(key, value)
		}

		// Handle non-200 status codes
		if statusCode != http.StatusOK {
			w.WriteHeader(statusCode)
			if response.Data != nil {
				w.Write([]byte(response.Data.(string)))
			}
			return
		}

		// Handle GraphQL response
		var data json.RawMessage
		if response.Data != nil {
			dataBytes, err := json.Marshal(response.Data)
			require.NoError(t, err)
			data = json.RawMessage(dataBytes)
		}

		graphqlResponse := GraphQLResponse{
			Data:   data,
			Errors: response.Errors,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(graphqlResponse)
	}))
}

// CreateTestServerWithHandler creates an HTTP test server with a custom handler
func CreateTestServerWithHandler(handler TestServerHandler) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(handler))
}

// CreateTestServerWithAPIKeyValidation creates a test server that validates API key
func CreateTestServerWithAPIKeyValidation(t *testing.T, expectedAPIKey string, response *TestServerResponse) *httptest.Server {
	t.Helper()

	return CreateTestServerWithHandler(func(w http.ResponseWriter, r *http.Request) {
		// Validate API key
		expectedAuth := "Bearer " + expectedAPIKey
		if r.Header.Get("Authorization") != expectedAuth {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		// Handle the response as normal
		CreateTestServer(t, response).Config.Handler.ServeHTTP(w, r)
	})
}

// Config represents a configuration (copied to avoid import cycle)
type Config struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

// GetAPIKey returns the API key (for interface compatibility)
func (c *Config) GetAPIKey() string {
	return c.APIKey
}

// GetBaseURL returns the base URL (for interface compatibility)
func (c *Config) GetBaseURL() string {
	return c.BaseURL
}

// SetupTestConfig creates a Config for testing
func SetupTestConfig(testCfg *TestConfig) *Config {
	if testCfg == nil {
		testCfg = DefaultTestConfig()
	}

	return &Config{
		APIKey:  testCfg.APIKey,
		BaseURL: testCfg.BaseURL,
	}
}

// WithConfigFunc represents the withConfig function signature from cmd package
type WithConfigFunc func(ctx context.Context, cfg interface{}) context.Context

// SetupTestCommand creates a cobra.Command with test context and config
func SetupTestCommand(t *testing.T, cfg *Config, withConfigFunc WithConfigFunc) (*cobra.Command, *bytes.Buffer) {
	t.Helper()

	cmd := &cobra.Command{}

	// Set up context with config if provided
	if cfg != nil && withConfigFunc != nil {
		ctx := withConfigFunc(context.Background(), cfg)
		cmd.SetContext(ctx)
	} else {
		cmd.SetContext(context.Background())
	}

	// Set up output capture
	var output bytes.Buffer
	cmd.SetOut(&output)

	return cmd, &output
}

// EnvironmentManager helps manage environment variables in tests
type EnvironmentManager struct {
	originalValues map[string]string
	t              *testing.T
}

// NewEnvironmentManager creates a new environment manager
func NewEnvironmentManager(t *testing.T) *EnvironmentManager {
	t.Helper()
	return &EnvironmentManager{
		originalValues: make(map[string]string),
		t:              t,
	}
}

// SetEnv sets an environment variable and remembers the original value
func (em *EnvironmentManager) SetEnv(key, value string) {
	em.t.Helper()

	// Store original value if not already stored
	if _, exists := em.originalValues[key]; !exists {
		em.originalValues[key] = os.Getenv(key)
	}

	os.Setenv(key, value)
}

// UnsetEnv unsets an environment variable and remembers the original value
func (em *EnvironmentManager) UnsetEnv(key string) {
	em.t.Helper()

	// Store original value if not already stored
	if _, exists := em.originalValues[key]; !exists {
		em.originalValues[key] = os.Getenv(key)
	}

	os.Unsetenv(key)
}

// Cleanup restores all environment variables to their original values
func (em *EnvironmentManager) Cleanup() {
	em.t.Helper()

	for key, originalValue := range em.originalValues {
		if originalValue == "" {
			os.Unsetenv(key)
		} else {
			os.Setenv(key, originalValue)
		}
	}
}

// TempDirManager helps manage temporary directories for config testing
type TempDirManager struct {
	tempDir      string
	originalHome string
	t            *testing.T
}

// NewTempDirManager creates a new temporary directory manager
func NewTempDirManager(t *testing.T) *TempDirManager {
	t.Helper()

	tempDir := t.TempDir()
	originalHome := os.Getenv("HOME")

	// Set HOME to temp directory
	os.Setenv("HOME", tempDir)

	return &TempDirManager{
		tempDir:      tempDir,
		originalHome: originalHome,
		t:            t,
	}
}

// GetTempDir returns the temporary directory path
func (tdm *TempDirManager) GetTempDir() string {
	return tdm.tempDir
}

// GetConfigPath returns the config file path in the temp directory
func (tdm *TempDirManager) GetConfigPath() string {
	return filepath.Join(tdm.tempDir, ".hardcover", "config.yaml")
}

// CreateConfig creates a config file in the temp directory
func (tdm *TempDirManager) CreateConfig(t *testing.T, cfg *Config) {
	t.Helper()

	// Create config directory
	configDir := filepath.Join(tdm.tempDir, ".hardcover")
	err := os.MkdirAll(configDir, 0o755)
	require.NoError(t, err)

	// Write config file
	configPath := filepath.Join(configDir, "config.yaml")
	data, err := yaml.Marshal(cfg)
	require.NoError(t, err)

	err = os.WriteFile(configPath, data, 0o600)
	require.NoError(t, err)
}

// Cleanup restores the original HOME environment variable
func (tdm *TempDirManager) Cleanup() {
	os.Setenv("HOME", tdm.originalHome)
}

// TestCase represents a common test case structure
type TestCase struct {
	Name    string
	Setup   func(t *testing.T) interface{}
	Execute func(t *testing.T, setupData interface{}) error
	Verify  func(t *testing.T, setupData interface{}, err error)
	Cleanup func(t *testing.T, setupData interface{})
}

// RunTestCases runs a slice of test cases
func RunTestCases(t *testing.T, testCases []TestCase) {
	t.Helper()

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var setupData interface{}

			if tc.Setup != nil {
				setupData = tc.Setup(t)
			}

			if tc.Cleanup != nil {
				defer tc.Cleanup(t, setupData)
			}

			var err error
			if tc.Execute != nil {
				err = tc.Execute(t, setupData)
			}

			if tc.Verify != nil {
				tc.Verify(t, setupData, err)
			}
		})
	}
}

// Common response builders for GraphQL tests
func SuccessResponse(data interface{}) *TestServerResponse {
	return &TestServerResponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}

func ErrorResponse(errors []GraphQLError) *TestServerResponse {
	return &TestServerResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}

func HTTPErrorResponse(statusCode int, body string) *TestServerResponse {
	return &TestServerResponse{
		Data:       body,
		StatusCode: statusCode,
	}
}
