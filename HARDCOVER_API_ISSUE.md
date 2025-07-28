# Hardcover API GraphQL Introspection Issues - Updated Analysis

## Summary

The GraphQL introspection schema at `https://api.hardcover.app/v1/graphql` has **fundamental mismatches** with the actual API structure documented at https://docs.hardcover.app/api/. While standard GraphQL auto-generation tools are unusable, we've successfully implemented a **custom type generation solution** that works around these issues.

## Current Status: Custom Solution Working ✅

After extensive investigation, we discovered that **standard GraphQL code generation tools** fail due to introspection schema issues, but we've successfully created a **custom Go-based type generator** that works around these problems.

## Complete List of Introspection Schema Issues

### 1. Search Field Parameters (Original Issue)

**Expected Schema (per API docs):**
```graphql
type Query {
  search(
    query: String!
    query_type: String
    per_page: Int
    page: Int
    sort: String
    fields: String
    weights: String
  ): SearchOutput
}
```

**Actual Introspection Result:**
```graphql
type Query {
  search: SearchOutput
}
```

**Impact**: Standard code generation tools cannot generate type-safe code for search functionality.

### 2. Book Query Structure

**Expected Schema (per API docs):**
```graphql
query GetBookDetails {
  editions(where: {id: {_eq: 21953653}}) {
    id
    title
    description
    slug
    isbn_10
    isbn_13
    release_date
    pages
    edition_format
    publisher { name }
    contributions { author { name } }
  }
}
```

**Actual Introspection Result**: Shows `book(id: ID!): Book` which doesn't work with the actual API structure.

**Impact**: Generated code for book queries fails at runtime.

### 3. User Profile Query Structure

**Expected Schema (per API docs):**
```graphql
query GetCurrentUser {
  me {
    id
    username
    email
    created_at
    updated_at
  }
}
```

**Actual Introspection Result**: Shows `me: users` returning a single object.

**Runtime Reality**: The actual API returns `me` as an array `[]users`, not a single object.

**Impact**: Generated code fails with "cannot unmarshal array into Go struct" errors.

### 4. Field Name Mismatches

**API Reality**: Uses snake_case fields (`release_date`, `isbn_10`, `isbn_13`, `edition_format`)

**Introspection Schema**: May show camelCase fields (`publicationYear`, `pageCount`, `editionFormat`)

**Impact**: Generated code references non-existent fields.

### 5. Complex Library Queries

**Issue**: Complex queries for user libraries use different structures than what's exposed in introspection.

**Reference**: https://docs.hardcover.app/api/guides/gettingallbooksinlibrary/

**Impact**: Library-related functionality cannot be generated with standard tools.

## Our Custom Solution: Working GraphQL Implementation ✅

### Custom Type Generator

We've implemented a custom Go script (`scripts/generate-types.go`) that:

1. **Fetches the GraphQL schema** via introspection
2. **Handles schema mismatches** by mapping problematic types
3. **Generates Go types** that work with the actual API
4. **Provides type safety** for GraphQL operations

### DRY GraphQL Architecture

Our current approach includes:

1. **Query Constants** (`internal/client/queries.go`): Centralized GraphQL queries
2. **Typed Responses** (`internal/client/responses.go`): Type-safe response structures
3. **Helper Functions** (`internal/client/helpers.go`): Clean API for query execution
4. **Generated Types** (`internal/client/types.go`): Auto-generated from schema

### Working Implementation

```go
// Clean, type-safe GraphQL operations
response, err := gqlClient.GetCurrentUser(ctx)
if err != nil {
    return fmt.Errorf("failed to get user profile: %w", err)
}

// Direct access to typed fields
user := response.Me
printToStdoutf(cmd.OutOrStdout(), "  ID: %d\n", user.ID)
printToStdoutf(cmd.OutOrStdout(), "  Username: %s\n", user.Username)
```

## Impact Assessment

### Standard Code Generation Tools ❌
- **genqlient**: Cannot generate working code due to schema mismatches
- **GraphQL codegen**: Same issues with introspection schema
- **gqlgenc**: Package naming conflicts and schema issues
- **Any off-the-shelf GraphQL tooling**: Fundamentally broken for this API

### Our Custom Solution ✅
- **Type Safety**: Achieved with custom type generation
- **Runtime Success**: All operations work correctly
- **Maintainability**: DRY approach with centralized queries
- **Developer Experience**: Clean, intuitive API

## Solution: Custom Type Generation + Manual HTTP

### Current Working Implementation

Our hybrid approach combines the best of both worlds:

1. **Custom Type Generation**: `scripts/generate-types.go` - Generates working Go types
2. **DRY GraphQL Architecture**: Centralized queries and typed responses
3. **Manual HTTP Client**: `internal/client/client.go` - Reliable HTTP operations
4. **Type-Safe Operations**: Helper functions with proper error handling

### Benefits of Our Approach

- ✅ **Actually works** with the real API
- ✅ **Type safe** with generated Go structs
- ✅ **DRY and maintainable** with centralized queries
- ✅ **Error handling** for real API responses
- ✅ **Extensible** for new GraphQL operations

## Technical Details

### Custom Type Generation Pattern

```go
// scripts/generate-types.go
func main() {
    // Fetch schema via introspection
    schema := fetchGraphQLSchema()
    
    // Generate Go types with custom mappings
    generateTypes(schema)
}
```

### DRY GraphQL Usage

```go
// internal/client/helpers.go
func (c *Client) GetCurrentUser(ctx context.Context) (*GetCurrentUserResponse, error) {
    var response GetCurrentUserResponse
    if err := c.Execute(ctx, GetCurrentUserQuery, nil, &response); err != nil {
        return nil, err
    }
    return &response, nil
}
```

### GraphQL Endpoint Usage

The GraphQL endpoint `https://api.hardcover.app/v1/graphql` is used successfully with:
- Custom type generation that handles schema issues
- Manual query construction based on API documentation
- Type-safe response parsing into Go structs
- Centralized query management

## Recommendations

### For This Project
1. **Keep custom type generation** - It works and provides type safety
2. **Maintain DRY GraphQL architecture** - Centralized and maintainable
3. **Use API documentation as source of truth** - Not introspection schema
4. **Extend with new queries** - Follow the established pattern

### For Hardcover.app API
1. **Fix introspection schema** to match actual API structure
2. **Update field names** to be consistent (snake_case vs camelCase)
3. **Document actual query structures** that work at runtime
4. **Test introspection schema** against real API responses

### For Other Developers
1. **Always test introspection schema** against actual API before using GraphQL tooling
2. **Don't assume introspection matches documentation** - verify both
3. **Consider custom type generation** when standard tools fail
4. **Use API documentation as primary source** of truth
5. **Implement DRY patterns** for maintainable GraphQL code

## Conclusion

While standard GraphQL auto-generation tools are unusable due to introspection schema issues, we've successfully implemented a custom solution that provides type safety and maintainability. Our approach demonstrates that GraphQL can work with the Hardcover.app API when using the right tooling and patterns.

The key insight is that **custom type generation + DRY architecture** can overcome introspection schema limitations while providing excellent developer experience and type safety.

## References

- [API Search Documentation](https://docs.hardcover.app/api/guides/searching/)
- [Book Details Guide](https://docs.hardcover.app/api/guides/gettingbookdetails/)
- [Getting All Books Guide](https://docs.hardcover.app/api/guides/gettingallbooksinlibrary/)
- GraphQL Endpoint: `https://api.hardcover.app/v1/graphql`
- API Documentation: https://docs.hardcover.app/api/
- Our Custom Type Generator: `scripts/generate-types.go`
- DRY GraphQL Architecture: `internal/client/queries.go`, `responses.go`, `helpers.go` 