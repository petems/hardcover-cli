package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"hardcover-cli/internal/config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings",
	Long: `Manage configuration settings for the Hardcover CLI application.

Available subcommands:
  set-api-key    Set your Hardcover.app API key
  get-api-key    Display your current API key
  show-path      Show the path to the configuration file`,
}

// configSetAPIKeyCmd represents the config set-api-key command
var configSetAPIKeyCmd = &cobra.Command{
	Use:   "set-api-key <api_key>",
	Short: "Set your Hardcover.app API key",
	Long: `Set your Hardcover.app API key for authenticating with the API.

The API key will be stored in a configuration file at:
  ~/.hardcover/config.yaml

Example:
  hardcover config set-api-key "your-api-key-here"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("API key is required")
		}
		if len(args) > 1 {
			return fmt.Errorf("too many arguments; expected exactly one API key")
		}

		apiKey := args[0]

		// Load existing config or create new one
		cfg, err := config.LoadConfig()
		if err != nil {
			cfg = config.DefaultConfig()
		}

		// Update the API key
		cfg.APIKey = apiKey

		// Save the configuration
		if err := config.SaveConfig(cfg); err != nil {
			return fmt.Errorf("failed to save configuration: %w", err)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "API key has been set and saved to configuration file.\n")

		configPath, err := config.GetConfigPath()
		if err == nil {
			fmt.Fprintf(cmd.OutOrStdout(), "Configuration file: %s\n", configPath)
		}

		return nil
	},
}

// configGetAPIKeyCmd represents the config get-api-key command
var configGetAPIKeyCmd = &cobra.Command{
	Use:   "get-api-key",
	Short: "Display your current API key",
	Long: `Display your current API key from the configuration file or environment variable.

The API key is loaded from:
1. HARDCOVER_API_KEY environment variable (if set)
2. Configuration file at ~/.hardcover/config.yaml

Example:
  hardcover config get-api-key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load configuration: %w", err)
		}

		if cfg.APIKey == "" {
			fmt.Fprintf(cmd.OutOrStdout(), "No API key is currently set.\n")
			fmt.Fprintf(cmd.OutOrStdout(), "Set it using:\n")
			fmt.Fprintf(cmd.OutOrStdout(), "  hardcover config set-api-key \"your-api-key\"\n")
			fmt.Fprintf(cmd.OutOrStdout(), "  or\n")
			fmt.Fprintf(cmd.OutOrStdout(), "  export HARDCOVER_API_KEY=\"your-api-key\"\n")
			return nil
		}

		// Show only the first and last few characters for security
		if len(cfg.APIKey) > 10 {
			masked := cfg.APIKey[:4] + "..." + cfg.APIKey[len(cfg.APIKey)-4:]
			fmt.Fprintf(cmd.OutOrStdout(), "API key: %s\n", masked)
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "API key: %s\n", cfg.APIKey)
		}

		// Show the source of the API key
		if os.Getenv("HARDCOVER_API_KEY") != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "Source: Environment variable (HARDCOVER_API_KEY)\n")
		} else {
			configPath, err := config.GetConfigPath()
			if err == nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Source: Configuration file (%s)\n", configPath)
			}
		}

		return nil
	},
}

// configShowPathCmd represents the config show-path command
var configShowPathCmd = &cobra.Command{
	Use:   "show-path",
	Short: "Show the path to the configuration file",
	Long: `Show the path to the configuration file where settings are stored.

Example:
  hardcover config show-path`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath, err := config.GetConfigPath()
		if err != nil {
			return fmt.Errorf("failed to get configuration path: %w", err)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Configuration file path: %s\n", configPath)

		// Check if file exists
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Fprintf(cmd.OutOrStdout(), "Configuration file does not exist yet.\n")
			fmt.Fprintf(cmd.OutOrStdout(), "It will be created when you set your API key.\n")
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "Configuration file exists.\n")
		}

		return nil
	},
}

func init() {
	configCmd.AddCommand(configSetAPIKeyCmd)
	configCmd.AddCommand(configGetAPIKeyCmd)
	configCmd.AddCommand(configShowPathCmd)
	rootCmd.AddCommand(configCmd)
}
