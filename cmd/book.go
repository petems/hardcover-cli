package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"hardcover-cli/internal/client"
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
	Use:   "get <book_id>",
	Short: "Get detailed information about a specific book",
	Long: `Retrieves and displays detailed information for a specific book by its ID.

The command will display:
- Book title and description
- Author(s) and contributors
- Publication details (year, page count, ISBN)
- Genres and categories
- Ratings and reviews summary
- Hardcover.app URL

Example:
  hardcover book get 12345
  hardcover book get "book-slug-or-id"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("book ID argument is required")
		}

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

		bookID := args[0]
		client := client.NewClient(cfg.BaseURL, cfg.APIKey)

		response, err := client.GetBook(context.Background(), bookID)
		if err != nil {
			return fmt.Errorf("failed to get book: %w", err)
		}

		book := response.GetBook()

		// Display detailed book information
		printToStdoutf(cmd.OutOrStdout(), "Book Details:\n")
		printToStdoutf(cmd.OutOrStdout(), "  Title: %s\n", book.GetTitle())
		printToStdoutf(cmd.OutOrStdout(), "  ID: %s\n", book.GetId())

		if book.GetDescription() != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Description: %s\n", book.GetDescription())
		}

		if book.GetSlug() != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Slug: %s\n", book.GetSlug())
		}

		if book.GetIsbn() != "" {
			printToStdoutf(cmd.OutOrStdout(), "  ISBN: %s\n", book.GetIsbn())
		}

		if book.GetPublicationYear() > 0 {
			printToStdoutf(cmd.OutOrStdout(), "  Publication Year: %d\n", book.GetPublicationYear())
		}

		if book.GetPageCount() > 0 {
			printToStdoutf(cmd.OutOrStdout(), "  Page Count: %d\n", book.GetPageCount())
		}

		// Display contributors
		contributors := book.GetCached_contributors()
		if len(contributors) > 0 {
			printToStdoutf(cmd.OutOrStdout(), "  Contributors:\n")
			for _, contributor := range contributors {
				role := contributor.GetRole()
				if role != "" {
					printToStdoutf(cmd.OutOrStdout(), "    - %s (%s)\n", contributor.GetName(), role)
				} else {
					printToStdoutf(cmd.OutOrStdout(), "    - %s\n", contributor.GetName())
				}
			}
		}

		// Display genres
		genres := book.GetCached_genres()
		if len(genres) > 0 {
			var genreNames []string
			for _, genre := range genres {
				genreNames = append(genreNames, genre.GetName())
			}
			printToStdoutf(cmd.OutOrStdout(), "  Genres: %s\n", strings.Join(genreNames, ", "))
		}

		// Display ratings
		if book.GetRatingsCount() > 0 {
			printToStdoutf(cmd.OutOrStdout(), "  Average Rating: %.2f (%d ratings)\n",
				book.GetAverageRating(), book.GetRatingsCount())
		}

		// Display image URL
		if book.GetImage() != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Image: %s\n", book.GetImage())
		}

		// Display creation/update timestamps
		if book.GetCreatedAt() != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Created: %s\n", book.GetCreatedAt())
		}

		if book.GetUpdatedAt() != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Updated: %s\n", book.GetUpdatedAt())
		}

		// Display Hardcover.app URL
		printToStdoutf(cmd.OutOrStdout(), "  Hardcover URL: https://hardcover.app/books/%s\n", book.GetSlug())

		return nil
	},
}

// setupBookCommands registers the book commands with the root command
func setupBookCommands() {
	bookCmd.AddCommand(bookGetCmd)
	rootCmd.AddCommand(bookCmd)
}
