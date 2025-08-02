package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"hardcover-cli/internal/testutil"
)

func TestInitConfig_UsesEnvVariable(t *testing.T) {
	ctm := testutil.NewConfigTestManager(t)
	defer ctm.Cleanup()

	expected := "env-api-key"
	ctm.SetEnv("HARDCOVER_API_KEY", expected)

	if rootCmd.PersistentFlags().Lookup("api-key") == nil {
		rootCmd.PersistentFlags().String("api-key", "", "Hardcover API key")
	}
	require.NoError(t, rootCmd.PersistentFlags().Set("api-key", ""))

	globalConfig = nil
	initConfig()

	require.NotNil(t, globalConfig)
	assert.Equal(t, expected, globalConfig.APIKey)
}

func TestInitConfig_FlagOverridesEnv(t *testing.T) {
	ctm := testutil.NewConfigTestManager(t)
	defer ctm.Cleanup()

	ctm.SetEnv("HARDCOVER_API_KEY", "env-api-key")

	if rootCmd.PersistentFlags().Lookup("api-key") == nil {
		rootCmd.PersistentFlags().String("api-key", "", "Hardcover API key")
	}
	require.NoError(t, rootCmd.PersistentFlags().Set("api-key", "flag-api-key"))

	globalConfig = nil
	initConfig()

	require.NotNil(t, globalConfig)
	assert.Equal(t, "flag-api-key", globalConfig.APIKey)

	// reset flag for other tests
	require.NoError(t, rootCmd.PersistentFlags().Set("api-key", ""))
}
