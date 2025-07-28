package client_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"hardcover-cli/internal/client"
	"hardcover-cli/internal/testutil"
)

// BenchmarkClient_Original_vs_Optimized compares original and optimized client performance.
func BenchmarkClient_Original_vs_Optimized(b *testing.B) {
	// Create test server with standard response
	server := testutil.CreateTestServerWithHandler(func(w http.ResponseWriter, _ *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{"search": {"results": {"hits": []}}}`),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	defer server.Close()

	query := testGraphQLQuery

	variables := map[string]interface{}{
		"query": "golang",
	}

	b.Run("Original", func(b *testing.B) {
		c := client.NewClient(server.URL, "test-api-key")
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var result map[string]interface{}
			err := c.Execute(context.Background(), query, variables, &result)
			if err != nil {
				b.Fatalf("Execute failed: %v", err)
			}
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		c := client.NewOptimizedClient(server.URL, "test-api-key")
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var result map[string]interface{}
			err := c.Execute(context.Background(), query, variables, &result)
			if err != nil {
				b.Fatalf("Execute failed: %v", err)
			}
		}
	})
}

// BenchmarkClient_LargeResponse_Optimized compares performance with large responses.
func BenchmarkClient_LargeResponse_Optimized(b *testing.B) {
	// Create a large response with many books
	largeData := make([]interface{}, 50)
	for i := 0; i < 50; i++ {
		largeData[i] = map[string]interface{}{
			"document": map[string]interface{}{
				"id":            "book" + string(rune(i)),
				"title":         "Programming Book " + string(rune(i)),
				"subtitle":      "A comprehensive guide to programming",
				"author_names":  []interface{}{"Author 1", "Author 2", "Author 3"},
				"release_year":  2020 + i,
				"slug":          "programming-book-" + string(rune(i)),
				"rating":        4.5,
				"ratings_count": 123 + i,
				"isbns":         []interface{}{"978-0134190440", "978-0134190457"},
				"series_names":  []interface{}{"Programming Series", "Tech Series"},
			},
		}
	}

	responseData := map[string]interface{}{
		"search": map[string]interface{}{
			"results": map[string]interface{}{
				"hits": largeData,
			},
		},
	}

	server := testutil.CreateTestServerWithHandler(func(w http.ResponseWriter, _ *http.Request) {
		response := client.GraphQLResponse{
			Data: mustMarshal(responseData),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	defer server.Close()

	query := `query SearchBooks($query: String!) { search(query: $query) { results } }`
	variables := map[string]interface{}{"query": "programming"}

	b.Run("Original-Large", func(b *testing.B) {
		c := client.NewClient(server.URL, "test-api-key")
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var result map[string]interface{}
			err := c.Execute(context.Background(), query, variables, &result)
			if err != nil {
				b.Fatalf("Execute failed: %v", err)
			}
		}
	})

	b.Run("Optimized-Large", func(b *testing.B) {
		c := client.NewOptimizedClient(server.URL, "test-api-key")
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var result map[string]interface{}
			err := c.Execute(context.Background(), query, variables, &result)
			if err != nil {
				b.Fatalf("Execute failed: %v", err)
			}
		}
	})

	b.Run("Optimized-Streaming", func(b *testing.B) {
		c := client.NewOptimizedClient(server.URL, "test-api-key")
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			err := c.ExecuteStreaming(context.Background(), query, variables, func(_ json.RawMessage) error {
				// Simulate processing the data without full unmarshaling
				return nil
			})
			if err != nil {
				b.Fatalf("ExecuteStreaming failed: %v", err)
			}
		}
	})
}

// BenchmarkClient_ConcurrentRequests_Optimized benchmarks concurrent performance.
func BenchmarkClient_ConcurrentRequests_Optimized(b *testing.B) {
	server := testutil.CreateTestServerWithHandler(func(w http.ResponseWriter, _ *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{"search": {"results": {"hits": []}}}`),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	defer server.Close()

	query := `query SearchBooks($query: String!) { search(query: $query) { results } }`
	variables := map[string]interface{}{"query": "golang"}

	b.Run("Original-Concurrent", func(b *testing.B) {
		c := client.NewClient(server.URL, "test-api-key")
		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				var result map[string]interface{}
				err := c.Execute(context.Background(), query, variables, &result)
				if err != nil {
					b.Fatalf("Execute failed: %v", err)
				}
			}
		})
	})

	b.Run("Optimized-Concurrent", func(b *testing.B) {
		c := client.NewOptimizedClient(server.URL, "test-api-key")
		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				var result map[string]interface{}
				err := c.Execute(context.Background(), query, variables, &result)
				if err != nil {
					b.Fatalf("Execute failed: %v", err)
				}
			}
		})
	})
}
