package env

import (
	"os"
	"strconv"
)

func GetString(key string, fallback string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return value
}

func GetInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valueAsInt, err := strconv.Atoi(value)

	if err != nil {
		return fallback
	}

	return valueAsInt

}
