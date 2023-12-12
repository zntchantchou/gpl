package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// interface value is the zero value for an interface
// it contains the type and the value
var w io.Writer

func main() {
	fmt.Println("w != nil :", w != nil) // false
	// panics => nil pointer error
	// w's interface value => TYPE: nil, VALUE nil
	// w.Write([]byte("Bonjour"))

	// equivalent to the explicit conversion io.Writer(os.Stdout)
	 w = os.Stdout
	// w is converted such as : io.Writer(os.Stdout)
	// the interface value of os.Stdout is captured

	w.Write([]byte("Salut\n")) // "salut"
	// is equivalent to
	// os.Stdout.Write([]byte("Salut"))

	w = new(bytes.Buffer)
	w.Write([]byte("Written in bytes buffen"))

	fmt.Println(w)
}

