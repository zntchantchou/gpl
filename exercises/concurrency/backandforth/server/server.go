package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:3001")
	fmt.Println("[Server]")

	if err != nil {
		fmt.Println("err")
	}
	// defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("err")
		}
		handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	s := bufio.NewScanner(c)
	for s.Scan() {
		fmt.Println(s.Text())
		time.Sleep(2 * time.Second)
	}
}
