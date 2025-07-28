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

package client

import (
	"encoding/json"
)

{{range .Types}}
// {{.Name}} represents the {{.Name}} GraphQL type
type {{toCamelCase .Name}} struct {
{{range .Fields}}{{if not (hasPrefix .Name "__")}}
	{{toCamelCase .Name}} {{getGoType .Type}} ` + "`json:\"{{.Name}}\"`" + `
{{end}}{{end}}}
{{end}}
`

func main() {
	apiKey := os.Getenv("HARDCOVER_API_KEY")
	if apiKey == "" {
		fmt.Println("HARDCOVER_API_KEY environment variable is required")
		os.Exit(1)
	}

	// Fetch GraphQL schema
	schema, err := fetchGraphQLSchema(apiKey)
	if err != nil {
		fmt.Printf("Failed to fetch GraphQL schema: %v\n", err)
		os.Exit(1)
	}

	// Generate types file
	if err := generateTypesFile(schema); err != nil {
		fmt.Printf("Failed to generate types file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully generated Go types from GraphQL schema")
}

func fetchGraphQLSchema(apiKey string) (*GraphQLResponse, error) {
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
		ctx, "POST", "https://api.hardcover.app/v1/graphql", bytes.NewBuffer(jsonData))
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

	return &gqlResp, nil
}

func generateTypesFile(schema *GraphQLResponse) error {
	// Filter out introspection types and conflicting types
	var filteredTypes []Type
	for _, t := range schema.Data.Schema.Types {
		if t.Name != "" && !strings.HasPrefix(t.Name, "__") && t.Name != "json" {
			filteredTypes = append(filteredTypes, t)
		}
	}

	// Create template
	tmpl, err := template.New("types").Funcs(template.FuncMap{
		"getGoType":   getGoType,
		"toCamelCase": toCamelCase,
		"hasPrefix":   strings.HasPrefix,
	}).Parse(typesTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create output directory
	if mkdirErr := os.MkdirAll("internal/client", dirPermissions); mkdirErr != nil {
		return fmt.Errorf("failed to create directory: %w", mkdirErr)
	}

	// Generate types file
	typesFile, err := os.Create("internal/client/types.go")
	if err != nil {
		return fmt.Errorf("failed to create types file: %w", err)
	}

	// Execute template
	data := struct {
		Types []Type
	}{
		Types: filteredTypes,
	}

	if err := tmpl.Execute(typesFile, data); err != nil {
		if closeErr := typesFile.Close(); closeErr != nil {
			return fmt.Errorf("failed to execute template: %w, and failed to close file: %w", err, closeErr)
		}
		return fmt.Errorf("failed to execute template: %w", err)
	}

	if err := typesFile.Close(); err != nil {
		return fmt.Errorf("failed to close types file: %w", err)
	}

	return nil
}

const (
	goStringType = "string"
	// ASCII uppercase conversion: clear bit 5 (32) to convert lowercase to uppercase
	asciiUppercaseMask = 32
	// Directory permissions: owner read/write/execute, group read/execute
	dirPermissions = 0o750
)

func getGoType(typeRef TypeRef) string {
	if typeRef.Kind == "NON_NULL" {
		return getGoType(*typeRef.OfType)
	}

	if typeRef.Kind == "LIST" {
		return "[]" + getGoType(*typeRef.OfType)
	}

	switch typeRef.Name {
	case "ID":
		return goStringType
	case "String":
		return goStringType
	case "Int":
		return "int"
	case "Float":
		return "float64"
	case "Boolean":
		return "bool"
	case "json":
		return "*json.RawMessage"
	case "jsonb":
		return "*json.RawMessage"
	case "citext":
		return goStringType
	default:
		return "*" + toCamelCase(typeRef.Name)
	}
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
