//go:build !dev

package client

import (
	"context"
	"fmt"

	"github.com/Khan/genqlient/graphql"
)

// Stub types for production builds

type GetBookBook struct{}
type GetBookBookCached_contributorsContributor struct{}
type GetBookBookCached_genresGenre struct{}
type GetCurrentUserMeUser struct{}
type SearchBooksSearchSearchResults interface{}
type SearchBooksSearchBookSearchResults struct{}
type SearchBooksSearchBookSearchResultsResultsBook struct{}
type SearchBooksSearchAuthorSearchResults struct{}
type SearchBooksSearchUserSearchResults struct{}

type GetBookResponse struct{}
type GetCurrentUserResponse struct{}
type SearchBooksResponse struct{}

type SearchType string

const (
	SearchTypeBooks   SearchType = "BOOKS"
	SearchTypeAuthors SearchType = "AUTHORS"
	SearchTypeUsers   SearchType = "USERS"
)

// Stub methods for production builds

func (v *GetBookBook) GetId() string           { return "" }
func (v *GetBookBook) GetTitle() string        { return "" }
func (v *GetBookBook) GetDescription() string  { return "" }
func (v *GetBookBook) GetSlug() string         { return "" }
func (v *GetBookBook) GetIsbn() string         { return "" }
func (v *GetBookBook) GetPublicationYear() int { return 0 }
func (v *GetBookBook) GetPageCount() int       { return 0 }
func (v *GetBookBook) GetCached_contributors() []GetBookBookCached_contributorsContributor {
	return nil
}
func (v *GetBookBook) GetCached_genres() []GetBookBookCached_genresGenre { return nil }
func (v *GetBookBook) GetImage() string                                  { return "" }
func (v *GetBookBook) GetAverageRating() float64                         { return 0 }
func (v *GetBookBook) GetRatingsCount() int                              { return 0 }

func (v *GetBookBookCached_contributorsContributor) GetName() string { return "" }
func (v *GetBookBookCached_contributorsContributor) GetRole() string { return "" }

func (v *GetBookBookCached_genresGenre) GetName() string { return "" }

func (v *GetCurrentUserMeUser) GetId() string        { return "" }
func (v *GetCurrentUserMeUser) GetUsername() string  { return "" }
func (v *GetCurrentUserMeUser) GetEmail() string     { return "" }
func (v *GetCurrentUserMeUser) GetCreatedAt() string { return "" }
func (v *GetCurrentUserMeUser) GetUpdatedAt() string { return "" }

func (v *GetBookResponse) GetBook() GetBookBook                          { return GetBookBook{} }
func (v *GetCurrentUserResponse) GetMe() GetCurrentUserMeUser            { return GetCurrentUserMeUser{} }
func (v *SearchBooksResponse) GetSearch() SearchBooksSearchSearchResults { return nil }

func (v *SearchBooksSearchBookSearchResults) GetTypename() string { return "" }
func (v *SearchBooksSearchBookSearchResults) GetTotalCount() int  { return 0 }
func (v *SearchBooksSearchBookSearchResults) GetResults() []SearchBooksSearchBookSearchResultsResultsBook {
	return nil
}

func (v *SearchBooksSearchBookSearchResultsResultsBook) GetId() string             { return "" }
func (v *SearchBooksSearchBookSearchResultsResultsBook) GetTitle() string          { return "" }
func (v *SearchBooksSearchBookSearchResultsResultsBook) GetDescription() string    { return "" }
func (v *SearchBooksSearchBookSearchResultsResultsBook) GetSlug() string           { return "" }
func (v *SearchBooksSearchBookSearchResultsResultsBook) GetIsbn() string           { return "" }
func (v *SearchBooksSearchBookSearchResultsResultsBook) GetPublicationYear() int   { return 0 }
func (v *SearchBooksSearchBookSearchResultsResultsBook) GetPageCount() int         { return 0 }
func (v *SearchBooksSearchBookSearchResultsResultsBook) GetAverageRating() float64 { return 0 }
func (v *SearchBooksSearchBookSearchResultsResultsBook) GetRatingsCount() int      { return 0 }

func (v *SearchBooksSearchAuthorSearchResults) GetTypename() string { return "" }
func (v *SearchBooksSearchUserSearchResults) GetTypename() string   { return "" }

// Stub functions for production builds

func GetBook(ctx context.Context, client graphql.Client, id string) (*GetBookResponse, error) {
	return nil, fmt.Errorf("GraphQL client not available in production builds")
}

func GetCurrentUser(ctx context.Context, client graphql.Client) (*GetCurrentUserResponse, error) {
	return nil, fmt.Errorf("GraphQL client not available in production builds")
}

func SearchBooks(ctx context.Context, client graphql.Client, query string, searchType SearchType) (*SearchBooksResponse, error) {
	return nil, fmt.Errorf("GraphQL client not available in production builds")
}
