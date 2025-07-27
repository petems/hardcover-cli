# Hardcover API GraphQL Introspection Issues - Complete Analysis

## Summary

The GraphQL introspection schema at `https://api.hardcover.app/v1/graphql` has **fundamental mismatches** with the actual API structure documented at https://docs.hardcover.app/api/. This makes GraphQL auto-generation tools completely unusable for this API.

## Critical Finding: GraphQL Auto-Generation is Not Viable

After extensive investigation, we discovered that **every single GraphQL query** has introspection schema issues that prevent code generation from working. The project has been refactored to abandon GraphQL auto-generation entirely in favor of manual HTTP implementations.

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

**Impact**: Code generation tools cannot generate type-safe code for search functionality.

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

**Impact**: Library-related functionality cannot be generated.

## Impact Assessment

### Code Generation Tools
- **genqlient**: Cannot generate working code due to schema mismatches
- **GraphQL codegen**: Same issues with introspection schema
- **Any GraphQL tooling**: Fundamentally broken for this API

### Developer Experience
- **Type Safety**: Impossible to achieve with generated code
- **Runtime Errors**: Generated code fails with unmarshaling errors
- **Documentation Mismatch**: API docs don't match introspection schema
- **Maintenance Overhead**: Constant fighting with broken tooling

## Solution: Manual HTTP Implementations

### Current Working Implementation

All functionality now uses manual HTTP implementations that work around the introspection issues:

1. **Search Functionality**: `cmd/search.go` - Direct HTTP requests with proper parameters
2. **Book Details**: `cmd/book_manual.go` - Uses correct `editions` query structure
3. **Book List**: `cmd/book_list.go` - Manual pagination and filtering
4. **User Profile**: `cmd/me_manual.go` - Handles array response from `me` field

### Benefits of Manual Approach

- ✅ **Actually works** with the real API
- ✅ **Matches documentation** exactly
- ✅ **Type safe** with proper Go structs
- ✅ **Error handling** for real API responses
- ✅ **Maintainable** and understandable code

## Technical Details

### Manual Implementation Pattern

```go
func performManualRequest(query string, variables map[string]interface{}) ([]byte, error) {
    // Direct HTTP POST to GraphQL endpoint
    // Proper authentication headers
    // Error handling and response parsing
    // Type-safe response structures
}
```

### GraphQL Endpoint Usage

The GraphQL endpoint `https://api.hardcover.app/v1/graphql` is still used, but:
- Queries are constructed manually based on API documentation
- Responses are parsed manually into Go structs
- No reliance on introspection schema or generated code

## Recommendations

### For This Project
1. **Remove all GraphQL generation infrastructure** - It's unused and adds complexity
2. **Keep manual HTTP implementations** - They work and are maintainable
3. **Use API documentation as source of truth** - Not introspection schema
4. **Simplify build system** - Remove dev/prod build tags for GraphQL

### For Hardcover.app API
1. **Fix introspection schema** to match actual API structure
2. **Update field names** to be consistent (snake_case vs camelCase)
3. **Document actual query structures** that work at runtime
4. **Test introspection schema** against real API responses

### For Other Developers
1. **Always test introspection schema** against actual API before using GraphQL tooling
2. **Don't assume introspection matches documentation** - verify both
3. **Consider manual HTTP implementations** when GraphQL tooling fails
4. **Use API documentation as primary source** of truth

## Conclusion

The GraphQL introspection schema issues make auto-generation completely unviable for the Hardcover.app API. Manual HTTP implementations are the only reliable approach. This represents a significant limitation of the API's GraphQL implementation and should be addressed by the API maintainers.

## References

- [API Search Documentation](https://docs.hardcover.app/api/guides/searching/)
- [Book Details Guide](https://docs.hardcover.app/api/guides/gettingbookdetails/)
- [Getting All Books Guide](https://docs.hardcover.app/api/guides/gettingallbooksinlibrary/)
- GraphQL Endpoint: `https://api.hardcover.app/v1/graphql`
- API Documentation: https://docs.hardcover.app/api/ 