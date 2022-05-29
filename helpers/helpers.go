package helpers

import (
	"fmt"
	"os"
)

// GetEnv get env variable or return default value
func GetEnv(key, value string) string {
	v := os.Getenv(key)
	if v == "" {
		return value
	}

	return v
}

// MustGetEnv gets env variable or return error
func MustGetEnv(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return "", fmt.Errorf("env `%s` not defined", name)
	}

	return val, nil
}
