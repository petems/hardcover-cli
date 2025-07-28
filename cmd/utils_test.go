package cmd

import (
	"bytes"
	"context"
	"testing"

	"hardcover-cli/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestMaskAPIKey(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"1234567890", "1234...7890"},
		{"12345678", "12345678"},
		{"short", "short"},
		{"", ""},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.out, maskAPIKey(tt.in))
	}
}

func TestPrintHelpers(t *testing.T) {
	var buf bytes.Buffer
	printToStdoutf(&buf, "Hello %s", "World")
	printToStdoutLn(&buf, "!")
	assert.Equal(t, "Hello World!\n", buf.String())
}

func TestGetConfig(t *testing.T) {
	orig := globalConfig
	defer func() { globalConfig = orig }()

	// Context provided config
	ctx := WithConfig(context.Background(), &config.Config{APIKey: "ctx"})
	cfg, ok := getConfig(ctx)
	assert.True(t, ok)
	assert.Equal(t, "ctx", cfg.APIKey)

	// Global config fallback
	globalConfig = &config.Config{APIKey: "global"}
	cfg, ok = getConfig(context.Background())
	assert.True(t, ok)
	assert.Equal(t, "global", cfg.APIKey)

	// Cancelled context
	globalConfig = nil
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	cfg, ok = getConfig(cancelCtx)
	assert.False(t, ok)
	assert.Nil(t, cfg)
}

func FuzzMaskAPIKey(f *testing.F) {
	seeds := []string{"", "short", "1234567890abcdef"}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		masked := maskAPIKey(s)
		if len(s) <= 8 {
			if masked != s {
				t.Fatalf("expected %q got %q", s, masked)
			}
			return
		}
		if len(masked) < 11 {
			t.Fatalf("masked string too short: %q", masked)
		}
		if !bytes.HasPrefix([]byte(masked), []byte(s[:4])) ||
			!bytes.HasSuffix([]byte(masked), []byte(s[len(s)-4:])) {
			t.Fatalf("masking incorrect: %q from %q", masked, s)
		}
	})
}
