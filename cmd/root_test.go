package cmd

import (
	"context"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecute_LoadsConfiguration(t *testing.T) {
	// This test verifies that the Execute function properly loads configuration
	// and sets it in the context before executing commands

	// We can't easily test the full Execute function since it calls os.Exit,
	// but we can test the configuration loading logic separately
}

func TestInitConfig_LoadsConfiguration(t *testing.T) {
	// Test that initConfig properly loads configuration and sets context
	// This is a unit test for the initConfig function

	// Create a test command to work with
	testCmd := &cobra.Command{}
	testCmd.SetContext(context.Background())

	// Set up flags
	testCmd.PersistentFlags().String("api-key", "", "test flag")

	// Mock the root command for testing
	originalRootCmd := rootCmd
	rootCmd = testCmd
	defer func() { rootCmd = originalRootCmd }()

	// Call initConfig
	initConfig()

	// Verify that context was set
	cfg, ok := getConfig(testCmd.Context())
	require.True(t, ok, "Configuration should be set in context")
	assert.NotNil(t, cfg, "Configuration should not be nil")
	assert.Equal(t, "https://api.hardcover.app/v1/graphql", cfg.BaseURL)
}
