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

	"hardcover-cli/internal/config"
)

func TestBookGetCmd_Success(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))

		// Verify GraphQL query
		var req map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)
		assert.Contains(t, req["query"], "query GetBook")
		assert.Contains(t, req["query"], "book")
		assert.Equal(t, "book123", req["variables"].(map[string]interface{})["id"])

		// Send response
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"book": map[string]interface{}{
					"id":              "book123",
					"title":           "The Go Programming Language",
					"description":     "A comprehensive guide to Go programming",
					"slug":            "go-programming-language",
					"isbn":            "978-0134190440",
					"publicationYear": 2015,
					"pageCount":       380,
					"cached_contributors": []map[string]interface{}{
						{
							"name": "Alan Donovan",
							"role": "author",
						},
						{
							"name": "Brian Kernighan",
							"role": "author",
						},
						{
							"name": "John Doe",
							"role": "editor",
						},
					},
					"cached_genres": []map[string]interface{}{
						{
							"name": "Programming",
						},
						{
							"name": "Technology",
						},
						{
							"name": "Computer Science",
						},
					},
					"image":         "https://example.com/book-cover.jpg",
					"averageRating": 4.5,
					"ratingsCount":  123,
					"createdAt":     "2023-01-01T00:00:00Z",
					"updatedAt":     "2023-01-02T00:00:00Z",
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
	assert.Contains(t, outputStr, "Contributors:")
	assert.Contains(t, outputStr, "Alan Donovan (author)")
	assert.Contains(t, outputStr, "Brian Kernighan (author)")
	assert.Contains(t, outputStr, "John Doe (editor)")
	assert.Contains(t, outputStr, "Publication Year: 2015")
	assert.Contains(t, outputStr, "Page Count: 380")
	assert.Contains(t, outputStr, "ISBN: 978-0134190440")
	assert.Contains(t, outputStr, "Genres: Programming, Technology, Computer Science")
	assert.Contains(t, outputStr, "Average Rating: 4.50 (123 ratings)")
	assert.Contains(t, outputStr, "Image: https://example.com/book-cover.jpg")
	assert.Contains(t, outputStr, "Hardcover URL: https://hardcover.app/books/go-programming-language")
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
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"book": nil,
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
		response := map[string]interface{}{
			"data": nil,
			"errors": []map[string]interface{}{
				{
					"message": "Book not found",
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
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"book": map[string]interface{}{
					"id":                  "book123",
					"title":               "Simple Book",
					"slug":                "simple-book",
					"cached_contributors": []map[string]interface{}{},
					"cached_genres":       []map[string]interface{}{},
					"averageRating":       0,
					"ratingsCount":        0,
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
	assert.NotContains(t, outputStr, "Contributors:")
	assert.NotContains(t, outputStr, "Publication Year:")
	assert.NotContains(t, outputStr, "Page Count:")
	assert.NotContains(t, outputStr, "ISBN:")
	assert.NotContains(t, outputStr, "Genres:")
	assert.NotContains(t, outputStr, "Average Rating:")
	assert.NotContains(t, outputStr, "Image:")
	assert.NotContains(t, outputStr, "Created:")
	assert.NotContains(t, outputStr, "Updated:")
}

func TestBookGetCmd_OnlyAuthors(t *testing.T) {
	// Create test server with book that has only authors, no other contributors
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"book": map[string]interface{}{
					"id":    "book123",
					"title": "Authors Only Book",
					"slug":  "authors-only-book",
					"cached_contributors": []map[string]interface{}{
						{
							"name": "Author One",
							"role": "author",
						},
						{
							"name": "Author Two",
							"role": "Author",
						},
						{
							"name": "Author Three",
							"role": "",
						},
					},
					"cached_genres": []map[string]interface{}{},
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

	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)

	// Execute command
	err := bookGetCmd.RunE(cmd, []string{"book123"})
	require.NoError(t, err)

	// Verify output shows authors but no contributors
	outputStr := output.String()
	assert.Contains(t, outputStr, "Contributors:")
	assert.Contains(t, outputStr, "Author One (author)")
	assert.Contains(t, outputStr, "Author Two (Author)")
	assert.Contains(t, outputStr, "Author Three")
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

	// Test with no arguments - this should fail validation before reaching RunE
	err := bookGetCmd.Args(cmd, []string{})
	require.Error(t, err)

	// Test with too many arguments - this should fail validation before reaching RunE
	err = bookGetCmd.Args(cmd, []string{"arg1", "arg2"})
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
			if subCmd.Use == "get <book_id>" {
				getFound = true
				break
			}
		}
		assert.True(t, getFound, "get subcommand should be registered")
		break
	}
	assert.True(t, found, "book command should be registered with root command")
}
