package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// bookCmd represents the book command
var bookCmd = &cobra.Command{
	Use:   "book",
	Short: "Manage and retrieve book information",
	Long: `Commands for managing and retrieving book information from Hardcover.app.

Available subcommands:
  get      Get detailed information about a specific book`,
}

// bookGetCmd represents the book get command
var bookGetCmd = &cobra.Command{
	Use:   "get [book-id]",
	Short: "Get detailed information about a specific book",
	Long: `Retrieves and displays detailed information about a specific book from Hardcover.app.

The command will display:
- Book ID, title, and subtitle
- Description and slug
- Pages and release information
- Rating and ratings count
- Contributors and tags
- Creation/update timestamps

Example:
  hardcover book get 123`,
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

		// Use book ID as string
		bookID := args[0]

		// NOTE: Book details functionality uses manual HTTP requests because the GraphQL introspection schema
		// doesn't match the actual API structure. The API uses 'editions' queries with where clauses instead
		// of direct 'book(id: ID!)' queries. See: https://docs.hardcover.app/api/guides/gettingbookdetails/
		//
		// DO NOT REMOVE THIS MANUAL IMPLEMENTATION - it's intentionally bypassing the generated code
		// to work around the schema mismatch issue.
		book, err := performManualBookDetails(cfg.BaseURL, cfg.APIKey, bookID)
		if err != nil {
			return fmt.Errorf("failed to get book: %w", err)
		}

		printToStdoutf(cmd.OutOrStdout(), "Book Details:\n\n")

		printToStdoutf(cmd.OutOrStdout(), "ID: %s\n", book.Id)
		printToStdoutf(cmd.OutOrStdout(), "Title: %s\n", book.Title)

		if book.Description != "" {
			printToStdoutf(cmd.OutOrStdout(), "Description: %s\n", book.Description)
		}

		if book.Slug != "" {
			printToStdoutf(cmd.OutOrStdout(), "Slug: %s\n", book.Slug)
		}

		if book.Isbn10 != "" {
			printToStdoutf(cmd.OutOrStdout(), "ISBN-10: %s\n", book.Isbn10)
		}

		if book.Isbn13 != "" {
			printToStdoutf(cmd.OutOrStdout(), "ISBN-13: %s\n", book.Isbn13)
		}

		if book.Pages > 0 {
			printToStdoutf(cmd.OutOrStdout(), "Pages: %d\n", book.Pages)
		}

		if book.EditionFormat != "" {
			printToStdoutf(cmd.OutOrStdout(), "Format: %s\n", book.EditionFormat)
		}

		if book.ReleaseDate != "" {
			printToStdoutf(cmd.OutOrStdout(), "Release Date: %s\n", book.ReleaseDate)
		}

		if book.Publisher.Name != "" {
			printToStdoutf(cmd.OutOrStdout(), "Publisher: %s\n", book.Publisher.Name)
		}

		// Display contributions
		if len(book.Contributions) > 0 {
			printToStdoutf(cmd.OutOrStdout(), "Contributors:\n")
			for _, contribution := range book.Contributions {
				printToStdoutf(cmd.OutOrStdout(), "  - %s\n", contribution.Author.Name)
			}
		}

		return nil
	},
}

// setupBookCommands registers the book commands with the root command
func setupBookCommands() {
	bookCmd.AddCommand(bookGetCmd)
	rootCmd.AddCommand(bookCmd)
}
