// Package client provides HTTP client functionality for interacting with the Hardcover API.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// OptimizedClient represents an optimized GraphQL client for the Hardcover API.
type OptimizedClient struct {
	endpoint   string
	apiKey     string
	httpClient *http.Client
	bufferPool sync.Pool
}

// NewOptimizedClient creates a new optimized GraphQL client with performance enhancements.
func NewOptimizedClient(endpoint, apiKey string) *OptimizedClient {
	return &OptimizedClient{
		endpoint: endpoint,
		apiKey:   apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
				WriteBufferSize:     4096,
				ReadBufferSize:      4096,
			},
		},
		bufferPool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 0, 1024))
			},
		},
	}
}

// Execute performs an optimized GraphQL query with better memory management.
func (c *OptimizedClient) Execute(
	ctx context.Context,
	query string,
	variables map[string]interface{},
	result interface{},
) error {
	// Get buffer from pool
	buf := c.bufferPool.Get().(*bytes.Buffer) //nolint:errcheck // sync.Pool.Get() always succeeds
	defer func() {
		buf.Reset()
		c.bufferPool.Put(buf)
	}()

	// Prepare the GraphQL request using streaming JSON encoder
	gqlReq := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	// Use the buffer for encoding
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(gqlReq); err != nil {
		return fmt.Errorf("failed to marshal GraphQL request: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, buf)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("User-Agent", "hardcover-cli/1.0.0")
	httpReq.Header.Set("Accept-Encoding", "gzip")

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
			_ = closeErr // explicitly ignore the error
		}
	}()

	// Check for HTTP errors early
	if httpResp.StatusCode != http.StatusOK {
		// Read limited error body to avoid memory issues
		limitedReader := io.LimitReader(httpResp.Body, 1024)
		body, _ := io.ReadAll(limitedReader) //nolint:errcheck // intentionally ignore error to always return HTTP error
		return fmt.Errorf("HTTP error %d: %s", httpResp.StatusCode, string(body))
	}

	// Use streaming decoder for response
	decoder := json.NewDecoder(httpResp.Body)

	var gqlResp GraphQLResponse
	if err := decoder.Decode(&gqlResp); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check for GraphQL errors
	if len(gqlResp.Errors) > 0 {
		return fmt.Errorf("GraphQL errors: %v", gqlResp.Errors)
	}

	// Unmarshal the data into the result if provided
	if result != nil {
		if err := json.Unmarshal(gqlResp.Data, result); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
	}

	return nil
}

// ExecuteStreaming performs a GraphQL query with streaming response processing.
// This is useful for large responses that can be processed incrementally.
func (c *OptimizedClient) ExecuteStreaming(
	ctx context.Context,
	query string,
	variables map[string]interface{},
	processor func(json.RawMessage) error,
) error {
	// Get buffer from pool
	buf := c.bufferPool.Get().(*bytes.Buffer) //nolint:errcheck // sync.Pool.Get() always succeeds
	defer func() {
		buf.Reset()
		c.bufferPool.Put(buf)
	}()

	// Prepare the GraphQL request
	gqlReq := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(gqlReq); err != nil {
		return fmt.Errorf("failed to marshal GraphQL request: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, buf)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("User-Agent", "hardcover-cli/1.0.0")
	httpReq.Header.Set("Accept-Encoding", "gzip")

	if c.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	// Make the HTTP request
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer func() {
		_ = httpResp.Body.Close() //nolint:errcheck // intentionally ignore error in defer
	}()

	if httpResp.StatusCode != http.StatusOK {
		limitedReader := io.LimitReader(httpResp.Body, 1024)
		body, _ := io.ReadAll(limitedReader) //nolint:errcheck // intentionally ignore error to always return HTTP error
		return fmt.Errorf("HTTP error %d: %s", httpResp.StatusCode, string(body))
	}

	// Stream decode the response
	decoder := json.NewDecoder(httpResp.Body)

	var gqlResp GraphQLResponse
	if err := decoder.Decode(&gqlResp); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(gqlResp.Errors) > 0 {
		return fmt.Errorf("GraphQL errors: %v", gqlResp.Errors)
	}

	// Process the data with the provided processor function
	return processor(gqlResp.Data)
}
