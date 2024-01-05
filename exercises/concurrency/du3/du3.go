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

var sema = make(chan struct{}, 20)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// print the result periodically if flag verbose is provided
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(time.Millisecond * 500)
	}
	fileSizes := make(chan int64)
	var nbFiles, totalBytes int64
	var wg sync.WaitGroup

	for _, root := range roots {
		wg.Add(1)
		go walkdir(root, fileSizes, &wg)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

loop:
	for {
		select {
		// a range loops does this implicitly
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nbFiles++
			totalBytes += size
		case <-tick:
			printDiskUsage(nbFiles, totalBytes)
		}
	}
	printDiskUsage(nbFiles, totalBytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	// round 3 decimal places
	fmt.Printf("%d files, %.3f MB \n", nfiles, float64(nbytes)/1e6)
}

func walkdir(dir string, fileSizes chan<- int64, n *sync.WaitGroup) {
	defer n.Done()
	for _, entry := range dirInfo(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkdir(subdir, fileSizes, n)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirInfo(dir string) []os.FileInfo {
	// can only be called 20 times concurrently
	// acquire token = adds a 'token' to the counting semaphore
	sema <- struct{}{}
	// release token = frees up spot for another token in the semaphore
	defer func() {
		<-sema
	}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error is %v", err)
		return nil
	}
	return entries
}
