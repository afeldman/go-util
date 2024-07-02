//go:build darwin
// +build darwin

// package filesystem contains functions for working with the filesystem
package filesystem

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

// osdir gets the home directory of the user
// Returns:
//
//	string -- the home directory of the user
//	error -- an error if the home directory could not be found
func osdir() (string, error) {

	// First prefer the HOME environmental variable
	if home := os.Getenv(homeEnv); home != "" {
		return home, nil
	}

	var stdout bytes.Buffer

	cmd := exec.Command("sh", "-c", `dscl -q . -read /Users/"$(whoami)" NFSHomeDirectory | sed 's/^[^ ]*: //'`)
	cmd.Stdout = &stdout
	if err := cmd.Run(); err == nil {
		result := strings.TrimSpace(stdout.String())
		if result != "" {
			return result, nil
		}
	}

	// If all else fails, try the shell
	stdout.Reset()
	cmd = exec.Command("sh", "-c", "cd && pwd")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}
