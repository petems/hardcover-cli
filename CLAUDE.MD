# Claude AI Configuration for Hardcover CLI

## Project Context

This is a Go-based CLI application for interacting with the Hardcover.app GraphQL API. The project follows modern Go development practices with comprehensive testing and documentation.

## Key Information for Claude

### Project Structure
- **Language**: Go 1.23+
- **Framework**: Cobra CLI framework
- **API**: GraphQL with Hardcover.app
- **Testing**: Testify framework with high coverage
- **Configuration**: YAML files and environment variables

### Current State
- ✅ User profile management (`hardcover me`)
- ✅ Book search (`hardcover search books <query>`)
- ✅ User search (`hardcover search users <query>`)
- ✅ Configuration management (`hardcover config`)
- ⚠️ Limited to read-only operations
- ⚠️ Some API schema inconsistencies

### Important Files
- `cmd/root.go` - Main CLI setup and command registration
- `cmd/search.go` - Search functionality implementation
- `cmd/me.go` - User profile commands
- `cmd/config.go` - Configuration management
- `internal/client/client.go` - GraphQL client
- `internal/config/config.go` - Configuration logic
- `internal/client/schema.graphql` - GraphQL schema
- `internal/client/queries.graphql` - GraphQL queries

## Claude Instructions

### When Working on This Codebase

1. **Follow Go Best Practices**
   - Use proper error handling with `fmt.Errorf` and `%w` verb
   - Follow Go naming conventions
   - Use interfaces appropriately
   - Write comprehensive tests

2. **Maintain Existing Patterns**
   - Follow the command structure in existing `cmd/` files
   - Use the same error handling patterns
   - Maintain consistent testing approaches
   - Follow the configuration management patterns

3. **Testing Requirements**
   - All new code must have unit tests
   - Use testify for assertions and mocking
   - Mock HTTP responses for API testing
   - Test both success and error scenarios
   - Aim for >80% test coverage

4. **Security Considerations**
   - Never hardcode API keys
   - Use secure file permissions (0600) for config files
   - Mask sensitive information in output
   - Validate all user inputs

5. **Documentation Standards**
   - Update README.md for new features
   - Include usage examples
   - Document configuration options
   - Keep code comments current

### Common Tasks for Claude

#### Adding New Commands
```go
// Follow this pattern in cmd/ directory
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
    // Get configuration
    cfg, err := cmd.Flags().GetString("config")
    if err != nil {
        return fmt.Errorf("failed to get config flag: %w", err)
    }
    
    // Implementation
    return nil
}
```

#### GraphQL Integration
```go
// Use the existing client pattern
client := internal.NewClient(apiKey)
result, err := client.ExecuteQuery(ctx, query, variables)
if err != nil {
    return fmt.Errorf("failed to execute query: %w", err)
}
```

#### Configuration Management
```go
// Use the existing config patterns
config, err := internal.LoadConfig(configPath)
if err != nil {
    return fmt.Errorf("failed to load config: %w", err)
}
```

### Testing Patterns

#### Unit Tests
```go
func TestCommand(t *testing.T) {
    // Setup
    cmd := newCommand()
    
    // Execute
    err := cmd.Execute()
    
    // Assert
    assert.NoError(t, err)
}
```

#### HTTP Mocking
```go
func TestAPI(t *testing.T) {
    // Setup mock server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Mock response
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"data": {...}}`))
    }))
    defer server.Close()
    
    // Test implementation
}
```

### Error Handling Patterns

#### Wrapping Errors
```go
if err != nil {
    return fmt.Errorf("failed to perform action: %w", err)
}
```

#### Validation
```go
if len(args) == 0 {
    return fmt.Errorf("query argument is required")
}
```

### Configuration Patterns

#### Loading Config
```go
config, err := internal.LoadConfig(configPath)
if err != nil {
    return fmt.Errorf("failed to load configuration: %w", err)
}
```

#### API Key Management
```go
apiKey := config.APIKey
if apiKey == "" {
    return fmt.Errorf("API key is required. Set it using:\n  export HARDCOVER_API_KEY=\"your-api-key\"\n  or\n  hardcover config set-api-key \"your-api-key\"")
}
```

## Specific Guidelines for Claude

### When Asked to Fix Issues
1. **Analyze the error** - Look at the specific error message and context
2. **Check existing patterns** - See how similar issues are handled in the codebase
3. **Write tests** - Include tests that would catch the issue
4. **Update documentation** - If the fix changes behavior, update relevant docs

### When Asked to Add Features
1. **Follow existing patterns** - Use the same structure as existing commands
2. **Add comprehensive tests** - Include unit tests for the new functionality
3. **Update documentation** - Add examples to README.md
4. **Consider security** - Ensure new features don't introduce security issues

### When Asked to Refactor
1. **Maintain functionality** - Ensure refactoring doesn't break existing features
2. **Update tests** - Modify tests to match the new structure
3. **Keep patterns consistent** - Follow the established patterns in the codebase
4. **Document changes** - Update comments and documentation

### When Asked to Debug
1. **Check configuration** - Verify API keys and config files
2. **Test with minimal examples** - Create simple test cases
3. **Check GraphQL schema** - Verify schema compatibility
4. **Look at error handling** - Ensure proper error propagation

## Common Commands for Development

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

## API Endpoints

- **Base URL**: `https://api.hardcover.app/v1/graphql`
- **Authentication**: Bearer token via API key
- **Rate Limiting**: Respect API limits

## Current Limitations

1. **Read-only operations** - No write operations implemented
2. **API schema inconsistencies** - Some endpoints may have issues
3. **Limited GraphQL client** - Using basic HTTP client, not genqlient
4. **No caching** - All requests hit the API directly

## Future Enhancements

1. **Complete genqlient integration** - Type-safe GraphQL client
2. **Additional API endpoints** - More Hardcover.app functionality
3. **Multiple output formats** - JSON, CSV, etc.
4. **Interactive mode** - TUI for better UX
5. **Batch operations** - Support for bulk operations

## Important Notes for Claude

- **Always test changes** - Run `make test` before suggesting changes
- **Check linting** - Ensure code passes `make lint`
- **Follow Go conventions** - Use proper Go idioms and patterns
- **Consider security** - Never expose sensitive information
- **Update documentation** - Keep README.md and comments current
- **Maintain compatibility** - Don't break existing functionality

This configuration ensures Claude can effectively assist with development while maintaining code quality and following project standards. 