//

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

var filename []string

type FileHashNode struct {
	Path string
	Hash []byte
}

type FileHash struct {
	Nodes []FileHashNode
}

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
		//fmt.Printf("Filename: %s \tHash: %x\n", filepath.Base(file.Name()), hash.Sum(nil))
		filehashnote := FileHashNode{infile.Name(), hash.Sum(nil)}
		fh.Nodes = append(fh.Nodes, filehashnote)
	}

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func countfiles(path string, info os.FileInfo, err error) error {
	check(err)
	abs, err := filepath.Abs(path)
	check(err)
	filename = append(filename, abs)
	return nil
}
