package log

import (
	"os"
	"strconv"
	"strings"
)

// envOr returns environment variable value or first non-empty default
func envOr(key string, def ...string) string {
	if env := os.Getenv(key); env != "" {
		return env
	}
	for _, d := range def {
		if d != "" {
			return d
		}
	}
	return ""
}

// envBoolOr returns environment variable as bool or default
func envBoolOr(key string, def bool) bool {
	if env := os.Getenv(key); env != "" {
		if val, err := strconv.ParseBool(env); err == nil {
			return val
		}
	}
	return def
}

// envIntOr returns environment variable as int or default
func envIntOr(key string, def int) int {
	if env := os.Getenv(key); env != "" {
		if val, err := strconv.Atoi(env); err == nil {
			return val
		}
	}
	return def
}

// envSliceOr returns environment variable as slice (comma-separated) or default
func envSliceOr(key string, def []string) []string {
	if env := os.Getenv(key); env != "" {
		return strings.Split(env, ",")
	}
	return def
}
