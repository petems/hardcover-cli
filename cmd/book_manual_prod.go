//go:build !dev

package cmd

import (
	"fmt"
)

// Stub types for production builds
type BookDetails struct{}
type Publisher struct{}
type Contribution struct{}
type Author struct{}

// performManualBookDetails is not available in production builds
func performManualBookDetails(endpoint, apiKey, bookID string) (*BookDetails, error) {
	return nil, fmt.Errorf("manual book details functionality not available in production builds")
}
