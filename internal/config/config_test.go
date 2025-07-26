package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	assert.NotNil(t, cfg)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
	assert.Empty(t, cfg.APIKey)
}

func TestLoadConfig_FromEnvironment(t *testing.T) {
	// Set environment variable
	expectedAPIKey := "test-api-key-from-env"
	os.Setenv("HARDCOVER_API_KEY", expectedAPIKey)
	defer os.Unsetenv("HARDCOVER_API_KEY")

	cfg, err := LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, expectedAPIKey, cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

func TestLoadConfig_FromFile(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, ".hardcover")
	configPath := filepath.Join(configDir, "config.yaml")

	// Create config directory
	err := os.MkdirAll(configDir, 0o755)
	require.NoError(t, err)

	// Create config file
	configContent := `api_key: test-api-key-from-file
base_url: https://api.hardcover.app/v1/graphql`
	err = os.WriteFile(configPath, []byte(configContent), 0o600)
	require.NoError(t, err)

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cfg, err := LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, "test-api-key-from-file", cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

func TestLoadConfig_NoFileExists(t *testing.T) {
	// Make sure no environment variable is set
	os.Unsetenv("HARDCOVER_API_KEY")

	// Create temporary directory for config
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cfg, err := LoadConfig()
	require.NoError(t, err)
	assert.Empty(t, cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

func TestSaveConfig(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cfg := &Config{
		APIKey:  "test-api-key",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}

	err := SaveConfig(cfg)
	require.NoError(t, err)

	// Verify file was created
	configPath := filepath.Join(tempDir, ".hardcover", "config.yaml")
	_, err = os.Stat(configPath)
	require.NoError(t, err)

	// Load the config back and verify
	loadedCfg, err := LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, cfg.APIKey, loadedCfg.APIKey)
	assert.Equal(t, cfg.BaseURL, loadedCfg.BaseURL)
}

func TestGetConfigPath(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	configPath, err := GetConfigPath()
	require.NoError(t, err)

	expectedPath := filepath.Join(tempDir, ".hardcover", "config.yaml")
	assert.Equal(t, expectedPath, configPath)
}

func TestLoadConfig_EnvironmentOverridesFile(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, ".hardcover")
	configPath := filepath.Join(configDir, "config.yaml")

	// Create config directory
	err := os.MkdirAll(configDir, 0o755)
	require.NoError(t, err)

	// Create config file with one API key
	configContent := `api_key: file-api-key
base_url: https://api.hardcover.app/v1/graphql`
	err = os.WriteFile(configPath, []byte(configContent), 0o600)
	require.NoError(t, err)

	// Set environment variable with different API key
	envAPIKey := "env-api-key"
	os.Setenv("HARDCOVER_API_KEY", envAPIKey)
	defer os.Unsetenv("HARDCOVER_API_KEY")

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cfg, err := LoadConfig()
	require.NoError(t, err)

	// Environment variable should override file
	assert.Equal(t, envAPIKey, cfg.APIKey)
}

func TestLoadConfig_InvalidYAML(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, ".hardcover")
	configPath := filepath.Join(configDir, "config.yaml")

	// Create config directory
	err := os.MkdirAll(configDir, 0o755)
	require.NoError(t, err)

	// Create config file with invalid YAML
	configContent := `api_key: test-api-key
base_url: https://api.hardcover.app/v1/graphql
invalid_yaml: [unclosed bracket`
	err = os.WriteFile(configPath, []byte(configContent), 0o600)
	require.NoError(t, err)

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	_, err = LoadConfig()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse config file")
}

// TestLoadConfig_EmptyEnvironmentVariable tests the specific fix for the issue
// where an empty environment variable was preventing config file loading
func TestLoadConfig_EmptyEnvironmentVariable(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, ".hardcover")
	configPath := filepath.Join(configDir, "config.yaml")

	// Create config directory
	err := os.MkdirAll(configDir, 0o755)
	require.NoError(t, err)

	// Create config file with API key
	configContent := `api_key: test-api-key-from-file
base_url: https://api.hardcover.app/v1/graphql`
	err = os.WriteFile(configPath, []byte(configContent), 0o600)
	require.NoError(t, err)

	// Set empty environment variable (this was the bug)
	os.Setenv("HARDCOVER_API_KEY", "")
	defer os.Unsetenv("HARDCOVER_API_KEY")

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cfg, err := LoadConfig()
	require.NoError(t, err)

	// Should load from config file even with empty environment variable
	assert.Equal(t, "test-api-key-from-file", cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

// TestLoadConfig_WhitespaceEnvironmentVariable tests that whitespace-only environment variables are treated as empty
func TestLoadConfig_WhitespaceEnvironmentVariable(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, ".hardcover")
	configPath := filepath.Join(configDir, "config.yaml")

	// Create config directory
	err := os.MkdirAll(configDir, 0o755)
	require.NoError(t, err)

	// Create config file with API key
	configContent := `api_key: test-api-key-from-file
base_url: https://api.hardcover.app/v1/graphql`
	err = os.WriteFile(configPath, []byte(configContent), 0o600)
	require.NoError(t, err)

	// Set whitespace-only environment variable
	os.Setenv("HARDCOVER_API_KEY", "   ")
	defer os.Unsetenv("HARDCOVER_API_KEY")

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cfg, err := LoadConfig()
	require.NoError(t, err)

	// Should load from config file even with whitespace-only environment variable
	assert.Equal(t, "test-api-key-from-file", cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

// TestLoadConfig_EnvironmentVariableOverridesFile tests that non-empty environment variables properly override config file
func TestLoadConfig_EnvironmentVariableOverridesFile(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, ".hardcover")
	configPath := filepath.Join(configDir, "config.yaml")

	// Create config directory
	err := os.MkdirAll(configDir, 0o755)
	require.NoError(t, err)

	// Create config file with API key
	configContent := `api_key: test-api-key-from-file
base_url: https://api.hardcover.app/v1/graphql`
	err = os.WriteFile(configPath, []byte(configContent), 0o600)
	require.NoError(t, err)

	// Set non-empty environment variable
	envAPIKey := "test-api-key-from-env"
	os.Setenv("HARDCOVER_API_KEY", envAPIKey)
	defer os.Unsetenv("HARDCOVER_API_KEY")

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cfg, err := LoadConfig()
	require.NoError(t, err)

	// Environment variable should override config file
	assert.Equal(t, envAPIKey, cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

// TestLoadConfig_NoEnvironmentVariableNoFile tests the default behavior when neither env var nor file exists
func TestLoadConfig_NoEnvironmentVariableNoFile(t *testing.T) {
	// Make sure no environment variable is set
	os.Unsetenv("HARDCOVER_API_KEY")

	// Create temporary directory for config
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cfg, err := LoadConfig()
	require.NoError(t, err)

	// Should return default config with empty API key
	assert.Empty(t, cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

// TestLoadConfig_ConfigFileOnly tests loading from config file when no environment variable is set
func TestLoadConfig_ConfigFileOnly(t *testing.T) {
	// Make sure no environment variable is set
	os.Unsetenv("HARDCOVER_API_KEY")

	// Create temporary directory for config
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, ".hardcover")
	configPath := filepath.Join(configDir, "config.yaml")

	// Create config directory
	err := os.MkdirAll(configDir, 0o755)
	require.NoError(t, err)

	// Create config file with API key
	configContent := `api_key: test-api-key-from-file
base_url: https://api.hardcover.app/v1/graphql`
	err = os.WriteFile(configPath, []byte(configContent), 0o600)
	require.NoError(t, err)

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cfg, err := LoadConfig()
	require.NoError(t, err)

	// Should load from config file
	assert.Equal(t, "test-api-key-from-file", cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

// TestLoadConfig_EnvironmentVariableOnly tests loading from environment variable when no config file exists
func TestLoadConfig_EnvironmentVariableOnly(t *testing.T) {
	// Set environment variable
	expectedAPIKey := "test-api-key-from-env"
	os.Setenv("HARDCOVER_API_KEY", expectedAPIKey)
	defer os.Unsetenv("HARDCOVER_API_KEY")

	// Create temporary directory for config (but don't create config file)
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cfg, err := LoadConfig()
	require.NoError(t, err)

	// Should load from environment variable
	assert.Equal(t, expectedAPIKey, cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

// TestLoadConfig_ConfigFileWithCustomBaseURL tests that custom base URL from config file is preserved
func TestLoadConfig_ConfigFileWithCustomBaseURL(t *testing.T) {
	// Make sure no environment variable is set
	os.Unsetenv("HARDCOVER_API_KEY")

	// Create temporary directory for config
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, ".hardcover")
	configPath := filepath.Join(configDir, "config.yaml")

	// Create config directory
	err := os.MkdirAll(configDir, 0o755)
	require.NoError(t, err)

	// Create config file with custom base URL
	configContent := `api_key: test-api-key
base_url: https://custom-api.example.com/graphql`
	err = os.WriteFile(configPath, []byte(configContent), 0o600)
	require.NoError(t, err)

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cfg, err := LoadConfig()
	require.NoError(t, err)

	// Should preserve custom base URL from config file
	assert.Equal(t, "test-api-key", cfg.APIKey)
	assert.Equal(t, "https://custom-api.example.com/graphql", cfg.BaseURL)
}

// TestLoadConfig_EnvironmentVariablePreservesBaseURL tests that environment variable doesn't affect base URL
func TestLoadConfig_EnvironmentVariablePreservesBaseURL(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, ".hardcover")
	configPath := filepath.Join(configDir, "config.yaml")

	// Create config directory
	err := os.MkdirAll(configDir, 0o755)
	require.NoError(t, err)

	// Create config file with custom base URL
	configContent := `api_key: file-api-key
base_url: https://custom-api.example.com/graphql`
	err = os.WriteFile(configPath, []byte(configContent), 0o600)
	require.NoError(t, err)

	// Set environment variable
	envAPIKey := "env-api-key"
	os.Setenv("HARDCOVER_API_KEY", envAPIKey)
	defer os.Unsetenv("HARDCOVER_API_KEY")

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cfg, err := LoadConfig()
	require.NoError(t, err)

	// Environment variable should override API key but preserve base URL
	assert.Equal(t, envAPIKey, cfg.APIKey)
	assert.Equal(t, "https://custom-api.example.com/graphql", cfg.BaseURL)
}
