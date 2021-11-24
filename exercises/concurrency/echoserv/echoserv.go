package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// echo server receives

func main() {
	fmt.Println("[Echo server] starting")
	serve()
	fmt.Println("[Echo server] done")
}

func serve() {
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	fmt.Println("[serve] listening on port 8000 ...")
	if err != nil {
		panic(err)
	}
	conn, err := l.Accept()
	fmt.Println("[serve] connexion received ...")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	handleConn(conn)
}

func handleConn(c net.Conn) {
	fmt.Println("[handleConn] start")
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 100*time.Millisecond)
	}
}
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
}
