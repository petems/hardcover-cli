package cmd

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"hardcover-cli/internal/testutil"
)

func TestSearchBooksCmd_Success(t *testing.T) {
	// Setup test data
	searchData := map[string]interface{}{
		"search": map[string]interface{}{
			"results": map[string]interface{}{
				"hits": []interface{}{
					map[string]interface{}{
						"document": map[string]interface{}{
							"id":            "book1",
							"title":         "Go Programming Language",
							"subtitle":      "A Book About Go",
							"author_names":  []interface{}{"Alan Donovan", "Brian Kernighan"},
							"release_year":  2015,
							"slug":          "go-programming-language",
							"rating":        4.5,
							"ratings_count": 123,
							"isbns":         []interface{}{"978-0134190440"},
							"series_names":  []interface{}{"Go Series"},
						},
					},
					map[string]interface{}{
						"document": map[string]interface{}{
							"id":            "book2",
							"title":         "Effective Go",
							"author_names":  []interface{}{"Go Team"},
							"release_year":  2020,
							"slug":          "effective-go",
							"rating":        4.2,
							"ratings_count": 89,
						},
					},
				},
			},
		},
	}

	// Create test server
	server := testutil.CreateTestServer(t, testutil.SuccessResponse(searchData))
	defer server.Close()

	// Setup config and command
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})

	// Set up context with config
	ctx := testutil.WithTestConfigAdapter(context.Background(), cfg)
	searchBooksCmd.SetContext(ctx)

	// Set up output capture
	var output bytes.Buffer
	searchBooksCmd.SetOut(&output)

	// Execute command
	err := searchBooksCmd.RunE(searchBooksCmd, []string{"golang"})
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
	// Setup config with empty API key
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "",
		BaseURL: "https://api.hardcover.app/v1/graphql",
	})

	// Set up context with config
	ctx := testutil.WithTestConfigAdapter(context.Background(), cfg)
	searchBooksCmd.SetContext(ctx)

	// Execute command
	err := searchBooksCmd.RunE(searchBooksCmd, []string{"golang"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "API key is required")
}

func TestSearchBooksCmd_NoResults(t *testing.T) {
	// Setup test data with no results
	searchData := map[string]interface{}{
		"search": map[string]interface{}{
			"results": map[string]interface{}{
				"hits": []interface{}{},
			},
		},
	}

	// Create test server
	server := testutil.CreateTestServer(t, testutil.SuccessResponse(searchData))
	defer server.Close()

	// Setup config and command
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})

	// Set up context with config
	ctx := testutil.WithTestConfigAdapter(context.Background(), cfg)
	searchBooksCmd.SetContext(ctx)

	// Set up output capture
	var output bytes.Buffer
	searchBooksCmd.SetOut(&output)

	// Execute command
	err := searchBooksCmd.RunE(searchBooksCmd, []string{"nonexistent"})
	require.NoError(t, err)

	// Verify output shows no results message
	outputStr := output.String()
	assert.Contains(t, outputStr, "No results found")
}

func TestSearchBooksCmd_APIError(t *testing.T) {
	// Create test server that returns error
	errors := []testutil.GraphQLError{
		{Message: "Search failed"},
	}
	server := testutil.CreateTestServer(t, testutil.ErrorResponse(errors))
	defer server.Close()

	// Setup config and command
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})

	// Set up context with config
	ctx := testutil.WithTestConfigAdapter(context.Background(), cfg)
	searchBooksCmd.SetContext(ctx)

	// Execute command
	err := searchBooksCmd.RunE(searchBooksCmd, []string{"golang"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to search books")
}

func TestSearchBooksCmd_MinimalData(t *testing.T) {
	// Setup test data with minimal book information
	searchData := map[string]interface{}{
		"search": map[string]interface{}{
			"results": map[string]interface{}{
				"hits": []interface{}{
					map[string]interface{}{
						"document": map[string]interface{}{
							"id":    "book1",
							"title": "Simple Book",
						},
					},
				},
			},
		},
	}

	// Create test server
	server := testutil.CreateTestServer(t, testutil.SuccessResponse(searchData))
	defer server.Close()

	// Setup config and command
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})

	// Set up context with config
	ctx := testutil.WithTestConfigAdapter(context.Background(), cfg)
	searchBooksCmd.SetContext(ctx)

	// Set up output capture
	var output bytes.Buffer
	searchBooksCmd.SetOut(&output)

	// Execute command
	err := searchBooksCmd.RunE(searchBooksCmd, []string{"simple"})
	require.NoError(t, err)

	// Verify output shows minimal book information
	outputStr := output.String()
	assert.Contains(t, outputStr, "Simple Book")
	assert.Contains(t, outputStr, "Edition ID: book1")
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
	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

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
		break
	}
	assert.True(t, found, "search command should be registered with root command")
}
