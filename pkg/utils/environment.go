package utils

import (
	"os"
)

// GetEnv get environment from os.
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == `` {
		return fallback
	}
	return value
}
