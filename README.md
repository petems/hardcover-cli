# Hardcover CLI

A comprehensive command-line interface for interacting with the Hardcover.app GraphQL API, built with custom type generation for type safety and maintainability.

## Features

- **User Profile Management**: Get your authenticated user profile information with type-safe GraphQL operations
- **Book Search**: Search for books by title, author, or other criteria
- **User Search**: Search for users by name, username, or location
- **Configuration Management**: Easy setup and management of API keys
- **Custom Type Generation**: Auto-generated Go types from GraphQL schema for compile-time safety
- **DRY GraphQL Architecture**: Centralized queries and typed responses for maintainability
- **High Test Coverage**: Comprehensive unit tests for all functionality
- **Well-Documented**: Clear help text and documentation for all commands

## Current Status

This CLI tool is actively developed and currently supports core functionality for interacting with the Hardcover.app API. The project uses a **custom type generation approach** to overcome GraphQL introspection schema issues while maintaining excellent developer experience.

### ✅ Working Features
- Book search with detailed results
- User search with profile information
- User profile retrieval (type-safe implementation)
- Configuration management
- Custom GraphQL type generation
- DRY GraphQL architecture with centralized queries
- Comprehensive test coverage

### ⚠️ Known Issues
- Limited to read-only operations (no write operations implemented)
- Standard GraphQL code generation tools don't work due to API schema inconsistencies
- Our custom type generation solution works around these limitations

For detailed API coverage information, see [API_COVERAGE.md](API_COVERAGE.md).

## Installation

### Prerequisites

- Go 1.23 or later
- A Hardcover.app account and API key

### Build from Source

```bash
git clone <repository-url>
cd hardcover-cli

# Generate types from GraphQL schema
make generate-types

# Build the application
make build
# or
go build -o bin/hardcover-cli main.go
```

### Available Make Commands

```bash
make build              # Build the application
make test               # Run all tests
make generate-types     # Regenerate GraphQL types
make clean              # Clean build artifacts
make install            # Install to $GOPATH/bin
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
- Email (if available)
- Name, bio, and location (if available)
- Creation and update timestamps

**Example Output:**
```
User Profile:
  ID: 12345
  Username: johndoe
  Email: john@example.com
  Name: John Doe
  Bio: Software developer and book lover
  Location: San Francisco, CA
  Created: 2023-01-15T10:30:00Z
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
│   ├── me.go              # User profile command (type-safe)
│   ├── search.go          # Search commands (books and users)
│   ├── config.go          # Configuration commands
│   └── *_test.go          # Unit tests
├── internal/
│   ├── client/            # GraphQL client implementation
│   │   ├── client.go      # HTTP client wrapper
│   │   ├── queries.go     # GraphQL query constants
│   │   ├── responses.go   # Typed response structures
│   │   ├── helpers.go     # Helper functions for query execution
│   │   ├── types.go       # Generated GraphQL types
│   │   └── queries.graphql # GraphQL query definitions
│   └── config/            # Configuration management
│       ├── config.go      # Configuration logic
│       └── config_test.go # Configuration tests
├── scripts/
│   └── generate-types.go  # Custom type generation script
├── main.go                # Application entry point
├── go.mod                 # Go module definition
├── Makefile               # Build and development commands
└── README.md             # This file
```

### GraphQL Type Generation System

Our project uses a custom Go script (`scripts/generate-types.go`) that:

1. **Fetches GraphQL Schema**: Performs introspection query to get the schema
2. **Handles Schema Issues**: Maps problematic types (e.g., `citext` → `string`)
3. **Generates Go Types**: Creates type-safe Go structs from GraphQL types
4. **Provides Type Safety**: Ensures compile-time type checking

### DRY GraphQL Architecture

We've implemented a maintainable GraphQL approach with:

1. **Query Constants** (`internal/client/queries.go`): Centralized GraphQL queries
2. **Typed Responses** (`internal/client/responses.go`): Type-safe response structures
3. **Helper Functions** (`internal/client/helpers.go`): Clean API for query execution
4. **Generated Types** (`internal/client/types.go`): Auto-generated from schema

### Type Generation Workflow

```bash
# Generate types from GraphQL schema
make generate-types

# This runs: go run scripts/generate-types.go
# Which creates: internal/client/types.go
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

### Implementing New Commands

For adding new GraphQL commands, follow our established pattern:

1. **Add Query** to `internal/client/queries.graphql`
2. **Regenerate Types** with `make generate-types`
3. **Add Response Type** to `internal/client/responses.go`
4. **Add Helper Function** to `internal/client/helpers.go`
5. **Add Query Constant** to `internal/client/queries.go`
6. **Implement CLI Command** following the established pattern
7. **Register Command** in `cmd/root.go`
8. **Add Tests** for the new command

See [DEVELOPMENT.md](DEVELOPMENT.md) for detailed implementation guide.

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

The application uses the following GraphQL queries with type-safe responses:

#### Get Current User

```graphql
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

### Type-Safe Usage Example

```go
// Clean, type-safe GraphQL operations
gqlClient := client.NewClient(cfg.BaseURL, cfg.APIKey)
response, err := gqlClient.GetCurrentUser(ctx)
if err != nil {
    return fmt.Errorf("failed to get user profile: %w", err)
}

// Direct access to typed fields
user := response.Me
printToStdoutf(cmd.OutOrStdout(), "  ID: %d\n", user.ID)
printToStdoutf(cmd.OutOrStdout(), "  Username: %s\n", user.Username)
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

### Development Guidelines

- **Use generated types** for type safety
- **Follow DRY patterns** with centralized queries and responses
- **Add comprehensive tests** with proper mocking
- **Update documentation** for new commands
- **Handle errors gracefully** with user-friendly messages

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

**Type Generation Issues:**
```bash
# If types are out of sync, regenerate them
make generate-types
```

**Search Errors:**
```bash
Error: failed to search books: GraphQL errors: [some error message]
```
This may indicate API schema changes or temporary service issues.

**Build Errors:**
```bash
# Clean and rebuild
make clean
make generate-types
make build
```

### Getting Help

For issues and questions:
- Check the help text: `hardcover --help`
- Review the documentation above
- Check [API_COVERAGE.md](API_COVERAGE.md) for implementation status
- Check [DEVELOPMENT.md](DEVELOPMENT.md) for development guidelines
- Check [HARDCOVER_API_ISSUE.md](HARDCOVER_API_ISSUE.md) for API limitations
- File an issue in the repository

## Changelog

### v1.1.0
- **Custom Type Generation**: Implemented custom GraphQL type generation system
- **DRY GraphQL Architecture**: Centralized queries and typed responses
- **Type-Safe Operations**: All GraphQL operations now use generated types
- **Improved Development Workflow**: Added Make commands and type generation
- **Enhanced Documentation**: Updated development guidelines and examples

### v1.0.0
- Initial release
- User profile management
- Book search functionality
- User search functionality
- Configuration management
- Comprehensive test coverage
