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

func TestBookGetCmd_Success(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))

		// Verify GraphQL query
		var req client.GraphQLRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)
		assert.Contains(t, req.Query, "query GetBook")
		assert.Contains(t, req.Query, "book")
		assert.Equal(t, "book123", req.Variables["id"])

		// Send response
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{
				"book": {
					"id": "book123",
					"title": "The Go Programming Language",
					"description": "A comprehensive guide to Go programming",
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
						},
						{
							"name": "John Doe",
							"role": "editor"
						}
					],
					"cached_genres": [
						{
							"name": "Programming"
						},
						{
							"name": "Technology"
						},
						{
							"name": "Computer Science"
						}
					],
					"image": "https://example.com/book-cover.jpg",
					"averageRating": 4.5,
					"ratingsCount": 123,
					"createdAt": "2023-01-01T00:00:00Z",
					"updatedAt": "2023-01-02T00:00:00Z"
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
	err := bookGetCmd.RunE(cmd, []string{"book123"})
	require.NoError(t, err)

	// Verify output
	outputStr := output.String()
	assert.Contains(t, outputStr, "Book Details:")
	assert.Contains(t, outputStr, "Title: The Go Programming Language")
	assert.Contains(t, outputStr, "ID: book123")
	assert.Contains(t, outputStr, "Description: A comprehensive guide to Go programming")
	assert.Contains(t, outputStr, "Authors: Alan Donovan, Brian Kernighan")
	assert.Contains(t, outputStr, "Contributors: John Doe (editor)")
	assert.Contains(t, outputStr, "Published: 2015")
	assert.Contains(t, outputStr, "Pages: 380")
	assert.Contains(t, outputStr, "ISBN: 978-0134190440")
	assert.Contains(t, outputStr, "Genres: Programming, Technology, Computer Science")
	assert.Contains(t, outputStr, "Rating: 4.5/5 (123 ratings)")
	assert.Contains(t, outputStr, "Cover Image: https://example.com/book-cover.jpg")
	assert.Contains(t, outputStr, "URL: https://hardcover.app/books/go-programming-language")
	assert.Contains(t, outputStr, "Created: 2023-01-01T00:00:00Z")
	assert.Contains(t, outputStr, "Updated: 2023-01-02T00:00:00Z")
}

func TestBookGetCmd_MissingAPIKey(t *testing.T) {
	// Create command with empty API key
	cfg := &config.Config{
		APIKey:  "",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	ctx := withConfig(context.Background(), cfg)

	cmd := &cobra.Command{}
	cmd.SetContext(ctx)

	// Execute command
	err := bookGetCmd.RunE(cmd, []string{"book123"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "API key is required")
}

func TestBookGetCmd_BookNotFound(t *testing.T) {
	// Create test server that returns null for book
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{
				"book": null
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
	err := bookGetCmd.RunE(cmd, []string{"nonexistent"})
	require.NoError(t, err)

	// Verify output handles null book gracefully
	outputStr := output.String()
	assert.Contains(t, outputStr, "Book Details:")
	assert.Contains(t, outputStr, "Title: ")
	assert.Contains(t, outputStr, "ID: ")
}

func TestBookGetCmd_APIError(t *testing.T) {
	// Create test server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`null`),
			Errors: []client.GraphQLError{
				{
					Message: "Book not found",
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
	err := bookGetCmd.RunE(cmd, []string{"book123"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get book")
}

func TestBookGetCmd_MinimalData(t *testing.T) {
	// Create test server with minimal book data
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{
				"book": {
					"id": "book123",
					"title": "Simple Book",
					"slug": "simple-book",
					"cached_contributors": [],
					"cached_genres": [],
					"averageRating": 0,
					"ratingsCount": 0
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
	err := bookGetCmd.RunE(cmd, []string{"book123"})
	require.NoError(t, err)

	// Verify output contains minimal information
	outputStr := output.String()
	assert.Contains(t, outputStr, "Title: Simple Book")
	assert.Contains(t, outputStr, "ID: book123")
	assert.NotContains(t, outputStr, "Description:")
	assert.NotContains(t, outputStr, "Authors:")
	assert.NotContains(t, outputStr, "Contributors:")
	assert.NotContains(t, outputStr, "Published:")
	assert.NotContains(t, outputStr, "Pages:")
	assert.NotContains(t, outputStr, "ISBN:")
	assert.NotContains(t, outputStr, "Genres:")
	assert.NotContains(t, outputStr, "Rating:")
	assert.NotContains(t, outputStr, "Cover Image:")
	assert.NotContains(t, outputStr, "Created:")
	assert.NotContains(t, outputStr, "Updated:")
}

func TestBookGetCmd_OnlyAuthors(t *testing.T) {
	// Create test server with book that has only authors, no other contributors
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{
				"book": {
					"id": "book123",
					"title": "Authors Only Book",
					"slug": "authors-only-book",
					"cached_contributors": [
						{
							"name": "Author One",
							"role": "author"
						},
						{
							"name": "Author Two",
							"role": "Author"
						},
						{
							"name": "Author Three",
							"role": ""
						}
					],
					"cached_genres": []
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
	err := bookGetCmd.RunE(cmd, []string{"book123"})
	require.NoError(t, err)

	// Verify output shows authors but no contributors
	outputStr := output.String()
	assert.Contains(t, outputStr, "Authors: Author One, Author Two, Author Three")
	assert.NotContains(t, outputStr, "Contributors:")
}

func TestBookGetCmd_CommandProperties(t *testing.T) {
	// Test command properties
	assert.Equal(t, "get <book_id>", bookGetCmd.Use)
	assert.Equal(t, "Get detailed information about a specific book", bookGetCmd.Short)
	assert.NotEmpty(t, bookGetCmd.Long)
	assert.Contains(t, bookGetCmd.Long, "Retrieves and displays detailed information")
	assert.Contains(t, bookGetCmd.Long, "hardcover book get")
}

func TestBookGetCmd_RequiresArgument(t *testing.T) {
	// Test that the command requires exactly one argument
	cfg := &config.Config{
		APIKey:  "test-api-key",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
	ctx := withConfig(context.Background(), cfg)

	cmd := &cobra.Command{}
	cmd.SetContext(ctx)

	// Test with no arguments
	err := bookGetCmd.RunE(cmd, []string{})
	require.Error(t, err)

	// Test with too many arguments
	err = bookGetCmd.RunE(cmd, []string{"arg1", "arg2"})
	require.Error(t, err)
}

func TestBookCmd_CommandProperties(t *testing.T) {
	// Test book command properties
	assert.Equal(t, "book", bookCmd.Use)
	assert.Equal(t, "Manage and retrieve book information", bookCmd.Short)
	assert.NotEmpty(t, bookCmd.Long)
	assert.Contains(t, bookCmd.Long, "get")
}

func TestBookCmd_Integration(t *testing.T) {
	// Setup commands for testing
	setupBookCommands()
	
	// Test the command is properly registered
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use != "book" {
			continue
		}
		found = true
		// Check that get subcommand is registered
		getFound := false
		for _, subCmd := range cmd.Commands() {
			if subCmd.Use != "get <book_id>" {
				continue
			}
			getFound = true
			break
		}
		assert.True(t, getFound, "get subcommand should be registered")
		break
	}
	assert.True(t, found, "book command should be registered with root command")
}

func TestGetBookResponse_JSONUnmarshal(t *testing.T) {
	// Test JSON unmarshaling
	jsonData := `{
		"book": {
			"id": "book123",
			"title": "Test Book",
			"description": "A test book",
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
			"ratingsCount": 50,
			"createdAt": "2023-01-01T00:00:00Z",
			"updatedAt": "2023-01-02T00:00:00Z"
		}
	}`

	var response GetBookResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	require.NoError(t, err)

	book := response.Book
	assert.Equal(t, "book123", book.ID)
	assert.Equal(t, "Test Book", book.Title)
	assert.Equal(t, "A test book", book.Description)
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
	assert.Equal(t, "2023-01-01T00:00:00Z", book.CreatedAt)
	assert.Equal(t, "2023-01-02T00:00:00Z", book.UpdatedAt)
}
