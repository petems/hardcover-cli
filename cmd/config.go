package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"hardcover-cli/internal/config"
)

// configCmd represents the config command.
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings",
	Long: `Manage configuration settings for the Hardcover CLI.

This command provides subcommands to manage API keys and configuration files:
- set-api-key: Set your Hardcover.app API key
- get-api-key: Display your current API key (masked)
- show-path: Show the path to the configuration file`,
}

// configSetAPIKeyCmd represents the config set-api-key command.
var configSetAPIKeyCmd = &cobra.Command{
	Use:   "set-api-key <api_key>",
	Short: "Set your Hardcover.app API key",
	Long: `Set your Hardcover.app API key and save it to the configuration file.

The API key will be saved to ~/.hardcover/config.yaml. You can also set the
API key using the HARDCOVER_API_KEY environment variable.

To get your API key, visit https://hardcover.app/account/developer

Example:
  hardcover config set-api-key your-api-key-here`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := args[0]

		// Load existing config or create new one
		cfg, err := config.LoadConfig()
		if err != nil {
			cfg = config.DefaultConfig()
		}

		// Set the API key
		cfg.APIKey = apiKey

		// Save the config
		if saveErr := config.SaveConfig(cfg); saveErr != nil {
			return fmt.Errorf("failed to save configuration: %w", saveErr)
		}

		printToStdoutLn(cmd.OutOrStdout(), "API key has been set and saved to configuration file.")

		// Show the configuration file path
		if configPath, pathErr := config.GetConfigPath(); pathErr == nil {
			printToStdoutf(cmd.OutOrStdout(), "Configuration file: %s\n", configPath)
		}

		return nil
	},
}

// configGetAPIKeyCmd represents the config get-api-key command.
var configGetAPIKeyCmd = &cobra.Command{
	Use:   "get-api-key",
	Short: "Display your current API key",
	Long: `Display your current API key (masked for security).

The API key can be set in two ways:
1. Configuration file: ~/.hardcover/config.yaml
2. Environment variable: HARDCOVER_API_KEY

Environment variables take precedence over configuration file settings.

Example:
  hardcover config get-api-key`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load configuration: %w", err)
		}

		if cfg.APIKey == "" {
			printToStdoutLn(cmd.OutOrStdout(), "No API key is currently set.")
			printToStdoutLn(cmd.OutOrStdout(), "")
			printToStdoutLn(cmd.OutOrStdout(), "You can set it using:")
			printToStdoutLn(cmd.OutOrStdout(), "  hardcover config set-api-key <your-api-key>")
			printToStdoutLn(cmd.OutOrStdout(), "  export HARDCOVER_API_KEY=<your-api-key>")
			return nil
		}

		// Mask the API key for security
		maskedKey := maskAPIKey(cfg.APIKey)
		printToStdoutf(cmd.OutOrStdout(), "API key: %s\n", maskedKey)

		// Determine source
		envKey := os.Getenv("HARDCOVER_API_KEY")
		if envKey != "" && envKey == cfg.APIKey {
			printToStdoutLn(cmd.OutOrStdout(), "Source: Environment variable (HARDCOVER_API_KEY)")
		} else {
			printToStdoutLn(cmd.OutOrStdout(), "Source: Configuration file")
		}

		return nil
	},
}

// configShowPathCmd represents the config show-path command.
var configShowPathCmd = &cobra.Command{
	Use:   "show-path",
	Short: "Show the path to the configuration file",
	Long: `Show the path to the configuration file.

The configuration file is located at ~/.hardcover/config.yaml

Example:
  hardcover config show-path`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		configPath, err := config.GetConfigPath()
		if err != nil {
			return fmt.Errorf("failed to get configuration path: %w", err)
		}

		printToStdoutf(cmd.OutOrStdout(), "Configuration file path: %s\n", configPath)

		// Check if file exists
		if _, statErr := os.Stat(configPath); os.IsNotExist(statErr) {
			printToStdoutLn(cmd.OutOrStdout(), "Configuration file does not exist yet.")
		} else {
			printToStdoutLn(cmd.OutOrStdout(), "Configuration file exists.")
		}

		return nil
	},
}

// setupConfigCommands registers the config commands with the root command.
func setupConfigCommands() {
	configCmd.AddCommand(configSetAPIKeyCmd)
	configCmd.AddCommand(configGetAPIKeyCmd)
	configCmd.AddCommand(configShowPathCmd)
	rootCmd.AddCommand(configCmd)
}

// maskAPIKey masks an API key for display, showing only first 4 and last 4 characters.
func maskAPIKey(apiKey string) string {
	if len(apiKey) <= 8 {
		return apiKey // Don't mask very short keys
	}
	return apiKey[:4] + "..." + apiKey[len(apiKey)-4:]
}
