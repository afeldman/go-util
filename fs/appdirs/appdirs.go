package appdirs

// ref: https://github.com/christopherobin/go-appdirs

import fs "github.com/afeldman/go-util/fs"

type AppConf struct {
	Name    string
	Version string
	Prefix  string
}

func (conf AppConf) Directories() (map[string]string, error) {
	homeDir, err := HomeDir()
	if err != nil {
		return nil, err
	}

	userDataDir, err := conf.UserDataDir()
	if err != nil {
		return nil, err
	}

	userConfigDir, err := conf.UserConfigDir()
	if err != nil {
		return nil, err
	}

	userCacheDir, err := conf.UserCacheDir()
	if err != nil {
		return nil, err
	}

	siteDataDir, err := conf.SiteDataDir()
	if err != nil {
		return nil, err
	}

	siteConfigDir, err := conf.SiteConfigDir()
	if err != nil {
		return nil, err
	}

	userLogDir, err := conf.UserLogDir()
	if err != nil {
		return nil, err
	}

	globalDataDir, err := GlobalDataDir()
	if err != nil {
		return nil, err
	}

	globalConfigDir, err := GlobalConfigDir()
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"Home":         homeDir,
		"Data":         userDataDir,
		"GlobalData":   globalDataDir,
		"SiteData":     siteDataDir,
		"Config":       userConfigDir,
		"GlobalConfig": globalConfigDir,
		"SiteConfig":   siteConfigDir,
		"Cache":        userCacheDir,
		"Log":          userLogDir,
	}, nil
}

func HomeDir() (string, error) {
	return fs.Dir()
}

func (conf AppConf) UserDataDir() (string, error) {
	return conf.userDataDir()
}

func (conf AppConf) SiteDataDir() (string, error) {
	return conf.siteDataDir()
}

func (conf AppConf) UserConfigDir() (string, error) {
	return conf.userConfigDir()
}

func (conf AppConf) SiteConfigDir() (string, error) {
	return conf.siteConfigDir()
}

func (conf AppConf) UserCacheDir() (string, error) {
	return conf.userCacheDir()
}

func (conf AppConf) UserLogDir() (string, error) {
	return conf.userLogDir()
}

func GlobalDataDir() (string, error) {
	return globalDataDir()
}

func GlobalConfigDir() (string, error) {
	return globalConfigDir()
}

func UserDataDir(appName string) (string, error) {
	return (AppConf{Name: appName}).UserDataDir()
}

func SiteDataDir(appName string) (string, error) {
	return (AppConf{Name: appName}).SiteDataDir()
}

func UserConfigDir(appName string) (string, error) {
	return (AppConf{Name: appName}).UserConfigDir()
}

func SiteConfigDir(appName string) (string, error) {
	return (AppConf{Name: appName}).SiteConfigDir()
}

func UserCacheDir(appName string) (string, error) {
	return (AppConf{Name: appName}).UserCacheDir()
}

func UserLogDir(appName string) (string, error) {
	return (AppConf{Name: appName}).UserLogDir()
}
