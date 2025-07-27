package cmd

import (
	"context"
	"fmt"

	"hardcover-cli/internal/config"
)

// printToStdoutf safely prints to stdout without checking errors (for CLI output)
func printToStdoutf(w interface{ Write([]byte) (int, error) }, format string, args ...interface{}) {
	_, _ = fmt.Fprintf(w, format, args...) //nolint:errcheck // CLI output errors are not critical
}

// printToStdoutLn safely prints a newline to stdout without checking errors
func printToStdoutLn(w interface{ Write([]byte) (int, error) }, args ...interface{}) {
	_, _ = fmt.Fprintln(w, args...) //nolint:errcheck // CLI output errors are not critical
}

// withConfig is a test helper function to inject configuration into context
func withConfig(ctx context.Context, cfg *config.Config) context.Context {
	// Store config globally for tests
	globalConfig = cfg
	return ctx
}
