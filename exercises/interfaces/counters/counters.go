package counters

import (
	"fmt"
	"io"
	"strings"
)

type ByteCounter struct {
	W       io.Writer
	Written int64
}

func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	n, err = bc.W.Write(p)
	bc.Written += int64(n)
	return
}

// -----------------------------------------------------------------------'

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	// convert to string
	asStr := string(p)
	fmt.Println("before: ", asStr)
	trimmed := strings.Trim(asStr, " ")
	fmt.Println("after: ", trimmed)

	numberOfWords := strings.Count(trimmed, " ") + 1
	*wc += WordCounter(numberOfWords)
	// find all spaces and add 1
	return numberOfWords, nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	n := 0
	fmt.Println("[Linecounter] Write")
	for _, char := range string(p) {
		if char == '\n' {
			n++
		}
	}
	return n + 1, nil
}
