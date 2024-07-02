//go:build windows
// +build windows

// package filesystem contains functions for working with the filesystem
package filesystem

import (
	"errors"
	"os"
)

// osdir gets the home directory of the user
// Returns:
//
//	 string -- the home directory of the user
//	error -- an error if the home directory could not be found
func osdir() (string, error) {
	if home := os.Getenv(homeEnv); home != "" {
		return home, nil
	}

	if home := os.Getenv("USERPROFILE"); home != "" {
		return home, nil
	}

	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, or USERPROFILE are blank")
	}

	return home, nil
}
