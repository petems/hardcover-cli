//go:build dev

package cmd

// setupSchemaCommands initializes schema-related commands
// This function is only available in development builds
func setupSchemaCommands() {
	initSchemaCmd()
	rootCmd.AddCommand(schemaCmd)
}
