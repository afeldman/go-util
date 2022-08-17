//go:build windows
// +build windows

package appdirs

import (
	"path"

	fs "github.com/afeldman/go-util/filesystem"
)

var (
	datahome string = "AppData\\Local"
)

func (conf AppConf) userDataDir() (string, error) {
	home, err := fs.Dir()
	if err != nil {
		return "", err
	}

	dataHome = path.Join(home, datahome)

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
	dataDir := "C:\\ProgramData\\"

	if conf.Prefix != "" {
		dataDir = path.Join(conf.Prefix, "share")
	}

	if conf.Name != "" {
		dataDir = path.Join(dataDir, conf.Name)

		if conf.Version != "" {
			dataDir = path.Join(dataDir, conf.Version)
		}
	}

	return dataDir, nil
}

func (conf AppConf) userConfigDir() (string, error) {
	home, err := fs.Dir()
	if err != nil {
		return "", err
	}

	configHome = path.Join(configHome, datahome)

	// add app information
	if conf.Name != "" {
		configHome = path.Join(configHome, conf.Name)

		if conf.Version != "" {
			configHome = path.Join(configHome, conf.Version)
		}
	}

	configHome = path.Join(configHome, "\\.config")

	return configHome, nil
}

func (conf AppConf) siteConfigDir() (string, error) {
	configDir := "C:\\ProgramData\\"

	if conf.Prefix != "" {
		configDir = path.Join(conf.Prefix, "config")
	}

	if conf.Name != "" {
		configDir = path.Join(configDir, conf.Name)

		if conf.Version != "" {
			configDir = path.Join(configDir, conf.Version)
		}
	}

	return configDir, nil
}

func (conf AppConf) userCacheDir() (string, error) {

	home, err := fs.Dir()
	if err != nil {
		return "", err
	}

	cacheDir = path.Join(home, ".cache")

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
	// first use the XDG_DATA_HOME env variable, otherwise fallback to a safe default
	cacheDir, err := conf.userCacheDir()
	if err != nil {
		return "", err
	}

	return path.Join(cacheDir, "logs"), nil
}

func globalConfigDir() (string, error) {
	return "C:\\ProgramData", nil
}

func globalDataDir() (string, error) {
	return "C:\\Program Files", nil
}
