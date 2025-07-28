# Hardcover CLI - AI Agent Configuration

## Project Overview

Hardcover CLI is a Go-based command-line interface for interacting with the Hardcover.app GraphQL API. This document provides essential information for AI agents working on this codebase.

## Architecture

### Core Structure
```
hardcover-cli/
├── cmd/                    # CLI command implementations
│   ├── root.go            # Root command and CLI setup
│   ├── me.go              # User profile command
│   ├── search.go          # Search commands (books and users)
│   ├── config.go          # Configuration commands
│   ├── context.go         # Context utilities
│   └── *_test.go          # Unit tests
├── internal/
│   ├── client/            # GraphQL client implementation
│   │   ├── client.go      # HTTP client wrapper
│   │   ├── schema.graphql # GraphQL schema
│   │   ├── queries.graphql # GraphQL queries
│   │   └── genqlient.yaml # genqlient configuration
│   ├── config/            # Configuration management
│   ├── contextutil/       # Context utilities
│   └── testutil/          # Testing utilities
├── main.go                # Application entry point
├── go.mod                 # Go module definition
└── README.md             # User documentation
```

### Key Technologies
- **Go 1.23+**: Core language
- **Cobra**: CLI framework
- **GraphQL**: API communication
- **YAML**: Configuration files
- **Testify**: Testing framework

## Development Guidelines

### Code Style
- Follow Go best practices and conventions
- Use comprehensive error handling
- Include unit tests for all new functionality
- Maintain high test coverage
- Use descriptive variable and function names

### Testing Requirements
- All new code must include unit tests
- Use testify for assertions and mocking
- Mock HTTP responses for API testing
- Test both success and error scenarios
- Maintain test coverage above 80%

### Configuration Management
- Support both environment variables and config files
- Store sensitive data securely (file permissions 0600)
- Provide clear error messages for missing configuration
- Support command-line overrides

### Error Handling
- Return meaningful error messages
- Log errors appropriately
- Handle network timeouts gracefully
- Validate user input thoroughly

## API Integration

### GraphQL Endpoint
- Base URL: `https://api.hardcover.app/v1/graphql`
- Authentication: Bearer token via API key
- Rate limiting: Respect API limits

### Current Implemented Features
- ✅ User profile retrieval (`hardcover me`)
- ✅ Book search (`hardcover search books <query>`)
- ✅ User search (`hardcover search users <query>`)
- ✅ Configuration management (`hardcover config`)

### GraphQL Schema
The application uses a custom GraphQL schema defined in `internal/client/schema.graphql` with queries in `internal/client/queries.graphql`.

## Build and Development

### Prerequisites
- Go 1.23 or later
- golangci-lint for code quality
- Hardcover.app API key for testing

### Common Commands
```bash
# Build the application
make build

# Run tests
make test

# Run linter
make lint

# Format code
make fmt

# Install binary
make install
```

### Testing
```bash
# Run all tests
go test ./...

# Run with coverage
go test ./... -cover

# Run specific package tests
go test ./internal/config -v
```

## Security Considerations

### API Key Management
- Never hardcode API keys in source code
- Use environment variables or secure config files
- Mask API keys in output (show only first few characters)
- Set appropriate file permissions (0600) for config files

### Input Validation
- Validate all user inputs
- Sanitize search queries
- Handle malformed GraphQL responses
- Implement proper error boundaries

## Performance Guidelines

### HTTP Client
- Use connection pooling
- Implement appropriate timeouts (30s default)
- Handle concurrent requests safely
- Cache responses when appropriate

### Memory Management
- Minimize memory allocations
- Use efficient JSON parsing
- Clean up resources properly
- Avoid memory leaks in long-running processes

## Extension Points

### Adding New Commands
1. Create new command file in `cmd/` directory
2. Follow existing command patterns
3. Add comprehensive tests
4. Update help documentation
5. Register command in `cmd/root.go`

### GraphQL Schema Updates
1. Update `internal/client/schema.graphql`
2. Add new queries to `internal/client/queries.graphql`
3. Update client implementation
4. Add corresponding tests
5. Update documentation

### Configuration Extensions
1. Extend config struct in `internal/config/config.go`
2. Add validation logic
3. Update command handlers
4. Add tests for new configuration options

## Common Patterns

### Command Structure
```go
func newCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "command-name",
        Short: "Brief description",
        Long:  "Detailed description",
        RunE:  runCommand,
    }
    
    // Add flags
    cmd.Flags().String("flag", "", "flag description")
    
    return cmd
}

func runCommand(cmd *cobra.Command, args []string) error {
    // Implementation
    return nil
}
```

### Error Handling
```go
if err != nil {
    return fmt.Errorf("failed to perform action: %w", err)
}
```

### Testing Pattern
```go
func TestCommand(t *testing.T) {
    // Setup
    // Execute
    // Assert
}
```

## Documentation Standards

### Code Comments
- Document exported functions and types
- Include examples for complex functions
- Explain business logic where necessary
- Keep comments up to date with code changes

### User Documentation
- Update README.md for new features
- Include usage examples
- Document configuration options
- Provide troubleshooting guidance

## Quality Assurance

### Linting
- Use golangci-lint for code quality
- Fix all linting issues before committing
- Maintain consistent code style
- Follow Go best practices

### Testing
- Aim for >80% test coverage
- Test edge cases and error conditions
- Use table-driven tests where appropriate
- Mock external dependencies

### Code Review
- Review all changes before merging
- Ensure tests pass
- Verify documentation is updated
- Check for security implications

## Troubleshooting

### Common Issues
1. **API Key Not Set**: Check environment variables and config file
2. **GraphQL Errors**: Verify schema compatibility
3. **Network Issues**: Check connectivity and timeouts
4. **Test Failures**: Ensure all dependencies are available

### Debugging
- Use verbose logging for troubleshooting
- Check GraphQL query responses
- Verify configuration file permissions
- Test with minimal examples

## Future Considerations

### Planned Enhancements
- Complete genqlient integration
- Add more API endpoints
- Support multiple output formats
- Implement interactive mode
- Add batch operations

### Scalability
- Design for concurrent usage
- Consider caching strategies
- Plan for API rate limiting
- Prepare for schema evolution

This document should be updated as the project evolves to ensure AI agents have the most current information for effective development assistance. 