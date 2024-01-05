package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, fileSizes chan<- int64, n *sync.WaitGroup) {
	defer n.Done()
	for _, entry := range dirEnts(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes, n)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// // we can return nil because return type is a slice
func dirEnts(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error is %v", err)
		return nil
	}
	return entries
}

func main() {
	var verbose = flag.Bool("v", false, "show verbose progress messages")
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fmt.Println("roots ", roots)

	// print the result periodically if flag verbose is provided
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(time.Millisecond * 100)
	}
	fileSizes := make(chan int64)
	var nbFiles, totalBytes int64
	var n sync.WaitGroup

	for _, root := range roots {
		walkDir(root, fileSizes, &n)
	}

	// limit concurrency using waitgroup (limits how many times walkdir can be called concurrently)
	go func() {
		n.Wait()
		close(fileSizes)
	}()

loop:
	for {
		select {
		// a range loops does this implicitly
		case size, ok := <-fileSizes:
			if !ok {
				fmt.Println("Ch has been closed")
				break loop
			}
			nbFiles++
			totalBytes += size
		case <-tick:
			printDU(nbFiles, totalBytes)
		}
	}
}

func printDU(nfiles, nbytes int64) {
	fmt.Printf("%d files, %1.f GB \n", nfiles, float64(nbytes)/1e9)
}
