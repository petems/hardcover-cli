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

		// Send response (new structure)
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{
				"search": {
					"results": {
						"hits": [
							{
								"document": {
									"id": "book1",
									"title": "Go Programming Language",
									"subtitle": "A Book About Go",
									"author_names": ["Alan Donovan", "Brian Kernighan"],
									"release_year": 2015,
									"slug": "go-programming-language",
									"rating": 4.5,
									"ratings_count": 123,
									"isbns": ["978-0134190440"],
									"series_names": ["Go Series"]
								}
							},
							{
								"document": {
									"id": "book2",
									"title": "Effective Go",
									"author_names": ["Go Team"],
									"release_year": 2020,
									"slug": "effective-go",
									"rating": 4.2,
									"ratings_count": 89
								}
							}
						]
					}
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
	assert.Contains(t, outputStr, "Go Programming Language")
	assert.Contains(t, outputStr, "A Book About Go")
	assert.Contains(t, outputStr, "Alan Donovan, Brian Kernighan")
	assert.Contains(t, outputStr, "2015")
	assert.Contains(t, outputStr, "Edition ID: book1")
	assert.Contains(t, outputStr, "https://hardcover.app/books/go-programming-language")
	assert.Contains(t, outputStr, "4.50/5 (123 ratings)")
	assert.Contains(t, outputStr, "978-0134190440")
	assert.Contains(t, outputStr, "Go Series")
	assert.Contains(t, outputStr, "Effective Go")
	assert.Contains(t, outputStr, "Go Team")
	assert.Contains(t, outputStr, "2020")
	assert.Contains(t, outputStr, "Edition ID: book2")
	assert.Contains(t, outputStr, "https://hardcover.app/books/effective-go")
	assert.Contains(t, outputStr, "4.20/5 (89 ratings)")
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
					"results": { "hits": [] }
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
	assert.Contains(t, outputStr, "No results found.")
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
					"results": {
						"hits": [
							{
								"document": {
									"id": "book1",
									"title": "Simple Book"
								}
							}
						]
					}
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

	// Test with no arguments - this should fail validation before reaching RunE
	err := searchBooksCmd.Args(cmd, []string{})
	require.Error(t, err)

	// Test with too many arguments - this should fail validation before reaching RunE
	err = searchBooksCmd.Args(cmd, []string{"arg1", "arg2"})
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
	// Setup commands for testing
	setupSearchCommands()

	// Test the command is properly registered
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use != "search" {
			continue
		}
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
	assert.True(t, found, "search command should be registered with root command")
}
