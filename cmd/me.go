package cmd

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"hardcover-cli/internal/client"
)

// meCmd represents the me command
var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get the current user's profile information based on the API key",
	Long: `Get detailed information about the current user account.

This command retrieves your user profile information using your API key,
including:
- User ID
- Username  
- Email (if available)
- Display name (if available)
- Registration date
- Last update date

Example:
  hardcover-cli me`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cfg, configFound := getConfig(cmd.Context())
		if !configFound {
			return errors.New("failed to get configuration")
		}

		if cfg.APIKey == "" {
			return errors.New("API key is required. Set it using:\n" +
				"  hardcover config set-api-key <your-api-key>\n" +
				"  or\n" +
				"  export HARDCOVER_API_KEY=<your-api-key>")
		}

		gqlClient := client.NewClient(cfg.BaseURL, cfg.APIKey)

		response, err := gqlClient.GetCurrentUser(context.Background())
		if err != nil {
			return fmt.Errorf("failed to get user profile: %w", err)
		}

		// Display the user information
		printToStdoutf(cmd.OutOrStdout(), "User Profile:\n")

		if response.Me == nil {
			return fmt.Errorf("no user data received")
		}

		user := response.Me
		printToStdoutf(cmd.OutOrStdout(), "  ID: %d\n", user.ID)
		if user.Username != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Username: %s\n", user.Username)
		}
		if user.Email != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Email: %s\n", user.Email)
		}
		if user.Name != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Name: %s\n", user.Name)
		}
		if user.Bio != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Bio: %s\n", user.Bio)
		}
		if user.Location != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Location: %s\n", user.Location)
		}
		if user.Created_at != nil {
			printToStdoutf(cmd.OutOrStdout(), "  Created: %s\n", time.Time(*user.Created_at).Format("2006-01-02 15:04:05"))
		}

		return nil
	},
}

// setupMeCommands registers the me command with the root command.
func setupMeCommands() {
	rootCmd.AddCommand(meCmd)
}
