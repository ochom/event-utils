package utils

import (
	"fmt"
	"os"
)

// GetEnvOrDefault get env variable or return default value
func GetEnvOrDefault(key, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}

// MustGetEnv gets env variable or return error
func MustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("Environment variable %s not set", key))
	}

	return val
}
