package env

import (
	"os"
	"strings"
)

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

func GetEnvOrDefault(name string, or string) *string {
	ret := os.Getenv(name)
	if (len(ret) == 0) || (ret == "") {
		ret = or
	}
	return &ret
}
