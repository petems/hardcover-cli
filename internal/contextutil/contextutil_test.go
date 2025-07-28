package contextutil_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"hardcover-cli/internal/config"
	"hardcover-cli/internal/contextutil"
)

// TestWithConfigAndGetConfig verifies that configuration can be stored in and
// retrieved from a context using the provided helpers.
func TestWithConfigAndGetConfig(t *testing.T) {
	cfg := &config.Config{APIKey: "key", BaseURL: "url"}

	ctx := context.Background()
	ctx = contextutil.WithConfig(ctx, cfg)

	got, ok := contextutil.GetConfig(ctx)
	assert.True(t, ok)
	assert.Equal(t, cfg, got)
}

// TestGetConfig_NoConfig ensures GetConfig returns false when no configuration
// is present in the context.
func TestGetConfig_NoConfig(t *testing.T) {
	ctx := context.Background()
	cfg, ok := contextutil.GetConfig(ctx)
	assert.False(t, ok)
	assert.Nil(t, cfg)
}
