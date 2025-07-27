# Hardcover CLI API Coverage

This document outlines the current API coverage of the Hardcover CLI implementation, highlighting what's implemented, what's missing, and what should be prioritized next.

## 📋 API Coverage Summary

### ✅ Implemented Features

#### 🔍 Search API
- ✅ **Book Search** (`hardcover search books <query>`)
  - Search by title, author, or other criteria
  - Returns book details with authors, ratings, genres
  - Supports pagination and result count
  - **Implementation**: `cmd/search.go` with GraphQL query

**Live Example:**
```bash
$ hardcover search books "golang"
1. Building RESTful Web services with Go: Learn how to build powerful RESTful APIs with Golang that scale gracefully
   Subtitle: Learn how to build powerful RESTful APIs with Golang that scale gracefully
   Authors: Naren Yellavula
   Edition ID: 1676108
   URL: https://hardcover.app/books/building-restful-web-services-with-go

-----------------------------
2. Mastering Go: Create Golang production applications using network libraries, concurrency, machine learning, and advanced data structures
   Subtitle: Create Golang production applications using network libraries, concurrency, machine learning, and advanced data structures
   Authors: Mihalis Tsoukalos
   Year: 2019
   Edition ID: 1863717
   URL: https://hardcover.app/books/mastering-go-create-golang-production-applications-using-network-libraries-concurrency-machine-learning-and-advanced-data-structures
   Rating: 4.00/5 (1 ratings)
   ISBNs: 1838555323, 9781838555320

-----------------------------
```

#### 📚 Book Management
- ✅ **Search Books** (`hardcover search books <query>`)
  - Search for books by title, author, or other criteria
  - Returns book details with authors, ratings, genres
  - Supports pagination and result count
  - **Implementation**: `cmd/search.go` with GraphQL query

**Live Example:**
```bash
$ hardcover search books "golang"
1. Building RESTful Web services with Go
   Subtitle: Learn how to build powerful RESTful APIs with Golang that scale gracefully
   Authors: Naren Yellavula
   Edition ID: 1676108
   URL: https://hardcover.app/books/building-restful-web-services-with-go

-----------------------------
2. Mastering Go
   Subtitle: Create Golang production applications using network libraries, concurrency, machine learning, and advanced data structures
   Authors: Mihalis Tsoukalos
   Year: 2019
   Edition ID: 1863717
   URL: https://hardcover.app/books/mastering-go-create-golang-production-applications-using-network-libraries-concurrency-machine-learning-and-advanced-data-structures
   Rating: 4.00/5 (1 ratings)
   ISBNs: 1838555323, 9781838555320

-----------------------------
```

#### 👤 User Management
- ✅ **Get Current User Profile** (`hardcover me`)
  - User ID, username
  - **Implementation**: `cmd/me.go` with GraphQL query

**Live Example:**
```bash
$ hardcover me
User Profile:
  ID: 12345
  Username: johndoe
```

- ✅ **Search Users** (`hardcover search users <query>`)
  - Search for users by name, username, or location
  - Returns user profiles with stats and metadata
  - **Implementation**: `cmd/search.go` with GraphQL query

**Live Example:**
```bash
$ hardcover search users "john"
1. johndoe
   Name: John Doe
   Location: New York, NY
   Books: 150
   Followers: 45
   Following: 23
   Pro: Yes
   Has Image: Yes

-----------------------------
```

#### ⚙️ Configuration
- ✅ **API Key Management**
  - Set/get API key via config file
  - Environment variable support
  - **Implementation**: `cmd/config.go`

**Live Examples:**
```bash
# Set API key
$ hardcover config set-api-key "your-api-key-here"

# Get API key (masked for security)
$ hardcover config get-api-key
Current API key: hc_************************

# Show config file path
$ hardcover config show-path
Configuration file: /home/user/.hardcover/config.yaml
```

### ❌ Missing Features

#### 🔍 Search API
- ❌ **Author Search** (`hardcover search authors <query>`)
  - Search for authors by name
  - Author details and bibliography
  - **Missing**: No implementation in search commands

#### 📚 Book Management
- ❌ **Get Book Details** (`hardcover book get <id>`)
  - Retrieve comprehensive book information
  - Includes authors, contributors, genres, ratings
  - Publication details (year, pages, ISBN)
  - Cover image and timestamps
  - **Missing**: No implementation in book commands

- ❌ **Book Listing** (`hardcover book list`)
  - List all books with pagination

- ❌ **Book Editions**
  - Different editions of the same book
  - Edition-specific details (format, publisher, etc.)
  - **Missing**: No edition-related commands

- ❌ **Book Activities**
  - Reading progress, reviews, ratings
  - User interactions with books
  - **Missing**: No activity tracking commands

#### 👥 Author Management
- ❌ **Author Details** (`hardcover author get <id>`)
  - Author biography and information
  - Published works and bibliography
  - **Missing**: No author-specific commands

#### 🎭 Character Management
- ❌ **Character Information**
  - Character details from books
  - Character relationships and appearances
  - **Missing**: No character-related commands

#### 📖 Edition Management
- ❌ **Edition Details**
  - Different book formats and editions
  - Publisher and publication information
  - **Missing**: No edition-specific commands

#### 📊 Activity Management
- ❌ **User Activities**
  - Reading progress tracking
  - Reviews and ratings management
  - Reading lists and shelves
  - **Missing**: No activity-related commands

## 🚧 Implementation Caveats

### GraphQL Schema Issues (Critical)
- **Fundamental Mismatches**: GraphQL introspection schema doesn't match actual API structure
- **Auto-Generation Abandoned**: Project has moved to manual HTTP implementations
- **Documentation Mismatch**: API docs don't match introspection schema
- **Reference**: See `HARDCOVER_API_ISSUE.md` for complete analysis

### Current Implementation Approach
- **Manual HTTP Requests**: All GraphQL queries use direct HTTP POST requests
- **Type-Safe Structs**: Custom Go structs for API responses
- **Error Handling**: Proper error handling for API failures
- **No Code Generation**: Abandoned GraphQL code generation tools

### API Coverage Gaps
- **Limited Search Types**: Only book search implemented, missing authors and users
- **No Write Operations**: All current operations are read-only
- **No Pagination**: Some endpoints lack proper pagination support
- **No Book Listing**: Book list functionality was removed

## 🎯 Next Steps Priority

### High Priority (Core Functionality)
1. **🔍 Complete Search API**
   - Implement author search (`hardcover search authors <query>`)
   - Add search result filtering and sorting

2. **📚 Book Details**
   - Implement `hardcover book get <id>` command
   - Display comprehensive book information
   - Use proper API structure from documentation

3. **👥 Author Management**
   - Add `hardcover author get <id>` command
   - Display author biography and bibliography
   - Link authors to their books

### Medium Priority (Enhanced Features)
4. **📖 Edition Support**
   - Add edition information to book details
   - Implement `hardcover edition get <id>` command
   - Show different formats and publishers

5. **📊 Activity Tracking**
   - Add reading progress commands
   - Implement review and rating management
   - Add reading list functionality

6. **🎭 Character Information**
   - Add character details to book information
   - Implement character search and details

### Low Priority (Nice to Have)
7. **📈 Analytics**
   - Reading statistics and trends
   - Genre preferences analysis
   - Reading goal tracking

8. **🔗 Social Features**
   - Friend/follower management
   - Social reading recommendations
   - Community features

## 📚 API Documentation References

### Guides
- [Getting All Books in Library](https://docs.hardcover.app/api/guides/gettingallbooksinlibrary/) ❌ **Not Implemented**
- [Getting Book Details](https://docs.hardcover.app/api/guides/gettingbookdetails/) ❌ **Not Implemented**
- [Searching](https://docs.hardcover.app/api/guides/searching/) ✅ **Partially Implemented**

### GraphQL Schemas
- [Authors Schema](https://docs.hardcover.app/api/graphql/schemas/authors/) ❌ **Not Implemented**
- [Books Schema](https://docs.hardcover.app/api/graphql/schemas/books/) ✅ **Partially Implemented**
- [Editions Schema](https://docs.hardcover.app/api/graphql/schemas/editions/) ❌ **Not Implemented**
- [Activities Schema](https://docs.hardcover.app/api/graphql/schemas/activities/) ❌ **Not Implemented**
- [Characters Schema](https://docs.hardcover.app/api/graphql/schemas/characters/) ❌ **Not Implemented**
- [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/) ✅ **Implemented**

## 🛠️ Technical Debt

### Immediate Fixes Needed
1. **GraphQL Schema Issues**: Fundamental mismatches prevent auto-generation
2. **Book Details Implementation**: Add book details retrieval functionality
3. **Error Handling**: Add proper error handling for API failures
4. **Rate Limiting**: Implement rate limit handling

### Code Quality Improvements
1. **Consistency**: Standardize command structure across all features
2. **Documentation**: Add inline documentation for complex queries
3. **Configuration**: Improve configuration validation
4. **Logging**: Add structured logging for debugging

### Architecture Decisions
1. **Manual HTTP**: Continue using manual HTTP implementations
2. **No Auto-Generation**: Abandon GraphQL code generation entirely
3. **API Documentation**: Use API docs as source of truth, not introspection
4. **Type Safety**: Maintain type-safe Go structs for all responses

## 📊 Coverage Statistics

- **Total API Endpoints**: ~15-20 estimated
- **Implemented**: 4 endpoints (20-25%)
- **Partially Implemented**: 0 endpoints (0%)
- **Missing**: 11-16 endpoints (70-80%)
- **Removed**: 1 endpoint (book listing)

## 🔄 Development Workflow

### For New Features
1. **Research**: Check API documentation for endpoint details
2. **Manual Implementation**: Use direct HTTP requests (no GraphQL generation)
3. **Type Safety**: Create proper Go structs for responses
4. **Testing**: Add comprehensive unit and integration tests
5. **Documentation**: Update this coverage document

### For Bug Fixes
1. **Reproduce**: Create minimal test case
2. **Fix**: Implement solution with tests
3. **Verify**: Ensure all existing functionality still works
4. **Document**: Update relevant documentation

### For API Integration
1. **Use API Docs**: Always reference API documentation, not introspection
2. **Manual HTTP**: Implement direct HTTP requests for GraphQL queries
3. **Error Handling**: Handle API-specific error responses
4. **Testing**: Test against real API endpoints

## 🚨 Critical Issues

### GraphQL Introspection Problems
- **Schema Mismatches**: Every GraphQL query has introspection issues
- **Code Generation Failure**: Auto-generation tools cannot produce working code
- **Documentation Discrepancies**: API docs don't match introspection schema
- **Runtime Errors**: Generated code fails with unmarshaling errors

### Recommended Approach
- **Manual HTTP**: Continue using manual HTTP implementations
- **API Documentation**: Use docs as source of truth
- **Type Safety**: Maintain custom Go structs
- **Testing**: Test against real API responses

---

*Last updated: July 27, 2025*
*Coverage: 25% of estimated API endpoints*
*Status: GraphQL auto-generation abandoned, using manual HTTP implementations* 