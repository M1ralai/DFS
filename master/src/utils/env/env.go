package env

import (
	"os"
	"strconv"
)

func GetEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func IntGetEnv(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		i, _ := strconv.Atoi(v)
		return i
	}
	return def
}
