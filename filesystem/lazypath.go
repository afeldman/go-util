package filesystem

import (
	"path/filepath"

	"github.com/afeldman/go-util/env"
)

type LazyPath struct {
	EnvironmentVariable string
	DefaultFn           func() string
}

func (lazypath LazyPath) Path(elem ...string) string {
	base := env.GetEnvOrDefault(lazypath.EnvironmentVariable,
		lazypath.DefaultFn())
	return filepath.Join(base, filepath.Join(elem...))
}
