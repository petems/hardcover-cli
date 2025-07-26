package cmd

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"hardcover-cli/internal/config"
)

func TestConfigSetAPIKeyCmd_Success(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err := configSetAPIKeyCmd.RunE(cmd, []string{"test-api-key-123"})
	require.NoError(t, err)

	// Verify output
	outputStr := output.String()
	assert.Contains(t, outputStr, "API key has been set and saved to configuration file")
	assert.Contains(t, outputStr, "Configuration file:")

	// Verify config file was created and contains correct data
	configPath := filepath.Join(tempDir, ".hardcover", "config.yaml")
	_, err = os.Stat(configPath)
	require.NoError(t, err)

	// Load config and verify
	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, "test-api-key-123", cfg.APIKey)
}

func TestConfigSetAPIKeyCmd_RequiresArgument(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	// Test with no arguments
	err := configSetAPIKeyCmd.RunE(cmd, []string{})
	require.Error(t, err)

	// Test with too many arguments
	err = configSetAPIKeyCmd.RunE(cmd, []string{"arg1", "arg2"})
	require.Error(t, err)
}

func TestConfigGetAPIKeyCmd_WithAPIKey(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	// Create config with API key
	cfg := &config.Config{
		APIKey:  "test-api-key-123456789",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	err := config.SaveConfig(cfg)
	require.NoError(t, err)

	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err = configGetAPIKeyCmd.RunE(cmd, []string{})
	require.NoError(t, err)

	// Verify output shows masked API key
	outputStr := output.String()
	assert.Contains(t, outputStr, "API key: test...9789")
	assert.Contains(t, outputStr, "Source: Configuration file")
}

func TestConfigGetAPIKeyCmd_WithEnvironmentVariable(t *testing.T) {
	// Set environment variable
	expectedAPIKey := "env-api-key-123456789"
	os.Setenv("HARDCOVER_API_KEY", expectedAPIKey)
	defer os.Unsetenv("HARDCOVER_API_KEY")

	// Create temporary directory for config
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err := configGetAPIKeyCmd.RunE(cmd, []string{})
	require.NoError(t, err)

	// Verify output shows masked API key from environment
	outputStr := output.String()
	assert.Contains(t, outputStr, "API key: env-...9789")
	assert.Contains(t, outputStr, "Source: Environment variable")
}

func TestConfigGetAPIKeyCmd_NoAPIKey(t *testing.T) {
	// Make sure no environment variable is set
	os.Unsetenv("HARDCOVER_API_KEY")

	// Create temporary directory for config
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err := configGetAPIKeyCmd.RunE(cmd, []string{})
	require.NoError(t, err)

	// Verify output shows no API key message
	outputStr := output.String()
	assert.Contains(t, outputStr, "No API key is currently set")
	assert.Contains(t, outputStr, "hardcover config set-api-key")
	assert.Contains(t, outputStr, "export HARDCOVER_API_KEY")
}

func TestConfigGetAPIKeyCmd_ShortAPIKey(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	// Create config with short API key
	cfg := &config.Config{
		APIKey:  "short",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	err := config.SaveConfig(cfg)
	require.NoError(t, err)

	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err = configGetAPIKeyCmd.RunE(cmd, []string{})
	require.NoError(t, err)

	// Verify output shows full API key for short keys
	outputStr := output.String()
	assert.Contains(t, outputStr, "API key: short")
}

func TestConfigShowPathCmd_Success(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err := configShowPathCmd.RunE(cmd, []string{})
	require.NoError(t, err)

	// Verify output
	outputStr := output.String()
	expectedPath := filepath.Join(tempDir, ".hardcover", "config.yaml")
	assert.Contains(t, outputStr, "Configuration file path: "+expectedPath)
	assert.Contains(t, outputStr, "Configuration file does not exist yet")
}

func TestConfigShowPathCmd_FileExists(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	// Create config file
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	err := config.SaveConfig(cfg)
	require.NoError(t, err)

	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err = configShowPathCmd.RunE(cmd, []string{})
	require.NoError(t, err)

	// Verify output
	outputStr := output.String()
	expectedPath := filepath.Join(tempDir, ".hardcover", "config.yaml")
	assert.Contains(t, outputStr, "Configuration file path: "+expectedPath)
	assert.Contains(t, outputStr, "Configuration file exists")
	assert.NotContains(t, outputStr, "does not exist yet")
}

func TestConfigCmd_CommandProperties(t *testing.T) {
	// Test config command properties
	assert.Equal(t, "config", configCmd.Use)
	assert.Equal(t, "Manage configuration settings", configCmd.Short)
	assert.NotEmpty(t, configCmd.Long)
	assert.Contains(t, configCmd.Long, "set-api-key")
	assert.Contains(t, configCmd.Long, "get-api-key")
	assert.Contains(t, configCmd.Long, "show-path")
}

func TestConfigSetAPIKeyCmd_CommandProperties(t *testing.T) {
	// Test command properties
	assert.Equal(t, "set-api-key <api_key>", configSetAPIKeyCmd.Use)
	assert.Equal(t, "Set your Hardcover.app API key", configSetAPIKeyCmd.Short)
	assert.NotEmpty(t, configSetAPIKeyCmd.Long)
	assert.Contains(t, configSetAPIKeyCmd.Long, "~/.hardcover/config.yaml")
	assert.Contains(t, configSetAPIKeyCmd.Long, "hardcover config set-api-key")
}

func TestConfigGetAPIKeyCmd_CommandProperties(t *testing.T) {
	// Test command properties
	assert.Equal(t, "get-api-key", configGetAPIKeyCmd.Use)
	assert.Equal(t, "Display your current API key", configGetAPIKeyCmd.Short)
	assert.NotEmpty(t, configGetAPIKeyCmd.Long)
	assert.Contains(t, configGetAPIKeyCmd.Long, "HARDCOVER_API_KEY")
	assert.Contains(t, configGetAPIKeyCmd.Long, "hardcover config get-api-key")
}

func TestConfigShowPathCmd_CommandProperties(t *testing.T) {
	// Test command properties
	assert.Equal(t, "show-path", configShowPathCmd.Use)
	assert.Equal(t, "Show the path to the configuration file", configShowPathCmd.Short)
	assert.NotEmpty(t, configShowPathCmd.Long)
	assert.Contains(t, configShowPathCmd.Long, "hardcover config show-path")
}

func TestConfigCmd_Integration(t *testing.T) {
	// Test the command is properly registered
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "config" {
			found = true

			// Check that subcommands are registered
			subCommands := []string{"set-api-key <api_key>", "get-api-key", "show-path"}
			for _, expectedSub := range subCommands {
				subFound := false
				for _, subCmd := range cmd.Commands() {
					if subCmd.Use == expectedSub {
						subFound = true
						break
					}
				}
				assert.True(t, subFound, "subcommand %s should be registered", expectedSub)
			}
			break
		}
	}
	assert.True(t, found, "config command should be registered with root command")
}

func TestConfigSetAPIKeyCmd_UpdatesExistingConfig(t *testing.T) {
	// Create temporary directory for config
	tempDir := t.TempDir()

	// Mock the home directory for testing
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	// Create initial config
	cfg := &config.Config{
		APIKey:  "old-api-key",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	err := config.SaveConfig(cfg)
	require.NoError(t, err)

	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	// Execute command to update API key
	err = configSetAPIKeyCmd.RunE(cmd, []string{"new-api-key"})
	require.NoError(t, err)

	// Verify config was updated
	updatedCfg, err := config.LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, "new-api-key", updatedCfg.APIKey)
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", updatedCfg.BaseURL)
}
