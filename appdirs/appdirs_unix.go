//go:build !windows && !darwin
// +build !windows,!darwin

package appdirs

import (
	"path"
	"strings"

	"github.com/afeldman/go-util/env"
	fs "github.com/afeldman/go-util/filesystem"
)

const PATHSEP = ":"

var (
	xgddatahome   string = "XDG_DATA_HOME"
	xdgdatadirs   string = "XDG_DATA_DIRS"
	xgdconfighome string = "XDG_CONFIG_HOME"
	xgdconfigdirs string = "XDG_CONFIG_DIRS"
	xgdcachehome  string = "XDG_CACHE_HOME"
)

func (conf AppConf) userDataDir() (string, error) {
	// first use the XDG_DATA_HOME env variable, otherwise fallback to a safe default

	dataHome, err := fs.Expand(env.Plan9GetEnv(xgddatahome))
	if err != nil {
		return "", err
	}

	if dataHome == "" {
		home, err := fs.Dir()
		if err != nil {
			return "", err
		}

		dataHome = path.Join(home, ".local/share")
	}

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
	dataDir := "/usr/local/share"

	if conf.Prefix != "" {
		dataDir = path.Join(conf.Prefix, "share")
	}

	// special case if the prefix is /, use /usr/share
	if conf.Prefix == "/" {
		dataDir = "/usr/share"
	}

	// first use the XDG_DATA_HOME env variable, otherwise fallback to a safe default
	if env.Plan9GetEnv(xdgdatadirs) != "" {
		var err error
		dataDir, err = fs.Expand(strings.Split(env.Plan9GetEnv(xdgdatadirs), PATHSEP)[0])
		if err != nil {
			return "", nil
		}
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
	// first use the XDG_DATA_HOME env variable, otherwise fallback to a safe default
	configHome, err := fs.Expand(env.Plan9GetEnv(xgdconfighome))
	if err != nil {
		return "", err
	}

	if configHome == "" {
		home, err := fs.Dir()
		if err != nil {
			return "", err
		}

		configHome = path.Join(home, ".config")
	}

	// add app information
	if conf.Name != "" {
		configHome = path.Join(configHome, conf.Name)

		if conf.Version != "" {
			configHome = path.Join(configHome, conf.Version)
		}
	}

	return configHome, nil
}

func (conf AppConf) siteConfigDir() (string, error) {
	configDir := "/etc/xdg"

	if conf.Prefix != "" {
		configDir = path.Join(conf.Prefix, "etc")
	}

	// first use the XDG_DATA_HOME env variable, otherwise fallback to a safe default
	if env.Plan9GetEnv(xgdconfigdirs) != "" {
		var err error
		configDir, err = fs.Expand(strings.Split(env.Plan9GetEnv(xgdconfigdirs), PATHSEP)[0])
		if err != nil {
			return "", nil
		}
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
	// first use the XDG_DATA_HOME env variable, otherwise fallback to a safe default
	cacheDir, err := fs.Expand(env.Plan9GetEnv(xgdcachehome))
	if err != nil {
		return "", err
	}

	if cacheDir == "" {
		home, err := fs.Dir()
		if err != nil {
			return "", err
		}

		cacheDir = path.Join(home, ".cache")
	}

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
	return "/etc/", nil
}

func globalDataDir() (string, error) {
	return "/usr/local", nil
}
