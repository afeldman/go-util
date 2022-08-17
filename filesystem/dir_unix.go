//go:build !windows && !darwin
// +build !windows,!darwin

package filesystem

import (
	"os/user"

	"github.com/afeldman/go-util/env"
)

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
