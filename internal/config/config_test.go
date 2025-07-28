package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"hardcover-cli/internal/config"
	"hardcover-cli/internal/testutil"
)

func TestDefaultConfig(t *testing.T) {
	cfg := config.DefaultConfig()
	assert.NotNil(t, cfg)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
	assert.Empty(t, cfg.APIKey)
}

func TestLoadConfig_FromEnvironment(t *testing.T) {
	// Setup environment
	envMgr := testutil.NewEnvironmentManager(t)
	defer envMgr.Cleanup()

	expectedAPIKey := "test-api-key-from-env"
	envMgr.SetEnv("HARDCOVER_API_KEY", expectedAPIKey)

	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, expectedAPIKey, cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

func TestLoadConfig_FromFile(t *testing.T) {
	// Setup environment and temp directory
	envMgr := testutil.NewEnvironmentManager(t)
	defer envMgr.Cleanup()
	envMgr.UnsetEnv("HARDCOVER_API_KEY")

	tempDirMgr := testutil.NewTempDirManager(t)
	defer tempDirMgr.Cleanup()

	configDir := filepath.Join(tempDirMgr.GetTempDir(), ".hardcover")
	configPath := filepath.Join(configDir, "config.yaml")

	// Create config directory
	err := os.MkdirAll(configDir, 0o755)
	require.NoError(t, err)

	// Create config file
	configContent := `api_key: test-api-key-from-file
base_url: https://api.hardcover.app/v1/graphql`
	err = os.WriteFile(configPath, []byte(configContent), 0o600)
	require.NoError(t, err)

	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, "test-api-key-from-file", cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

func TestLoadConfig_NoFileExists(t *testing.T) {
	// Setup environment and temp directory
	envMgr := testutil.NewEnvironmentManager(t)
	defer envMgr.Cleanup()
	envMgr.UnsetEnv("HARDCOVER_API_KEY")

	tempDirMgr := testutil.NewTempDirManager(t)
	defer tempDirMgr.Cleanup()

	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	assert.Empty(t, cfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}

func TestSaveConfig(t *testing.T) {
	// Setup environment and temp directory
	envMgr := testutil.NewEnvironmentManager(t)
	defer envMgr.Cleanup()
	envMgr.UnsetEnv("HARDCOVER_API_KEY")

	tempDirMgr := testutil.NewTempDirManager(t)
	defer tempDirMgr.Cleanup()

	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}

	err := config.SaveConfig(cfg)
	require.NoError(t, err)

	// Verify file was created
	configPath := tempDirMgr.GetConfigPath()
	_, err = os.Stat(configPath)
	require.NoError(t, err)

	// Load the config back and verify
	loadedCfg, err := config.LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, cfg.APIKey, loadedCfg.APIKey)
	assert.Equal(t, cfg.BaseURL, loadedCfg.BaseURL)
}

func TestGetConfigPath(t *testing.T) {
	// Setup temp directory
	tempDirMgr := testutil.NewTempDirManager(t)
	defer tempDirMgr.Cleanup()

	configPath, err := config.GetConfigPath()
	require.NoError(t, err)

	expectedPath := tempDirMgr.GetConfigPath()
	assert.Equal(t, expectedPath, configPath)
}

func TestLoadConfig_EnvironmentOverridesFile(t *testing.T) {
	// Setup environment and temp directory
	envMgr := testutil.NewEnvironmentManager(t)
	defer envMgr.Cleanup()

	tempDirMgr := testutil.NewTempDirManager(t)
	defer tempDirMgr.Cleanup()

	configDir := filepath.Join(tempDirMgr.GetTempDir(), ".hardcover")
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
	envMgr.SetEnv("HARDCOVER_API_KEY", envAPIKey)

	cfg, err := config.LoadConfig()
	require.NoError(t, err)

	// Environment variable should override file
	assert.Equal(t, envAPIKey, cfg.APIKey)
}

func TestLoadConfig_InvalidYAML(t *testing.T) {
	// Setup environment and temp directory
	envMgr := testutil.NewEnvironmentManager(t)
	defer envMgr.Cleanup()
	envMgr.UnsetEnv("HARDCOVER_API_KEY")

	tempDirMgr := testutil.NewTempDirManager(t)
	defer tempDirMgr.Cleanup()

	configDir := filepath.Join(tempDirMgr.GetTempDir(), ".hardcover")
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

	_, err = config.LoadConfig()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse config file")
}
