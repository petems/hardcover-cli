//go:build dev

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"hardcover-cli/internal/config"

	"github.com/spf13/cobra"
)

const (
	unknownType = "Unknown"
	dirPerm     = 0o750
	filePerm    = 0o600
)

// schemaWarningComment is automatically added to all generated schemas to warn about
// known issues with the GraphQL introspection schema not matching the actual API structure
const schemaWarningComment = `#
# IMPORTANT: This schema has known issues with the GraphQL introspection schema not matching
# the actual API structure documented at https://docs.hardcover.app/api/
#
# Known Issues:
# 1. SEARCH FIELD: The search field parameters (query_type, per_page, page, sort, fields, weights)
#    are not properly exposed in introspection but work at runtime.
#    See: https://docs.hardcover.app/api/guides/searching/
#
# 2. BOOK QUERIES: The API uses 'editions' queries with where clauses instead of direct 'book(id: ID!)'
#    queries. The introspection schema may show different field names and structures.
#    See: https://docs.hardcover.app/api/guides/gettingbookdetails/
#
# 3. FIELD NAMES: The API uses snake_case fields (release_date, isbn_10, isbn_13) but introspection
#    may show camelCase fields (publicationYear, pageCount).
#
# 4. LIBRARY QUERIES: Complex queries for user libraries may use different structures than
#    what's exposed in introspection.
#    See: https://docs.hardcover.app/api/guides/gettingallbooksinlibrary/
#
# WORKAROUNDS:
# - Search functionality is implemented manually in cmd/search.go using direct HTTP requests
# - Book details queries may need manual implementation to match the actual API structure
# - Always test queries against the actual API before relying on generated code
#
# DO NOT regenerate this schema without verifying that the introspection schema matches
# the actual API structure documented at https://docs.hardcover.app/api/
`

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Fetch the latest GraphQL schema from the remote endpoint",
	Long: `Fetch the latest GraphQL schema from the Hardcover.app GraphQL API endpoint.
This command downloads the schema using introspection and saves it to the local schema file.
Authentication is handled via the config file or command line arguments.`,
	RunE: runSchemaFetch,
}

func initSchemaCmd() {
	// Check if flags are already added to prevent re-registration
	if schemaCmd.Flags().Lookup("output") == nil {
		schemaCmd.Flags().StringP("output", "o", "", "Output file path (default: internal/client/schema.graphql)")
	}
	if schemaCmd.Flags().Lookup("endpoint") == nil {
		schemaCmd.Flags().StringP("endpoint", "e", "", "GraphQL endpoint URL (default: from config)")
	}
}

func runSchemaFetch(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	cfg, ok := getConfig(ctx)
	if !ok {
		return fmt.Errorf("failed to get configuration")
	}

	// Get and validate flags
	outputPath, endpoint, err := getSchemaFlags(cmd)
	if err != nil {
		return err
	}

	// Set defaults
	outputPath = getDefaultOutputPath(outputPath)
	endpoint = getDefaultEndpoint(endpoint, cfg)

	// Ensure output directory exists
	if dirErr := ensureOutputDirectory(outputPath); dirErr != nil {
		return dirErr
	}

	// Display status
	displaySchemaStatus(endpoint, cfg.APIKey)

	// Fetch schema
	schemaSDL, err := fetchSchema(ctx, endpoint, cfg.APIKey)
	if err != nil {
		return err
	}

	// Write to file
	if err := os.WriteFile(outputPath, []byte(schemaSDL), filePerm); err != nil {
		return fmt.Errorf("failed to write schema file: %w", err)
	}

	fmt.Printf("Schema successfully written to: %s\n", outputPath)
	return nil
}

func introspectionToSDL(introspection map[string]interface{}) (string, error) {
	// Convert introspection JSON to GraphQL SDL
	data, ok := introspection["data"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid introspection response format")
	}

	schema, ok := data["__schema"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid schema in introspection response")
	}

	var sdl strings.Builder
	sdl.WriteString("# Generated GraphQL Schema from Introspection\n")
	sdl.WriteString(fmt.Sprintf("# Fetched on: %s\n", time.Now().Format("2006-01-02 15:04:05 UTC")))
	sdl.WriteString(schemaWarningComment)
	sdl.WriteString("\n")

	// Extract types
	types, ok := schema["types"].([]interface{})
	if !ok {
		return "", fmt.Errorf("invalid types in schema")
	}

	// Process each type
	for _, typeInterface := range types {
		if err := processType(&sdl, typeInterface); err != nil {
			return "", err
		}
	}

	return sdl.String(), nil
}

func processType(sdl *strings.Builder, typeInterface interface{}) error {
	typeMap, ok := typeInterface.(map[string]interface{})
	if !ok {
		return nil
	}

	kind, ok := typeMap["kind"].(string)
	if !ok {
		return nil
	}
	name, ok := typeMap["name"].(string)
	if !ok {
		return nil
	}
	description, _ := typeMap["description"].(string) //nolint:errcheck // description is optional

	// Skip introspection types
	if strings.HasPrefix(name, "__") {
		return nil
	}

	// Add description as comment
	if description != "" {
		fmt.Fprintf(sdl, "\"\"\"\n%s\n\"\"\"\n", description)
	}

	switch kind {
	case "OBJECT":
		return processObjectType(sdl, typeMap, name)
	case "SCALAR":
		fmt.Fprintf(sdl, "scalar %s\n\n", name)
	case "ENUM":
		return processEnumType(sdl, typeMap)
	case "INPUT_OBJECT":
		return processInputObjectType(sdl, typeMap)
	}

	return nil
}

func processObjectType(sdl *strings.Builder, typeMap map[string]interface{}, name string) error {
	// Map query_root to Query for standard GraphQL compatibility
	typeName := name
	if name == "query_root" {
		typeName = "Query"
	}
	fmt.Fprintf(sdl, "type %s {\n", typeName)

	if fields, ok := typeMap["fields"].([]interface{}); ok {
		for _, fieldInterface := range fields {
			field, ok := fieldInterface.(map[string]interface{})
			if !ok {
				continue
			}
			fieldName, ok := field["name"].(string)
			if !ok {
				continue
			}
			fieldType := getTypeString(field["type"])
			fmt.Fprintf(sdl, "  %s: %s\n", fieldName, fieldType)
		}
	}
	sdl.WriteString("}\n\n")
	return nil
}

func processEnumType(sdl *strings.Builder, typeMap map[string]interface{}) error {
	name, ok := typeMap["name"].(string)
	if !ok {
		return nil
	}
	fmt.Fprintf(sdl, "enum %s {\n", name)

	if enumValues, ok := typeMap["enumValues"].([]interface{}); ok {
		for _, valueInterface := range enumValues {
			value, ok := valueInterface.(map[string]interface{})
			if !ok {
				continue
			}
			valueName, ok := value["name"].(string)
			if !ok {
				continue
			}
			fmt.Fprintf(sdl, "  %s\n", valueName)
		}
	}
	sdl.WriteString("}\n\n")
	return nil
}

func processInputObjectType(sdl *strings.Builder, typeMap map[string]interface{}) error {
	name, ok := typeMap["name"].(string)
	if !ok {
		return nil
	}
	fmt.Fprintf(sdl, "input %s {\n", name)

	if inputFields, ok := typeMap["inputFields"].([]interface{}); ok {
		for _, fieldInterface := range inputFields {
			field, ok := fieldInterface.(map[string]interface{})
			if !ok {
				continue
			}
			fieldName, ok := field["name"].(string)
			if !ok {
				continue
			}
			fieldType := getTypeString(field["type"])
			fmt.Fprintf(sdl, "  %s: %s\n", fieldName, fieldType)
		}
	}
	sdl.WriteString("}\n\n")
	return nil
}

func getTypeString(typeInterface interface{}) string {
	typeMap, ok := typeInterface.(map[string]interface{})
	if !ok {
		return unknownType
	}

	kind, ok := typeMap["kind"].(string)
	if !ok {
		return unknownType
	}
	name, ok := typeMap["name"].(string)
	if !ok {
		return unknownType
	}

	switch kind {
	case "NON_NULL":
		return getTypeString(typeMap["ofType"]) + "!"
	case "LIST":
		return "[" + getTypeString(typeMap["ofType"]) + "]"
	case "SCALAR", "OBJECT", "ENUM", "INPUT_OBJECT":
		return name
	default:
		return unknownType
	}
}

func getIntrospectionQuery() string {
	return `query IntrospectionQuery { 
		__schema { 
			queryType { name } 
			mutationType { name } 
			subscriptionType { name } 
			types { ...FullType } 
			directives { 
				name 
				description 
				locations 
				args { ...InputValue } 
			} 
		} 
	} 
	
	fragment FullType on __Type { 
		kind 
		name 
		description 
		fields(includeDeprecated: true) { 
			name 
			description 
			args { ...InputValue } 
			type { ...TypeRef } 
			isDeprecated 
			deprecationReason 
		} 
		inputFields { ...InputValue } 
		interfaces { ...TypeRef } 
		enumValues(includeDeprecated: true) { 
			name 
			description 
			isDeprecated 
			deprecationReason 
		} 
		possibleTypes { ...TypeRef } 
	} 
	
	fragment InputValue on __InputValue { 
		name 
		description 
		type { ...TypeRef } 
		defaultValue 
	} 
	
	fragment TypeRef on __Type { 
		kind 
		name 
		ofType { 
			kind 
			name 
			ofType { 
				kind 
				name 
				ofType { 
					kind 
					name 
					ofType { 
						kind 
						name 
						ofType { 
							kind 
							name 
						} 
					} 
				} 
			} 
		} 
	}`
}

func getSchemaFlags(cmd *cobra.Command) (outputPath, endpoint string, err error) {
	outputPath, err = cmd.Flags().GetString("output")
	if err != nil {
		return "", "", fmt.Errorf("failed to get output flag: %w", err)
	}
	endpoint, err = cmd.Flags().GetString("endpoint")
	if err != nil {
		return "", "", fmt.Errorf("failed to get endpoint flag: %w", err)
	}
	return outputPath, endpoint, nil
}

func getDefaultOutputPath(outputPath string) string {
	if outputPath == "" {
		return "internal/client/schema.graphql"
	}
	return outputPath
}

func getDefaultEndpoint(endpoint string, cfg *config.Config) string {
	if endpoint == "" {
		endpoint = cfg.BaseURL
	}
	if endpoint == "" {
		endpoint = "https://api.hardcover.app/v1/graphql"
	}
	return endpoint
}

func ensureOutputDirectory(outputPath string) error {
	outputDir := filepath.Dir(outputPath)
	if mkdirErr := os.MkdirAll(outputDir, dirPerm); mkdirErr != nil {
		return fmt.Errorf("failed to create output directory: %w", mkdirErr)
	}
	return nil
}

func displaySchemaStatus(endpoint, apiKey string) {
	fmt.Printf("Fetching schema from: %s\n", endpoint)
	if apiKey != "" {
		fmt.Println("Using API key for authentication")
	} else {
		fmt.Println("No API key found - trying without authentication")
	}
}

func fetchSchema(ctx context.Context, endpoint, apiKey string) (string, error) {
	// Create the GraphQL request body
	requestBody := map[string]string{
		"query": getIntrospectionQuery(),
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(string(requestJSON)))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close response body: %v\n", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return "", fmt.Errorf("HTTP %d: failed to read response body: %w", resp.StatusCode, readErr)
		}
		return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var result map[string]interface{}
	if decodeErr := json.NewDecoder(resp.Body).Decode(&result); decodeErr != nil {
		return "", fmt.Errorf("failed to decode response: %w", decodeErr)
	}

	// Check for GraphQL errors
	if errors, ok := result["errors"].([]interface{}); ok && len(errors) > 0 {
		return "", fmt.Errorf("GraphQL errors: %v", errors)
	}

	// Convert introspection result to GraphQL SDL
	schemaSDL, err := introspectionToSDL(result)
	if err != nil {
		return "", fmt.Errorf("failed to convert introspection to SDL: %w", err)
	}

	return schemaSDL, nil
}
