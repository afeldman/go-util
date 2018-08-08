package filesystem

import "os"

func MkDir(path string, permission os.FileMode) error{
	//choose your permissions well
	pathErr := os.MkdirAll(path, permission)
	return pathErr
}
