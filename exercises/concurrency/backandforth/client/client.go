package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3001")
	if err != nil {
		fmt.Println(err)
	}
	go mustCopy(conn, os.Stdin)
	r := bufio.NewReader(os.Stdin)
	err, input := r.ReadString('\n')
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
}
