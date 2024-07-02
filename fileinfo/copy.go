// Description: This file contains the function to check the content type of a file.
package fileinfo

import (
	"io"
	"log"
	"os"
)

/*
	Copy a file

Arguments:

		src -- the source file to copy
	    dest -- the destination file to copy to

Returns:

	error -- an error if the file could not be copied
*/
func Fcopy(src, dest string) error {
	// Open original file
	originalFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer originalFile.Close()

	// Create new file
	newFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		return err
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
