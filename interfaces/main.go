package main

import (
	"bytes"
	"fmt"

	"github.com/zclmk/interfaces/countingwriter"
)

func main() {
	b := &bytes.Buffer{}
	wr, written := countingwriter.Countingwriter(b)
	wr.Write([]byte("This is content"))
	fmt.Println("written", *written)
	wr.Write([]byte("This is more content dhdzhihizdhizhdidz"))
	fmt.Println("written", *written)
}
