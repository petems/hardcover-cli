package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"hardcover-cli/internal/config"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hardcover",
	Short: "A CLI tool for interacting with the Hardcover.app GraphQL API",
	Long: `Hardcover CLI is a command-line interface for interacting with the Hardcover.app GraphQL API.
It allows you to search for books, get book details, and manage your profile.

Before using the CLI, you need to set your Hardcover.app API key:

  # Set via environment variable
  export HARDCOVER_API_KEY="your-api-key-here"
  
  # Or set via config file
  hardcover config set-api-key "your-api-key-here"

Examples:
  hardcover me                           # Get your user profile
  hardcover search books "golang"        # Search for books about golang
  hardcover book get 12345               # Get details for book with ID 12345`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hardcover/config.yaml)")
	rootCmd.PersistentFlags().StringP("api-key", "k", "", "Hardcover.app API key (can also be set via HARDCOVER_API_KEY environment variable)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to load config: %v\n", err)
		return
	}

	// Override with command-line flag if provided
	if apiKey, _ := rootCmd.PersistentFlags().GetString("api-key"); apiKey != "" {
		cfg.APIKey = apiKey
	}

	// Store config in a way that subcommands can access it
	rootCmd.SetContext(withConfig(rootCmd.Context(), cfg))
}