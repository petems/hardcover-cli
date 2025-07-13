package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"hardcover-cli/internal/client"
)

// Contributor represents a book contributor
type Contributor struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

// Genre represents a book genre
type Genre struct {
	Name string `json:"name"`
}

// Book represents a book from the API
type Book struct {
	ID                  string        `json:"id"`
	Title               string        `json:"title"`
	Slug                string        `json:"slug"`
	ISBN                string        `json:"isbn"`
	PublicationYear     int           `json:"publicationYear"`
	PageCount           int           `json:"pageCount"`
	CachedContributors  []Contributor `json:"cached_contributors"`
	CachedGenres        []Genre       `json:"cached_genres"`
	Image               string        `json:"image"`
	AverageRating       float64       `json:"averageRating"`
	RatingsCount        int           `json:"ratingsCount"`
	Description         string        `json:"description"`
	CreatedAt           string        `json:"createdAt"`
	UpdatedAt           string        `json:"updatedAt"`
}

// BookSearchResults represents the search results for books
type BookSearchResults struct {
	Results    []Book `json:"results"`
	TotalCount int    `json:"totalCount"`
}

// SearchBooksResponse represents the response from the SearchBooks query
type SearchBooksResponse struct {
	Search BookSearchResults `json:"search"`
}

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

		const gqlQuery = `
			query SearchBooks($query: String!) {
				search(query: $query, type: BOOKS) {
					... on BookSearchResults {
						totalCount
						results {
							... on Book {
								id
								title
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
							}
						}
					}
				}
			}
		`

		variables := map[string]interface{}{
			"query": query,
		}

		var response SearchBooksResponse
		if err := client.Execute(context.Background(), gqlQuery, variables, &response); err != nil {
			return fmt.Errorf("failed to search books: %w", err)
		}

		// Display the search results
		fmt.Printf("Search Results for \"%s\":\n", query)
		fmt.Printf("Found %d books\n\n", response.Search.TotalCount)

		for i, book := range response.Search.Results {
			fmt.Printf("%d. %s\n", i+1, book.Title)
			
			// Display authors
			if len(book.CachedContributors) > 0 {
				var authors []string
				for _, contributor := range book.CachedContributors {
					if contributor.Role == "" || contributor.Role == "author" || contributor.Role == "Author" {
						authors = append(authors, contributor.Name)
					}
				}
				if len(authors) > 0 {
					fmt.Printf("   Authors: %s\n", strings.Join(authors, ", "))
				}
			}

			// Display publication year
			if book.PublicationYear > 0 {
				fmt.Printf("   Published: %d\n", book.PublicationYear)
			}

			// Display page count
			if book.PageCount > 0 {
				fmt.Printf("   Pages: %d\n", book.PageCount)
			}

			// Display rating
			if book.AverageRating > 0 {
				fmt.Printf("   Rating: %.1f/5 (%d ratings)\n", book.AverageRating, book.RatingsCount)
			}

			// Display genres
			if len(book.CachedGenres) > 0 {
				var genres []string
				for _, genre := range book.CachedGenres {
					genres = append(genres, genre.Name)
				}
				fmt.Printf("   Genres: %s\n", strings.Join(genres, ", "))
			}

			// Display Hardcover URL
			fmt.Printf("   URL: https://hardcover.app/books/%s\n", book.Slug)
			
			// Display ID for further queries
			fmt.Printf("   ID: %s\n", book.ID)
			
			if i < len(response.Search.Results)-1 {
				fmt.Println()
			}
		}

		return nil
	},
}

func init() {
	searchCmd.AddCommand(searchBooksCmd)
	rootCmd.AddCommand(searchCmd)
}