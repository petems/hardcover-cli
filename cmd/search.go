package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	defaultPerPage = 10
)

// SearchResult represents the response from the manual search API
// This is used because the GraphQL introspection schema doesn't properly expose search parameters
type SearchResult struct {
	Error     string `json:"error"`
	Ids       []int  `json:"ids"`
	Page      int    `json:"page"`
	PerPage   int    `json:"per_page"`
	Query     string `json:"query"`
	QueryType string `json:"query_type"`
	Results   string `json:"results"`
}

// performManualSearch performs a search using direct HTTP requests to work around
// the GraphQL introspection schema issue where search parameters are not properly exposed
func performManualSearch(endpoint, apiKey, query, queryType string, perPage, page int) (*SearchResult, error) {
	// Create the GraphQL query with the actual parameters that work at runtime
	graphqlQuery := fmt.Sprintf(`{
		"query": "query SearchBooks($query: String!, $query_type: String, $per_page: Int, $page: Int) { search(query: $query, query_type: $query_type, per_page: $per_page, page: $page) { error ids page per_page query query_type results } }",
		"variables": {
			"query": "%s",
			"query_type": "%s",
			"per_page": %d,
			"page": %d
		}
	}`, query, queryType, perPage, page)

	// Create HTTP request
	req, err := http.NewRequestWithContext(context.Background(), "POST", endpoint, strings.NewReader(graphqlQuery))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}

	// Make request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, fmt.Errorf("HTTP %d: failed to read response body: %w", resp.StatusCode, readErr)
		}
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var result map[string]interface{}
	if decodeErr := json.NewDecoder(resp.Body).Decode(&result); decodeErr != nil {
		return nil, fmt.Errorf("failed to decode response: %w", decodeErr)
	}

	// Check for GraphQL errors
	if errors, ok := result["errors"].([]interface{}); ok && len(errors) > 0 {
		return nil, fmt.Errorf("GraphQL errors: %v", errors)
	}

	// Extract search data
	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format: missing data")
	}

	searchData, ok := data["search"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format: missing search data")
	}

	// Convert to SearchResult
	searchResult := &SearchResult{}
	if errorVal, ok := searchData["error"].(string); ok {
		searchResult.Error = errorVal
	}
	if queryVal, ok := searchData["query"].(string); ok {
		searchResult.Query = queryVal
	}
	if queryTypeVal, ok := searchData["query_type"].(string); ok {
		searchResult.QueryType = queryTypeVal
	}
	if pageVal, ok := searchData["page"].(float64); ok {
		searchResult.Page = int(pageVal)
	}
	if perPageVal, ok := searchData["per_page"].(float64); ok {
		searchResult.PerPage = int(perPageVal)
	}
	if resultsVal, ok := searchData["results"].(string); ok {
		searchResult.Results = resultsVal
	}
	if idsVal, ok := searchData["ids"].([]interface{}); ok {
		for _, id := range idsVal {
			if idFloat, ok := id.(float64); ok {
				searchResult.Ids = append(searchResult.Ids, int(idFloat))
			}
		}
	}

	return searchResult, nil
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
	Use:   "books [query]",
	Short: "Search for books",
	Long: `Search for books on Hardcover.app.
	
The search will return search results with:
- Error information if any
- Book IDs found
- Pagination information
- Query details

Example:
  hardcover search books "lord of the rings"
  hardcover search books "harry potter" --query-type "Book" --per-page 5`,
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

		// Get search parameters
		query := args[0]

		// Initialize flags if not already done
		if cmd.Flags().Lookup("query-type") == nil {
			cmd.Flags().String("query-type", "Book",
				"Type of content to search for (Book, Author, Character, List, Prompt, Publisher, Series, User)")
			cmd.Flags().Int("per-page", defaultPerPage, "Number of results per page")
			cmd.Flags().Int("page", 1, "Page number to return")
		}

		queryType, err := cmd.Flags().GetString("query-type")
		if err != nil {
			return fmt.Errorf("failed to get query-type flag: %w", err)
		}
		perPage, err := cmd.Flags().GetInt("per-page")
		if err != nil {
			return fmt.Errorf("failed to get per-page flag: %w", err)
		}
		page, err := cmd.Flags().GetInt("page")
		if err != nil {
			return fmt.Errorf("failed to get page flag: %w", err)
		}

		// Set defaults
		if queryType == "" {
			queryType = "Book"
		}
		if perPage == 0 {
			perPage = 10
		}
		if page == 0 {
			page = 1
		}

		// NOTE: Search functionality uses manual HTTP requests because the GraphQL introspection schema
		// doesn't properly expose the search field parameters. The runtime API supports search with
		// parameters like query_type, per_page, page, sort, fields, weights, but these are not
		// reflected in the introspection schema. See: https://docs.hardcover.app/api/guides/searching/
		//
		// DO NOT REMOVE THIS MANUAL IMPLEMENTATION - it's intentionally bypassing the generated code
		// to work around the schema mismatch issue.
		searchResults, err := performManualSearch(cfg.BaseURL, cfg.APIKey, query, queryType, perPage, page)
		if err != nil {
			return fmt.Errorf("failed to search books: %w", err)
		}

		// Display search results
		printToStdoutf(cmd.OutOrStdout(), "Search Results:\n")

		if searchResults.Error != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Error: %s\n", searchResults.Error)
		}

		if searchResults.Query != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Query: %s\n", searchResults.Query)
		}

		if searchResults.QueryType != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Query Type: %s\n", searchResults.QueryType)
		}

		if searchResults.Page > 0 {
			printToStdoutf(cmd.OutOrStdout(), "  Page: %d\n", searchResults.Page)
		}

		if searchResults.PerPage > 0 {
			printToStdoutf(cmd.OutOrStdout(), "  Per Page: %d\n", searchResults.PerPage)
		}

		if len(searchResults.Ids) > 0 {
			printToStdoutf(cmd.OutOrStdout(), "  Book IDs found: %v\n", searchResults.Ids)
		} else {
			printToStdoutf(cmd.OutOrStdout(), "  No book IDs found\n")
		}

		if searchResults.Results != "" {
			printToStdoutf(cmd.OutOrStdout(), "  Results: %s\n", searchResults.Results)
		}

		return nil
	},
}

// setupSearchCommands registers the search commands with the root command
func setupSearchCommands() {
	// Check if flags are already added to prevent re-registration
	if searchBooksCmd.Flags().Lookup("query-type") == nil {
		// Add flags to searchBooksCmd
		searchBooksCmd.Flags().String("query-type", "Book",
			"Type of content to search for (Book, Author, Character, List, Prompt, Publisher, Series, User)")
		searchBooksCmd.Flags().Int("per-page", defaultPerPage, "Number of results per page")
		searchBooksCmd.Flags().Int("page", 1, "Page number to return")
	}

	// Check if command is already added to prevent re-registration
	searchAdded := false
	for _, cmd := range searchCmd.Commands() {
		if cmd.Use == searchBooksCmd.Use {
			searchAdded = true
			break
		}
	}
	if !searchAdded {
		searchCmd.AddCommand(searchBooksCmd)
	}

	rootAdded := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == searchCmd.Use {
			rootAdded = true
			break
		}
	}
	if !rootAdded {
		rootCmd.AddCommand(searchCmd)
	}
}
