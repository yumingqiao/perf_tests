package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

//only reads files one level from the given directory.
func main() {

	directory := flag.String("directory", "configs", "directory relative to root, from which all files will be read")
	flag.Parse()
	fmt.Println("directory:", *directory)

	// test start
	readStart := time.Now()
	totalFiles := 0
	var totalSize int64
	if err := filepath.Walk(*directory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			panic(err)
		}
		totalFiles += 1
		totalSize += info.Size()
		_, err = os.ReadFile(fmt.Sprintf("%s/%s", *directory, info.Name()))
		return err
	}); err != nil {
		panic(err)
	}
	duration := time.Since(readStart)
	fmt.Println(fmt.Sprintf("spend %v to read %d files from direcotry %v, total size %dKB", duration, totalFiles, *directory, totalSize/1e3))
}
