package client

// Response types for GraphQL queries

// GetCurrentUserResponse represents the response from the GetCurrentUser query.
type GetCurrentUserResponse struct {
	Me *Users `json:"me"`
}

// SearchBooksResponse represents the response from the SearchBooks query.
type SearchBooksResponse struct {
	Search *BookSearchResults `json:"search"`
}

// BookSearchResults represents the search results for books.
type BookSearchResults struct {
	Results    []*Books `json:"results"`
	TotalCount int      `json:"totalCount"`
}

// GetBookResponse represents the response from the GetBook query.
type GetBookResponse struct {
	Book *Books `json:"book"`
}
