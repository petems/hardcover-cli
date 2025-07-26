package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"hardcover-cli/internal/config"
)

func TestConfigSetAPIKeyCmd_Success(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "hardcover-test-")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Set the config directory to the temp directory
	configDir := filepath.Join(tempDir, ".hardcover")
	os.Setenv("XDG_CONFIG_HOME", tempDir)
	defer os.Unsetenv("XDG_CONFIG_HOME")

	// Create command
	cmd := &cobra.Command{}

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err = configSetAPIKeyCmd.RunE(cmd, []string{"test-api-key"})
	require.NoError(t, err)

	// Check output
	outputStr := output.String()
	assert.Contains(t, outputStr, "API key has been set and saved to configuration file")
	assert.Contains(t, outputStr, "Configuration file:")

	// Verify the API key was saved
	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, "test-api-key", cfg.APIKey)

	// Verify the config file was created
	configPath := filepath.Join(configDir, "config.yaml")
	_, err = os.Stat(configPath)
	require.NoError(t, err)
}

func TestConfigGetAPIKeyCmd_NoAPIKey(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "hardcover-test-")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Set the config directory to the temp directory
	os.Setenv("XDG_CONFIG_HOME", tempDir)
	defer os.Unsetenv("XDG_CONFIG_HOME")

	// Create command
	cmd := &cobra.Command{}

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err = configGetAPIKeyCmd.RunE(cmd, []string{})
	require.NoError(t, err)

	// Check output
	outputStr := output.String()
	assert.Contains(t, outputStr, "No API key is currently set")
	assert.Contains(t, outputStr, "Set it using:")
}

func TestConfigGetAPIKeyCmd_WithAPIKey(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "hardcover-test-")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Set the config directory to the temp directory
	os.Setenv("XDG_CONFIG_HOME", tempDir)
	defer os.Unsetenv("XDG_CONFIG_HOME")

	// Clear any existing environment variable
	os.Unsetenv("HARDCOVER_API_KEY")

	// Save a test API key
	cfg := config.DefaultConfig()
	cfg.APIKey = "test-api-key-123456"
	err = config.SaveConfig(cfg)
	require.NoError(t, err)

	// Create command
	cmd := &cobra.Command{}

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err = configGetAPIKeyCmd.RunE(cmd, []string{})
	require.NoError(t, err)

	// Check output contains masked API key
	outputStr := output.String()
	assert.Contains(t, outputStr, "API key: test***********3456")
	assert.Contains(t, outputStr, "Source: Configuration file")
}

func TestConfigGetAPIKeyCmd_WithEnvVar(t *testing.T) {
	// Set environment variable
	os.Setenv("HARDCOVER_API_KEY", "env-api-key-654321")
	defer os.Unsetenv("HARDCOVER_API_KEY")

	// Create command
	cmd := &cobra.Command{}

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err := configGetAPIKeyCmd.RunE(cmd, []string{})
	require.NoError(t, err)

	// Check output contains masked API key from environment
	outputStr := output.String()
	assert.Contains(t, outputStr, "API key: env-**********4321")
	assert.Contains(t, outputStr, "Source: Environment variable")
}

func TestConfigPathCmd_Success(t *testing.T) {
	// Create command
	cmd := &cobra.Command{}

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err := configShowPathCmd.RunE(cmd, []string{})
	require.NoError(t, err)

	// Check output
	outputStr := output.String()
	assert.Contains(t, outputStr, "Configuration file path:")
}

func TestConfigSetAPIKeyCmd_CommandProperties(t *testing.T) {
	assert.Equal(t, "set-api-key <api_key>", configSetAPIKeyCmd.Use)
	assert.Equal(t, "Set your Hardcover.app API key", configSetAPIKeyCmd.Short)
	assert.NotEmpty(t, configSetAPIKeyCmd.Long)
	assert.NotNil(t, configSetAPIKeyCmd.RunE)
}

func TestConfigGetAPIKeyCmd_CommandProperties(t *testing.T) {
	assert.Equal(t, "get-api-key", configGetAPIKeyCmd.Use)
	assert.Equal(t, "Display your current API key", configGetAPIKeyCmd.Short)
	assert.NotEmpty(t, configGetAPIKeyCmd.Long)
	assert.NotNil(t, configGetAPIKeyCmd.RunE)
}

func TestConfigPathCmd_CommandProperties(t *testing.T) {
	assert.Equal(t, "show-path", configShowPathCmd.Use)
	assert.Equal(t, "Show the path to the configuration file", configShowPathCmd.Short)
	assert.NotEmpty(t, configShowPathCmd.Long)
	assert.NotNil(t, configShowPathCmd.RunE)
}

func TestConfigCmd_CommandProperties(t *testing.T) {
	assert.Equal(t, "config", configCmd.Use)
	assert.Equal(t, "Manage configuration settings", configCmd.Short)
	assert.NotEmpty(t, configCmd.Long)
}
