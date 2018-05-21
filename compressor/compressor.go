package compress

import (
	"os"
	"path/filepath"
	"archive/tar"
	"strings"
	"io"
	"compress/gzip"
	"log"
)

func check(e error) {
 	if e != nil {
		log.Println(e)
 		panic(e)
 	}
}

func TarIt(source, target string) string{	
	target = target + ".tar"
	
	tarfile, err := os.Create(target)
	check(err)
	
	defer tarfile.Close()

	tarball := tar.NewWriter(tarfile)
	defer tarball.Close()

	info, err := os.Stat(source)
	check(err)

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, 
	func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if err := tarball.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(tarball, file)
		return err
	})	
	return target
}

func GzipIt(source, target string, hcomment string){
	log.Println(source)
	log.Println(target)
	
	reader, err := os.Open(source)
	check(err)

	filename := filepath.Base(source)
	target = target + ".gz"
	
	writer, err := os.Create(target)
	check(err)
	
	defer writer.Close()

	archiver := gzip.NewWriter(writer)
	archiver.Name = filename
	archiver.Comment = hcomment
	defer archiver.Close()

	_, err = io.Copy(archiver, reader)
}
