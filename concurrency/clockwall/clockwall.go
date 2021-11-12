package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
)

var newyork = flag.Int("newyork", 0, "city and corresponding clock server port")
var tokyo = flag.Int("tokyo", 0, "city and corresponding clock server port")
var london = flag.Int("london", 0, "city and corresponding clock server port")

var londonTime = ""
var tokyoTime = ""
var newyorkTime = ""

func main() {
	// listen to servers if flag is provided
	flag.Parse()
	fmt.Println(*newyork)
	fmt.Println(*tokyo)
	fmt.Println(*london)
	go handleConnection(*london, &londonTime)
	go handleConnection(*newyork, &newyorkTime)
	go handleConnection(*tokyo, &tokyoTime)
	for {
		fmt.Print("Time in Tokyo: ", tokyoTime)
		fmt.Print("Time in NewYork: ", newyorkTime)
		fmt.Print("Time in London: ", londonTime)
		fmt.Println(" --------------------------------------")
		time.Sleep(1 * time.Second)
	}
}

func createConnectionString(p int) string {
	formatted := strconv.FormatInt(int64(p), 10)
	return "localhost:" + formatted
}

func handleConnection(p int, curTime *string) {
	if p == 0 {
		fmt.Println("Closing connection")
		return
	}
	conn, err := net.Dial("tcp", createConnectionString(p))
	if err != nil {
		fmt.Println("[handleConnection] an error occured", err, createConnectionString(p))
	}
	defer conn.Close()
	for {
		newReader := bufio.NewReader(conn)
		v, err := newReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		// store current time in curTime
		*curTime = v
		time.Sleep(1 * time.Second)
	}

}
