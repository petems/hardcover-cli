package cmd

import (
	"context"
	"fmt"
	"strings"

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
			return fmt.Errorf("API key is required. Set it using:\n  export HARDCOVER_API_KEY=\"your-api-key\"\n  or\n  hardcover config set-api-key \"your-api-key\"")
		}

		query := args[0]
		client := client.NewClient(cfg.BaseURL, cfg.APIKey)

		response, err := client.SearchBooks(context.Background(), query)
		if err != nil {
			return fmt.Errorf("failed to search books: %w", err)
		}

		// Handle the union type - check typename field
		searchResults := response.GetSearch()
		
		// Use typename to determine the concrete type
		typename := searchResults.GetTypename()
		
		if typename == "BookSearchResults" {
			if bookResults, ok := searchResults.(*client.SearchBooksSearchBookSearchResults); ok {
				totalCount := bookResults.GetTotalCount()
				books := bookResults.GetResults()

				if totalCount == 0 {
					fmt.Printf("No books found for query: %s\n", query)
					return nil
				}

				fmt.Printf("Found %d books for query: %s\n\n", totalCount, query)

				// Display each book
				for i, book := range books {
					fmt.Printf("%d. %s\n", i+1, book.GetTitle())
					fmt.Printf("   ID: %s\n", book.GetId())
					
					// Display contributors
					contributors := book.GetCached_contributors()
					if len(contributors) > 0 {
						var contributorNames []string
						for _, contributor := range contributors {
							role := contributor.GetRole()
							if role != "" {
								contributorNames = append(contributorNames, fmt.Sprintf("%s (%s)", contributor.GetName(), role))
							} else {
								contributorNames = append(contributorNames, contributor.GetName())
							}
						}
						fmt.Printf("   Contributors: %s\n", strings.Join(contributorNames, ", "))
					}

					// Display publication details
					if book.GetPublicationYear() > 0 {
						fmt.Printf("   Published: %d\n", book.GetPublicationYear())
					}

					if book.GetPageCount() > 0 {
						fmt.Printf("   Pages: %d\n", book.GetPageCount())
					}

					if book.GetIsbn() != "" {
						fmt.Printf("   ISBN: %s\n", book.GetIsbn())
					}

					// Display genres
					genres := book.GetCached_genres()
					if len(genres) > 0 {
						var genreNames []string
						for _, genre := range genres {
							genreNames = append(genreNames, genre.GetName())
						}
						fmt.Printf("   Genres: %s\n", strings.Join(genreNames, ", "))
					}

					// Display ratings
					if book.GetRatingsCount() > 0 {
						fmt.Printf("   Rating: %.1f/5 (%d ratings)\n", book.GetAverageRating(), book.GetRatingsCount())
					}

					// Display URL
					if book.GetSlug() != "" {
						fmt.Printf("   URL: https://hardcover.app/books/%s\n", book.GetSlug())
					}

					// Add separator between books (except for the last one)
					if i < len(books)-1 {
						fmt.Println()
					}
				}
			} else {
				return fmt.Errorf("failed to cast to BookSearchResults")
			}
		} else {
			return fmt.Errorf("unexpected search result type: %s", typename)
		}

		return nil
	},
}

func init() {
	searchCmd.AddCommand(searchBooksCmd)
	rootCmd.AddCommand(searchCmd)
}