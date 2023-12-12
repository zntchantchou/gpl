package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	wordPtr := flag.String("sha", "256", "type of sha hashing to use")
	allowedFlags := []string{"256", "384", "512"}
	flag.Parse()
	isValidFlag := slices.Contains[[]string, string](allowedFlags, *wordPtr)
	fmt.Printf("isValidFlag %v\n ", isValidFlag)
	if !isValidFlag {
		log.Fatal("Flag is invalid")
	}
	if isValidFlag {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if *wordPtr == "256" {
				hashed := sha256.Sum256([]byte(scanner.Text()))
				fmt.Print("Text : ", scanner.Text(), " hashed: ", hashed, " \n")
			}
			if *wordPtr == "384" {
				hashed := sha512.Sum384([]byte(scanner.Text()))
				fmt.Print("Text : ", scanner.Text(), " hashed: ", hashed, " \n")
			}
			if *wordPtr == "512" {
				hashed := sha512.Sum512([]byte(scanner.Text()))
				fmt.Print("Text : ", scanner.Text(), " hashed: ", hashed, "\n")
			}
		}
	}
}
