//go:build !dev

package cmd

// setupSchemaCommands is a no-op in production builds
// This function is only available in production builds
func setupSchemaCommands() {
	// Schema commands are not available in production builds
}
