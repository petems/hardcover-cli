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
		cfg, ok := getConfig(cmd.Context())
		if !ok {
			return fmt.Errorf("failed to get configuration")
		}

		if cfg.APIKey == "" {
			return fmt.Errorf("API key is required. Set it using:\n  export HARDCOVER_API_KEY=\"your-api-key\"\n  or\n  hardcover config set-api-key \"your-api-key\"")
		}

		bookID := args[0]
		client := client.NewClient(cfg.BaseURL, cfg.APIKey)

		response, err := client.GetBook(context.Background(), bookID)
		if err != nil {
			return fmt.Errorf("failed to get book: %w", err)
		}

		book := response.GetBook()

		// Display detailed book information
		fmt.Printf("Book Details:\n")
		fmt.Printf("  Title: %s\n", book.GetTitle())
		fmt.Printf("  ID: %s\n", book.GetId())
		
		if book.GetDescription() != "" {
			fmt.Printf("  Description: %s\n", book.GetDescription())
		}
		
		if book.GetSlug() != "" {
			fmt.Printf("  Slug: %s\n", book.GetSlug())
		}
		
		if book.GetIsbn() != "" {
			fmt.Printf("  ISBN: %s\n", book.GetIsbn())
		}
		
		if book.GetPublicationYear() > 0 {
			fmt.Printf("  Publication Year: %d\n", book.GetPublicationYear())
		}
		
		if book.GetPageCount() > 0 {
			fmt.Printf("  Page Count: %d\n", book.GetPageCount())
		}

		// Display contributors
		contributors := book.GetCached_contributors()
		if len(contributors) > 0 {
			fmt.Printf("  Contributors:\n")
			for _, contributor := range contributors {
				role := contributor.GetRole()
				if role != "" {
					fmt.Printf("    - %s (%s)\n", contributor.GetName(), role)
				} else {
					fmt.Printf("    - %s\n", contributor.GetName())
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
			fmt.Printf("  Genres: %s\n", strings.Join(genreNames, ", "))
		}

		// Display ratings
		if book.GetRatingsCount() > 0 {
			fmt.Printf("  Average Rating: %.2f (%d ratings)\n", book.GetAverageRating(), book.GetRatingsCount())
		}

		// Display image URL
		if book.GetImage() != "" {
			fmt.Printf("  Image: %s\n", book.GetImage())
		}

		// Display creation/update timestamps
		if book.GetCreatedAt() != "" {
			fmt.Printf("  Created: %s\n", book.GetCreatedAt())
		}
		
		if book.GetUpdatedAt() != "" {
			fmt.Printf("  Updated: %s\n", book.GetUpdatedAt())
		}

		// Display Hardcover.app URL
		fmt.Printf("  Hardcover URL: https://hardcover.app/books/%s\n", book.GetSlug())

		return nil
	},
}

func init() {
	bookCmd.AddCommand(bookGetCmd)
	rootCmd.AddCommand(bookCmd)
}