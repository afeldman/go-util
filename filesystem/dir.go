package filesystem

//ref: https://github.com/mitchellh/go-homedir

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var (
	DisableCache bool
	homedirCache string
	cacheLock    sync.RWMutex
	homeEnv      string = "HOME"
)

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

func MkDir(path string, permission os.FileMode) error {
	//choose your permissions well
	pathErr := os.MkdirAll(path, permission)
	return pathErr
}

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

func Reset() {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	homedirCache = ""
}
