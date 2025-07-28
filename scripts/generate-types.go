package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"hardcover-cli/internal/config"
)

const introspectionQuery = `
query IntrospectionQuery {
  __schema {
    queryType {
      name
    }
    mutationType {
      name
    }
    subscriptionType {
      name
    }
    types {
      ...FullType
    }
    directives {
      name
      description
      locations
      args {
        ...InputValue
      }
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
    args {
      ...InputValue
    }
    type {
      ...TypeRef
    }
    isDeprecated
    deprecationReason
  }
  inputFields {
    ...InputValue
  }
  interfaces {
    ...TypeRef
  }
  enumValues(includeDeprecated: true) {
    name
    description
    isDeprecated
    deprecationReason
  }
  possibleTypes {
    ...TypeRef
  }
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
    }
  }
}
`

const testQuery = `
query {
  me {
    id
    username
  }
}
`

type GraphQLRequest struct {
	Query string `json:"query"`
}

type GraphQLResponse struct {
	Data struct {
		Schema Schema `json:"__schema"`
	} `json:"data"`
}

type Schema struct {
	Types []Type `json:"types"`
}

type Type struct {
	Kind        string  `json:"kind"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Fields      []Field `json:"fields,omitempty"`
	EnumValues  []Enum  `json:"enumValues,omitempty"`
}

type Field struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        TypeRef `json:"type"`
	Args        []Input `json:"args"`
}

type Input struct {
	Type        TypeRef `json:"type"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

type Enum struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TypeRef struct {
	OfType *TypeRef `json:"ofType"`
	Kind   string   `json:"kind"`
	Name   string   `json:"name"`
}

const typesTemplate = `// Code generated from GraphQL schema, DO NOT EDIT.
// Generated at: {{.Timestamp}}

package client

import (
	"encoding/json"
	"time"
)

// Scalar type definitions
type Date time.Time
type Timestamp time.Time
type Timestamptz time.Time
type Numeric float64
type Float8 float64
type Bigint int64
type Smallint int16

{{range .Types}}
// {{.Name}} represents the {{.Name}} GraphQL type
type {{toCamelCase .Name}} struct {
{{range .Fields}}{{if not (hasPrefix .Name "__")}}
	{{toCamelCase .Name}} {{getGoType .Type}} ` + "`json:\"{{.Name}}\"`" + `
{{end}}{{end}}}
{{end}}
`

const schemaTemplate = `# Generated from GraphQL introspection, DO NOT EDIT.
# Generated at: {{.Timestamp}}

{{.Schema}}
`

func main() {
	// Try to get API key from environment variable first
	apiKey := os.Getenv("HARDCOVER_API_KEY")
	configSource := "environment variable"

	// If not set in environment, try to load from config file
	if apiKey == "" {
		configPath, err := config.GetConfigPath()
		if err != nil {
			fmt.Printf("Failed to get config path: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Loading config from: %s\n", configPath)

		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Printf("Failed to load config: %v\n", err)
			fmt.Println("Please set HARDCOVER_API_KEY environment variable or configure via 'hardcover config set-api-key'")
			os.Exit(1)
		}

		if cfg.APIKey == "" {
			fmt.Println("No API key found in environment variable HARDCOVER_API_KEY or config file")
			fmt.Println("Please set your API key using one of these methods:")
			fmt.Println("  export HARDCOVER_API_KEY=\"your-api-key-here\"")
			fmt.Println("  hardcover config set-api-key \"your-api-key-here\"")
			os.Exit(1)
		}

		apiKey = cfg.APIKey
		configSource = "config file"
	}

	fmt.Printf("Using API key from %s\n", configSource)

	// Test the API key with a simple query first
	fmt.Println("Testing API key with simple query...")
	if err := testAPIKey(apiKey); err != nil {
		fmt.Printf("API key test failed: %v\n", err)
		fmt.Println("This suggests the API key might be invalid or expired")
		os.Exit(1)
	}
	fmt.Println("API key test successful!")

	// Fetch GraphQL schema
	fmt.Println("Fetching GraphQL schema...")
	schema, err := fetchGraphQLSchema(apiKey)
	if err != nil {
		fmt.Printf("Failed to fetch GraphQL schema: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully fetched schema with %d types\n", len(schema.Data.Schema.Types))

	// Generate types file
	fmt.Println("Generating Go types...")
	if err := generateTypesFile(schema); err != nil {
		fmt.Printf("Failed to generate types file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully generated Go types from GraphQL schema")
}

func testAPIKey(apiKey string) error {
	// Load config to get base URL
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Create simple test request
	req := GraphQLRequest{
		Query: testQuery,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal test request: %w", err)
	}

	// Create HTTP request
	ctx := context.Background()
	httpReq, err := http.NewRequestWithContext(
		ctx, "POST", cfg.BaseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create test request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("User-Agent", "hardcover-cli/1.0.0")
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to execute test request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read test response: %w", err)
	}
	if closeErr := resp.Body.Close(); closeErr != nil {
		return fmt.Errorf("failed to close test response body: %w", closeErr)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("test request failed with HTTP error %d: %s", resp.StatusCode, string(body))
	}

	// Try to parse response to check for GraphQL errors
	var testResp struct {
		Data   json.RawMessage `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.Unmarshal(body, &testResp); err != nil {
		return fmt.Errorf("failed to parse test response: %w", err)
	}

	if len(testResp.Errors) > 0 {
		return fmt.Errorf("GraphQL errors in test query: %v", testResp.Errors)
	}

	return nil
}

func fetchGraphQLSchema(apiKey string) (*GraphQLResponse, error) {
	// Load config to get base URL
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Create GraphQL request
	req := GraphQLRequest{
		Query: introspectionQuery,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	ctx := context.Background()
	httpReq, err := http.NewRequestWithContext(
		ctx, "POST", cfg.BaseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("User-Agent", "hardcover-cli/1.0.0")
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	if closeErr := resp.Body.Close(); closeErr != nil {
		return nil, fmt.Errorf("failed to close response body: %w", closeErr)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var gqlResp GraphQLResponse
	if unmarshalErr := json.Unmarshal(body, &gqlResp); unmarshalErr != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", unmarshalErr)
	}

	// Validate response structure
	if gqlResp.Data.Schema.Types == nil {
		return nil, fmt.Errorf("invalid GraphQL response: missing schema types")
	}

	return &gqlResp, nil
}

func generateTypesFile(schema *GraphQLResponse) error {
	// Filter out introspection types and conflicting types
	var filteredTypes []Type
	for _, t := range schema.Data.Schema.Types {
		if shouldIncludeType(&t) {
			filteredTypes = append(filteredTypes, t)
		}
	}

	if len(filteredTypes) == 0 {
		return fmt.Errorf("no valid types found in schema")
	}

	// Create output directory
	if mkdirErr := os.MkdirAll("internal/client", dirPermissions); mkdirErr != nil {
		return fmt.Errorf("failed to create directory: %w", mkdirErr)
	}

	// Generate timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05 MST")

	// Generate types file
	if err := generateTypesGoFile(filteredTypes, timestamp); err != nil {
		return fmt.Errorf("failed to generate types.go: %w", err)
	}

	// Generate schema file
	if err := generateSchemaFile(schema, timestamp); err != nil {
		return fmt.Errorf("failed to generate schema.graphql: %w", err)
	}

	fmt.Printf("Generated %d types from GraphQL schema\n", len(filteredTypes))
	return nil
}

func shouldIncludeType(t *Type) bool {
	// Skip introspection types
	if strings.HasPrefix(t.Name, "__") {
		return false
	}

	// Skip empty names
	if t.Name == "" {
		return false
	}

	// Skip certain problematic types
	excludedTypes := map[string]bool{
		"json":   true,
		"jsonb":  true,
		"citext": true,
	}

	if excludedTypes[strings.ToLower(t.Name)] {
		return false
	}

	// Only include types that have fields or are enums
	return len(t.Fields) > 0 || len(t.EnumValues) > 0
}

func generateTypesGoFile(types []Type, timestamp string) error {
	// Create template
	tmpl, err := template.New("types").Funcs(template.FuncMap{
		"getGoType":   getGoType,
		"toCamelCase": toCamelCase,
		"hasPrefix":   strings.HasPrefix,
	}).Parse(typesTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Generate types file
	typesFile, err := os.Create("internal/client/types.go")
	if err != nil {
		return fmt.Errorf("failed to create types file: %w", err)
	}
	defer func() {
		if closeErr := typesFile.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close types file: %v\n", closeErr)
		}
	}()

	// Execute template
	data := struct {
		Timestamp string
		Types     []Type
	}{
		Timestamp: timestamp,
		Types:     types,
	}

	if err := tmpl.Execute(typesFile, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

func generateSchemaFile(schema *GraphQLResponse, timestamp string) error {
	// Create template
	tmpl, err := template.New("schema").Parse(schemaTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse schema template: %w", err)
	}

	// Generate schema file
	schemaFile, err := os.Create("internal/client/schema.graphql")
	if err != nil {
		return fmt.Errorf("failed to create schema file: %w", err)
	}
	defer func() {
		if closeErr := schemaFile.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close schema file: %v\n", closeErr)
		}
	}()

	// Convert schema to string representation
	schemaStr := convertSchemaToString(schema)

	// Execute template
	data := struct {
		Schema    string
		Timestamp string
	}{
		Schema:    schemaStr,
		Timestamp: timestamp,
	}

	if err := tmpl.Execute(schemaFile, data); err != nil {
		return fmt.Errorf("failed to execute schema template: %w", err)
	}

	return nil
}

func convertSchemaToString(schema *GraphQLResponse) string {
	var result strings.Builder

	writeScalarDefinitions(&result)

	// Group types by kind
	objects, enums, inputs, unions := groupTypesByKind(schema.Data.Schema.Types)

	// Generate object types
	writeObjectTypes(&result, objects)

	// Generate enum types
	writeEnumTypes(&result, enums)

	// Generate input types
	writeInputTypes(&result, inputs)

	// Generate union types
	writeUnionTypes(&result, unions)

	return result.String()
}

func writeScalarDefinitions(result *strings.Builder) {
	scalars := []string{
		"ID", graphQLString, graphQLInt, graphQLFloat, graphQLBoolean,
		graphQLDate, graphQLTimestamp, graphQLTimestamptz, graphQLNumeric,
		graphQLFloat8, graphQLBigint, graphQLSmallint, "Citext", "Json", "Jsonb",
	}

	for _, scalar := range scalars {
		fmt.Fprintf(result, "scalar %s\n", scalar)
	}
	fmt.Fprint(result, "\n")
}

func groupTypesByKind(types []Type) (objects, enums, inputs, unions []Type) {
	for _, t := range types {
		if strings.HasPrefix(t.Name, "__") {
			continue // Skip introspection types
		}

		switch t.Kind {
		case "OBJECT":
			objects = append(objects, t)
		case "ENUM":
			enums = append(enums, t)
		case "INPUT_OBJECT":
			inputs = append(inputs, t)
		case "UNION":
			unions = append(unions, t)
		}
	}
	return
}

func writeObjectTypes(result *strings.Builder, objects []Type) {
	for _, obj := range objects {
		if len(obj.Fields) == 0 {
			continue
		}

		fmt.Fprintf(result, "type %s {\n", obj.Name)
		for _, field := range obj.Fields {
			if strings.HasPrefix(field.Name, "__") {
				continue
			}
			fmt.Fprintf(result, "  %s: %s\n", field.Name, getGraphQLType(field.Type))
		}
		fmt.Fprint(result, "}\n\n")
	}
}

func writeEnumTypes(result *strings.Builder, enums []Type) {
	for _, enum := range enums {
		if len(enum.EnumValues) == 0 {
			continue
		}

		fmt.Fprintf(result, "enum %s {\n", enum.Name)
		for _, value := range enum.EnumValues {
			fmt.Fprintf(result, "  %s\n", value.Name)
		}
		fmt.Fprint(result, "}\n\n")
	}
}

func writeInputTypes(result *strings.Builder, inputs []Type) {
	for _, input := range inputs {
		if len(input.Fields) == 0 {
			continue
		}

		fmt.Fprintf(result, "input %s {\n", input.Name)
		for _, field := range input.Fields {
			if strings.HasPrefix(field.Name, "__") {
				continue
			}
			fmt.Fprintf(result, "  %s: %s\n", field.Name, getGraphQLType(field.Type))
		}
		fmt.Fprint(result, "}\n\n")
	}
}

func writeUnionTypes(result *strings.Builder, unions []Type) {
	for _, union := range unions {
		if len(union.Fields) == 0 {
			continue
		}

		fmt.Fprintf(result, "union %s = ", union.Name)
		types := make([]string, 0)
		for _, field := range union.Fields {
			if strings.HasPrefix(field.Name, "__") {
				continue
			}
			types = append(types, field.Type.Name)
		}
		fmt.Fprintf(result, "%s\n\n", strings.Join(types, " | "))
	}
}

func getGraphQLType(typeRef TypeRef) string {
	if typeRef.Kind == "NON_NULL" {
		return getGraphQLType(*typeRef.OfType) + "!"
	}

	if typeRef.Kind == "LIST" {
		return "[" + getGraphQLType(*typeRef.OfType) + "]"
	}

	return getGraphQLTypeName(typeRef.Name)
}

var graphQLTypeMap = map[string]string{
	"ID":               "ID",
	graphQLString:      graphQLString,
	graphQLInt:         graphQLInt,
	graphQLFloat:       graphQLFloat,
	graphQLBoolean:     graphQLBoolean,
	graphQLDate:        graphQLDate,
	graphQLTimestamp:   graphQLTimestamp,
	graphQLTimestamptz: graphQLTimestamptz,
	graphQLNumeric:     graphQLNumeric,
	graphQLFloat8:      graphQLFloat8,
	graphQLBigint:      graphQLBigint,
	graphQLSmallint:    graphQLSmallint,
	"Citext":           "Citext",
	"Json":             "Json",
	"Jsonb":            "Jsonb",
}

func getGraphQLTypeName(typeName string) string {
	if mapped, exists := graphQLTypeMap[typeName]; exists {
		return mapped
	}
	return typeName
}

const (
	goStringType = "string"
	// ASCII uppercase conversion: clear bit 5 (32) to convert lowercase to uppercase
	asciiUppercaseMask = 32
	// Directory permissions: owner read/write/execute, group read/execute
	dirPermissions = 0o750

	// GraphQL scalar type names
	graphQLString      = "String"
	graphQLInt         = "Int"
	graphQLFloat       = "Float"
	graphQLBoolean     = "Boolean"
	graphQLDate        = "Date"
	graphQLTimestamp   = "Timestamp"
	graphQLTimestamptz = "Timestamptz"
	graphQLNumeric     = "Numeric"
	graphQLFloat8      = "Float8"
	graphQLBigint      = "Bigint"
	graphQLSmallint    = "Smallint"
)

func getGoType(typeRef TypeRef) string {
	if typeRef.Kind == "NON_NULL" {
		return getGoType(*typeRef.OfType)
	}

	if typeRef.Kind == "LIST" {
		return "[]" + getGoType(*typeRef.OfType)
	}

	return getGoTypeName(typeRef.Name)
}

var goTypeMap = map[string]string{
	"ID":               goStringType,
	graphQLString:      goStringType,
	graphQLInt:         "int",
	graphQLFloat:       "float64",
	graphQLBoolean:     "bool",
	graphQLDate:        "*Date",
	graphQLTimestamp:   "*Timestamp",
	graphQLTimestamptz: "*Timestamptz",
	graphQLNumeric:     "*Numeric",
	graphQLFloat8:      "*Float8",
	graphQLBigint:      "*Bigint",
	graphQLSmallint:    "*Smallint",
	"json":             "*json.RawMessage",
	"jsonb":            "*json.RawMessage",
	"citext":           goStringType,
}

func getGoTypeName(typeName string) string {
	if mapped, exists := goTypeMap[typeName]; exists {
		return mapped
	}
	return "*" + toCamelCase(typeName)
}

func toCamelCase(s string) string {
	if s == "" {
		return s
	}

	// Handle common cases
	switch s {
	case "id":
		return "ID"
	case "url":
		return "URL"
	case "api":
		return "API"
	}

	// Simple camel case conversion
	if len(s) > 1 {
		return string(rune(s[0])&^asciiUppercaseMask) + s[1:]
	}
	return string(rune(s[0]) & ^asciiUppercaseMask)
}
