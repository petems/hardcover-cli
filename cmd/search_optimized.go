package cmd

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// OptimizedStringBuilder provides memory-efficient string concatenation.
type OptimizedStringBuilder struct {
	builder strings.Builder
}

// NewOptimizedStringBuilder creates a new string builder with initial capacity.
func NewOptimizedStringBuilder(capacity int) *OptimizedStringBuilder {
	osb := &OptimizedStringBuilder{}
	osb.builder.Grow(capacity)
	return osb
}

// WriteStrings efficiently concatenates multiple strings with a separator.
func (osb *OptimizedStringBuilder) WriteStrings(strs []interface{}, separator string) {
	first := true
	for _, s := range strs {
		if str, ok := s.(string); ok && str != "" {
			if !first {
				osb.builder.WriteString(separator)
			}
			osb.builder.WriteString(str)
			first = false
		}
	}
}

// WriteFormattedField writes a formatted field to the builder.
func (osb *OptimizedStringBuilder) WriteFormattedField(label, value string) {
	if value != "" {
		osb.builder.WriteString("   ")
		osb.builder.WriteString(label)
		osb.builder.WriteString(": ")
		osb.builder.WriteString(value)
		osb.builder.WriteString("\n")
	}
}

// WriteFormattedIntField writes a formatted integer field to the builder.
func (osb *OptimizedStringBuilder) WriteFormattedIntField(label string, value float64) {
	if value > 0 {
		osb.builder.WriteString("   ")
		osb.builder.WriteString(label)
		osb.builder.WriteString(": ")
		osb.builder.WriteString(strconv.Itoa(int(value)))
		osb.builder.WriteString("\n")
	}
}

// WriteFormattedFloatField writes a formatted float field to the builder.
func (osb *OptimizedStringBuilder) WriteFormattedFloatField(label string, value float64, precision int) {
	if value > 0 {
		osb.builder.WriteString("   ")
		osb.builder.WriteString(label)
		osb.builder.WriteString(": ")
		osb.builder.WriteString(strconv.FormatFloat(value, 'f', precision, 64))
		osb.builder.WriteString("\n")
	}
}

// Reset resets the builder for reuse.
func (osb *OptimizedStringBuilder) Reset() {
	osb.builder.Reset()
}

// String returns the built string.
func (osb *OptimizedStringBuilder) String() string {
	return osb.builder.String()
}

// WriteTo writes the built string to a writer.
func (osb *OptimizedStringBuilder) WriteTo(w io.Writer) (int64, error) {
	content := osb.builder.String()
	n, err := w.Write([]byte(content))
	return int64(n), err
}

// optimizedFormatBookResult formats a single book result with optimized string operations.
func optimizedFormatBookResult(book map[string]interface{}, index int) string {
	// Pre-allocate with estimated capacity to reduce reallocations
	builder := NewOptimizedStringBuilder(512)

	formatBookTitle(builder, book, index)
	formatBookMetadata(builder, book)
	formatBookIdentifiers(builder, book)
	formatBookCollections(builder, book)

	// Add separator
	builder.builder.WriteString("\n-----------------------------\n")

	return builder.String()
}

// formatBookTitle formats the title and subtitle of a book.
func formatBookTitle(builder *OptimizedStringBuilder, book map[string]interface{}, index int) {
	// Title
	if title, ok := book["title"].(string); ok {
		builder.builder.WriteString(strconv.Itoa(index + 1))
		builder.builder.WriteString(". ")
		builder.builder.WriteString(title)
		builder.builder.WriteString("\n")
	}

	// Subtitle
	if subtitle, ok := book["subtitle"].(string); ok {
		builder.WriteFormattedField("Subtitle", subtitle)
	}
}

// formatBookMetadata formats the metadata fields of a book.
func formatBookMetadata(builder *OptimizedStringBuilder, book map[string]interface{}) {
	// Authors - optimized string concatenation
	if authors, ok := book["author_names"].([]interface{}); ok && len(authors) > 0 {
		authorsBuilder := NewOptimizedStringBuilder(len(authors) * 20) // estimate 20 chars per author
		authorsBuilder.WriteStrings(authors, ", ")
		if authorsStr := authorsBuilder.String(); authorsStr != "" {
			builder.WriteFormattedField("Authors", authorsStr)
		}
	}

	// Year
	if year, ok := book["release_year"].(float64); ok {
		builder.WriteFormattedIntField("Year", year)
	}

	// Rating with optimized formatting
	if rating, ok := book["rating"].(float64); ok && rating > 0 {
		if ratingsCount, ok := book["ratings_count"].(float64); ok {
			ratingStr := fmt.Sprintf("%.2f/5 (%d ratings)", rating, int(ratingsCount))
			builder.WriteFormattedField("Rating", ratingStr)
		}
	}
}

// formatBookIdentifiers formats the identifier fields of a book.
func formatBookIdentifiers(builder *OptimizedStringBuilder, book map[string]interface{}) {
	// Edition ID
	if id, ok := book["id"].(string); ok {
		builder.WriteFormattedField("Edition ID", id)
	}

	// URL
	if slug, ok := book["slug"].(string); ok && slug != "" {
		url := "https://hardcover.app/books/" + slug
		builder.WriteFormattedField("URL", url)
	}

	// ISBNs
	if isbns, ok := book["isbns"].([]interface{}); ok && len(isbns) > 0 {
		isbnsBuilder := NewOptimizedStringBuilder(len(isbns) * 15) // estimate 15 chars per ISBN
		isbnsBuilder.WriteStrings(isbns, ", ")
		if isbnsStr := isbnsBuilder.String(); isbnsStr != "" {
			builder.WriteFormattedField("ISBNs", isbnsStr)
		}
	}
}

// formatBookCollections formats the collection fields of a book.
func formatBookCollections(builder *OptimizedStringBuilder, book map[string]interface{}) {
	// Series
	if series, ok := book["series_names"].([]interface{}); ok && len(series) > 0 {
		seriesBuilder := NewOptimizedStringBuilder(len(series) * 25) // estimate 25 chars per series
		seriesBuilder.WriteStrings(series, ", ")
		if seriesStr := seriesBuilder.String(); seriesStr != "" {
			builder.WriteFormattedField("Series", seriesStr)
		}
	}
}

// optimizedFormatUserResult formats a single user result with optimized string operations.
func optimizedFormatUserResult(user map[string]interface{}, index int) string {
	builder := NewOptimizedStringBuilder(256)

	// Username
	if username, ok := user["username"].(string); ok {
		builder.builder.WriteString(strconv.Itoa(index + 1))
		builder.builder.WriteString(". ")
		builder.builder.WriteString(username)
		builder.builder.WriteString("\n")
	}

	// Name
	if name, ok := user["name"].(string); ok {
		builder.WriteFormattedField("Name", name)
	}

	// Location
	if location, ok := user["location"].(string); ok {
		builder.WriteFormattedField("Location", location)
	}

	// Flair
	if flair, ok := user["flair"].(string); ok {
		builder.WriteFormattedField("Flair", flair)
	}

	// Books count
	if booksCount, ok := user["books_count"].(float64); ok {
		builder.WriteFormattedIntField("Books", booksCount)
	}

	// Followers count
	if followersCount, ok := user["followers_count"].(float64); ok {
		builder.WriteFormattedIntField("Followers", followersCount)
	}

	// Following count
	if followedUsersCount, ok := user["followed_users_count"].(float64); ok {
		builder.WriteFormattedIntField("Following", followedUsersCount)
	}

	// Pro status
	if pro, ok := user["pro"].(bool); ok && pro {
		builder.WriteFormattedField("Pro", "Yes")
	}

	// Image
	if image, ok := user["image"]; ok && image != nil {
		builder.WriteFormattedField("Has Image", "Yes")
	}

	// Add separator
	builder.builder.WriteString("\n-----------------------------\n")

	return builder.String()
}

// BatchProcessor provides optimized batch processing for search results.
type BatchProcessor struct {
	batchSize int
	buffer    *OptimizedStringBuilder
}

// NewBatchProcessor creates a new batch processor with specified batch size.
func NewBatchProcessor(batchSize int) *BatchProcessor {
	return &BatchProcessor{
		batchSize: batchSize,
		buffer:    NewOptimizedStringBuilder(batchSize * 512), // estimate 512 chars per result
	}
}

// ProcessBooksBatch processes a batch of book results efficiently.
func (bp *BatchProcessor) ProcessBooksBatch(hits []interface{}, startIndex int, writer io.Writer) error {
	bp.buffer.Reset()

	for i, hit := range hits {
		if hitMap, ok := hit.(map[string]interface{}); ok {
			if book, ok := hitMap["document"].(map[string]interface{}); ok {
				result := optimizedFormatBookResult(book, startIndex+i)
				bp.buffer.builder.WriteString(result)
			}
		}
	}

	_, err := bp.buffer.WriteTo(writer)
	return err
}

// ProcessUsersBatch processes a batch of user results efficiently.
func (bp *BatchProcessor) ProcessUsersBatch(hits []interface{}, startIndex int, writer io.Writer) error {
	bp.buffer.Reset()

	for i, hit := range hits {
		if hitMap, ok := hit.(map[string]interface{}); ok {
			if user, ok := hitMap["document"].(map[string]interface{}); ok {
				result := optimizedFormatUserResult(user, startIndex+i)
				bp.buffer.builder.WriteString(result)
			}
		}
	}

	_, err := bp.buffer.WriteTo(writer)
	return err
}
