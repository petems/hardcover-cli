# Hardcover CLI

A comprehensive command-line interface for interacting with the Hardcover.app GraphQL API.

## Features

- **User Profile Management**: Get your authenticated user profile information
- **Book Search**: Search for books by title, author, or other criteria
- **Book Details**: Retrieve detailed information about specific books
- **Configuration Management**: Easy setup and management of API keys
- **High Test Coverage**: Comprehensive unit tests for all functionality
- **Well-Documented**: Clear help text and documentation for all commands

## Installation

### Prerequisites

- Go 1.19 or later
- A Hardcover.app account and API key

### Build from Source

#### Production Build (Recommended for End Users)

```bash
git clone <repository-url>
cd hardcover-cli
make build-prod
# or
go build -o hardcover .
```

#### Development Build (Includes Schema Management)

```bash
git clone <repository-url>
cd hardcover-cli
make build-dev
# or
go build -tags dev -o hardcover-dev .
```

The development build includes additional features for schema management and code generation, while the production build excludes these development-only features for a smaller, cleaner binary.

## Configuration

Before using the CLI, you need to set your Hardcover.app API key. You can do this in two ways:

### Option 1: Environment Variable

```bash
export HARDCOVER_API_KEY="your-api-key-here"
```

### Option 2: Configuration File

```bash
hardcover config set-api-key "your-api-key-here"
```

The configuration file is stored at `~/.hardcover/config.yaml`.

## Usage

### Basic Commands

#### Get Your Profile

```bash
hardcover me
```

Displays your user profile information including:
- User ID
- Username
- Email address
- Account creation date
- Last updated date

#### Search for Books

```bash
hardcover search books "golang programming"
hardcover search books "tolkien"
hardcover search books "machine learning"
```

Returns matching books with:
- Title and author information
- Publication details
- Ratings and genres
- Hardcover.app URL

#### Get Book Details

```bash
hardcover book get 12345
hardcover book get "book-slug-or-id"
```

Displays comprehensive book information including:
- Title and description
- Author(s) and contributors
- Publication details (year, page count, ISBN)
- Genres and categories
- Ratings and reviews summary
- Cover image URL
- Hardcover.app URL

### Configuration Commands

#### Set API Key

```bash
hardcover config set-api-key "your-api-key-here"
```

#### Get Current API Key

```bash
hardcover config get-api-key
```

#### Show Configuration File Path

```bash
hardcover config show-path
```

### Global Options

- `--config`: Specify a custom config file path
- `--api-key`: Override the API key for a single command
- `--help`: Show help for any command

## Examples

### Search and Get Book Details

```bash
# Search for Go programming books
hardcover search books "golang"

# Get details for a specific book (use the ID from search results)
hardcover book get 67890
```

### Profile Management

```bash
# Get your profile
hardcover me

# Set up your API key
hardcover config set-api-key "your-api-key"

# Check your current API key
hardcover config get-api-key
```

## Development

### Project Structure

```
hardcover-cli/
├── cmd/                    # CLI command implementations
│   ├── root.go            # Root command and CLI setup
│   ├── me.go              # User profile command
│   ├── search.go          # Search commands
│   ├── book.go            # Book-related commands
│   ├── config.go          # Configuration commands
│   └── *_test.go          # Unit tests
├── internal/
│   ├── client/            # GraphQL client implementation
│   │   ├── client.go      # HTTP client wrapper
│   │   ├── schema.graphql # GraphQL schema
│   │   └── queries.graphql # GraphQL queries
│   └── config/            # Configuration management
│       ├── config.go      # Configuration logic
│       └── config_test.go # Configuration tests
├── main.go                # Application entry point
├── go.mod                 # Go module definition
└── README.md             # This file
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -cover

# Run tests for specific package
go test ./internal/config -v
```

### GraphQL Schema

The application uses a GraphQL schema based on the Hardcover.app API. The schema is defined in `internal/client/schema.graphql` and includes types for:

- User profile information
- Book details and metadata
- Search results
- Contributors and genres

### Key Dependencies

- **Cobra**: CLI framework for Go
- **Testify**: Testing toolkit with assertions and mocks
- **YAML**: Configuration file parsing

## API Reference

### GraphQL Queries

The application uses the following GraphQL queries:

#### Get Current User

```graphql
query GetCurrentUser {
  me {
    id
    username
    email
    createdAt
    updatedAt
  }
}
```

#### Search Books

```graphql
query SearchBooks($query: String!) {
  search(query: $query, type: BOOKS) {
    ... on BookSearchResults {
      totalCount
      results {
        ... on Book {
          id
          title
          slug
          cached_contributors {
            name
            role
          }
          cached_genres {
            name
          }
          averageRating
          ratingsCount
        }
      }
    }
  }
}
```

#### Get Book Details

```graphql
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
```

## Development

### GraphQL Schema Management

For developers working on the GraphQL client, see [GRAPHQL_SCHEMA_UPDATE.md](GRAPHQL_SCHEMA_UPDATE.md) for detailed instructions on:

- Fetching the latest schema from the live endpoint
- Regenerating Go client code
- Troubleshooting schema updates
- Best practices for schema management

### API Coverage

See [API_COVERAGE.md](API_COVERAGE.md) for a comprehensive overview of:

- Currently implemented features vs available API capabilities
- Missing features and implementation priorities
- Technical considerations and limitations
- Links to official API documentation

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For issues and questions:
- Check the help text: `hardcover --help`
- Review the documentation above
- File an issue in the repository

## Changelog

### v1.0.0
- Initial release
- User profile management
- Book search functionality
- Book details retrieval
- Configuration management
- Comprehensive test coverage
