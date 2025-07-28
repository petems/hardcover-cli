package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"hardcover-cli/internal/client"
)

// GetCurrentUserResponse represents the response from the GetCurrentUser query
type GetCurrentUserResponse struct {
	Me *client.Users `json:"me"`
}

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

		query := `
			query GetCurrentUser {
				me {
					id
					username
					email
					name
					created_at
					updated_at
				}
			}
		`

		var response GetCurrentUserResponse
		if err := gqlClient.Execute(context.Background(), query, nil, &response); err != nil {
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
			printToStdoutf(cmd.OutOrStdout(), "  Created: %s\n", user.Created_at)
		}

		return nil
	},
}

// setupMeCommands registers the me command with the root command.
func setupMeCommands() {
	rootCmd.AddCommand(meCmd)
}
