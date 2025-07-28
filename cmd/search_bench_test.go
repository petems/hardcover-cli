package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"hardcover-cli/internal/testutil"
)

// BenchmarkSearchBooksCmd_ResponseProcessing benchmarks the processing of search results
func BenchmarkSearchBooksCmd_ResponseProcessing(b *testing.B) {
	// Setup test data with multiple books to simulate real API response
	searchData := map[string]interface{}{
		"search": map[string]interface{}{
			"results": map[string]interface{}{
				"hits": []interface{}{
					map[string]interface{}{
						"document": map[string]interface{}{
							"id":            "book1",
							"title":         "Go Programming Language",
							"subtitle":      "A Comprehensive Guide to Go",
							"author_names":  []interface{}{"Alan Donovan", "Brian Kernighan"},
							"release_year":  2015,
							"slug":          "go-programming-language",
							"rating":        4.5,
							"ratings_count": 123,
							"isbns":         []interface{}{"978-0134190440", "978-0134190457"},
							"series_names":  []interface{}{"Programming Series", "Go Series"},
						},
					},
					map[string]interface{}{
						"document": map[string]interface{}{
							"id":            "book2",
							"title":         "Effective Go",
							"subtitle":      "Best Practices for Go Development",
							"author_names":  []interface{}{"Go Team", "Rob Pike", "Ken Thompson"},
							"release_year":  2020,
							"slug":          "effective-go",
							"rating":        4.2,
							"ratings_count": 89,
							"isbns":         []interface{}{"978-1234567890"},
							"series_names":  []interface{}{"Go Series"},
						},
					},
					map[string]interface{}{
						"document": map[string]interface{}{
							"id":            "book3",
							"title":         "Concurrency in Go",
							"subtitle":      "Tools and Techniques for Developers",
							"author_names":  []interface{}{"Katherine Cox-Buday"},
							"release_year":  2017,
							"slug":          "concurrency-in-go",
							"rating":        4.7,
							"ratings_count": 234,
							"isbns":         []interface{}{"978-1491941195", "978-1491941201", "978-1491941218"},
							"series_names":  []interface{}{"O'Reilly Series", "Concurrency Series"},
						},
					},
					map[string]interface{}{
						"document": map[string]interface{}{
							"id":            "book4",
							"title":         "Learning Go",
							"subtitle":      "An Idiomatic Approach to Real-World Go Programming",
							"author_names":  []interface{}{"Jon Bodner"},
							"release_year":  2021,
							"slug":          "learning-go",
							"rating":        4.3,
							"ratings_count": 156,
							"isbns":         []interface{}{"978-1492077213"},
							"series_names":  []interface{}{"O'Reilly Series"},
						},
					},
					map[string]interface{}{
						"document": map[string]interface{}{
							"id":            "book5",
							"title":         "Go in Action",
							"subtitle":      "Programming Language in Practice",
							"author_names":  []interface{}{"William Kennedy", "Brian Ketelsen", "Erik St. Martin"},
							"release_year":  2015,
							"slug":          "go-in-action",
							"rating":        4.1,
							"ratings_count": 178,
							"isbns":         []interface{}{"978-1617291784", "978-1617291791"},
							"series_names":  []interface{}{"Manning Series", "Go Series"},
						},
					},
				},
			},
		},
	}

	// Create test server
	server := testutil.CreateTestServerWithHandler(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"data": searchData,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	defer server.Close()

	// Setup config and command
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})

	// Set up context with config
	ctx := testutil.WithTestConfigAdapter(context.Background(), cfg)

	// Reset timer to exclude setup time
	b.ResetTimer()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		// Set up fresh command for each iteration
		cmd := &cobra.Command{}
		cmd.SetContext(ctx)

		var output bytes.Buffer
		cmd.SetOut(&output)

		// Execute the actual function being benchmarked
		err := searchBooksCmd.RunE(cmd, []string{"golang"})
		if err != nil {
			b.Fatalf("Benchmark failed: %v", err)
		}
	}
}

// BenchmarkSearchUsersCmd_ResponseProcessing benchmarks user search response processing
func BenchmarkSearchUsersCmd_ResponseProcessing(b *testing.B) {
	// Setup test data with multiple users
	searchData := map[string]interface{}{
		"search": map[string]interface{}{
			"results": map[string]interface{}{
				"hits": []interface{}{
					map[string]interface{}{
						"document": map[string]interface{}{
							"id":                   "user1",
							"username":             "gopher123",
							"name":                 "John Doe",
							"location":             "San Francisco, CA",
							"flair":                "Go Developer",
							"books_count":          245,
							"followers_count":      1023,
							"followed_users_count": 456,
							"pro":                  true,
							"image":                "https://example.com/avatar1.jpg",
						},
					},
					map[string]interface{}{
						"document": map[string]interface{}{
							"id":                   "user2",
							"username":             "bookworm",
							"name":                 "Jane Smith",
							"location":             "New York, NY",
							"flair":                "Tech Book Enthusiast",
							"books_count":          412,
							"followers_count":      789,
							"followed_users_count": 234,
							"pro":                  false,
							"image":                nil,
						},
					},
					map[string]interface{}{
						"document": map[string]interface{}{
							"id":                   "user3",
							"username":             "programmer",
							"name":                 "Alex Johnson",
							"location":             "Seattle, WA",
							"flair":                "Software Engineer",
							"books_count":          156,
							"followers_count":      345,
							"followed_users_count": 123,
							"pro":                  true,
							"image":                "https://example.com/avatar3.jpg",
						},
					},
				},
			},
		},
	}

	// Create test server
	server := testutil.CreateTestServerWithHandler(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"data": searchData,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	defer server.Close()

	// Setup config and command
	cfg := testutil.SetupTestConfig(&testutil.TestConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})

	// Set up context with config
	ctx := testutil.WithTestConfigAdapter(context.Background(), cfg)

	// Reset timer to exclude setup time
	b.ResetTimer()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		// Set up fresh command for each iteration
		cmd := &cobra.Command{}
		cmd.SetContext(ctx)

		var output bytes.Buffer
		cmd.SetOut(&output)

		// Execute the actual function being benchmarked
		err := searchUsersCmd.RunE(cmd, []string{"john"})
		if err != nil {
			b.Fatalf("Benchmark failed: %v", err)
		}
	}
}

// BenchmarkStringBuilding benchmarks string operations used in search results
func BenchmarkStringBuilding(b *testing.B) {
	authors := []interface{}{"Alan Donovan", "Brian Kernighan", "Rob Pike", "Ken Thompson"}
	
	b.Run("Current-strings.Join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			authorStrs := make([]string, 0, len(authors))
			for _, a := range authors {
				if s, ok := a.(string); ok {
					authorStrs = append(authorStrs, s)
				}
			}
			_ = strings.Join(authorStrs, ", ")
		}
	})
	
	b.Run("Optimized-strings.Builder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var builder strings.Builder
			first := true
			for _, a := range authors {
				if s, ok := a.(string); ok {
					if !first {
						builder.WriteString(", ")
					}
					builder.WriteString(s)
					first = false
				}
			}
			_ = builder.String()
		}
	})
}