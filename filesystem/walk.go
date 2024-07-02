// package filesystem contains functions for working with the filesystem
package filesystem

import (
	"os"
	"path/filepath"
)

// FilePathWalkDir walks a directory and returns a list of files
// Arguments:
//
//	root -- the directory to walk
//
// Returns:
//
//	 []string -- a list of files
//		error -- an error if the directory could not be walked
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
