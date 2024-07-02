// package filesystem contains functions for working with the filesystem
package filesystem

import (
	"crypto/sha256"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
)

// 8KB
const filechunk = 8192

// filenames
var filename []string

// hashes
type FileHashNode struct {
	Path string
	Hash []byte
}

// filehash
type FileHash struct {
	Nodes []FileHashNode
}

// Hash hashes the files in a directory
// Arguments:
// dir -- the directory to hash
func (fh *FileHash) Hash(dir string) {

	err := filepath.Walk(dir, countfiles)
	check(err)

	//go through all files
	for _, file := range filename {
		// Open the file for reading
		infile, err := os.Open(file)
		check(err)

		defer infile.Close()

		// Get file info
		info, err := infile.Stat()
		check(err)

		// Get the filesize
		filesize := info.Size()

		// Calculate the number of blocks
		blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

		// Start hash
		hash := sha256.New()

		// Check each block
		for i := uint64(0); i < blocks; i++ {
			// Calculate block size
			blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))

			// Make a buffer
			buf := make([]byte, blocksize)

			// Make a buffer
			infile.Read(buf)

			// Write to the buffer
			io.WriteString(hash, string(buf))
		}

		// Output the results
		filehashnote := FileHashNode{infile.Name(), hash.Sum(nil)}
		fh.Nodes = append(fh.Nodes, filehashnote)
	}

}

// check checks for errors
// Arguments:
// e -- the error to check
func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

// countfiles counts the files in a directory
// Arguments:
// path -- the path to the directory
// info -- the file info
// err -- an error
//
// Returns:
//
// error -- an error
func countfiles(path string, info os.FileInfo, err error) error {
	check(err)
	abs, err := filepath.Abs(path)
	check(err)
	filename = append(filename, abs)
	return nil
}
