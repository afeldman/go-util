//go:build windows
// +build windows

package filesystem

import (
	"errors"
	"os"
)

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
