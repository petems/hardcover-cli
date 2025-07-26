package cmd

import (
	"context"
	"fmt"

	"hardcover-cli/internal/config"
)

type contextKey string

const (
	configKey contextKey = "config"
)

// withConfig adds the configuration to the context
func withConfig(ctx context.Context, cfg *config.Config) context.Context {
	return context.WithValue(ctx, configKey, cfg)
}

// getConfig retrieves the configuration from the context
func getConfig(ctx context.Context) (*config.Config, bool) {
	cfg, ok := ctx.Value(configKey).(*config.Config)
	return cfg, ok
}

// printToStdoutf safely prints to stdout without checking errors (for CLI output)
func printToStdoutf(w interface{ Write([]byte) (int, error) }, format string, args ...interface{}) {
	_, _ = fmt.Fprintf(w, format, args...) //nolint:errcheck // CLI output errors are not critical
}
