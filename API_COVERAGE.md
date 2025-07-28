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

## 🔍 Documented vs Implemented GraphQL Commands

### 📚 Documented GraphQL Commands (from API Documentation)

Based on the documentation URLs and CLI analysis, here are the GraphQL commands documented in the Hardcover API:

#### 🔍 Search Commands

**🔧 Technical Implementation:**
- **Search Engine**: [Typesense](https://typesense.org/) - Lightning-fast, open-source search engine
- **Shared Index**: Same Typesense index used on the Hardcover website is used for this API endpoint
- **Current Limitations**: Search API does not currently support filtering by parameters besides `query`
- **Available Options**: Can change which attributes (columns) are searched and modify sorting

1. **`search`** - Generic search endpoint with multiple query types
   - **Documented**: ✅ [Searching Guide](https://docs.hardcover.app/api/guides/searching/)
   - **Implemented**: ✅ `hardcover search books <query>` and `hardcover search users <query>`
   - **Missing**: ❌ Author, Character, List, Prompt, Publisher, Series search
   - **GraphQL Query** (from documentation):
     ```graphql
     query SearchBooks($query: String!) {
       search(query: $query, query_type: "Book", per_page: 25, page: 1) {
         ids
         results
         query
         query_type
         page
         per_page
       }
     }
     ```

2. **`search`** - Author search
   - **Documented**: ✅ [Searching Guide](https://docs.hardcover.app/api/guides/searching/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query** (from documentation):
     ```graphql
     query BooksByRowling {
       search(
         query: "rowling",
         query_type: "Author",
         per_page: 5,
         page: 1
       ) {
         results
       }
     }
     ```
   - **Available Fields**: `alternate_names`, `books`, `books_count`, `image`, `name`, `name_personal`, `series_names`, `slug`
   - **Default Fields**: `name,name_personal,alternate_names,series_names,books`
   - **Default Sort**: `_text_match:desc,books_count:desc`
   - **Default Weights**: `3,3,3,2,1`

3. **`search`** - Book search (enhanced)
   - **Documented**: ✅ [Searching Guide](https://docs.hardcover.app/api/guides/searching/)
   - **Implemented**: ✅ `hardcover search books <query>`
   - **GraphQL Query** (from documentation):
     ```graphql
     query LordOfTheRingsBooks {
       search(
         query: "lord of the rings",
         query_type: "Book",
         per_page: 5,
         page: 1
       ) {
         results
       }
     }
     ```
   - **Available Fields**: `activities_count`, `alternative_titles`, `audio_seconds`, `author_names`, `compilation`, `content_warnings`, `contribution_types`, `contributions`, `cover_color`, `description`, `featured_series`, `featured_series_position`, `genres`, `isbns`, `lists_count`, `has_audiobook`, `has_ebook`, `moods`, `pages`, `prompts_count`, `rating`, `ratings_count`, `release_date_i`, `release_year`, `reviews_count`, `series_names`, `slug`, `subtitle`, `tags`, `title`, `users_count`, `users_read_count`
   - **Default Fields**: `title,isbns,series_names,author_names,alternative_titles`
   - **Default Sort**: `_text_match:desc,users_count:desc`
   - **Default Weights**: `5,5,3,1,1`

4. **`search`** - Character search
   - **Documented**: ✅ [Searching Guide](https://docs.hardcover.app/api/guides/searching/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query** (from documentation):
     ```graphql
     query CharactersNamedPeter {
       search(
         query: "peter",
         query_type: "Character",
         per_page: 5,
         page: 1
       ) {
         results
       }
     }
     ```
   - **Available Fields**: `author_names`, `books`, `books_count`, `name`, `object_type`, `slug`
   - **Default Fields**: `name,books,author_names`
   - **Default Sort**: `_text_match:desc,books_count:desc`
   - **Default Weights**: `4,2,2`

5. **`search`** - List search
   - **Documented**: ✅ [Searching Guide](https://docs.hardcover.app/api/guides/searching/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query** (from documentation):
     ```graphql
     query ListsNamedBest {
       search(
         query: "best",
         query_type: "List",
         per_page: 5,
         page: 1
       ) {
         results
       }
     }
     ```
   - **Available Fields**: `description`, `books`, `books_count`, `likes_count`, `object_type`, `name`, `slug`, `user`
   - **Default Fields**: `name,description,books`
   - **Default Sort**: `_text_match:desc,likes_count:desc`
   - **Default Weights**: `3,2,1`

6. **`search`** - Prompt search
   - **Documented**: ✅ [Searching Guide](https://docs.hardcover.app/api/guides/searching/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query** (from documentation):
     ```graphql
     query PromptsAboutLearning {
       search(
         query: "learn from",
         query_type: "Prompt",
         per_page: 5,
         page: 1
       ) {
         results
       }
     }
     ```
   - **Available Fields**: `answers_count`, `books`, `books_count`, `question`, `slug`, `user`, `users_count`
   - **Default Fields**: `question,books`
   - **Default Sort**: `_text_match:desc`
   - **Default Weights**: `2,1`

7. **`search`** - Publisher search
   - **Documented**: ✅ [Searching Guide](https://docs.hardcover.app/api/guides/searching/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query** (from documentation):
     ```graphql
     query PublishersNamedPenguin {
       search(
         query: "penguin",
         query_type: "Publisher",
         per_page: 5,
         page: 1
       ) {
         results
       }
     }
     ```
   - **Available Fields**: `editions_count`, `name`, `object_type`, `slug`
   - **Default Fields**: `name`
   - **Default Sort**: `_text_match:desc,editions_count:desc`
   - **Default Weights**: `1`

8. **`search`** - Series search
   - **Documented**: ✅ [Searching Guide](https://docs.hardcover.app/api/guides/searching/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query** (from documentation):
     ```graphql
     query SeriesNamedHarryPotter {
       search(
         query: "harry potter",
         query_type: "Series",
         per_page: 7,
         page: 1
       ) {
         results
       }
     }
     ```
   - **Available Fields**: `author_name`, `author`, `books_count`, `books`, `name`, `primary_books_count`, `readers_count`, `slug`
   - **Default Fields**: `name,books,author_name`
   - **Default Sort**: `_text_match:desc,readers_count:desc`
   - **Default Weights**: `2,1,1`

9. **`search`** - User search (enhanced)
   - **Documented**: ✅ [Searching Guide](https://docs.hardcover.app/api/guides/searching/)
   - **Implemented**: ✅ `hardcover search users <query>`
   - **GraphQL Query** (from documentation):
     ```graphql
     query UsersNamedAdam {
       search(
         query: "adam",
         query_type: "User",
         per_page: 5,
         page: 1
       ) {
         results
       }
     }
     ```
   - **Available Fields**: `books_count`, `flair`, `followed_users_count`, `followers_count`, `image`, `location`, `name`, `pro`, `username`
   - **Default Fields**: `name,username,location`
   - **Default Sort**: `_text_match:desc,followers_count:desc`
   - **Default Weights**: `2,2,1`

#### 📚 Book Commands
10. **`list_books`** - Get user's book library
   - **Documented**: ✅ [Getting All Books Guide](https://docs.hardcover.app/api/guides/gettingallbooksinlibrary/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query** (from documentation):
     ```graphql
     {
       list_books(
         where: {
           user_books: {
             user_id: {_eq: ##USER_ID##}
           }
         },
         distinct_on: book_id
         limit: 5
         offset: 0
       ) {
         book {
           title
           pages
           release_date
         }
       }
     }
     ```

11. **`books`** - Get books by author (from Books Schema)
   - **Documented**: ✅ [Books Schema](https://docs.hardcover.app/api/graphql/schemas/books/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query** (from documentation):
     ```graphql
     query BooksByUserCount {
       books(
         where: {
           contributions: {
             author: {
               name: {_eq: "Brandon Sanderson"}
             }
           }
         }
         limit: 10
         order_by: {users_count: desc}
       ) {
         pages
         title
         id
       }
     }
     ```

12. **`editions`** - Get book editions (from Books Schema)
   - **Documented**: ✅ [Books Schema](https://docs.hardcover.app/api/graphql/schemas/books/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query** (from documentation):
     ```graphql
     query GetEditionsFromTitle {
       editions(where: {title: {_eq: "Oathbringer"}}) {
         id
         title
         edition_format
         pages
         release_date
         isbn_10
         isbn_13
         publisher {
           name
         }
       }
     }
     ```

13. **`createBook`** - Create a new book (Mutation)
   - **Documented**: ✅ [Books Schema](https://docs.hardcover.app/api/graphql/schemas/books/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Mutation** (from documentation):
     ```graphql
     mutation {
       createBook(input: {
         title: "My First Book",
         pages: 300,
         release_date: "2024-09-07"
         description: "This is my first book."
       }) {
         book {
           title
           pages
           release_date
           description
         }
       }
     }
     ```

#### 👤 User Commands
14. **`me`** - Get current user profile
   - **Documented**: ✅ [Users Schema](https://docs.hardcover.app/api/graphql/schemas/users/)
   - **Implemented**: ✅ `hardcover me`
   - **GraphQL Query**:
     ```graphql
     query GetCurrentUser {
       me {
         id
         username
       }
     }
     ```

#### 👥 Author Commands
15. **`list_authors`** - Get author information
   - **Documented**: ✅ [Authors Schema](https://docs.hardcover.app/api/graphql/schemas/authors/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query**: Not shown in documentation

#### 📖 Edition Commands
16. **`editions`** - Get edition details by ISBN
   - **Documented**: ✅ [Editions Schema](https://docs.hardcover.app/api/graphql/schemas/editions/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query** (from documentation):
     ```graphql
     query GetEditionByISBN {
       editions(where: {isbn_13: {_eq: "9780547928227"}}) {
           id
           title
           subtitle
           isbn_13
           isbn_10
           asin
           pages
           release_date
           physical_format
           publisher {
               name
           }
           book {
               id
               title
               rating
               contributions {
                   author {
                       name
                   }
               }
           }
           language {
               language
           }
           reading_format {
               format
           }
       }
     }
     ```

17. **`editions`** - Get all editions of a book
   - **Documented**: ✅ [Editions Schema](https://docs.hardcover.app/api/graphql/schemas/editions/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query** (from documentation):
     ```graphql
     query GetBookEditions {
       editions(
           where: {book_id: {_eq: 328491}}
           order_by: {release_date: desc}
       ) {
           id
           title
           isbn_13
           pages
           release_date
           physical_format
           publisher {
               name
           }
           language {
               language
           }
           reading_format {
               format
           }
           users_count
           rating
       }
     }
     ```

18. **`editions`** - Find editions by publisher
    - **Documented**: ✅ [Editions Schema](https://docs.hardcover.app/api/graphql/schemas/editions/)
    - **Implemented**: ❌ Not implemented in CLI
    - **GraphQL Query** (from documentation):
      ```graphql
      query GetPublisherEditions {
        editions(
            where: {publisher_id: {_eq: 1}}
            order_by: {release_date: desc}
            limit: 10
        ) {
            id
            title
            isbn_13
            release_date
            physical_format
            book {
                title
                rating
                contributions {
                    author {
                        name
                    }
                }
            }
        }
      }
      ```

19. **`editions`** - Search editions by format (audiobooks)
    - **Documented**: ✅ [Editions Schema](https://docs.hardcover.app/api/graphql/schemas/editions/)
    - **Implemented**: ❌ Not implemented in CLI
    - **GraphQL Query** (from documentation):
      ```graphql
      query GetAudiobookEditions {
        editions(
            where: {reading_format_id: {_eq: 2}, audio_seconds: {_gt: 0}}
            order_by: {users_count: desc}
            limit: 10
        ) {
            id
            title
            asin
            audio_seconds
            publisher {
                name
            }
            cached_contributors
            book {
                title
                rating
            }
        }
      }
      ```

#### 📊 Activity Commands
20. **`activities`** - Get user's activities
    - **Documented**: ✅ [Activities Schema](https://docs.hardcover.app/api/graphql/schemas/activities/)
    - **Implemented**: ❌ Not implemented in CLI
    - **GraphQL Query** (from documentation):
      ```graphql
      {
        activities(where: {user_id: {_eq: ##USER_ID##}}, limit: 10) {
            event
            likes_count
            book_id
            created_at
        }
      }
      ```

21. **`activities`** - Get activities for a specific book
     - **Documented**: ✅ [Activities Schema](https://docs.hardcover.app/api/graphql/schemas/activities/)
     - **Implemented**: ❌ Not implemented in CLI
     - **GraphQL Query** (from documentation):
       ```graphql
       {
           activities(
                 order_by: {created_at: desc}
                 where: {book_id: {_eq: 10257}, event: {_eq: "UserBookActivity"}}
                 limit: 10
           ) {
                 data
                 event
                 object_type
                 book_id
           }
       }
       ```

#### 🎭 Character Commands
22. **`characters`** - Get character by ID
     - **Documented**: ✅ [Characters Schema](https://docs.hardcover.app/api/graphql/schemas/characters/)
     - **Implemented**: ❌ Not implemented in CLI
     - **GraphQL Query** (from documentation):
       ```graphql
       query {
         characters(where: {id: {_eq: "1"}}, limit: 1) {
           id,
           name
         }
       }
       ```

23. **`characters`** - Get character by name
     - **Documented**: ✅ [Characters Schema](https://docs.hardcover.app/api/graphql/schemas/characters/)
     - **Implemented**: ❌ Not implemented in CLI
     - **GraphQL Query** (from documentation):
       ```graphql
       query {
         characters(where: {name: {_eq: "Harry Potter"}}) {
           biography
           slug
           state
           name
         }
       }
       ```

24. **`characters`** - Get all characters
     - **Documented**: ✅ [Characters Schema](https://docs.hardcover.app/api/graphql/schemas/characters/)
     - **Implemented**: ❌ Not implemented in CLI
     - **GraphQL Query** (from documentation):
       ```graphql
       query {
         characters(limit: 10) {
         id,
         name
       }
     }
     ```

13. **`characters`** - Get books featuring a character
   - **Documented**: ✅ [Characters Schema](https://docs.hardcover.app/api/graphql/schemas/characters/)
   - **Implemented**: ❌ Not implemented in CLI
   - **GraphQL Query** (from documentation):
     ```graphql
     query GetCharacterBooks {
       characters(where: {name: {_eq: "Harry Potter"}}) {
         name
         book_characters {
           book {
             title
           }
         }
         contributions {
           author {
             name
           }
         }
       }
     }
     ```

### 📊 Implementation Status Summary

| Command Category | Documented | Implemented | CLI Command | Status |
|------------------|------------|-------------|-------------|---------|
| **Search Books** | ✅ | ✅ | `hardcover search books <query>` | Complete |
| **Search Users** | ✅ | ✅ | `hardcover search users <query>` | Complete |
| **Search Authors** | ✅ | ❌ | `hardcover search authors <query>` | Missing |
| **Search Characters** | ✅ | ❌ | `hardcover search characters <query>` | Missing |
| **Search Lists** | ✅ | ❌ | `hardcover search lists <query>` | Missing |
| **Search Prompts** | ✅ | ❌ | `hardcover search prompts <query>` | Missing |
| **Search Publishers** | ✅ | ❌ | `hardcover search publishers <query>` | Missing |
| **Search Series** | ✅ | ❌ | `hardcover search series <query>` | Missing |
| **User Profile** | ✅ | ✅ | `hardcover me` | Complete |
| **Book Library** | ✅ | ❌ | `hardcover book list` | Missing |
| **Books by Author** | ✅ | ❌ | `hardcover book by-author <author>` | Missing |
| **Book Editions** | ✅ | ❌ | `hardcover edition list <book>` | Missing |
| **Edition by ISBN** | ✅ | ❌ | `hardcover edition isbn <isbn>` | Missing |
| **Edition by Publisher** | ✅ | ❌ | `hardcover edition publisher <id>` | Missing |
| **Edition by Format** | ✅ | ❌ | `hardcover edition format <format>` | Missing |
| **Create Book** | ✅ | ❌ | `hardcover book create` | Missing |
| **Author Search** | ✅ | ❌ | `hardcover search authors` | Missing |
| **Author Details** | ✅ | ❌ | `hardcover author get <id>` | Missing |
| **Edition Details** | ✅ | ❌ | `hardcover edition get <id>` | Missing |
| **User Activities** | ✅ | ❌ | `hardcover activity list` | Missing |
| **Book Activities** | ✅ | ❌ | `hardcover activity book <id>` | Missing |
| **Character by ID** | ✅ | ❌ | `hardcover character get <id>` | Missing |
| **Character by Name** | ✅ | ❌ | `hardcover character search <name>` | Missing |
| **Character List** | ✅ | ❌ | `hardcover character list` | Missing |
| **Character Books** | ✅ | ❌ | `hardcover character books <name>` | Missing |

### 🎯 Coverage Statistics

- **Total Documented Commands**: 27 major command categories
- **Implemented in CLI**: 3 commands (11%)
- **Missing Implementation**: 24 commands (89%)
- **Configuration Commands**: 3 additional commands (100% implemented)

## 🚧 Implementation Caveats

### API Limitations & Requirements (Critical)
- **Rate Limiting**: 60 requests per minute (1 request per second)
- **Query Timeout**: Maximum 30 seconds per query
- **Token Expiration**: API tokens expire after 1 year and reset on January 1st
- **Query Depth**: Maximum depth of 3 levels
- **Data Access**: Limited to own user data, public data, and followed users' data
- **Disabled Queries**: `_like`, `_nlike`, `_ilike`, `_niregex`, `_nregex`, `_iregex`, `_regex`, `_nsimilar`, `_similar`
- **Security**: Tokens must be kept secure, not used in browser environments
- **User Agent**: Recommended to include user-agent header with script description

### Search Architecture
- **Search Engine**: Hardcover uses [Typesense](https://typesense.org/) as the underlying search engine
- **Shared Infrastructure**: API search endpoint uses the same Typesense index as the website
- **Search Limitations**: Currently only supports `query` parameter filtering, no additional filter parameters
- **Flexible Options**: Supports customizing searchable attributes and sorting criteria

### GraphQL Schema Issues (Critical)
- **Fundamental Mismatches**: GraphQL introspection schema doesn't match actual API structure
- **Auto-Generation Abandoned**: Project has moved to manual HTTP implementations
- **Documentation Mismatch**: API docs don't match introspection schema
- **Reference**: See `HARDCOVER_API_ISSUE.md` for complete analysis

### Current Implementation Approach
- **Manual HTTP Requests**: All GraphQL queries use direct HTTP POST requests
- **Type-Safe Structs**: Custom Go structs for API responses
- **Error Handling**: Basic error handling for API failures
- **No Code Generation**: Abandoned GraphQL code generation tools

### Required Implementation Updates (Based on API Documentation)
- **Rate Limiting**: ❌ Not implemented - need to add 60 req/min limit
- **Timeout Handling**: ❌ Not implemented - need 30-second timeout
- **User Agent**: ❌ Not implemented - need to add user-agent header
- **Token Security**: ⚠️ Partially implemented - need better token validation
- **Query Validation**: ❌ Not implemented - need to check query depth and operators
- **Retry Logic**: ❌ Not implemented - need exponential backoff for 429s

### API Coverage Gaps
- **Limited Search Types**: Only book search implemented, missing authors and users
- **No Write Operations**: All current operations are read-only
- **No Pagination**: Some endpoints lack proper pagination support
- **No Book Listing**: Book list functionality was removed

## 🎯 Next Steps Priority

### 🔧 Critical Infrastructure Improvements (Based on API Limitations)

1. **Rate Limiting Implementation**
   - Add rate limiting to respect 60 requests/minute limit
   - Implement exponential backoff for 429 responses
   - Add request queuing for high-frequency operations
   - **Suggested**: Use `golang.org/x/time/rate` package

2. **Timeout & Error Handling**
   - Set 30-second timeout for all GraphQL queries
   - Implement proper error handling for 401, 403, 404, 429, 500 responses
   - Add retry logic for transient failures
   - **Suggested**: Add `--timeout` flag for custom timeouts

3. **User Agent & Security**
   - Add user-agent header: `hardcover-cli/v1.0.0`
   - Ensure API tokens are never logged or exposed
   - Add token validation and expiration checking
   - **Suggested**: Add `--user-agent` flag for custom user agents

4. **Query Optimization**
   - Ensure all queries respect 3-level depth limit
   - Avoid disabled query operators (`_like`, `_regex`, etc.)
   - Implement query caching to reduce API calls
   - **Suggested**: Add query validation before sending

### High Priority (Core Functionality)
5. **🔍 Complete Search API**
   - Implement author search (`hardcover search authors <query>`)
   - Add search result filtering and sorting
   - Respect rate limiting for search operations

6. **📚 Book Details**
   - Implement `hardcover book get <id>` command
   - Display comprehensive book information
   - Use proper API structure from documentation

7. **👥 Author Management**
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

## 🔧 Implementation Recommendations

### Rate Limiting Implementation
```go
// Suggested implementation using golang.org/x/time/rate
import "golang.org/x/time/rate"

var limiter = rate.NewLimiter(rate.Every(time.Second), 60) // 60 requests per minute

func makeAPIRequest(query string) error {
    if err := limiter.Wait(context.Background()); err != nil {
        return fmt.Errorf("rate limit exceeded: %w", err)
    }
    // ... make API request
}
```

### Timeout & Error Handling
```go
// Suggested timeout implementation
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

// Handle specific error codes
switch resp.StatusCode {
case 429:
    return fmt.Errorf("rate limited, retry after %s", resp.Header.Get("Retry-After"))
case 401:
    return fmt.Errorf("invalid or expired token")
case 403:
    return fmt.Errorf("access denied to requested resource")
}
```

### User Agent Header
```go
// Add to all API requests
req.Header.Set("User-Agent", "hardcover-cli/v1.0.0")
```

## 📚 API Documentation References

### Core Documentation
- [Getting Started](https://docs.hardcover.app/api/getting-started/) ✅ **Referenced for limitations**
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

*Last updated: July 28, 2025*
*Coverage: 25% of estimated API endpoints*
*Status: GraphQL auto-generation abandoned, using manual HTTP implementations* 