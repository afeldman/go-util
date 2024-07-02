/*
package crypto
This package provides functions to calculate the checksum of a file
*/
package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

/*
	Calculate the SHA256 checksum of a file

Arguments:

	filePath -- the path to the file to calculate the checksum for

Returns:

	string -- the SHA256 checksum of the file
	error -- an error if the file could not be read
*/
func SHA256sum(filePath string) (result string, err error) {
	// Open the file
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	// Create a new SHA256 hash
	h := sha256.New()

	// Copy the file contents to the hash
	if _, err = io.Copy(h, f); err != nil {
		return
	}

	// Get the checksum
	result = hex.EncodeToString(h.Sum(nil))
	return
}

/*
	md5 file hash

Arguments:

	filePath -- the path to the file to calculate the checksum for

Returns:

	string -- the md5 checksum of the file
	error -- an error if the file could not be read
*/
func MD5sum(filePath string) (result string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return
	}

	result = hex.EncodeToString(hash.Sum(nil))
	return
}
