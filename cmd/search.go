package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"hardcover-cli/internal/client"

	"github.com/spf13/cobra"
)

// Note: Using types from internal/client package instead of duplicating them here

// SearchUser represents a user from the search API.
type SearchUser struct {
	Image              interface{} `json:"image"`
	ID                 string      `json:"id"`
	Username           string      `json:"username"`
	Name               string      `json:"name"`
	Location           string      `json:"location"`
	Flair              string      `json:"flair"`
	BooksCount         int         `json:"books_count"`
	FollowersCount     int         `json:"followers_count"`
	FollowedUsersCount int         `json:"followed_users_count"`
	Pro                bool        `json:"pro"`
}

// SearchUsersResponse represents the response from the SearchUsers query
type SearchUsersResponse struct {
	Search struct {
		Query     string       `json:"query"`
		QueryType string       `json:"query_type"`
		IDs       []string     `json:"ids"`
		Results   []SearchUser `json:"results"`
		Page      int          `json:"page"`
		PerPage   int          `json:"per_page"`
	} `json:"search"`
}

// searchCmd represents the search command.
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for content on Hardcover.app",
	Long: `Search for books and users on Hardcover.app.
	
Available subcommands:
  books    Search for books by title, author, or other criteria
  users    Search for users by name, username, or location`,
}

// searchBooksCmd represents the search books command.
var searchBooksCmd = &cobra.Command{
	Use:   "books <query>",
	Short: "Search for books",
	Long: `Search for books based on title, author, or other criteria.
	
The search will return matching books with their:
- Title and subtitle
- Author names
- Publication year
- Edition ID and URL
- Rating and ratings count
- ISBNs and series information

Example:
  hardcover search books "golang programming"
  hardcover search books "tolkien"
  hardcover search books "machine learning"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, ok := getConfig(cmd.Context())
		if !ok {
			return errors.New("failed to get configuration")
		}

		if cfg.APIKey == "" {
			return errors.New("API key is required. Set it using:\n" +
				"  export HARDCOVER_API_KEY=\"your-api-key\"\n" +
				"  or\n" +
				"  hardcover config set-api-key \"your-api-key\"")
		}

		query := args[0]
		client := client.NewClient(cfg.BaseURL, cfg.APIKey)

		// Use GraphQL query as shown in the documentation
		const gqlQuery = `
			query SearchBooks($query: String!) {
				search(query: $query, query_type: "Book", per_page: 25, page: 1) {
					ids
					results
					query
					query_type
					page
					per_page
				}
			}
		`

		variables := map[string]interface{}{
			"query": query,
		}

		var response map[string]interface{}
		if err := client.Execute(context.Background(), gqlQuery, variables, &response); err != nil {
			return fmt.Errorf("failed to search books: %w", err)
		}

		if searchData, ok := response["search"].(map[string]interface{}); ok {
			if resultsMap, ok := searchData["results"].(map[string]interface{}); ok {
				if hits, ok := resultsMap["hits"].([]interface{}); ok && len(hits) > 0 {
					for i, hit := range hits {
						hitMap, ok := hit.(map[string]interface{})
						if !ok {
							continue
						}
						book, ok := hitMap["document"].(map[string]interface{})
						if !ok {
							continue
						}
						// Title
						if title, ok := book["title"].(string); ok {
							printToStdoutf(cmd.OutOrStdout(), "%d. %s\n", i+1, title)
						}
						// Subtitle
						if subtitle, ok := book["subtitle"].(string); ok && subtitle != "" {
							printToStdoutf(cmd.OutOrStdout(), "   Subtitle: %s\n", subtitle)
						}
						// Author names
						if authors, ok := book["author_names"].([]interface{}); ok && len(authors) > 0 {
							authorStrs := make([]string, 0, len(authors))
							for _, a := range authors {
								if s, ok := a.(string); ok {
									authorStrs = append(authorStrs, s)
								}
							}
							printToStdoutf(cmd.OutOrStdout(), "   Authors: %s\n", strings.Join(authorStrs, ", "))
						}
						// Year
						if year, ok := book["release_year"].(float64); ok && year > 0 {
							printToStdoutf(cmd.OutOrStdout(), "   Year: %.0f\n", year)
						}
						// Edition ID
						if id, ok := book["id"].(string); ok {
							printToStdoutf(cmd.OutOrStdout(), "   Edition ID: %s\n", id)
						}
						// Slug (URL)
						if slug, ok := book["slug"].(string); ok && slug != "" {
							printToStdoutf(cmd.OutOrStdout(), "   URL: https://hardcover.app/books/%s\n", slug)
						}
						// Rating
						if rating, ok := book["rating"].(float64); ok && rating > 0 {
							if ratingsCount, ok := book["ratings_count"].(float64); ok {
								printToStdoutf(
									cmd.OutOrStdout(),
									"   Rating: %.2f/5 (%.0f ratings)\n",
									rating,
									ratingsCount,
								)
							}
						}
						// ISBNs
						if isbns, ok := book["isbns"].([]interface{}); ok && len(isbns) > 0 {
							isbnStrs := make([]string, 0, len(isbns))
							for _, s := range isbns {
								if str, ok := s.(string); ok {
									isbnStrs = append(isbnStrs, str)
								}
							}
							printToStdoutf(cmd.OutOrStdout(), "   ISBNs: %s\n", strings.Join(isbnStrs, ", "))
						}
						// Series names
						if series, ok := book["series_names"].([]interface{}); ok && len(series) > 0 {
							seriesStrs := make([]string, 0, len(series))
							for _, s := range series {
								if str, ok := s.(string); ok {
									seriesStrs = append(seriesStrs, str)
								}
							}
							printToStdoutf(cmd.OutOrStdout(), "   Series: %s\n", strings.Join(seriesStrs, ", "))
						}
						printToStdoutLn(cmd.OutOrStdout())
						printToStdoutf(cmd.OutOrStdout(), "-----------------------------\n")
					}
					return nil
				}
			}
			printToStdoutf(cmd.OutOrStdout(), "No results found.\n")
			return nil
		}
		printToStdoutf(cmd.OutOrStdout(), "No results found or unexpected response.\n")
		return nil
	},
}

// searchUsersCmd represents the search users command.
var searchUsersCmd = &cobra.Command{
	Use:   "users <query>",
	Short: "Search for users",
	Long: `Search for users based on name, username, or location.
	
The search will return matching users with their:
- Username and name
- Location and custom flair
- Books count, followers, and following counts
- Pro supporter status
- Profile image availability

Example:
  hardcover search users "adam"
  hardcover search users "john smith"
  hardcover search users "new york"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, ok := getConfig(cmd.Context())
		if !ok {
			return errors.New("failed to get configuration")
		}

		if cfg.APIKey == "" {
			return errors.New("API key is required. Set it using:\n" +
				"  export HARDCOVER_API_KEY=\"your-api-key\"\n" +
				"  or\n" +
				"  hardcover config set-api-key \"your-api-key\"")
		}

		query := args[0]
		client := client.NewClient(cfg.BaseURL, cfg.APIKey)

		// Use GraphQL query for user search
		const gqlQuery = `
			query SearchUsers($query: String!) {
				search(query: $query, query_type: "User", per_page: 25, page: 1) {
					ids
					results
					query
					query_type
					page
					per_page
				}
			}
		`

		variables := map[string]interface{}{
			"query": query,
		}

		var response map[string]interface{}
		if err := client.Execute(context.Background(), gqlQuery, variables, &response); err != nil {
			return fmt.Errorf("failed to search users: %w", err)
		}

		if searchData, ok := response["search"].(map[string]interface{}); ok {
			if resultsMap, ok := searchData["results"].(map[string]interface{}); ok {
				if hits, ok := resultsMap["hits"].([]interface{}); ok && len(hits) > 0 {
					for i, hit := range hits {
						hitMap, ok := hit.(map[string]interface{})
						if !ok {
							continue
						}
						user, ok := hitMap["document"].(map[string]interface{})
						if !ok {
							continue
						}

						// Username
						if username, ok := user["username"].(string); ok {
							printToStdoutf(cmd.OutOrStdout(), "%d. %s\n", i+1, username)
						}

						// Name
						if name, ok := user["name"].(string); ok && name != "" {
							printToStdoutf(cmd.OutOrStdout(), "   Name: %s\n", name)
						}

						// Location
						if location, ok := user["location"].(string); ok && location != "" {
							printToStdoutf(cmd.OutOrStdout(), "   Location: %s\n", location)
						}

						// Flair
						if flair, ok := user["flair"].(string); ok && flair != "" {
							printToStdoutf(cmd.OutOrStdout(), "   Flair: %s\n", flair)
						}

						// Books count
						if booksCount, ok := user["books_count"].(float64); ok {
							printToStdoutf(cmd.OutOrStdout(), "   Books: %d\n", int(booksCount))
						}

						// Followers count
						if followersCount, ok := user["followers_count"].(float64); ok {
							printToStdoutf(cmd.OutOrStdout(), "   Followers: %d\n", int(followersCount))
						}

						// Following count
						if followedUsersCount, ok := user["followed_users_count"].(float64); ok {
							printToStdoutf(cmd.OutOrStdout(), "   Following: %d\n", int(followedUsersCount))
						}

						// Pro status
						if pro, ok := user["pro"].(bool); ok {
							if pro {
								printToStdoutf(cmd.OutOrStdout(), "   Pro: Yes\n")
							}
						}

						// Image
						if image, ok := user["image"]; ok && image != nil {
							printToStdoutf(cmd.OutOrStdout(), "   Has Image: Yes\n")
						}

						printToStdoutLn(cmd.OutOrStdout())
						printToStdoutf(cmd.OutOrStdout(), "-----------------------------\n")
					}
					return nil
				}
			}
			printToStdoutf(cmd.OutOrStdout(), "No results found.\n")
			return nil
		}
		printToStdoutf(cmd.OutOrStdout(), "No results found or unexpected response.\n")
		return nil
	},
}

// setupSearchCommands registers the search commands with the root command.
func setupSearchCommands() {
	searchCmd.AddCommand(searchBooksCmd)
	searchCmd.AddCommand(searchUsersCmd)
	rootCmd.AddCommand(searchCmd)
}
