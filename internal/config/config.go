package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

const (
	configDirName  = ".hardcover"
	configFileName = "config.yaml"
	configFilePerm = 0o600
	configDirPerm  = 0o755
)

// DefaultConfig returns a config with default values
func DefaultConfig() *Config {
	return &Config{
		BaseURL: "https://api.hardcover.app/v1/graphql",
	}
}

// LoadConfig loads configuration from file and environment variables
func LoadConfig() (*Config, error) {
	cfg := DefaultConfig()

	// Environment variables override everything
	if apiKey := os.Getenv("HARDCOVER_API_KEY"); apiKey != "" {
		cfg.APIKey = apiKey
		return cfg, nil
	}

	// Try to load from config file
	configPath, err := GetConfigPath()
	if err != nil {
		// If we can't get the config path, continue with default config
		// This allows the CLI to work even if home directory is not accessible
		return cfg, nil
	}

	if _, statErr := os.Stat(configPath); statErr == nil {
		data, readErr := os.ReadFile(configPath) // configPath is constructed from user home directory
		if readErr != nil {
			return nil, fmt.Errorf("failed to read config file: %w", readErr)
		}

		if yamlErr := yaml.Unmarshal(data, cfg); yamlErr != nil {
			return nil, fmt.Errorf("failed to parse config file: %w", yamlErr)
		}
	}

	return cfg, nil
}

// SaveConfig saves the configuration to file
func SaveConfig(cfg *Config) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get config path: %w", err)
	}

	// Create config directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	if mkdirErr := os.MkdirAll(configDir, configDirPerm); mkdirErr != nil {
		return fmt.Errorf("failed to create config directory: %w", mkdirErr)
	}

	// Marshal config to YAML
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Write to file
	if writeErr := os.WriteFile(configPath, data, configFilePerm); writeErr != nil {
		return fmt.Errorf("failed to write config file: %w", writeErr)
	}

	return nil
}

// GetConfigPath returns the path to the configuration file
func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	return filepath.Join(homeDir, configDirName, configFileName), nil
}
