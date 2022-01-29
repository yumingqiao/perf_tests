package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func main() {
	files := flag.Int("files", 100, "number of files to create")
	sizeKB := flag.Int64("size-kb", 1e2, "size of a file in KB")
	directory := flag.String("o", "", "where to save the output files")
	flag.Parse()

	fmt.Println("files:", *files)
	fmt.Println("size-kb:", *sizeKB)
	fmt.Println("directory:", *directory)
	if *directory == "" {
		*directory = randomString(5)
	}
	fmt.Println(fmt.Sprintf("writing to %s", *directory))
	if err := os.Mkdir(*directory, 0777); err != nil {
		panic(err)
	}

	// test start
	writeStart := time.Now()
	for i := 0; i < *files; i++ {
		f, err := os.Create(fmt.Sprintf("%v/dummy%d", *directory, i))
		if err != nil {
			panic(err)
		}
		if err := f.Truncate(*sizeKB * 1e3); err != nil {
			panic(err)
		}
	}
	duration := time.Since(writeStart)
	fmt.Println(fmt.Sprintf("spend %v to write %d files to direcotry %v, each has %dKB", duration, *files, *directory, *sizeKB))
}
