package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conn", conn)
	// defer conn.Close()

	// go fns.HandleConn(conn)
	// // print echo server response
	// go mustCopy(os.Stdout, conn)
	// // send input to connection
	// mustCopy(os.Stdin, conn)
}

// func mustCopy(dst io.Writer, src io.Reader) {
// 	if _, err := io.Copy(dst, src); err != nil {
// 		log.Fatal(err)
// 	}
// }
