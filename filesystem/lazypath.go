// package filesystem contains functions for working with the filesystem
package filesystem

import (
	"path/filepath"

	"github.com/afeldman/go-util/env"
)

// LazyPath is a struct that contains a function to get a path
type LazyPath struct {
	EnvironmentVariable string
	DefaultFn           func() string
}

// Path gets a path
// Arguments:
//
//	elem -- the elements of the path
//
// Returns:
//
//	string -- the path
func (lazypath LazyPath) Path(elem ...string) string {
	base := env.GetEnvOrDefault(lazypath.EnvironmentVariable,
		lazypath.DefaultFn())
	return filepath.Join(base, filepath.Join(elem...))
}
