package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		if(strings.ToUpper(input.Text()) == "EXIT") {
			return
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("appeared %d times :%s\n", n, line)
			 }
			}
	}

// NOTE: ignoring potential errors from input.Err()

}