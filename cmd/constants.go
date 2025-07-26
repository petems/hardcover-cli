package cmd

const (
	// API key required error message used across multiple commands
	apiKeyRequiredMsg = "API key is required. Set it using:\n  export HARDCOVER_API_KEY=\"your-api-key\"\n" +
		"  or\n  hardcover config set-api-key \"your-api-key\""
	
	// Minimum API key length for masking
	apiKeyMaskingThreshold = 10
)