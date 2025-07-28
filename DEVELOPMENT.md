# Hardcover CLI Development Summary

## Project Overview

This project implements a comprehensive GoLang command-line interface (CLI) application for interacting with the Hardcover.app GraphQL API. The application uses a **custom type generation approach** to overcome GraphQL introspection schema issues while maintaining type safety and developer experience.

## Architecture & Design

### Core Components

1. **CLI Framework**: Built using the Cobra library for robust command-line interface management
2. **Custom GraphQL Client**: Custom HTTP client with type-safe GraphQL query execution using generated types
3. **Type Generation System**: Custom Go script that generates types from GraphQL schema
4. **DRY GraphQL Architecture**: Centralized queries, typed responses, and helper functions
5. **Configuration Management**: Flexible configuration system supporting both file and environment variable sources
6. **Comprehensive Testing**: Unit tests with mocking and high coverage using testify

### Directory Structure

```
hardcover-cli/
├── cmd/                    # CLI command implementations
│   ├── root.go            # Root command and initialization
│   ├── me.go              # User profile command (uses generated types)
│   ├── search.go          # Book search functionality
│   ├── book.go            # Book details retrieval
│   ├── config.go          # Configuration management commands
│   ├── context.go         # Context utilities for config passing
│   └── *_test.go          # Comprehensive unit tests
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
├── README.md             # User documentation
└── DEVELOPMENT.md        # This file
```

## GraphQL Type Generation System

### Custom Type Generator

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

## Implementation Details

### Commands Implemented

#### 1. `hardcover me` (Type-Safe Implementation)
- Uses generated `client.Users` type for type safety
- Fetches authenticated user profile information
- Displays user ID, username, email, and timestamps
- Includes comprehensive error handling and validation

#### 2. `hardcover search books <query>`
- Searches for books using GraphQL API
- Supports filtering by title, author, and other criteria
- Displays formatted results with book details, ratings, and genres
- Includes pagination support and result counting

#### 3. `hardcover book get <book_id>`
- Retrieves detailed information for a specific book
- Shows comprehensive book metadata including:
  - Title, description, and publication details
  - Author and contributor information
  - Genre classification
  - Ratings and review statistics
  - Cover image and external URLs

#### 4. `hardcover config` subcommands
- `set-api-key`: Store API key securely
- `get-api-key`: Display current API key (masked for security)
- `show-path`: Show configuration file location

### GraphQL Integration

The application uses a well-defined GraphQL schema with the following key queries:

```graphql
# User profile retrieval (from queries.graphql)
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

# Book search functionality
query SearchBooks($query: String!) {
  search(query: $query, type: BOOKS) {
    ... on BookSearchResults {
      totalCount
      results {
        ... on Book {
          id
          title
          slug
          cached_contributors { name role }
          cached_genres { name }
          averageRating
          ratingsCount
        }
      }
    }
  }
}

# Book details retrieval
query GetBook($id: ID!) {
  book(id: $id) {
    id
    title
    description
    # ... additional fields
  }
}
```

### Type-Safe GraphQL Usage

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

## Implementing New Commands

### Step-by-Step Guide for New GraphQL Commands

#### 1. Add Query to `internal/client/queries.graphql`

```graphql
query GetUserLibrary($userId: ID!) {
  user(id: $userId) {
    id
    username
    library {
      books {
        id
        title
        author
        rating
      }
    }
  }
}
```

#### 2. Update Type Generation

```bash
# Regenerate types to include new GraphQL types
make generate-types
```

#### 3. Add Response Type to `internal/client/responses.go`

```go
// GetUserLibraryResponse represents the response from the GetUserLibrary query
type GetUserLibraryResponse struct {
    User *Users `json:"user"`
}
```

#### 4. Add Helper Function to `internal/client/helpers.go`

```go
// GetUserLibrary executes the GetUserLibrary query with the given user ID
func (c *Client) GetUserLibrary(ctx context.Context, userID string) (*GetUserLibraryResponse, error) {
    variables := map[string]interface{}{
        "userId": userID,
    }
    var response GetUserLibraryResponse
    if err := c.Execute(ctx, GetUserLibraryQuery, variables, &response); err != nil {
        return nil, err
    }
    return &response, nil
}
```

#### 5. Add Query Constant to `internal/client/queries.go`

```go
const (
    // ... existing queries ...
    
    // GetUserLibraryQuery fetches a user's library
    GetUserLibraryQuery = `
query GetUserLibrary($userId: ID!) {
  user(id: $userId) {
    id
    username
    library {
      books {
        id
        title
        author
        rating
      }
    }
  }
}
`
)
```

#### 6. Implement CLI Command

```go
// cmd/library.go
package cmd

import (
    "context"
    "fmt"
    "github.com/spf13/cobra"
    "hardcover-cli/internal/client"
)

var libraryCmd = &cobra.Command{
    Use:   "library [user-id]",
    Short: "Get a user's library",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        cfg, ok := getConfig(cmd.Context())
        if !ok {
            return fmt.Errorf("failed to get configuration")
        }

        if cfg.APIKey == "" {
            return fmt.Errorf("API key is required")
        }

        gqlClient := client.NewClient(cfg.BaseURL, cfg.APIKey)
        response, err := gqlClient.GetUserLibrary(context.Background(), args[0])
        if err != nil {
            return fmt.Errorf("failed to get user library: %w", err)
        }

        // Display results using typed response
        if response.User == nil {
            return fmt.Errorf("no user data received")
        }

        user := response.User
        printToStdoutf(cmd.OutOrStdout(), "User Library:\n")
        printToStdoutf(cmd.OutOrStdout(), "  User: %s\n", user.Username)
        
        // Access library data through generated types
        // ... display library information ...

        return nil
    },
}

func setupLibraryCommands() {
    rootCmd.AddCommand(libraryCmd)
}
```

#### 7. Register Command in `cmd/root.go`

```go
func init() {
    // ... existing commands ...
    setupLibraryCommands()
}
```

#### 8. Add Tests

```go
// cmd/library_test.go
func TestLibraryCmd_Success(t *testing.T) {
    // Test implementation using typed responses
    // Mock HTTP responses and verify output
}
```

### Best Practices for New Commands

1. **Always use generated types** for type safety
2. **Follow the DRY pattern** with centralized queries and responses
3. **Use helper functions** for clean, maintainable code
4. **Add comprehensive tests** with proper mocking
5. **Update documentation** for new commands
6. **Handle errors gracefully** with user-friendly messages

### Configuration System

The application implements a flexible configuration system:

1. **Environment Variables**: `HARDCOVER_API_KEY`
2. **Configuration File**: `~/.hardcover/config.yaml`
3. **Command-line Flags**: `--api-key` for one-time overrides

Configuration precedence: Command-line flags > Environment variables > Configuration file

### Testing Strategy

The project includes comprehensive unit tests with:

- **Package Coverage**: All major packages have dedicated test files
- **Mocking**: HTTP server mocking for API interactions
- **Error Scenarios**: Testing of failure cases and edge conditions
- **Integration Tests**: Command registration and interaction testing
- **Configuration Tests**: File system mocking and environment variable testing
- **Type Safety Tests**: Verification of generated type usage

### Test Files Overview

- `internal/config/config_test.go`: Configuration management tests
- `internal/client/client_test.go`: GraphQL client tests
- `cmd/me_test.go`: User profile command tests (using generated types)
- `cmd/search_test.go`: Book search command tests
- `cmd/book_test.go`: Book retrieval command tests
- `cmd/config_test.go`: Configuration command tests

### Error Handling

The application implements comprehensive error handling:

- **API Errors**: GraphQL error parsing and user-friendly messages
- **Network Errors**: Connection failure handling and timeouts
- **Configuration Errors**: Missing API keys and invalid configurations
- **Validation Errors**: Input validation and argument checking
- **Type Errors**: Compile-time type safety with generated types

### Security Considerations

1. **API Key Storage**: Secure file permissions (0600) for configuration files
2. **API Key Display**: Masking of sensitive information in output
3. **HTTP Headers**: Proper authentication header handling
4. **Input Validation**: Sanitization of user inputs
5. **Type Safety**: Compile-time validation of API responses

## Dependencies

- **github.com/spf13/cobra**: CLI framework
- **github.com/stretchr/testify**: Testing framework with assertions and mocks
- **gopkg.in/yaml.v3**: YAML configuration file parsing

## Build and Development Commands

### Local Development

```bash
# Clone the repository
git clone <repository-url>
cd hardcover-cli

# Install dependencies
go mod tidy

# Generate types from GraphQL schema
make generate-types

# Build the application
go build -o bin/hardcover-cli .

# Run tests
go test ./...

# Run with coverage
go test ./... -cover

# Clean and rebuild
make clean
make build
```

### Available Make Commands

```bash
make build              # Build the application
make test               # Run all tests
make generate-types     # Regenerate GraphQL types
make clean              # Clean build artifacts
make install            # Install to $GOPATH/bin
```

### Production Deployment

```bash
# Build for production
go build -ldflags="-s -w" -o hardcover .

# Cross-compilation examples
GOOS=linux GOARCH=amd64 go build -o hardcover-linux .
GOOS=windows GOARCH=amd64 go build -o hardcover-windows.exe .
GOOS=darwin GOARCH=amd64 go build -o hardcover-macos .
```

## Code Quality

The project follows Go best practices:

- **Package Structure**: Clear separation of concerns
- **Error Handling**: Comprehensive error checking and reporting
- **Documentation**: Extensive inline documentation and help text
- **Testing**: High test coverage with meaningful test cases
- **Code Style**: Consistent formatting and naming conventions
- **Type Safety**: Generated types ensure compile-time validation

## API Integration

The application integrates with the Hardcover.app GraphQL API:

- **Endpoint**: `https://api.hardcover.app/v1/graphql`
- **Authentication**: Bearer token authentication
- **Rate Limiting**: Respectful API usage with appropriate timeouts
- **Error Handling**: Proper GraphQL error parsing and user feedback
- **Type Safety**: Generated types ensure API response validation

## Performance Considerations

1. **HTTP Client**: Reusable HTTP client with connection pooling
2. **Timeouts**: Configurable request timeouts (30 seconds default)
3. **Memory Usage**: Efficient JSON parsing and minimal memory footprint
4. **Concurrent Safety**: Thread-safe configuration and client usage
5. **Type Generation**: Fast compile-time type checking

## Future Enhancements

1. **Additional Commands**: Support for more Hardcover.app API endpoints
2. **Output Formats**: JSON and CSV output options
3. **Batch Operations**: Support for bulk operations
4. **Interactive Mode**: TUI for enhanced user experience
5. **Configuration Profiles**: Support for multiple API configurations
6. **GraphQL Schema Evolution**: Automated handling of schema changes

## Conclusion

This Hardcover CLI application demonstrates a professional-grade Go application with:

- **Custom type generation** that overcomes GraphQL introspection issues
- **DRY GraphQL architecture** for maintainable and extensible code
- **Type safety** through generated Go structs
- **Clean architecture** and separation of concerns
- **Comprehensive testing** and error handling
- **User-friendly CLI interface** with extensive help documentation
- **Secure configuration management**
- **Extensible design** for future enhancements

The project serves as an excellent example of modern Go development practices and demonstrates how to work around GraphQL API limitations while maintaining excellent developer experience and type safety.