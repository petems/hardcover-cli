package client_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"hardcover-cli/internal/client"
	"hardcover-cli/internal/testutil"
)

// BenchmarkClient_Execute_JSONMarshaling benchmarks JSON marshaling operations
func BenchmarkClient_Execute_JSONMarshaling(b *testing.B) {
	// Create test server with simple response
	server := testutil.CreateTestServerWithHandler(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{"search": {"results": {"hits": []}}}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	defer server.Close()

	c := client.NewClient(server.URL, "test-api-key")

	query := `
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
		"query": "golang programming",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result map[string]interface{}
		err := c.Execute(context.Background(), query, variables, &result)
		if err != nil {
			b.Fatalf("Execute failed: %v", err)
		}
	}
}

// BenchmarkClient_Execute_LargeResponse benchmarks processing large API responses
func BenchmarkClient_Execute_LargeResponse(b *testing.B) {
	// Create a large response with many books
	largeData := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
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

	// Create test server with large response
	server := testutil.CreateTestServerWithHandler(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: mustMarshal(responseData),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	defer server.Close()

	c := client.NewClient(server.URL, "test-api-key")

	query := `
		query SearchBooks($query: String!) {
			search(query: $query, query_type: "Book", per_page: 100, page: 1) {
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
		"query": "programming",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result map[string]interface{}
		err := c.Execute(context.Background(), query, variables, &result)
		if err != nil {
			b.Fatalf("Execute failed: %v", err)
		}
	}
}

// BenchmarkClient_Execute_ConcurrentRequests benchmarks concurrent API calls
func BenchmarkClient_Execute_ConcurrentRequests(b *testing.B) {
	// Create test server
	server := testutil.CreateTestServerWithHandler(func(w http.ResponseWriter, r *http.Request) {
		response := client.GraphQLResponse{
			Data: json.RawMessage(`{"search": {"results": {"hits": []}}}`),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	defer server.Close()

	c := client.NewClient(server.URL, "test-api-key")

	query := `
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
		"query": "golang",
	}

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
}

// BenchmarkJSON_Operations benchmarks different JSON processing approaches
func BenchmarkJSON_Operations(b *testing.B) {
	sampleData := map[string]interface{}{
		"search": map[string]interface{}{
			"results": map[string]interface{}{
				"hits": []interface{}{
					map[string]interface{}{
						"document": map[string]interface{}{
							"id":            "book1",
							"title":         "Go Programming Language",
							"author_names":  []interface{}{"Alan Donovan", "Brian Kernighan"},
							"release_year":  2015,
							"rating":        4.5,
							"ratings_count": 123,
						},
					},
				},
			},
		},
	}

	b.Run("Marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := json.Marshal(sampleData)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Unmarshal", func(b *testing.B) {
		data, _ := json.Marshal(sampleData)
		for i := 0; i < b.N; i++ {
			var result map[string]interface{}
			err := json.Unmarshal(data, &result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// mustMarshal is a helper function for benchmark setup
func mustMarshal(v interface{}) json.RawMessage {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return json.RawMessage(data)
}