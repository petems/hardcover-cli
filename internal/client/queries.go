package client

// GraphQL query constants
const (
	// GetCurrentUserQuery fetches the current user's profile information
	GetCurrentUserQuery = `
query GetCurrentUser {
  me {
    id
    username
    email
    name
    bio
    location
    createdAt
    updatedAt
  }
}
`

	// SearchBooksQuery searches for books by query string
	SearchBooksQuery = `
query SearchBooks($query: String!) {
  search(query: $query, type: BOOKS) {
    ... on BookSearchResults {
      totalCount
      results {
        ... on Book {
          id
          title
          slug
          isbn
          publicationYear
          pageCount
          cached_contributors {
            name
            role
          }
          cached_genres {
            name
          }
          image
          averageRating
          ratingsCount
        }
      }
    }
  }
}
`

	// GetBookQuery fetches a specific book by ID
	GetBookQuery = `
query GetBook($id: ID!) {
  book(id: $id) {
    id
    title
    description
    slug
    isbn
    publicationYear
    pageCount
    cached_contributors {
      name
      role
    }
    cached_genres {
      name
    }
    image
    averageRating
    ratingsCount
    createdAt
    updatedAt
  }
}
`
)
