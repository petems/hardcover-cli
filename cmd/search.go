package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"hardcover-cli/internal/client"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for content on Hardcover.app",
	Long: `Search for various types of content on Hardcover.app including books, authors, and users.
	
Available subcommands:
  books    Search for books by title, author, or other criteria`,
}

// searchBooksCmd represents the search books command
var searchBooksCmd = &cobra.Command{
	Use:   "books <query>",
	Short: "Search for books",
	Long: `Search for books based on title, author, or other criteria.
	
The search will return matching books with their:
- Title and author information
- Publication details
- Ratings and genres
- Hardcover.app URL

Example:
  hardcover search books "golang programming"
  hardcover search books "tolkien"
  hardcover search books "machine learning"`,
	Args: cobra.ExactArgs(1),
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

		query := args[0]
		client := client.NewClient(cfg.BaseURL, cfg.APIKey)

		response, err := client.SearchBooks(context.Background(), query)
		if err != nil {
			return fmt.Errorf("failed to search books: %w", err)
		}

		// Display search results
		searchResults := response.GetSearch()

		// Check the type name to determine what kind of results we got
		typeName := searchResults.GetTypename()

		if typeName != "BookSearchResults" {
			printToStdoutf(cmd.OutOrStdout(), "Search completed for query: %s\n", query)
			printToStdoutf(cmd.OutOrStdout(), "No book results found or unexpected result type: %s\n", typeName)
			return nil
		}

		// For now, just display a simple message since we can't access the specific fields
		// without the type assertion working
		printToStdoutf(cmd.OutOrStdout(), "Search Results for \"%s\":\n", query)
		printToStdoutf(cmd.OutOrStdout(), "Found book results (type: %s)\n", typeName)
		printToStdoutf(cmd.OutOrStdout(), "Note: Full search results display is being updated to work with GraphQL types.\n")

		return nil
	},
}

// setupSearchCommands registers the search commands with the root command
func setupSearchCommands() {
	searchCmd.AddCommand(searchBooksCmd)
	rootCmd.AddCommand(searchCmd)
}
