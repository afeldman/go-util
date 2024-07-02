// package filesystem contains functions for working with the filesystem
package filesystem

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"sync"
)

// variables to be used in the package
var (
	DisableCache bool
	homedirCache string
	cacheLock    sync.RWMutex
	homeEnv      string = "HOME"
)

// is emptx checks if a directory is empty
// Arguments:
//
//	path -- the path to the directory to check
//
// Returns:
//
//	bool -- true if the directory is empty, false otherwise
//	error -- an error if the directory could not be read
func IsEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

// mk dir creates a directory
// Arguments:
//
//	path -- the path to the directory to create
//	permission -- the permissions to set on the directory
//
// Returns:
//
//	error -- an error if the directory could not be created
func MkDir(path string, permission os.FileMode) error {
	//choose your permissions well
	pathErr := os.MkdirAll(path, permission)
	return pathErr
}

// Dir gets the home directory of the user
// Returns:
//
//	string -- the home directory of the user
//	error -- an error if the home directory could not be found
func Dir() (string, error) {
	if !DisableCache {
		cacheLock.RLock()
		cached := homedirCache
		cacheLock.RUnlock()
		if cached != "" {
			return cached, nil
		}
	}

	cacheLock.Lock()
	defer cacheLock.Unlock()

	result, err := osdir()
	if err != nil {
		return "", err
	}

	homedirCache = result
	return result, nil
}

// Expand expands the path to include the home directory
// Arguments:
//
//	path -- the path to expand
//
// Returns:
//
//	string -- the expanded path
//	error -- an error if the path could not be expanded
func Expand(path string) (string, error) {
	if len(path) == 0 {
		return path, nil
	}

	if path[0] != '~' {
		return path, nil
	}

	if len(path) > 1 && path[1] != '/' && path[1] != '\\' {
		return "", errors.New("cannot expand user-specific home dir")
	}

	dir, err := Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, path[1:]), nil
}

// set the home dir to empty path
func Reset() {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	homedirCache = ""
}
