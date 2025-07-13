package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"hardcover-cli/internal/client"
)

// meCmd represents the me command
var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get your user profile information",
	Long: `Fetches and displays the authenticated user's profile information including:
- User ID
- Username
- Email address
- Account creation date
- Last updated date

Example:
  hardcover me`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, ok := getConfig(cmd.Context())
		if !ok {
			return fmt.Errorf("failed to get configuration")
		}

		if cfg.APIKey == "" {
			return fmt.Errorf("API key is required. Set it using:\n  export HARDCOVER_API_KEY=\"your-api-key\"\n  or\n  hardcover config set-api-key \"your-api-key\"")
		}

		client := client.NewClient(cfg.BaseURL, cfg.APIKey)

		response, err := client.GetCurrentUser(context.Background())
		if err != nil {
			return fmt.Errorf("failed to get user profile: %w", err)
		}

		// Display the user information using the generated types
		user := response.GetMe()
		fmt.Fprintf(cmd.OutOrStdout(), "User Profile:\n")
		fmt.Fprintf(cmd.OutOrStdout(), "  ID: %s\n", user.GetId())
		fmt.Fprintf(cmd.OutOrStdout(), "  Username: %s\n", user.GetUsername())
		if user.GetEmail() != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "  Email: %s\n", user.GetEmail())
		}
		if user.GetCreatedAt() != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "  Created: %s\n", user.GetCreatedAt())
		}
		if user.GetUpdatedAt() != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "  Updated: %s\n", user.GetUpdatedAt())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(meCmd)
}