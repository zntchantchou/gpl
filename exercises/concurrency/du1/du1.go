package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walkdir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkdir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// // we can return nil because return type is a slice
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error is %v", err)
		return nil
	}
	return entries
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fmt.Println("roots ", roots)
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkdir(root, fileSizes)
		}
		close(fileSizes)
	}()
	var nbFiles, totalSize int64
	for size := range fileSizes {
		nbFiles++
		totalSize += size
	}
	printDiskUsage(nbFiles, totalSize)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %1.f GB \n", nfiles, float64(nbytes)/1e9)
}
