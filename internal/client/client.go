// Package client provides HTTP client functionality for interacting with the Hardcover API.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents a GraphQL client for the Hardcover API.
type Client struct {
	endpoint   string
	apiKey     string
	httpClient *http.Client
}

// GraphQLRequest represents a GraphQL request.
type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// GraphQLResponse represents a GraphQL response.
type GraphQLResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []GraphQLError  `json:"errors,omitempty"`
}

// GraphQLError represents a GraphQL error.
type GraphQLError struct {
	Message   string                 `json:"message"`
	Locations []GraphQLErrorLocation `json:"locations,omitempty"`
}

// GraphQLErrorLocation represents the location of a GraphQL error.
type GraphQLErrorLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// Error implements the error interface for GraphQLError.
func (e GraphQLError) Error() string {
	return e.Message
}

// NewClient creates a new GraphQL client.
func NewClient(endpoint, apiKey string) *Client {
	return &Client{
		endpoint: endpoint,
		apiKey:   apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second, // 30 seconds is a reasonable timeout for HTTP requests
		},
	}
}

// Execute performs a GraphQL query and unmarshals the result.
func (c *Client) Execute(
	ctx context.Context,
	query string,
	variables map[string]interface{},
	result interface{},
) error {
	// Prepare the GraphQL request
	gqlReq := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	jsonData, err := json.Marshal(gqlReq)
	if err != nil {
		return fmt.Errorf("failed to marshal GraphQL request: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("User-Agent", "hardcover-cli/1.0.0")

	// Set authorization header if API key is provided
	if c.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	// Make the HTTP request
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer func() {
		if closeErr := httpResp.Body.Close(); closeErr != nil {
			// Log the error but don't fail the request
			// This is a common pattern for defer statements
			_ = closeErr // explicitly ignore the error
		}
	}()

	// Read response body
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for HTTP errors
	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error %d: %s", httpResp.StatusCode, string(body))
	}

	// Parse GraphQL response
	var gqlResp GraphQLResponse
	if unmarshalErr := json.Unmarshal(body, &gqlResp); unmarshalErr != nil {
		return fmt.Errorf("failed to unmarshal response: %w", unmarshalErr)
	}

	// Check for GraphQL errors
	if len(gqlResp.Errors) > 0 {
		return fmt.Errorf("GraphQL errors: %v", gqlResp.Errors)
	}

	// Unmarshal the data into the result if provided
	if result != nil {
		if dataErr := json.Unmarshal(gqlResp.Data, result); dataErr != nil {
			return fmt.Errorf("failed to unmarshal data: %w", dataErr)
		}
	}

	return nil
}
