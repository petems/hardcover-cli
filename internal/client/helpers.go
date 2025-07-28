package client

import (
	"context"
)

// GetCurrentUser executes the GetCurrentUser query and returns the response.
func (c *Client) GetCurrentUser(ctx context.Context) (*GetCurrentUserResponse, error) {
	var response GetCurrentUserResponse
	if err := c.Execute(ctx, GetCurrentUserQuery, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SearchBooks executes the SearchBooks query with the given search term.
func (c *Client) SearchBooks(ctx context.Context, query string) (*SearchBooksResponse, error) {
	variables := map[string]interface{}{
		"query": query,
	}
	var response SearchBooksResponse
	if err := c.Execute(ctx, SearchBooksQuery, variables, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetBook executes the GetBook query with the given book ID.
func (c *Client) GetBook(ctx context.Context, id string) (*GetBookResponse, error) {
	variables := map[string]interface{}{
		"id": id,
	}
	var response GetBookResponse
	if err := c.Execute(ctx, GetBookQuery, variables, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
