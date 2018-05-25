package fileinfo

import (
	"net/http"
	"os"
)

func CheckContent(out *os.File) (string, err) {
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
