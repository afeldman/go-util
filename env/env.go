package env

import (
	"os"
	"strings"
)

func Plan9GetEnv(env_var string) string {
	return GetEnvOrDefault(env_var, strings.ToLower(env_var))
}

func IsWindows() bool {
	return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}

func GetEnv(env_var string) []string {
	if IsWindows() {
		return strings.Split(os.Getenv(env_var), ";")
	} else {
		return strings.Split(os.Getenv(env_var), ":")
	}
}

func GetEnvOrDefault(name, or string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	return or
}
