package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"hardcover-cli/internal/config"
	"hardcover-cli/internal/contextutil"
)

var cfgFile string
var globalConfig *config.Config // Global config storage

// WithConfig injects configuration into context for testing.
func WithConfig(ctx context.Context, cfg *config.Config) context.Context {
	return contextutil.WithConfig(ctx, cfg)
}

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "hardcover-cli",
	Short: "A CLI tool for interacting with the Hardcover.app GraphQL API",
	Long: `Hardcover CLI is a command-line interface for interacting with the Hardcover.app GraphQL API.
It allows you to search for books and users, manage your profile, and configure API settings.

Before using the CLI, you need to set your Hardcover.app API key:

  # Set via environment variable
  export HARDCOVER_API_KEY="your-api-key-here"
  
  # Or set via config file
  hardcover config set-api-key "your-api-key-here"

Get your API key from: https://hardcover.app/account/developer

Available Commands:
  config    Manage configuration settings
  me        Get your user profile information
  search    Search for books and users
  help      Help about any command`,
}

// SetupCommands initializes all commands and their relationships.
func SetupCommands() {
	setupRootCommand()
	setupConfigCommands()
	setupMeCommands()
	setupSearchCommands()
}

// Execute runs the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// setupRootCommand configures the root command with flags and initialization.
func setupRootCommand() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hardcover/config.yaml)")
	rootCmd.PersistentFlags().String("api-key", "", "Hardcover API key (overrides config file)")

	// Run the setup commands when the package is loaded
	SetupCommands()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Check for API key flag override
	if apiKeyFlag, flagErr := rootCmd.PersistentFlags().GetString("api-key"); flagErr == nil && apiKeyFlag != "" {
		cfg.APIKey = apiKeyFlag
	}

	// Store globally for access in commands
	globalConfig = cfg
}

// getConfig retrieves the configuration with context support.
func getConfig(ctx context.Context) (*config.Config, bool) {
	// Check if context is cancelled
	select {
	case <-ctx.Done():
		return nil, false
	default:
	}

	// Add a timeout for config retrieval (useful for future async operations)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Check for cancellation again after timeout setup
	select {
	case <-ctx.Done():
		return nil, false
	default:
	}

	// First check if config is available in context (for testing)
	if cfg, ok := ctx.Value(contextutil.ConfigContextKey{}).(*config.Config); ok && cfg != nil {
		return cfg, true
	}

	// Return global config
	if globalConfig != nil {
		return globalConfig, true
	}
	return nil, false
}
