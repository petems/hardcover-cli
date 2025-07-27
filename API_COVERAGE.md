# Hardcover CLI API Coverage

This document tracks the current implementation status of the Hardcover CLI against the full Hardcover.app GraphQL API.

## Current Implementation Status

| Feature | Status | CLI Command | Documentation | Notes |
|---------|--------|-------------|---------------|-------|
| **User Profile** | ✅ | `hardcover me` | [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) | Basic user info (id, username, email, timestamps) |
| **Book Search** | ✅ | `hardcover search books` | [Searching Guide](https://docs.hardcover.app/api/guides/searching/) | Basic book search with pagination |
| **Book Details** | ✅ | `hardcover book get` | [Book Details Guide](https://docs.hardcover.app/api/guides/gettingbookdetails/) | Basic book information |
| **Book Listing** | ✅ | `hardcover book list` | [Getting All Books Guide](https://docs.hardcover.app/api/guides/gettingallbooksinlibrary/) | List all books |
| **Configuration** | ✅ | `hardcover config` | - | API key management |

## Missing API Features

### Core Entities

| Feature | Status | Documentation | Implementation Priority | Notes |
|---------|--------|---------------|------------------------|-------|
| **Authors** | ❌ | [Authors Schema](https://docs.hardcover.app/api/graphql/schemas/authors/) | Medium | Author search, details, books by author |
| **Editions** | ❌ | [Editions Schema](https://docs.hardcover.app/api/graphql/schemas/editions/) | Low | Different book editions |
| **Characters** | ❌ | [Characters Schema](https://docs.hardcover.app/api/graphql/schemas/characters/) | Low | Character information |
| **Activities** | ❌ | [Activities Schema](https://docs.hardcover.app/api/graphql/schemas/activities/) | Medium | User activity feed |

### User Library Management

| Feature | Status | Documentation | Implementation Priority | Notes |
|---------|--------|---------------|------------------------|-------|
| **User Books** | ❌ | [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) | High | Personal library management |
| **Reading Status** | ❌ | [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) | High | Want to read, currently reading, read |
| **Reading Journals** | ❌ | [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) | Medium | Reading progress and notes |
| **User Goals** | ❌ | [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) | Low | Reading goals and challenges |

### Social Features

| Feature | Status | Documentation | Implementation Priority | Notes |
|---------|--------|---------------|------------------------|-------|
| **Lists** | ❌ | [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) | Medium | Create and manage book lists |
| **Following** | ❌ | [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) | Low | Follow users and lists |
| **Likes** | ❌ | [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) | Low | Like books, reviews, lists |
| **Reviews** | ❌ | [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) | Medium | Write and read reviews |

### Advanced Search

| Feature | Status | Documentation | Implementation Priority | Notes |
|---------|--------|---------------|------------------------|-------|
| **Author Search** | ❌ | [Searching Guide](https://docs.hardcover.app/api/guides/searching/) | Medium | Search by author name |
| **User Search** | ❌ | [Searching Guide](https://docs.hardcover.app/api/guides/searching/) | Low | Search for users |
| **Advanced Filters** | ❌ | [Searching Guide](https://docs.hardcover.app/api/guides/searching/) | Medium | Filter by genre, year, rating, etc. |
| **Series Search** | ❌ | [Searching Guide](https://docs.hardcover.app/api/guides/searching/) | Low | Search book series |

### Recommendations

| Feature | Status | Documentation | Implementation Priority | Notes |
|---------|--------|---------------|------------------------|-------|
| **Book Recommendations** | ❌ | [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) | Medium | Personalized recommendations |
| **Trending Books** | ❌ | [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) | Low | Popular books |

## Current Implementation Details

### ✅ Implemented Features

#### User Profile (`hardcover me`)
- **Fields**: id, username, email, created_at, updated_at
- **Missing**: profile image, bio, location, website, social links
- **Caveat**: Only basic user information is displayed

#### Book Search (`hardcover search books`)
- **Features**: Basic text search with pagination
- **Parameters**: query, query_type, per_page, page
- **Missing**: Advanced filters, sorting options, search suggestions
- **Caveat**: Search results return IDs only, requires additional queries for details

#### Book Details (`hardcover book get`)
- **Fields**: id, title, description, slug, pages, rating, ratings_count, release_year, release_date, subtitle, timestamps, cached_contributors, cached_tags, cached_image
- **Missing**: authors, genres, series, reviews, recommendations, similar books
- **Caveat**: Uses cached fields instead of real relationships

#### Book Listing (`hardcover book list`)
- **Features**: List all books with basic information
- **Missing**: Filtering, sorting, pagination controls
- **Caveat**: Limited to basic book information

## Implementation Priorities

### High Priority (Core Functionality)
1. **User Books Management** - Add/remove books from personal library
2. **Reading Status** - Mark books as want to read, currently reading, or read
3. **Enhanced Book Details** - Include authors, genres, and real relationships

### Medium Priority (Enhanced Features)
1. **Author Search and Details** - Search and view author information
2. **Advanced Search Filters** - Filter by genre, year, rating, etc.
3. **Reading Journals** - Track reading progress and notes
4. **Lists Management** - Create and manage book lists

### Low Priority (Social Features)
1. **Reviews and Ratings** - Write and read book reviews
2. **Following System** - Follow users and lists
3. **Recommendations** - Personalized book recommendations
4. **Activity Feed** - View user activity

## Technical Considerations

### Current Limitations
- **Cached Fields**: Many book details use cached fields instead of real GraphQL relationships
- **Search Results**: Search returns IDs only, requiring additional queries
- **Pagination**: Basic pagination without advanced controls
- **Error Handling**: Limited error handling for API failures

### Required Schema Updates
- Add queries for user books and reading status
- Add mutations for library management
- Add queries for authors and enhanced book details
- Add queries for lists and social features

### Performance Considerations
- Implement caching for frequently accessed data
- Add batch queries to reduce API calls
- Implement connection pooling for GraphQL client

## Contributing

To implement missing features:

1. **Update GraphQL Queries**: Add new queries to `internal/client/queries.graphql`
2. **Regenerate Client**: Run `make graphql-update` to generate new Go code
3. **Add CLI Commands**: Implement new commands in the `cmd/` directory
4. **Add Tests**: Include comprehensive tests for new functionality
5. **Update Documentation**: Update this coverage table and README

## References

- [Getting All Books Guide](https://docs.hardcover.app/api/guides/gettingallbooksinlibrary/)
- [Getting Book Details Guide](https://docs.hardcover.app/api/guides/gettingbookdetails/)
- [Searching Guide](https://docs.hardcover.app/api/guides/searching/)
- [Authors Schema](https://docs.hardcover.app/api/graphql/schemas/authors/)
- [Books Schema](https://docs.hardcover.app/api/graphql/schemas/books/)
- [Editions Schema](https://docs.hardcover.app/api/graphql/schemas/editions/)
- [Activities Schema](https://docs.hardcover.app/api/graphql/schemas/activities/)
- [Characters Schema](https://docs.hardcover.app/api/graphql/schemas/characters/)
- [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) 