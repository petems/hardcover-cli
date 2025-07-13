package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"hardcover-cli/internal/client"
	"hardcover-cli/internal/config"
)

func TestSearchBooksCmd_Success(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		
		// Verify GraphQL query
		var req client.GraphQLRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)
		assert.Contains(t, req.Query, "query SearchBooks")
		assert.Contains(t, req.Query, "search")
		assert.Equal(t, "golang", req.Variables["query"])
		
		// Send response
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{
				"search": {
					"totalCount": 2,
					"results": [
						{
							"id": "book1",
							"title": "Go Programming Language",
							"slug": "go-programming-language",
							"isbn": "978-0134190440",
							"publicationYear": 2015,
							"pageCount": 380,
							"cached_contributors": [
								{
									"name": "Alan Donovan",
									"role": "author"
								},
								{
									"name": "Brian Kernighan",
									"role": "author"
								}
							],
							"cached_genres": [
								{
									"name": "Programming"
								},
								{
									"name": "Technology"
								}
							],
							"image": "https://example.com/book1.jpg",
							"averageRating": 4.5,
							"ratingsCount": 123
						},
						{
							"id": "book2",
							"title": "Effective Go",
							"slug": "effective-go",
							"isbn": "",
							"publicationYear": 2020,
							"pageCount": 250,
							"cached_contributors": [
								{
									"name": "Go Team",
									"role": "author"
								}
							],
							"cached_genres": [
								{
									"name": "Programming"
								}
							],
							"image": "",
							"averageRating": 4.2,
							"ratingsCount": 89
						}
					]
				}
			}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()
	
	// Create command with test context
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}
	ctx := withConfig(context.Background(), cfg)
	
	cmd := &cobra.Command{}
	cmd.SetContext(ctx)
	
	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)
	
	// Execute command
	err := searchBooksCmd.RunE(cmd, []string{"golang"})
	require.NoError(t, err)
	
	// Verify output
	outputStr := output.String()
	assert.Contains(t, outputStr, "Search Results for \"golang\":")
	assert.Contains(t, outputStr, "Found 2 books")
	assert.Contains(t, outputStr, "Go Programming Language")
	assert.Contains(t, outputStr, "Alan Donovan, Brian Kernighan")
	assert.Contains(t, outputStr, "Published: 2015")
	assert.Contains(t, outputStr, "Pages: 380")
	assert.Contains(t, outputStr, "Rating: 4.5/5 (123 ratings)")
	assert.Contains(t, outputStr, "Genres: Programming, Technology")
	assert.Contains(t, outputStr, "URL: https://hardcover.app/books/go-programming-language")
	assert.Contains(t, outputStr, "ID: book1")
	assert.Contains(t, outputStr, "Effective Go")
}

func TestSearchBooksCmd_MissingAPIKey(t *testing.T) {
	// Create command with empty API key
	cfg := &config.Config{
		APIKey:  "",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	ctx := withConfig(context.Background(), cfg)
	
	cmd := &cobra.Command{}
	cmd.SetContext(ctx)
	
	// Execute command
	err := searchBooksCmd.RunE(cmd, []string{"golang"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "API key is required")
}

func TestSearchBooksCmd_NoResults(t *testing.T) {
	// Create test server that returns no results
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{
				"search": {
					"totalCount": 0,
					"results": []
				}
			}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()
	
	// Create command with test context
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}
	ctx := withConfig(context.Background(), cfg)
	
	cmd := &cobra.Command{}
	cmd.SetContext(ctx)
	
	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)
	
	// Execute command
	err := searchBooksCmd.RunE(cmd, []string{"nonexistent"})
	require.NoError(t, err)
	
	// Verify output
	outputStr := output.String()
	assert.Contains(t, outputStr, "Found 0 books")
}

func TestSearchBooksCmd_APIError(t *testing.T) {
	// Create test server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`null`),
			Errors: []client.GraphQLError{
				{
					Message: "Search failed",
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()
	
	// Create command with test context
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}
	ctx := withConfig(context.Background(), cfg)
	
	cmd := &cobra.Command{}
	cmd.SetContext(ctx)
	
	// Execute command
	err := searchBooksCmd.RunE(cmd, []string{"golang"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to search books")
}

func TestSearchBooksCmd_MinimalData(t *testing.T) {
	// Create test server with minimal book data
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{
				"search": {
					"totalCount": 1,
					"results": [
						{
							"id": "book1",
							"title": "Simple Book",
							"slug": "simple-book",
							"cached_contributors": [],
							"cached_genres": [],
							"averageRating": 0,
							"ratingsCount": 0
						}
					]
				}
			}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()
	
	// Create command with test context
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}
	ctx := withConfig(context.Background(), cfg)
	
	cmd := &cobra.Command{}
	cmd.SetContext(ctx)
	
	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)
	
	// Execute command
	err := searchBooksCmd.RunE(cmd, []string{"simple"})
	require.NoError(t, err)
	
	// Verify output contains minimal information
	outputStr := output.String()
	assert.Contains(t, outputStr, "Simple Book")
	assert.Contains(t, outputStr, "ID: book1")
	assert.NotContains(t, outputStr, "Authors:")
	assert.NotContains(t, outputStr, "Published:")
	assert.NotContains(t, outputStr, "Pages:")
	assert.NotContains(t, outputStr, "Rating:")
	assert.NotContains(t, outputStr, "Genres:")
}

func TestSearchBooksCmd_CommandProperties(t *testing.T) {
	// Test command properties
	assert.Equal(t, "books <query>", searchBooksCmd.Use)
	assert.Equal(t, "Search for books", searchBooksCmd.Short)
	assert.NotEmpty(t, searchBooksCmd.Long)
	assert.Contains(t, searchBooksCmd.Long, "Search for books based on title, author")
	assert.Contains(t, searchBooksCmd.Long, "hardcover search books")
}

func TestSearchBooksCmd_RequiresArgument(t *testing.T) {
	// Test that the command requires exactly one argument
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	ctx := withConfig(context.Background(), cfg)
	
	cmd := &cobra.Command{}
	cmd.SetContext(ctx)
	
	// Test with no arguments
	err := searchBooksCmd.RunE(cmd, []string{})
	require.Error(t, err)
	
	// Test with too many arguments
	err = searchBooksCmd.RunE(cmd, []string{"arg1", "arg2"})
	require.Error(t, err)
}

func TestSearchCmd_CommandProperties(t *testing.T) {
	// Test search command properties
	assert.Equal(t, "search", searchCmd.Use)
	assert.Equal(t, "Search for content on Hardcover.app", searchCmd.Short)
	assert.NotEmpty(t, searchCmd.Long)
	assert.Contains(t, searchCmd.Long, "books")
}

func TestSearchCmd_Integration(t *testing.T) {
	// Test the command is properly registered
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "search" {
			found = true
			// Check that books subcommand is registered
			booksFound := false
			for _, subCmd := range cmd.Commands() {
				if subCmd.Use == "books <query>" {
					booksFound = true
					break
				}
			}
			assert.True(t, booksFound, "books subcommand should be registered")
			break
		}
	}
	assert.True(t, found, "search command should be registered with root command")
}

func TestSearchBooksResponse_JSONUnmarshal(t *testing.T) {
	// Test JSON unmarshaling
	jsonData := `{
		"search": {
			"totalCount": 1,
			"results": [
				{
					"id": "book1",
					"title": "Test Book",
					"slug": "test-book",
					"isbn": "978-1234567890",
					"publicationYear": 2023,
					"pageCount": 200,
					"cached_contributors": [
						{
							"name": "Test Author",
							"role": "author"
						}
					],
					"cached_genres": [
						{
							"name": "Test Genre"
						}
					],
					"image": "https://example.com/image.jpg",
					"averageRating": 4.0,
					"ratingsCount": 50
				}
			]
		}
	}`
	
	var response SearchBooksResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	require.NoError(t, err)
	
	assert.Equal(t, 1, response.Search.TotalCount)
	assert.Len(t, response.Search.Results, 1)
	
	book := response.Search.Results[0]
	assert.Equal(t, "book1", book.ID)
	assert.Equal(t, "Test Book", book.Title)
	assert.Equal(t, "test-book", book.Slug)
	assert.Equal(t, "978-1234567890", book.ISBN)
	assert.Equal(t, 2023, book.PublicationYear)
	assert.Equal(t, 200, book.PageCount)
	assert.Equal(t, "Test Author", book.CachedContributors[0].Name)
	assert.Equal(t, "author", book.CachedContributors[0].Role)
	assert.Equal(t, "Test Genre", book.CachedGenres[0].Name)
	assert.Equal(t, "https://example.com/image.jpg", book.Image)
	assert.Equal(t, 4.0, book.AverageRating)
	assert.Equal(t, 50, book.RatingsCount)
}