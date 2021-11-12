package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

var portNumber = flag.Int("port", 0, "port number for the clock server")

func main() {
	flag.Parse()
	fmt.Println("pFlag", *portNumber)
	p := strconv.FormatInt(int64(*portNumber), 10)
	listener, err := net.Listen("tcp", "localhost:"+p)
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("[clock1] listening")
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
		fmt.Println("Closed a connection")
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	fmt.Println(*portNumber)
	for {
		fmt.Println("clock Printing")
		_, err := io.WriteString(c, createTime(portNumber))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func createTime(p *int) string {
	t := time.Now()
	switch *(p) {
	case 8000:
		fmt.Println("PORT IS 8000")
		t = t.Add(1 * time.Hour)
		fmt.Println("t after add", t)
	case 8010:
		fmt.Println("PORT IS 8010")
		t = t.Add(5 * time.Hour)
	case 8020:
		fmt.Println("PORT IS 8020")
		t = t.Add(8 * time.Hour)
	default:
		fmt.Println("PORT IS DEFAULT")
	}
	return t.Format("15:04:05\n")
}
