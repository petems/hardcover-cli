# Hardcover CLI - Claude AI Guide

## Project Context for Claude

This is a **Go-based CLI application** for the Hardcover.app GraphQL API. The project has **significant GraphQL introspection schema issues** that require special handling and manual implementations.

## Critical Knowledge for Claude

### 🚨 GraphQL Schema Issues (CRITICAL)
The GraphQL introspection schema at `https://api.hardcover.app/v1/graphql` **does NOT match** the actual API structure documented at https://docs.hardcover.app/api/. This means:

1. **Code generation tools fail** - genqlient cannot generate working code
2. **Manual HTTP implementations required** - See `cmd/search.go` and `cmd/book_manual.go`
3. **Always check API docs first** - Don't trust the introspection schema
4. **Test against actual API** - Not just the generated code

### Key Schema Issues:
- Search field parameters not exposed in introspection
- Book queries use `editions` structure, not direct `book(id:)`
- Field name mismatches (snake_case vs camelCase)
- Complex library queries not exposed

## Repository Architecture

### Build System
- **Development builds**: `go build -tags dev` (includes manual implementations)
- **Production builds**: `go build` (excludes dev features)
- **Makefile**: Comprehensive build and development tasks

### Core Dependencies
```go
github.com/spf13/cobra     // CLI framework
github.com/stretchr/testify // Testing framework
github.com/Khan/genqlient  // GraphQL code generation
gopkg.in/yaml.v3          // Configuration
```

### Directory Structure
```
cmd/           # CLI commands (Cobra-based)
internal/      # Private packages
  client/      # GraphQL client (manual + generated)
  config/      # Configuration management
```

## Development Patterns for Claude

### When Adding New Features

1. **Check API documentation first** - https://docs.hardcover.app/api/
2. **Test against actual API** - Don't rely on introspection
3. **Use manual HTTP if needed** - Follow patterns in `cmd/search.go`
4. **Add comprehensive tests** - All packages have test files
5. **Update documentation** - Keep `API_COVERAGE.md` current

### GraphQL Integration Pattern

```go
// Manual HTTP implementation (when introspection fails)
func performManualRequest(query string, variables map[string]interface{}) ([]byte, error) {
    // Direct HTTP POST to GraphQL endpoint
    // Proper authentication headers
    // Error handling and response parsing
}

// Generated client usage (when introspection works)
client := internal/client.NewClient()
response, err := client.Search(ctx, query)
```

### Configuration Pattern

```go
// Precedence: CLI flags > Environment > Config file
config := internal/config.Load()
apiKey := config.GetAPIKey()

// Secure storage: ~/.hardcover/config.yaml (0600 permissions)
// API keys masked in output for security
```

## Key Files for Claude to Understand

### Implementation Files
- `cmd/root.go` - CLI setup and command registration
- `cmd/search.go` - **Manual search implementation** (critical)
- `cmd/book_manual.go` - **Manual book details** (critical)
- `cmd/me.go` - User profile (uses generated client)
- `cmd/config.go` - Configuration management
- `internal/client/client.go` - HTTP client wrapper
- `internal/config/config.go` - Configuration logic

### Documentation Files
- `API_SCHEMA_ISSUES.md` - **CRITICAL** - Complete schema issues documentation
- `API_COVERAGE.md` - Current vs. available API features
- `DEVELOPMENT.md` - Architecture and development guide
- `GRAPHQL_SCHEMA_UPDATE.md` - Schema management instructions

### Build Files
- `Makefile` - Build tasks, testing, GraphQL generation
- `go.mod` - Dependencies
- `.golangci.yml` - Linting configuration

## Common Tasks for Claude

### Adding New Commands
1. Create file in `cmd/` directory
2. Implement using Cobra framework
3. Add to root command in `cmd/root.go`
4. Write tests (follow existing patterns)
5. Update documentation

### GraphQL Integration
1. **Check API docs first** - https://docs.hardcover.app/api/
2. **Test actual API** - Don't trust introspection
3. **Use manual HTTP if needed** - Follow `cmd/search.go` pattern
4. **Document workarounds** - Update `API_SCHEMA_ISSUES.md`

### Testing
- All packages have dedicated test files
- Use testify for assertions and mocking
- Mock HTTP server for API interactions
- Test error scenarios and edge cases

## Important Constraints

### GraphQL Schema Limitations
- **Search parameters**: Not in introspection schema
- **Book queries**: Use `editions` structure, not `book(id:)`
- **Field names**: API uses snake_case, introspection may show camelCase
- **Complex queries**: Library queries not exposed

### Build Constraints
- **Development builds**: Include manual implementations
- **Production builds**: Exclude dev features
- **Code generation**: genqlient for GraphQL client

### Security Requirements
- **API key storage**: Secure file permissions (0600)
- **API key display**: Masking in output
- **Input validation**: Sanitization required

## Development Workflows

### Building
```bash
make build-prod    # Production build
make build-dev     # Development build (includes manual implementations)
```

### Testing
```bash
make test          # Run all tests
go test ./... -cover  # With coverage
```

### GraphQL Management
```bash
make graphql-update  # Fetch schema and regenerate
make graphql-info    # Show schema information
```

### Code Quality
```bash
make lint          # Run linter
make fmt           # Format code
```

## API Integration Examples

### Manual HTTP Implementation (Critical Pattern)
```go
func performManualSearch(query string) (*SearchResponse, error) {
    // Direct HTTP request to GraphQL endpoint
    // Proper parameter handling (query_type, per_page, page)
    // Error handling and response parsing
    // Type-safe response structure
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

## Troubleshooting for Claude

### Common Issues
1. **GraphQL schema mismatches**: Use manual HTTP implementations
2. **API key configuration**: Check precedence order
3. **Build failures**: Ensure Go 1.23+
4. **Test failures**: Check API connectivity and mocking

### Debugging Steps
- Use development build for schema management
- Check `API_SCHEMA_ISSUES.md` for known problems
- Test against actual API, not just introspection schema
- Review manual implementations in `cmd/search.go` and `cmd/book_manual.go`

## Key Takeaways for Claude

1. **GraphQL introspection is unreliable** - Always check API docs first
2. **Manual HTTP implementations are necessary** - Follow existing patterns
3. **Test against actual API** - Don't trust generated code alone
4. **Comprehensive testing required** - All packages have test files
5. **Documentation is critical** - Keep `API_SCHEMA_ISSUES.md` updated
6. **Security matters** - API keys, input validation, secure storage
7. **Build tags strategy** - Development vs production builds

## References
- [API Schema Issues](API_SCHEMA_ISSUES.md) - **CRITICAL READING**
- [API Coverage](API_COVERAGE.md) - Current implementation status
- [Development Guide](DEVELOPMENT.md) - Architecture overview
- [GraphQL Schema Management](GRAPHQL_SCHEMA_UPDATE.md) - Schema handling
- [Hardcover API Documentation](https://docs.hardcover.app/api/) - Official API docs 