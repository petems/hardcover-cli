package contextutil_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"hardcover-cli/internal/config"
	"hardcover-cli/internal/contextutil"
)

func TestWithConfigAndGetConfig(t *testing.T) {
	cfg := &config.Config{APIKey: "k", BaseURL: "u"}
	ctx := context.Background()

	ctx = contextutil.WithConfig(ctx, cfg)
	got, ok := contextutil.GetConfig(ctx)

	require.True(t, ok)
	assert.Equal(t, cfg, got)
}

func TestGetConfig_NoConfig(t *testing.T) {
	cfg, ok := contextutil.GetConfig(context.Background())
	assert.False(t, ok)
	assert.Nil(t, cfg)
}

func TestGetConfig_WrongType(t *testing.T) {
	type otherKey struct{}
	ctx := context.WithValue(context.Background(), otherKey{}, "value")

	cfg, ok := contextutil.GetConfig(ctx)
	assert.False(t, ok)
	assert.Nil(t, cfg)
}
