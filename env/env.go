// package env
// This package provides functions to get environment variables

package env

import (
	"os"
	"strings"
)

/*
	plan9 format environment variable

Plan 9 uses a different format for environment variables.
It uses a single string with the format key=value separated by null bytes.
This function returns the value of the environment variable in Plan 9 format.

Arguments:

	env_var -- the name of the environment variable to get

Returns:

	string -- the value of the environment variable
*/
func Plan9GetEnv(env_var string) string {
	return GetEnvOrDefault(env_var, strings.ToLower(env_var))
}

/*
	are we running on Windows?

Returns:

	bool -- true if the OS is Windows, false otherwise
*/
func IsWindows() bool {
	return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}

/*
	Get the value of an environment variable

Arguments:

	env_var -- the name of the environment variable to get

Returns:

	[]string -- the value of the environment variable
*/
func GetEnv(env_var string) []string {
	if IsWindows() {
		return strings.Split(os.Getenv(env_var), ";")
	} else {
		return strings.Split(os.Getenv(env_var), ":")
	}
}

/*
	get the env value or default if not exists

Arguments:

	name -- the name of the environment variable to get
	or -- the default value to return if the environment variable does not exist

Returns:

	string -- the value of the environment variable or the default value if it does not exist
*/
func GetEnvOrDefault(name, or string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	return or
}
