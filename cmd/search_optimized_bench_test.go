package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// BenchmarkStringProcessing compares current vs optimized string processing
func BenchmarkStringProcessing(b *testing.B) {
	// Sample book data for benchmarking
	sampleBook := map[string]interface{}{
		"id":            "book1",
		"title":         "Go Programming Language",
		"subtitle":      "A Comprehensive Guide to Go Programming",
		"author_names":  []interface{}{"Alan Donovan", "Brian Kernighan", "Rob Pike"},
		"release_year":  2015,
		"slug":          "go-programming-language",
		"rating":        4.5,
		"ratings_count": 123,
		"isbns":         []interface{}{"978-0134190440", "978-0134190457", "978-0134190464"},
		"series_names":  []interface{}{"Programming Series", "Go Series", "Computer Science"},
	}

	// Sample user data for benchmarking
	sampleUser := map[string]interface{}{
		"id":                   "user1",
		"username":             "gopher123",
		"name":                 "John Doe",
		"location":             "San Francisco, CA",
		"flair":                "Go Developer Extraordinaire",
		"books_count":          245,
		"followers_count":      1023,
		"followed_users_count": 456,
		"pro":                  true,
		"image":                "https://example.com/avatar.jpg",
	}

	b.Run("Book-Current", func(b *testing.B) {
		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			buf.Reset()
			// Simulate current book formatting logic
			formatBookCurrent(sampleBook, i, &buf)
		}
	})

	b.Run("Book-Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = optimizedFormatBookResult(sampleBook, i)
		}
	})

	b.Run("User-Current", func(b *testing.B) {
		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			buf.Reset()
			formatUserCurrent(sampleUser, i, &buf)
		}
	})

	b.Run("User-Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = optimizedFormatUserResult(sampleUser, i)
		}
	})
}

// BenchmarkBatchProcessing compares batch processing approaches
func BenchmarkBatchProcessing(b *testing.B) {
	// Create multiple sample hits
	hits := make([]interface{}, 10)
	for i := 0; i < 10; i++ {
		hits[i] = map[string]interface{}{
			"document": map[string]interface{}{
				"id":            "book" + string(rune(i+'0')),
				"title":         "Programming Book " + string(rune(i+'0')),
				"subtitle":      "A comprehensive guide to programming",
				"author_names":  []interface{}{"Author 1", "Author 2", "Author 3"},
				"release_year":  2020 + i,
				"slug":          "programming-book-" + string(rune(i+'0')),
				"rating":        4.5,
				"ratings_count": 123 + i,
				"isbns":         []interface{}{"978-0134190440", "978-0134190457"},
				"series_names":  []interface{}{"Programming Series", "Tech Series"},
			},
		}
	}

	b.Run("Sequential-Current", func(b *testing.B) {
		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			buf.Reset()
			for j, hit := range hits {
				if hitMap, ok := hit.(map[string]interface{}); ok {
					if book, ok := hitMap["document"].(map[string]interface{}); ok {
						formatBookCurrent(book, j, &buf)
					}
				}
			}
		}
	})

	b.Run("Batch-Optimized", func(b *testing.B) {
		processor := NewBatchProcessor(10)
		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			buf.Reset()
			processor.ProcessBooksBatch(hits, 0, &buf)
		}
	})
}

// BenchmarkOptimizedStringBuilder compares different string building approaches
func BenchmarkOptimizedStringBuilder(b *testing.B) {
	fields := []string{"Author 1", "Author 2", "Author 3", "Author 4", "Author 5"}
	
	b.Run("strings.Join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			interfaces := make([]interface{}, len(fields))
			for j, field := range fields {
				interfaces[j] = field
			}
			
			strs := make([]string, 0, len(interfaces))
			for _, iface := range interfaces {
				if s, ok := iface.(string); ok {
					strs = append(strs, s)
				}
			}
			_ = strings.Join(strs, ", ")
		}
	})
	
	b.Run("OptimizedStringBuilder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			interfaces := make([]interface{}, len(fields))
			for j, field := range fields {
				interfaces[j] = field
			}
			
			builder := NewOptimizedStringBuilder(100)
			builder.WriteStrings(interfaces, ", ")
			_ = builder.String()
		}
	})
	
	b.Run("OptimizedStringBuilder-Reuse", func(b *testing.B) {
		builder := NewOptimizedStringBuilder(100)
		for i := 0; i < b.N; i++ {
			builder.Reset()
			interfaces := make([]interface{}, len(fields))
			for j, field := range fields {
				interfaces[j] = field
			}
			
			builder.WriteStrings(interfaces, ", ")
			_ = builder.String()
		}
	})
}

// formatBookCurrent simulates the current book formatting logic for comparison
func formatBookCurrent(book map[string]interface{}, index int, buf *bytes.Buffer) {
	// Title
	if title, ok := book["title"].(string); ok {
		buf.WriteString(fmt.Sprintf("%d. %s\n", index+1, title))
	}
	
	// Subtitle
	if subtitle, ok := book["subtitle"].(string); ok && subtitle != "" {
		buf.WriteString(fmt.Sprintf("   Subtitle: %s\n", subtitle))
	}
	
	// Authors
	if authors, ok := book["author_names"].([]interface{}); ok && len(authors) > 0 {
		authorStrs := make([]string, 0, len(authors))
		for _, a := range authors {
			if s, ok := a.(string); ok {
				authorStrs = append(authorStrs, s)
			}
		}
		buf.WriteString(fmt.Sprintf("   Authors: %s\n", strings.Join(authorStrs, ", ")))
	}
	
	// Year
	if year, ok := book["release_year"].(float64); ok && year > 0 {
		buf.WriteString(fmt.Sprintf("   Year: %.0f\n", year))
	}
	
	// Other fields...
	buf.WriteString("-----------------------------\n")
}

// formatUserCurrent simulates the current user formatting logic for comparison
func formatUserCurrent(user map[string]interface{}, index int, buf *bytes.Buffer) {
	// Username
	if username, ok := user["username"].(string); ok {
		buf.WriteString(fmt.Sprintf("%d. %s\n", index+1, username))
	}
	
	// Name
	if name, ok := user["name"].(string); ok && name != "" {
		buf.WriteString(fmt.Sprintf("   Name: %s\n", name))
	}
	
	// Other fields...
	buf.WriteString("-----------------------------\n")
}