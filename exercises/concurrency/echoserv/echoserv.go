package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// echo server receives
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func HandleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	// listen for user input
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func main() {
	ln, err := net.Listen("tcp", "localhost:3001")
	if err != nil {
		fmt.Println(err)
	}
	for {
		fmt.Println("listening...")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		HandleConn(conn)
	}
}
