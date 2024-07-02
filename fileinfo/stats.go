// Description: This file contains the function to check the content type of a file.

package fileinfo

import "os"

// DirExists checks if a directory exists
// Arguments:
//
//	path -- the path to the directory
func DirExists(path string) bool {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
