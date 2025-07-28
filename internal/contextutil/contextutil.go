// Package contextutil provides utilities for working with context-based configuration.
package contextutil

import (
	"context"

	"hardcover-cli/internal/config"
)

// ConfigContextKey is a type-safe key for context values.
type ConfigContextKey struct{}

// WithConfig injects configuration into context.
func WithConfig(ctx context.Context, cfg *config.Config) context.Context {
	return context.WithValue(ctx, ConfigContextKey{}, cfg)
}

// GetConfig retrieves configuration from context.
func GetConfig(ctx context.Context) (*config.Config, bool) {
	if cfg, ok := ctx.Value(ConfigContextKey{}).(*config.Config); ok && cfg != nil {
		return cfg, true
	}
	return nil, false
}
