package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"hardcover-cli/internal/client"
)

// GetBookResponse represents the response from the GetBook query
type GetBookResponse struct {
	Book Book `json:"book"`
}

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

		const query = `
			query GetBook($id: ID!) {
				book(id: $id) {
					id
					title
					description
					slug
					isbn
					publicationYear
					pageCount
					cached_contributors {
						name
						role
					}
					cached_genres {
						name
					}
					image
					averageRating
					ratingsCount
					createdAt
					updatedAt
				}
			}
		`

		variables := map[string]interface{}{
			"id": bookID,
		}

		var response GetBookResponse
		if err := client.Execute(context.Background(), query, variables, &response); err != nil {
			return fmt.Errorf("failed to get book: %w", err)
		}

		book := response.Book

		// Display detailed book information
		fmt.Printf("Book Details:\n")
		fmt.Printf("  Title: %s\n", book.Title)
		fmt.Printf("  ID: %s\n", book.ID)
		
		if book.Description != "" {
			fmt.Printf("  Description: %s\n", book.Description)
		}

		// Display authors and contributors
		if len(book.CachedContributors) > 0 {
			var authors []string
			var otherContributors []string
			
			for _, contributor := range book.CachedContributors {
				if contributor.Role == "" || contributor.Role == "author" || contributor.Role == "Author" {
					authors = append(authors, contributor.Name)
				} else {
					otherContributors = append(otherContributors, fmt.Sprintf("%s (%s)", contributor.Name, contributor.Role))
				}
			}
			
			if len(authors) > 0 {
				fmt.Printf("  Authors: %s\n", strings.Join(authors, ", "))
			}
			
			if len(otherContributors) > 0 {
				fmt.Printf("  Contributors: %s\n", strings.Join(otherContributors, ", "))
			}
		}

		// Display publication details
		if book.PublicationYear > 0 {
			fmt.Printf("  Published: %d\n", book.PublicationYear)
		}

		if book.PageCount > 0 {
			fmt.Printf("  Pages: %d\n", book.PageCount)
		}

		if book.ISBN != "" {
			fmt.Printf("  ISBN: %s\n", book.ISBN)
		}

		// Display genres
		if len(book.CachedGenres) > 0 {
			var genres []string
			for _, genre := range book.CachedGenres {
				genres = append(genres, genre.Name)
			}
			fmt.Printf("  Genres: %s\n", strings.Join(genres, ", "))
		}

		// Display rating information
		if book.AverageRating > 0 {
			fmt.Printf("  Rating: %.1f/5 (%d ratings)\n", book.AverageRating, book.RatingsCount)
		}

		// Display image URL
		if book.Image != "" {
			fmt.Printf("  Cover Image: %s\n", book.Image)
		}

		// Display Hardcover URL
		fmt.Printf("  URL: https://hardcover.app/books/%s\n", book.Slug)

		// Display timestamps
		if book.CreatedAt != "" {
			fmt.Printf("  Created: %s\n", book.CreatedAt)
		}
		if book.UpdatedAt != "" {
			fmt.Printf("  Updated: %s\n", book.UpdatedAt)
		}

		return nil
	},
}

func init() {
	bookCmd.AddCommand(bookGetCmd)
	rootCmd.AddCommand(bookCmd)
}