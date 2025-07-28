package cmd

import (
	"context"
	"fmt"

	"hardcover-cli/internal/config"
)

// printToStdoutf safely prints to stdout without checking errors (for CLI output).
func printToStdoutf(w interface{ Write([]byte) (int, error) }, format string, args ...interface{}) {
	_, _ = fmt.Fprintf(w, format, args...) // CLI output errors are not critical
}

// printToStdoutLn safely prints a newline to stdout without checking errors.
func printToStdoutLn(w interface{ Write([]byte) (int, error) }, args ...interface{}) {
	_, _ = fmt.Fprintln(w, args...) // CLI output errors are not critical
}

// withConfig is a test helper function to inject configuration into context.
func withConfig(ctx context.Context, cfg *config.Config) context.Context {
	// Store config globally for tests
	globalConfig = cfg
	return ctx
}

// withTestConfig is a test helper function that converts testutil.Config to config.Config.
func withTestConfig(ctx context.Context, cfg interface{}) context.Context {
	// Convert testutil.Config to config.Config
	if testCfg, ok := cfg.(testConfigConvertible); ok {
		realCfg := &config.Config{
			APIKey:  testCfg.GetAPIKey(),
			BaseURL: testCfg.GetBaseURL(),
		}
		return withConfig(ctx, realCfg)
	}
	return ctx
}

// testConfigConvertible interface to avoid import cycle.
type testConfigConvertible interface {
	GetAPIKey() string
	GetBaseURL() string
}
