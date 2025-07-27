//go:build !dev

package cmd

import (
	"context"
	"fmt"
	"hardcover-cli/internal/config"
	"strings"

	"github.com/spf13/cobra"
)

// schemaCmd is not available in production builds
var schemaCmd *cobra.Command

func initSchemaCmd() {
	// Schema command is not available in production builds
}

func runSchemaFetch(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("schema command not available in production builds")
}

func introspectionToSDL(introspection map[string]interface{}) (string, error) {
	return "", fmt.Errorf("schema functionality not available in production builds")
}

func processType(sdl *strings.Builder, typeInterface interface{}) error {
	return fmt.Errorf("schema functionality not available in production builds")
}

func processObjectType(sdl *strings.Builder, typeMap map[string]interface{}, name string) error {
	return fmt.Errorf("schema functionality not available in production builds")
}

func processEnumType(sdl *strings.Builder, typeMap map[string]interface{}) error {
	return fmt.Errorf("schema functionality not available in production builds")
}

func processInputObjectType(sdl *strings.Builder, typeMap map[string]interface{}) error {
	return fmt.Errorf("schema functionality not available in production builds")
}

func getTypeString(typeInterface interface{}) string {
	return "Unknown"
}

func getIntrospectionQuery() string {
	return ""
}

func getSchemaFlags(cmd *cobra.Command) (outputPath, endpoint string, err error) {
	return "", "", fmt.Errorf("schema functionality not available in production builds")
}

func getDefaultOutputPath(outputPath string) string {
	return ""
}

func getDefaultEndpoint(endpoint string, cfg *config.Config) string {
	return ""
}

func ensureOutputDirectory(outputPath string) error {
	return fmt.Errorf("schema functionality not available in production builds")
}

func displaySchemaStatus(endpoint, apiKey string) {
	// No-op in production builds
}

func fetchSchema(ctx context.Context, endpoint, apiKey string) (string, error) {
	return "", fmt.Errorf("schema functionality not available in production builds")
}
