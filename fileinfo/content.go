// Description: This file contains the function to check the content type of a file.
package fileinfo

import (
	"net/http"
	"os"
)

// CheckContent checks the content type of a file
// Arguments:
//
//	out -- the file to check
//
// Returns:
//
//	string -- the content type of the file
//	error -- an error if the file could not be read
func CheckContent(out *os.File) (string, error) {
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
