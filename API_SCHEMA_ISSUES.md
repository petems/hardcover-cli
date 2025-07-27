# Hardcover API Schema Issues and Workarounds

This document describes known issues with the GraphQL introspection schema not matching the actual Hardcover API structure, and the workarounds implemented in this codebase.

## Overview

The GraphQL introspection schema at `https://api.hardcover.app/v1/graphql` has several discrepancies with the actual API structure documented at https://docs.hardcover.app/api/. This creates challenges for code generation tools like `genqlient` that rely on the introspection schema.

## Known Issues

### 1. Search Field Parameters

**Issue**: The search field parameters are not properly exposed in the introspection schema.

**Expected Schema** (per API docs):
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

**Actual Introspection Result**:
```graphql
type Query {
  search: SearchOutput
}
```

**Impact**: Code generation tools cannot generate type-safe code for search functionality.

**Workaround**: Manual HTTP implementation in `cmd/search.go` using direct GraphQL queries.

**Reference**: https://docs.hardcover.app/api/guides/searching/

### 2. Book Query Structure

**Issue**: The API uses `editions` queries with where clauses instead of direct `book(id: ID!)` queries.

**Expected Schema** (per API docs):
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

**Actual Introspection Result**: Shows `book(id: ID!): Book` which may not work as expected.

**Impact**: Generated code may not work with the actual API structure.

**Workaround**: Manual HTTP implementation in `cmd/book_manual.go` using the correct query structure.

**Reference**: https://docs.hardcover.app/api/guides/gettingbookdetails/

### 3. Field Name Mismatches

**Issue**: The API uses snake_case field names but introspection may show camelCase.

**Examples**:
- API: `release_date`, `isbn_10`, `isbn_13`, `edition_format`
- Introspection: `publicationYear`, `pageCount`, `editionFormat`

**Impact**: Generated code may reference non-existent fields.

**Workaround**: Manual implementations use the correct snake_case field names.

### 4. Library Queries

**Issue**: Complex queries for user libraries may use different structures than what's exposed in introspection.

**Reference**: https://docs.hardcover.app/api/guides/gettingallbooksinlibrary/

**Impact**: Library-related functionality may not work with generated code.

**Workaround**: Not yet implemented - would need manual HTTP implementation.

## Implemented Workarounds

### 1. Search Functionality

**File**: `cmd/search.go`
**Implementation**: `performManualSearch()`
**Features**:
- Direct HTTP requests to GraphQL endpoint
- Proper parameter handling (query_type, per_page, page)
- Error handling and response parsing
- Type-safe response structure

**Usage**:
```bash
./bin/hardcover-cli-dev search books "lord of the rings"
```

### 2. Book Details Functionality

**File**: `cmd/book_manual.go`
**Implementation**: `performManualBookDetails()`
**Features**:
- Uses correct `editions` query structure
- Proper field name mapping (snake_case)
- Publisher and contribution data extraction
- Error handling for missing books

**Usage**:
```bash
./bin/hardcover-cli-dev book get 12345
```

### 3. Schema Warning Comments

**File**: `cmd/schema.go`
**Implementation**: `schemaWarningComment` constant
**Features**:
- Automatically added to all generated schemas
- Documents all known issues
- Provides references to API documentation
- Warns against accidental regeneration

## Code Generation Strategy

### Development Builds (`-tags dev`)
- Includes manual implementations
- Uses generated code where possible
- Falls back to manual HTTP requests for problematic queries

### Production Builds
- Excludes manual implementations
- Uses stub implementations that return errors
- Prevents accidental use of development-only features

## Testing Strategy

### Manual Testing
- Test search functionality with various queries
- Test book details with different book IDs
- Verify error handling for invalid inputs

### Schema Validation
- Compare generated schema with API documentation
- Identify new discrepancies
- Update manual implementations as needed

## Future Improvements

### 1. Additional Manual Implementations
- Library queries (user books)
- Author queries
- Publisher queries
- Activity queries

### 2. Schema Monitoring
- Automated comparison of introspection vs. documentation
- Alert when new discrepancies are found
- Track API changes over time

### 3. API Documentation Integration
- Generate manual implementations from API docs
- Validate manual implementations against docs
- Keep implementations in sync with API changes

## Maintenance Guidelines

### When Regenerating Schema
1. **Always test** the generated code against the actual API
2. **Check for new issues** by comparing with API documentation
3. **Update manual implementations** if API structure changes
4. **Preserve warning comments** in generated schemas

### When Adding New Features
1. **Check API documentation** first
2. **Test against actual API** before implementing
3. **Use manual HTTP requests** if introspection schema is insufficient
4. **Document any workarounds** in this file

### When Updating Dependencies
1. **Test all manual implementations** after updates
2. **Verify schema generation** still works
3. **Check for new introspection issues**
4. **Update warning comments** if needed

## References

- [Hardcover API Documentation](https://docs.hardcover.app/api/)
- [Search API Guide](https://docs.hardcover.app/api/guides/searching/)
- [Book Details API Guide](https://docs.hardcover.app/api/guides/gettingbookdetails/)
- [Library API Guide](https://docs.hardcover.app/api/guides/gettingallbooksinlibrary/)
- [GraphQL Schemas](https://docs.hardcover.app/api/graphql/schemas/)

## Contributing

When contributing to this codebase:

1. **Read this document** before making changes
2. **Test against the actual API** not just the introspection schema
3. **Use manual implementations** when introspection is insufficient
4. **Update this documentation** when new issues are discovered
5. **Preserve existing workarounds** unless the underlying issue is fixed 