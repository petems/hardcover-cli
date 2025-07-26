# GraphQL Code Generation

This document describes the Make tasks available for managing GraphQL schema and code generation in the hardcover-cli project.

## Overview

The project uses [genqlient](https://github.com/Khan/genqlient) to generate type-safe Go code from GraphQL schemas and queries. This ensures that your Go code is always in sync with the GraphQL API schema.

## Make Tasks

### `make graphql-info`
Shows information about the current GraphQL configuration:
- Schema file location
- Configuration file location
- Generated code file location
- Current GraphQL endpoint
- List of available types in the schema

### `make graphql-generate`
Generates Go code from the current GraphQL schema and queries:
- Ensures all dependencies are installed
- Runs genqlient to generate type-safe Go code
- Updates `internal/client/generated.go`

### `make graphql-fetch`
Fetches the latest GraphQL schema from a remote endpoint:
- Downloads the schema using GraphQL introspection
- Converts the introspection result to GraphQL SDL format
- Updates `internal/client/schema.graphql`

**Usage:**
```bash
# Using the default endpoint
make graphql-fetch

# Using a custom endpoint
make graphql-fetch GRAPHQL_ENDPOINT=https://your-api.com/graphql

# With authentication (if required)
make graphql-fetch GRAPHQL_ENDPOINT=https://your-api.com/graphql
```

**Requirements:**
- The GraphQL endpoint must support introspection queries
- If authentication is required, you may need to modify the curl command in the Makefile
- `jq` is recommended for JSON processing (optional)
- `go` command must be available

### `make graphql-update`
Complete workflow that fetches the latest schema and regenerates code:
- Runs `graphql-fetch` to get the latest schema
- Runs `graphql-generate` to create updated Go code

## Configuration

### GraphQL Endpoint
The default GraphQL endpoint is set in the Makefile:
```makefile
GRAPHQL_ENDPOINT ?= https://api.hardcover.app/v1/graphql
```

You can override this by setting the environment variable or passing it as a parameter to the make commands.

### Files
- **Schema**: `internal/client/schema.graphql` - GraphQL schema definition
- **Queries**: `internal/client/queries.graphql` - GraphQL queries used by the CLI
- **Config**: `internal/client/genqlient.yaml` - genqlient configuration
- **Generated**: `internal/client/generated.go` - Generated Go code (do not edit manually)

## Workflow

### Development Workflow
1. Make changes to `internal/client/queries.graphql` if you need new queries
2. Run `make graphql-generate` to regenerate the Go code
3. Update your Go code to use the new generated types
4. Test your changes

### Schema Update Workflow
1. Run `make graphql-update` to fetch the latest schema and regenerate code
2. Review the changes in `internal/client/generated.go`
3. Update your Go code if the API has changed
4. Test your changes

### Adding New Queries
1. Add your GraphQL query to `internal/client/queries.graphql`
2. Run `make graphql-generate` to generate the corresponding Go types
3. Use the generated types in your Go code

## Troubleshooting

### Schema Fetch Issues
If `make graphql-fetch` fails:
1. Check that the endpoint URL is correct
2. Verify the endpoint is accessible
3. Ensure the endpoint supports introspection queries
4. Check if authentication is required

### Code Generation Issues
If `make graphql-generate` fails:
1. Run `go mod tidy` to ensure dependencies are up to date
2. Check that the schema file is valid GraphQL
3. Verify that queries reference valid schema types

### Authentication
If the GraphQL endpoint requires authentication, you'll need to modify the `graphql-fetch` task in the Makefile to include the appropriate headers:

```makefile
@curl -s -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_TOKEN" \
     -d '{"query":"..."}' \
     $(GRAPHQL_ENDPOINT) > /tmp/schema_response.json
```

## Examples

### Fetch schema from a different endpoint
```bash
make graphql-fetch GRAPHQL_ENDPOINT=https://staging-api.hardcover.app/v1/graphql
```

### Generate code only (without fetching schema)
```bash
make graphql-generate
```

### Complete update workflow
```bash
make graphql-update
```

### Check current configuration
```bash
make graphql-info
``` 