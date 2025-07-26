package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"hardcover-cli/internal/client"
)

// User represents the user structure from the API
type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// GetCurrentUserResponse represents the response from the GetCurrentUser query
type GetCurrentUserResponse struct {
	Me User `json:"me"`
}

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
			return fmt.Errorf("%s", apiKeyRequiredMsg)
		}

		client := client.NewClient(cfg.BaseURL, cfg.APIKey)

		const query = `
			query GetCurrentUser {
				me {
					id
					username
					email
					createdAt
					updatedAt
				}
			}
		`

		var response GetCurrentUserResponse
		if err := client.Execute(context.Background(), query, nil, &response); err != nil {
			return fmt.Errorf("failed to get user profile: %w", err)
		}

		// Display the user information
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "User Profile:\n")
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "  ID: %s\n", response.Me.ID)
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "  Username: %s\n", response.Me.Username)
		if response.Me.Email != "" {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "  Email: %s\n", response.Me.Email)
		}
		if response.Me.CreatedAt != "" {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "  Created: %s\n", response.Me.CreatedAt)
		}
		if response.Me.UpdatedAt != "" {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "  Updated: %s\n", response.Me.UpdatedAt)
		}

		return nil
	},
}

// setupMeCommand initializes the me command
func setupMeCommand() {
	rootCmd.AddCommand(meCmd)
}
