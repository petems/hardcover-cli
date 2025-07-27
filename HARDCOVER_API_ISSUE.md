# Hardcover API GraphQL Introspection Issue

## Summary

The GraphQL introspection schema at `https://api.hardcover.app/v1/graphql` does not expose the parameters for the `search` field, even though the API documentation and runtime behavior support them.

## Issue Details

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

## Impact

- **Code Generation Tools**: Tools like `genqlient` cannot generate type-safe code for the search functionality because the introspection schema doesn't match the runtime API.
- **Developer Experience**: Developers need to manually patch their local schema files or use raw HTTP requests for search functionality.
- **Documentation Mismatch**: The API documentation at https://docs.hardcover.app/api/guides/searching/ shows parameters that aren't reflected in the introspection schema.

## Workaround

Currently, developers need to manually patch their local `schema.graphql` files to add the missing parameters, then regenerate their client code.

## Request

Could the GraphQL introspection schema be updated to include the search field parameters? This would enable proper code generation and improve the developer experience.

## References

- [API Search Documentation](https://docs.hardcover.app/api/guides/searching/)
- GraphQL Endpoint: `https://api.hardcover.app/v1/graphql`

## Example Query That Works at Runtime

```graphql
query SearchBooks {
  search(
    query: "lord of the rings"
    query_type: "Book"
    per_page: 5
    page: 1
  ) {
    results
    ids
    query
    query_type
  }
}
```

This query works perfectly at runtime but cannot be validated or code-generated from the current introspection schema. 