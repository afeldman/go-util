//go:build !windows && !darwin
// +build !windows,!darwin

// package filesystem contains functions for working with the filesystem
package filesystem

import (
	"os/user"

	"github.com/afeldman/go-util/env"
)

// osdir gets the home directory of the user
// Returns:
//
//	string -- the home directory of the user
//	error -- an error if the home directory could not be found
func osdir() (string, error) {

	// First prefer the HOME environmental variable
	if home := env.Plan9GetEnv(homeEnv); home != "" {
		return home, nil
	}

	var usr *user.User
	var err error
	if usr, err = user.Current(); err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}
