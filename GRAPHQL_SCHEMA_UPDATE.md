# GraphQL Schema Update Guide

This guide explains how to fetch the latest GraphQL schema from the Hardcover.app live endpoint and regenerate the Go client code.

## Prerequisites

- A valid Hardcover.app API key
- Go 1.19 or later
- The `genqlient` tool (automatically installed via Makefile)

## Quick Start

### 1. Set Your API Key

First, ensure you have your Hardcover.app API key configured:

```bash
# Option 1: Environment variable
export HARDCOVER_API_KEY="your-api-key-here"

# Option 2: Using the CLI config
./bin/hardcover-cli config set-api-key "your-api-key-here"
```

### 2. Build the Development Version

The schema fetching functionality is only available in the development build:

```bash
make build-dev
```

This creates `bin/hardcover-cli-dev` with schema management features.

### 3. Fetch Latest Schema

Fetch the latest GraphQL schema from the live endpoint:

```bash
make graphql-fetch
```

Or manually:

```bash
./bin/hardcover-cli-dev schema --endpoint https://api.hardcover.app/v1/graphql
```

### 4. Regenerate Go Code

Generate the updated Go client code from the new schema:

```bash
make graphql-generate
```

### 5. Complete Update (Recommended)

Or do both steps at once:

```bash
make graphql-update
```

## Detailed Steps

### Step 1: Verify API Key

Test that your API key is working:

```bash
./bin/hardcover-cli-dev me
```

If this returns your user profile, your API key is valid.

### Step 2: Fetch Schema with Custom Options

You can customize the schema fetch:

```bash
# Use a different output location
./bin/hardcover-cli-dev schema --output ./custom-schema.graphql

# Use a different endpoint (if needed)
./bin/hardcover-cli-dev schema --endpoint https://staging-api.hardcover.app/v1/graphql

# Combine both
./bin/hardcover-cli-dev schema --output ./staging-schema.graphql --endpoint https://staging-api.hardcover.app/v1/graphql
```

### Step 3: Verify Schema Changes

After fetching, you can inspect the schema:

```bash
# Show schema information
make graphql-info

# View the schema file
cat internal/client/schema.graphql

# Check for new types
grep "^type " internal/client/schema.graphql | sort
```

### Step 4: Regenerate Code

Generate the updated Go client:

```bash
# Using the Makefile (recommended)
make graphql-generate

# Or manually with latest genqlient
go run github.com/Khan/genqlient@latest internal/client/genqlient.yaml
```

### Step 5: Test the Changes

Build and test the updated client:

```bash
# Build development version with new code
make build-dev

# Test a basic query
./bin/hardcover-cli-dev me

# Test search functionality
./bin/hardcover-cli-dev search books "test"
```

## Troubleshooting

### API Key Issues

If you get authentication errors:

```bash
# Check if API key is set
echo $HARDCOVER_API_KEY

# Or check config
./bin/hardcover-cli-dev config get-api-key

# Test with explicit API key
HARDCOVER_API_KEY="your-key" ./bin/hardcover-cli-dev schema
```

### Schema Fetch Errors

If schema fetching fails:

```bash
# Check network connectivity
curl -X POST https://api.hardcover.app/v1/graphql \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -d '{"query":"query { __schema { types { name } } }"}'

# Try without authentication (some endpoints allow this)
./bin/hardcover-cli-dev schema --endpoint https://api.hardcover.app/v1/graphql
```

### Code Generation Issues

If code generation fails:

```bash
# Check genqlient version
go run github.com/Khan/genqlient@latest --version

# Verify schema file exists
ls -la internal/client/schema.graphql

# Check genqlient config
cat internal/client/genqlient.yaml

# Try with verbose output
go run github.com/Khan/genqlient@latest --verbose internal/client/genqlient.yaml
```

### Build Issues

If the development build fails:

```bash
# Clean and rebuild
make clean
make build-dev

# Check for build tag issues
go build -tags dev -v .
```

## File Locations

- **Schema file**: `internal/client/schema.graphql`
- **Queries file**: `internal/client/queries.graphql`
- **Generated code**: `internal/client/generated.go`
- **Config file**: `internal/client/genqlient.yaml`

## Schema Endpoints

- **Production**: `https://api.hardcover.app/v1/graphql`
- **Staging** (if available): `https://staging-api.hardcover.app/v1/graphql`

## Best Practices

1. **Always test your API key** before fetching schema
2. **Review schema changes** before regenerating code
3. **Test the updated client** with basic queries
4. **Commit schema changes** to version control
5. **Update queries** if the schema has breaking changes

## Common Workflows

### Daily Development

```bash
# Quick schema update
make graphql-update
make build-dev
```

### Before Releases

```bash
# Fetch latest schema
make graphql-fetch

# Review changes
git diff internal/client/schema.graphql

# Regenerate if needed
make graphql-generate

# Test thoroughly
make test
make build-dev
./bin/hardcover-cli-dev me
```

### Schema Investigation

```bash
# Fetch schema
make graphql-fetch

# Explore available types
grep "^type " internal/client/schema.graphql | grep -i book

# Check specific type
grep -A 20 "^type Book" internal/client/schema.graphql
```

## Integration with CI/CD

For automated schema updates, you can use:

```bash
# In your CI pipeline
export HARDCOVER_API_KEY="$HARDCOVER_API_KEY"
make graphql-update
make test
make build-prod
```

## Notes

- The schema command is only available in development builds (`-tags dev`)
- Production builds exclude schema management features
- Schema fetching requires a valid API key
- Generated code is automatically excluded from production builds
- Always test changes before committing to version control 