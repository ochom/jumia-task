package utils

import "os"

// GetEnv gets an environment variable with provided name or
// returns the default value `alternative`
func GetEnv(name, alternative string) string {
	value := os.Getenv(name)
	if value == "" {
		return alternative
	}

	return value
}
