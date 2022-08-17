//go:build darwin
// +build darwin

package appdirs

import (
	"path"
	"strings"

	"github.com/afeldman/go-util/env"
	fs "github.com/afeldman/go-util/fs"
)

func (conf AppConf) userDataDir() (string, error) {
	home, err := fs.Dir()
	if err != nil {
		return "", err
	}

	dataHome = path.Join(home, "Library/Application Support")

	// add app information
	if conf.Name != "" {
		dataHome = path.Join(dataHome, conf.Name)

		if conf.Version != "" {
			dataHome = path.Join(dataHome, conf.Version)
		}
	}

	return dataHome, nil
}

func (conf AppConf) siteDataDir() (string, error) {
	dataDir := "/Library/Application Support"

	if conf.Name != "" {
		dataDir = path.Join(dataDir, conf.Name)

		if conf.Version != "" {
			dataDir = path.Join(dataDir, conf.Version)
		}
	}

	return dataDir, nil
}

func (conf AppConf) userConfigDir() (string, error) {
	return conf.userDataDir()
}

func (conf AppConf) siteConfigDir() (string, error) {
	return conf.siteDataDir()
}

func (conf AppConf) userCacheDir() (string, error) {
	home, err := fs.Dir()
	if err != nil {
		return "", err
	}

	cacheDir = path.Join(home, "Library/Caches")

	// add app information
	if conf.Name != "" {
		cacheDir = path.Join(cacheDir, conf.Name)

		if conf.Version != "" {
			cacheDir = path.Join(cacheDir, conf.Version)
		}
	}

	return cacheDir, nil
}

func (conf AppConf) userLogDir() (string, error) {
	home, err := fs.Dir()
	if err != nil {
		return "", err
	}

	logDir = path.Join(home, "Library/Logs")

	// add app information
	if conf.Name != "" {
		logDir = path.Join(logDir, conf.Name)

		if conf.Version != "" {
			logDir = path.Join(logDir, conf.Version)
		}
	}

	return logDir, nil
}

func globalConfigDir() (string, error) {
	return "/etc/", nil
}

func globalDataDir() (string, error) {
	return "/usr", nil
}
