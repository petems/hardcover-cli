# Hardcover CLI - Agent Documentation

## Project Overview

This is a Go-based command-line interface (CLI) for the Hardcover.app GraphQL API. The project demonstrates professional-grade Go development with comprehensive testing, documentation, and a clean architecture.

**Key Characteristics:**
- **Language**: Go 1.23+
- **Architecture**: CLI application using Cobra framework
- **API Integration**: GraphQL client for Hardcover.app
- **Testing**: High coverage with testify framework
- **Configuration**: YAML-based with environment variable support

## Repository Structure

```
hardcover-cli/
├── cmd/                    # CLI command implementations
│   ├── root.go            # Root command and CLI setup
│   ├── me.go              # User profile command
│   ├── search.go          # Search functionality
│   ├── book.go            # Book-related commands
│   ├── config.go          # Configuration management
│   ├── context.go         # Context utilities
│   └── *_test.go          # Unit tests
├── internal/
│   ├── client/            # GraphQL client implementation
│   │   ├── client.go      # HTTP client wrapper
│   │   ├── schema.graphql # GraphQL schema
│   │   ├── queries.graphql # GraphQL queries
│   │   └── genqlient.yaml # Code generation config
│   └── config/            # Configuration management
│       ├── config.go      # Configuration logic
│       └── config_test.go # Configuration tests
├── main.go                # Application entry point
├── Makefile               # Build and development tasks
└── Documentation files    # Various .md files
```

## Critical Architecture Patterns

### 1. GraphQL Schema Issues Workaround
**CRITICAL**: The project has significant GraphQL introspection schema issues. The actual API structure doesn't match the introspection schema, requiring manual HTTP implementations.

**Key Files:**
- `cmd/search.go` - Manual search implementation
- `cmd/book_manual.go` - Manual book details implementation
- `API_SCHEMA_ISSUES.md` - Complete documentation of issues

**Pattern**: Use manual HTTP requests when introspection schema is insufficient.

### 2. Build Tags Strategy
- **Development builds** (`-tags dev`): Include manual implementations and schema management
- **Production builds**: Exclude development features for smaller binaries

### 3. Configuration Management
- **Precedence**: Command-line flags > Environment variables > Configuration file
- **Storage**: `~/.hardcover/config.yaml` with secure permissions (0600)
- **Security**: API keys are masked in output

## Key Files and Their Purposes

### Core Implementation Files
- `cmd/root.go` - CLI setup and command registration
- `cmd/me.go` - User profile retrieval
- `cmd/search.go` - Book search (manual HTTP implementation)
- `cmd/book.go` - Book details (manual HTTP implementation)
- `cmd/config.go` - Configuration management
- `internal/client/client.go` - GraphQL HTTP client
- `internal/config/config.go` - Configuration logic

### Documentation Files
- `API_COVERAGE.md` - Current vs. available API features
- `API_SCHEMA_ISSUES.md` - GraphQL schema problems and workarounds
- `DEVELOPMENT.md` - Development guide and architecture overview
- `GRAPHQL_SCHEMA_UPDATE.md` - Schema management instructions

### Build and Development
- `Makefile` - Build tasks, testing, and GraphQL code generation
- `go.mod` - Dependencies (Cobra, testify, genqlient, yaml.v3)
- `.golangci.yml` - Linting configuration

## Development Workflows

### Adding New Commands
1. Create new file in `cmd/` directory
2. Implement command using Cobra framework
3. Add to root command in `cmd/root.go`
4. Write comprehensive tests
5. Update documentation

### GraphQL Integration
1. **Check API documentation first** - Don't rely on introspection schema
2. **Use manual HTTP implementation** if introspection is insufficient
3. **Test against actual API** before implementing
4. **Document workarounds** in `API_SCHEMA_ISSUES.md`

### Testing Strategy
- **Unit tests**: All packages have dedicated test files
- **Mocking**: HTTP server mocking for API interactions
- **Coverage**: High test coverage with meaningful test cases
- **Error scenarios**: Comprehensive failure case testing

## Important Constraints and Considerations

### GraphQL Schema Issues
- **Search parameters**: Not properly exposed in introspection
- **Book query structure**: Uses `editions` queries instead of direct `book(id:)`
- **Field name mismatches**: API uses snake_case, introspection may show camelCase
- **Library queries**: Complex structures not exposed in introspection

### Build System
- **Development builds**: Include manual implementations and schema features
- **Production builds**: Exclude development features
- **Code generation**: Uses genqlient for GraphQL client generation

### Security
- **API key storage**: Secure file permissions
- **API key display**: Masking in output
- **Input validation**: Sanitization of user inputs

## Common Tasks and Commands

### Building
```bash
make build-prod    # Production build
make build-dev     # Development build with schema features
```

### Testing
```bash
make test          # Run all tests
go test ./... -cover  # With coverage
```

### GraphQL Schema Management
```bash
make graphql-update  # Fetch schema and regenerate code
make graphql-info    # Show schema information
```

### Code Quality
```bash
make lint          # Run linter
make fmt           # Format code
```

## API Integration Patterns

### Manual HTTP Implementation
When GraphQL introspection is insufficient:

```go
func performManualSearch(query string) (*SearchResponse, error) {
    // Direct HTTP request to GraphQL endpoint
    // Proper parameter handling
    // Error handling and response parsing
}
```

### Configuration Management
```go
// Precedence: CLI flags > Environment > Config file
config := internal/config.Load()
apiKey := config.GetAPIKey()
```

### Error Handling
```go
// Comprehensive error handling for API failures
// User-friendly error messages
// Proper GraphQL error parsing
```

## Future Development Priorities

### High Priority
1. **User Books Management** - Add/remove books from personal library
2. **Reading Status** - Mark books as want to read, currently reading, or read
3. **Enhanced Book Details** - Include authors, genres, and real relationships

### Medium Priority
1. **Author Search and Details** - Search and view author information
2. **Advanced Search Filters** - Filter by genre, year, rating, etc.
3. **Reading Journals** - Track reading progress and notes

### Low Priority
1. **Reviews and Ratings** - Write and read book reviews
2. **Following System** - Follow users and lists
3. **Recommendations** - Personalized book recommendations

## Troubleshooting Guide

### Common Issues
1. **GraphQL schema mismatches**: Use manual HTTP implementations
2. **API key configuration**: Check precedence order
3. **Build failures**: Ensure correct Go version (1.23+)
4. **Test failures**: Check API connectivity and mocking

### Debugging
- Use development build for schema management
- Check `API_SCHEMA_ISSUES.md` for known problems
- Test against actual API, not just introspection schema
- Review manual implementations in `cmd/search.go` and `cmd/book_manual.go`

## References
- [API Coverage Documentation](API_COVERAGE.md)
- [GraphQL Schema Issues](API_SCHEMA_ISSUES.md)
- [Development Guide](DEVELOPMENT.md)
- [GraphQL Schema Management](GRAPHQL_SCHEMA_UPDATE.md)
- [Hardcover API Documentation](https://docs.hardcover.app/api/) 