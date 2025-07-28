package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"hardcover-cli/internal/client"
)

// meCmd represents the me command.
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

		// Handle response data
		if response.Me != nil {
			// Try to extract user data
			if userList, ok := response.Me.([]interface{}); ok && len(userList) > 0 {
				if userProfile, userOk := userList[0].(map[string]interface{}); userOk {
					printUserProfile(cmd.OutOrStdout(), userProfile)
					return nil
				}
			}

			// Try direct object approach
			if userMap, ok := response.Me.(map[string]interface{}); ok {
				printUserProfile(cmd.OutOrStdout(), userMap)
				return nil
			}
		}

		return errors.New("unexpected response format for user data")
	},
}

// GetCurrentUserResponse represents the GraphQL response for the current user query.
type GetCurrentUserResponse struct {
	Me interface{} `json:"me"`
}

// printUserProfile formats and prints user profile information.
func printUserProfile(w interface{ Write([]byte) (int, error) }, user map[string]interface{}) {
	printToStdoutLn(w, "User Profile:")
	printToStdoutLn(w, "=============")

	if id, ok := user["id"].(string); ok {
		printToStdoutf(w, "ID: %s\n", id)
	}

	if username, ok := user["username"].(string); ok {
		printToStdoutf(w, "Username: %s\n", username)
	}

	if email, ok := user["email"].(string); ok && email != "" {
		printToStdoutf(w, "Email: %s\n", email)
	}

	if name, ok := user["name"].(string); ok && name != "" {
		printToStdoutf(w, "Display Name: %s\n", name)
	}

	if createdAt, ok := user["created_at"].(string); ok && createdAt != "" {
		printToStdoutf(w, "Created: %s\n", createdAt)
	}

	if updatedAt, ok := user["updated_at"].(string); ok && updatedAt != "" {
		printToStdoutf(w, "Updated: %s\n", updatedAt)
	}
}

// setupMeCommands registers the me command with the root command.
func setupMeCommands() {
	rootCmd.AddCommand(meCmd)
}
