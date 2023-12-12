package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Counting from input")
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v \n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	fmt.Printf("\n ----- Filename: %s ----- \n\n", f.Name())
	for scanner.Scan() {
		if strings.ToUpper(scanner.Text()) == "EXIT" {
			return
		}
		if(len(scanner.Text()) > 1) {
			counts[scanner.Text()]++
		}
		if counts[scanner.Text()] > 1 {
			fmt.Printf("%s appeared %d times \n", scanner.Text(), counts[scanner.Text()])
		}
	}
}
