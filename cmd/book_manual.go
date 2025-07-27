//go:build dev

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// BookDetails represents the response from manual book details API
// This is used because the GraphQL introspection schema doesn't match the actual API structure
type BookDetails struct {
	Id            string         `json:"id"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	Slug          string         `json:"slug"`
	Isbn10        string         `json:"isbn_10"`
	Isbn13        string         `json:"isbn_13"`
	ReleaseDate   string         `json:"release_date"`
	Pages         int            `json:"pages"`
	EditionFormat string         `json:"edition_format"`
	Publisher     Publisher      `json:"publisher"`
	Contributions []Contribution `json:"contributions"`
}

type Publisher struct {
	Name string `json:"name"`
}

type Contribution struct {
	Author Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
}

// performManualBookDetails performs a book details query using direct HTTP requests to work around
// the GraphQL introspection schema issue where the API structure doesn't match what's exposed
func performManualBookDetails(endpoint, apiKey, bookID string) (*BookDetails, error) {
	// Create the GraphQL query based on the actual API structure from the documentation
	// The API uses editions queries with where clauses, not direct book(id:) queries
	graphqlQuery := fmt.Sprintf(`{
		"query": "query GetBookDetails($bookId: Int!) { editions(where: {id: {_eq: $bookId}}) { id title description slug isbn_10 isbn_13 release_date pages edition_format publisher { name } contributions { author { name } } } }",
		"variables": {
			"bookId": %s
		}
	}`, bookID)

	// Create HTTP request
	req, err := http.NewRequestWithContext(context.Background(), "POST", endpoint, strings.NewReader(graphqlQuery))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}

	// Make request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, fmt.Errorf("HTTP %d: failed to read response body: %w", resp.StatusCode, readErr)
		}
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var result map[string]interface{}
	if decodeErr := json.NewDecoder(resp.Body).Decode(&result); decodeErr != nil {
		return nil, fmt.Errorf("failed to decode response: %w", decodeErr)
	}

	// Check for GraphQL errors
	if errors, ok := result["errors"].([]interface{}); ok && len(errors) > 0 {
		return nil, fmt.Errorf("GraphQL errors: %v", errors)
	}

	// Extract data
	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format: missing data")
	}

	editionsData, ok := data["editions"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format: missing editions")
	}

	if len(editionsData) == 0 {
		return nil, fmt.Errorf("no book found with ID: %s", bookID)
	}

	// Convert the first edition to BookDetails
	edition := editionsData[0].(map[string]interface{})

	bookDetails := &BookDetails{}

	if id, ok := edition["id"].(float64); ok {
		bookDetails.Id = fmt.Sprintf("%.0f", id)
	}
	if title, ok := edition["title"].(string); ok {
		bookDetails.Title = title
	}
	if description, ok := edition["description"].(string); ok {
		bookDetails.Description = description
	}
	if slug, ok := edition["slug"].(string); ok {
		bookDetails.Slug = slug
	}
	if isbn10, ok := edition["isbn_10"].(string); ok {
		bookDetails.Isbn10 = isbn10
	}
	if isbn13, ok := edition["isbn_13"].(string); ok {
		bookDetails.Isbn13 = isbn13
	}
	if releaseDate, ok := edition["release_date"].(string); ok {
		bookDetails.ReleaseDate = releaseDate
	}
	if pages, ok := edition["pages"].(float64); ok {
		bookDetails.Pages = int(pages)
	}
	if editionFormat, ok := edition["edition_format"].(string); ok {
		bookDetails.EditionFormat = editionFormat
	}

	// Extract publisher
	if publisherData, ok := edition["publisher"].(map[string]interface{}); ok {
		if publisherName, ok := publisherData["name"].(string); ok {
			bookDetails.Publisher = Publisher{Name: publisherName}
		}
	}

	// Extract contributions
	if contributionsData, ok := edition["contributions"].([]interface{}); ok {
		for _, contrib := range contributionsData {
			if contribMap, ok := contrib.(map[string]interface{}); ok {
				if authorData, ok := contribMap["author"].(map[string]interface{}); ok {
					if authorName, ok := authorData["name"].(string); ok {
						bookDetails.Contributions = append(bookDetails.Contributions, Contribution{
							Author: Author{Name: authorName},
						})
					}
				}
			}
		}
	}

	return bookDetails, nil
}
