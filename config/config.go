package config

import (
	"os"
)

// GetOpenAIKey returns the OpenAI API key from environment variable
func GetOpenAIKey() string {
	return os.Getenv("OPENAI_API_KEY")
}
