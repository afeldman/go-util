package filesystem

import "os"

func MkDir(path string, permission os.FileMode){
	//choose your permissions well
	pathErr := os.MkdirAll(path, permission)

	//check if you need to panic, fallback or report
	if pathErr != nil {
		fmt.Println(pathErr)
	}

}
