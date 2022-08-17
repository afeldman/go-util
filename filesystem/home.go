package filesystem

import (
	"path"
	"path/filepath"

	"github.com/afeldman/go-util/env"
)

const DefaultHomeEnvVar = "GOPIZZA_HOME"

var (
	HomePrefix string = ""
	homepath          = LazyPath{
		EnvironmentVariable: DefaultHomeEnvVar,
		DefaultFn:           defaultHome,
	}

	userpath = LazyPath{
		EnvironmentVariable: "HOME",
		DefaultFn:           defaultUserHome,
	}

	binpath = LazyPath{
		EnvironmentVariable: "GOPIZZA_BINPATH",
		DefaultFn:           defaultBinPath,
	}
)

func defaultHome() string {
	return filepath.Join(HomePrefix, "gopizza")
}

func defaultUserHome() string {
	if home := env.Plan9GetEnv("HOME"); len(home) > 0 {
		return home
	}
	return env.Plan9GetEnv("USERPROFILE")
}

func defaultBinPath() string {
	return filepath.Join(HomePrefix, "bin")
}

// Barrel returns the path to the fish barrel.
func Barrel() string {
	return homepath.Path("Barrel")
}

// Rigs returns the path to the fishing rigs.
func Rigs() string {
	return homepath.Path("Rigs")
}

// String returns Home as a string.
//
// Implements fmt.Stringer.
func String() string {
	return homepath.Path("")
}

// UserHome returns the home path.
func UserHome() string {
	return userpath.Path("")
}

// BinPath is the path where executables should be installed by gofish.
func BinPath() string {
	return binpath.Path("")
}

// DefaultRig returns the name of the default fishing rig.
func DefaultRig() string {
	return path.Join("github.com", "fishworks", "fish-food")
}
