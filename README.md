# Hardcover CLI

A comprehensive command-line interface for interacting with the Hardcover.app GraphQL API.

## Features

- **User Profile Management**: Get your authenticated user profile information
- **Book Search**: Search for books by title, author, or other criteria
- **User Search**: Search for users by name, username, or location
- **Configuration Management**: Easy setup and management of API keys
- **High Test Coverage**: Comprehensive unit tests for all functionality
- **Well-Documented**: Clear help text and documentation for all commands

## Current Status

This CLI tool is actively developed and currently supports core functionality for interacting with the Hardcover.app API. Some features are still in development due to API schema inconsistencies.

### ✅ Working Features
- Book search with detailed results
- User search with profile information
- User profile retrieval
- Configuration management
- Comprehensive test coverage

### ⚠️ Known Issues
- Limited to read-only operations (no write operations implemented)
- Some API endpoints may have GraphQL schema inconsistencies

For detailed API coverage information, see [API_COVERAGE.md](API_COVERAGE.md).

## Installation

### Prerequisites

- Go 1.23 or later
- A Hardcover.app account and API key

### Build from Source

```bash
git clone <repository-url>
cd hardcover-cli
go build -o hardcover main.go
```

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

**Example Output:**
```
User Profile:
  ID: 12345
  Username: johndoe
```

#### Search for Books

```bash
hardcover search books "golang programming"
hardcover search books "tolkien"
hardcover search books "machine learning"
```

Returns matching books with:
- Title and subtitle
- Author names
- Publication year
- Edition ID and URL
- Rating and ratings count
- ISBNs and series information

**Example Output:**
```
1. Building RESTful Web services with Go
   Subtitle: Learn how to build powerful RESTful APIs with Golang that scale gracefully
   Authors: Naren Yellavula
   Edition ID: 1676108
   URL: https://hardcover.app/books/building-restful-web-services-with-go

-----------------------------
2. Mastering Go
   Subtitle: Create Golang production applications using network libraries, concurrency, machine learning, and advanced data structures
   Authors: Mihalis Tsoukalos
   Year: 2019
   Edition ID: 1863717
   URL: https://hardcover.app/books/mastering-go-create-golang-production-applications-using-network-libraries-concurrency-machine-learning-and-advanced-data-structures
   Rating: 4.00/5 (1 ratings)
   ISBNs: 1838555323, 9781838555320

-----------------------------
```

#### Search for Users

```bash
hardcover search users "john"
hardcover search users "adam smith"
hardcover search users "new york"
```

Returns matching users with:
- Username and name
- Location and custom flair
- Books count, followers, and following counts
- Pro supporter status
- Profile image availability

**Example Output:**
```
1. johndoe
   Name: John Doe
   Location: New York, NY
   Books: 150
   Followers: 45
   Following: 23
   Pro: Yes
   Has Image: Yes

-----------------------------
2. adam_smith
   Name: Adam Smith
   Location: San Francisco, CA
   Books: 89
   Followers: 12
   Following: 67
   Has Image: Yes

-----------------------------
```

### Configuration Commands

#### Set API Key

```bash
hardcover config set-api-key "your-api-key-here"
```

#### Get API Key

```bash
hardcover config get-api-key
```

**Example Output:**
```
Current API key: hc_************************
```

#### Show Config Path

```bash
hardcover config show-path
```

**Example Output:**
```
Configuration file: /home/user/.hardcover/config.yaml
```



### Global Options

- `--config`: Specify a custom config file path
- `--api-key`: Override the API key for a single command
- `--help`: Show help for any command

## Examples

### Search for Books and Users

```bash
# Search for Go programming books
hardcover search books "golang"

# Search for users named John
hardcover search users "john"
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
│   ├── search.go          # Search commands (books and users)
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

#### Search Users

```graphql
query SearchUsers($query: String!) {
  search(query: $query, query_type: "User", per_page: 25, page: 1) {
    ids
    results
    query
    query_type
    page
    per_page
  }
}
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Troubleshooting

### Common Issues

**API Key Not Set:**
```bash
Error: API key is required. Set it using:
  export HARDCOVER_API_KEY="your-api-key"
  or
  hardcover config set-api-key "your-api-key"
```

**API Key Not Set:**
```bash
Error: API key is required. Set it using:
  export HARDCOVER_API_KEY="your-api-key"
  or
  hardcover config set-api-key "your-api-key"
```

**Search Errors:**
```bash
Error: failed to search books: GraphQL errors: [some error message]
```
This may indicate API schema changes or temporary service issues.

### Getting Help

For issues and questions:
- Check the help text: `hardcover --help`
- Review the documentation above
- Check [API_COVERAGE.md](API_COVERAGE.md) for implementation status
- File an issue in the repository

## Changelog

### v1.0.0
- Initial release
- User profile management
- Book search functionality
- User search functionality
- Configuration management
- Comprehensive test coverage
