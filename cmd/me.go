package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"hardcover-cli/internal/client"
)

// GetCurrentUserResponse represents the response from the GetCurrentUser query
type GetCurrentUserResponse struct {
	Me interface{} `json:"me"` // Use interface{} to handle both single User and []User
}

// meCmd represents the me command
var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get the current user's profile information based on the API key",
	Long: `Fetches and displays the authenticated user's profile information including:
- User ID
- Username

Example:
  hardcover-cli me`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, ok := getConfig(cmd.Context())
		if !ok {
			return fmt.Errorf("failed to get configuration")
		}

		if cfg.APIKey == "" {
			return fmt.Errorf("API key is required. Set it using:\n" +
				"  export HARDCOVER_API_KEY=\"your-api-key\"\n" +
				"  or\n" +
				"  hardcover config set-api-key \"your-api-key\"")
		}

		client := client.NewClient(cfg.BaseURL, cfg.APIKey)

		const query = `
			query GetCurrentUser {
				me {
					id
					username
				}
			}
		`

		var response GetCurrentUserResponse
		if err := client.Execute(context.Background(), query, nil, &response); err != nil {
			return fmt.Errorf("failed to get user profile: %w", err)
		}

		// Display the user information
		printToStdoutf(cmd.OutOrStdout(), "User Profile:\n")

		// Handle the response which could be a single user or array
		switch me := response.Me.(type) {
		case map[string]interface{}:
			printToStdoutf(cmd.OutOrStdout(), "  ID: %v\n", me["id"])
			printToStdoutf(cmd.OutOrStdout(), "  Username: %s\n", me["username"])
		case []interface{}:
			if len(me) > 0 {
				if user, ok := me[0].(map[string]interface{}); ok {
					printToStdoutf(cmd.OutOrStdout(), "  ID: %v\n", user["id"])
					printToStdoutf(cmd.OutOrStdout(), "  Username: %s\n", user["username"])
				}
			}
		default:
			return fmt.Errorf("unexpected response format for user data")
		}

		return nil
	},
}

// setupMeCommands registers the me command with the root command
func setupMeCommands() {
	rootCmd.AddCommand(meCmd)
}
