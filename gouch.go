package main

import (
	"fmt"
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: gouch file\n")
		return
	}
	newFileName := os.Args[1]
	newFile, err = os.Create(newFileName)
	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()
}
