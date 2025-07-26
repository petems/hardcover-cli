package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
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

	// Check environment variable first
	if apiKey := os.Getenv("HARDCOVER_API_KEY"); apiKey != "" {
		cfg.APIKey = apiKey
		return cfg, nil
	}

	// Try to load from config file
	configPath, err := getConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get config path: %w", err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Config file doesn't exist, return default config
		return cfg, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
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
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0600); err != nil {
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
