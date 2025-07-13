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
	err := os.MkdirAll(configDir, 0755)
	require.NoError(t, err)
	
	// Create config file
	configContent := `api_key: test-api-key-from-file
base_url: https://api.hardcover.app/v1/graphql`
	err = os.WriteFile(configPath, []byte(configContent), 0600)
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
	err := os.MkdirAll(configDir, 0755)
	require.NoError(t, err)
	
	// Create config file with one API key
	configContent := `api_key: file-api-key
base_url: https://api.hardcover.app/v1/graphql`
	err = os.WriteFile(configPath, []byte(configContent), 0600)
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
	err := os.MkdirAll(configDir, 0755)
	require.NoError(t, err)
	
	// Create config file with invalid YAML
	configContent := `api_key: test-api-key
base_url: https://api.hardcover.app/v1/graphql
invalid_yaml: [unclosed bracket`
	err = os.WriteFile(configPath, []byte(configContent), 0600)
	require.NoError(t, err)
	
	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)
	
	_, err = LoadConfig()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse config file")
}