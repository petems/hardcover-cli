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

func TestSearchBooksCmd_Success(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))

		// Verify GraphQL query
		var req map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)
		assert.Contains(t, req["query"], "query SearchBooks")
		assert.Contains(t, req["query"], "search")
		variables := req["variables"].(map[string]interface{})
		assert.Equal(t, "golang", variables["query"])

		// Send response
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"search": map[string]interface{}{
					"error":      "",
					"ids":        []int{1, 2},
					"page":       1,
					"per_page":   10,
					"query":      "golang",
					"query_type": "Book",
					"results":    "[]",
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
	err := searchBooksCmd.RunE(cmd, []string{"golang"})
	require.NoError(t, err)

	// Verify output
	outputStr := output.String()
	assert.Contains(t, outputStr, "Search Results:")
	assert.Contains(t, outputStr, "Query: golang")
	assert.Contains(t, outputStr, "Query Type: Book")
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
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"search": map[string]interface{}{
					"error":      "",
					"ids":        []int{},
					"page":       1,
					"per_page":   10,
					"query":      "nonexistent",
					"query_type": "Book",
					"results":    "[]",
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
	err := searchBooksCmd.RunE(cmd, []string{"nonexistent"})
	require.NoError(t, err)

	// Verify output
	outputStr := output.String()
	assert.Contains(t, outputStr, "Search Results:")
	assert.Contains(t, outputStr, "Query: nonexistent")
}

func TestSearchBooksCmd_APIError(t *testing.T) {
	// Create test server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"data": nil,
			"errors": []map[string]interface{}{
				{
					"message": "Search failed",
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
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"search": map[string]interface{}{
					"error":      "",
					"ids":        []int{1},
					"page":       1,
					"per_page":   10,
					"query":      "simple",
					"query_type": "Book",
					"results":    "[]",
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
	err := searchBooksCmd.RunE(cmd, []string{"simple"})
	require.NoError(t, err)

	// Verify output contains minimal information
	outputStr := output.String()
	assert.Contains(t, outputStr, "Search Results:")
	assert.Contains(t, outputStr, "Query: simple")
}

func TestSearchBooksCmd_CommandProperties(t *testing.T) {
	// Test command properties
	assert.Equal(t, "books [query]", searchBooksCmd.Use)
	assert.Equal(t, "Search for books", searchBooksCmd.Short)
	assert.NotEmpty(t, searchBooksCmd.Long)
	assert.Contains(t, searchBooksCmd.Long, "hardcover search books")
}

func TestSearchBooksCmd_RequiresArgument(t *testing.T) {
	// Test that the command requires exactly one argument
	cmd := &cobra.Command{}

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
	SetupCommands()

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
			if subCmd.Use == "books [query]" {
				booksFound = true
				break
			}
		}
		assert.True(t, booksFound, "books subcommand should be registered")
		break
	}
	assert.True(t, found, "search command should be registered with root command")
}
