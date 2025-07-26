package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

// Config holds the application configuration
type Config struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
}

// LoadConfig loads configuration from environment variables and config file
func LoadConfig() (*Config, error) {
	cfg := DefaultConfig()

	// Try to load from config file first
	configPath, err := getConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get config path: %w", err)
	}

	if _, statErr := os.Stat(configPath); !os.IsNotExist(statErr) {
		// Config file exists, load it
		data, err := os.ReadFile(configPath) //nolint:gosec // configPath is constructed from user home directory
		if err != nil {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}

		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("failed to parse config file: %w", err)
		}
	}

	// Check environment variable - it overrides config file
	if apiKey := os.Getenv("HARDCOVER_API_KEY"); apiKey != "" {
		// Trim whitespace and check if it's still non-empty
		if trimmedAPIKey := strings.TrimSpace(apiKey); trimmedAPIKey != "" {
			cfg.APIKey = trimmedAPIKey
		}
	}

	return cfg, nil
}

// SaveConfig saves the configuration to a file
func SaveConfig(cfg *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get config path: %w", err)
	}

	// Create config directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	const configDirPerm = 0o750
	if mkdirErr := os.MkdirAll(configDir, configDirPerm); mkdirErr != nil {
		return fmt.Errorf("failed to create config directory: %w", mkdirErr)
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	const configFilePerm = 0o600
	if err := os.WriteFile(configPath, data, configFilePerm); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// getConfigPath returns the path to the configuration file
func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	return filepath.Join(homeDir, ".hardcover", "config.yaml"), nil
}

// GetConfigPath returns the configuration file path for external access
func GetConfigPath() (string, error) {
	return getConfigPath()
}
