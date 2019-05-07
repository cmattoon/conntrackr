package config

import (
	"os"
	"strconv"
)

// Returns a value from the environment, or default
func GetEnv(key string, defval string) (val string) {
	if val = os.Getenv(key); val != "" {
		return val
	}
	return defval
}

// Same as GetEnv, but casts to an int
func GetEnvInt(key string, defval int) (val int) {
	if sval := GetEnv(key, ""); sval != "" {
		i, _ := strconv.Atoi(sval)
		return int(i)
	}
	return defval
}
