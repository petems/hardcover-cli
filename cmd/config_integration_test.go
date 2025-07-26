package cmd

import (
	"context"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"hardcover-cli/internal/config"
)

func TestWithConfig_AddsConfigurationToContext(t *testing.T) {
	// Test that withConfig properly adds configuration to context
	ctx := context.Background()
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: "https://test.example.com/graphql",
	}

	newCtx := withConfig(ctx, cfg)

	// Verify configuration can be retrieved
	retrievedCfg, ok := getConfig(newCtx)
	require.True(t, ok, "Configuration should be retrievable from context")
	assert.Equal(t, cfg.APIKey, retrievedCfg.APIKey)
	assert.Equal(t, cfg.BaseURL, retrievedCfg.BaseURL)
}

func TestGetConfig_RetrievesConfigurationFromContext(t *testing.T) {
	// Test that getConfig properly retrieves configuration from context
	ctx := context.Background()
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: "https://test.example.com/graphql",
	}

	ctxWithConfig := withConfig(ctx, cfg)

	// Test successful retrieval
	retrievedCfg, ok := getConfig(ctxWithConfig)
	require.True(t, ok, "Should successfully retrieve configuration")
	assert.Equal(t, cfg.APIKey, retrievedCfg.APIKey)
	assert.Equal(t, cfg.BaseURL, retrievedCfg.BaseURL)

	// Test retrieval from context without configuration
	emptyCtx := context.Background()
	_, ok = getConfig(emptyCtx)
	assert.False(t, ok, "Should return false when no configuration in context")
}

func TestGetConfig_ReturnsNilWhenNoConfiguration(t *testing.T) {
	// Test that getConfig returns nil when no configuration is in context
	ctx := context.Background()

	cfg, ok := getConfig(ctx)
	assert.False(t, ok, "Should return false when no configuration in context")
	assert.Nil(t, cfg, "Should return nil configuration when not found")
}

func TestSetupCommands_RegistersAllCommands(t *testing.T) {
	// Test that SetupCommands properly registers all commands
	// This ensures the command structure is set up correctly

	// Create a fresh root command for testing
	testRootCmd := &cobra.Command{Use: "test"}

	// Store original and restore after test
	originalRootCmd := rootCmd
	rootCmd = testRootCmd
	defer func() { rootCmd = originalRootCmd }()

	// Setup commands
	SetupCommands()

	// Verify that commands are registered
	expectedCommands := []string{"me", "search", "book", "config"}
	for _, expectedCmd := range expectedCommands {
		found := false
		for _, cmd := range testRootCmd.Commands() {
			if cmd.Use == expectedCmd {
				found = true
				break
			}
		}
		assert.True(t, found, "Command %s should be registered", expectedCmd)
	}
}

func TestSetupRootCommand_ConfiguresFlagsAndInitialization(t *testing.T) {
	// Test that setupRootCommand properly configures flags and initialization

	// Create a fresh root command for testing
	testRootCmd := &cobra.Command{Use: "test"}

	// Store original and restore after test
	originalRootCmd := rootCmd
	rootCmd = testRootCmd
	defer func() { rootCmd = originalRootCmd }()

	// Setup root command
	setupRootCommand()

	// Verify that persistent flags are configured
	configFlag := testRootCmd.PersistentFlags().Lookup("config")
	assert.NotNil(t, configFlag, "config flag should be configured")

	apiKeyFlag := testRootCmd.PersistentFlags().Lookup("api-key")
	assert.NotNil(t, apiKeyFlag, "api-key flag should be configured")

	// Verify that local flags are configured
	toggleFlag := testRootCmd.Flags().Lookup("toggle")
	assert.NotNil(t, toggleFlag, "toggle flag should be configured")
}
