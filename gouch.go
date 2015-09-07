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
		fmt.Fprintf(os.Stderr, "Usage: gouch file ...\n")
		return
	}

	for _, arg := range os.Args[1:] {
		newFileName := arg
		newFile, err = os.Create(newFileName)
		if err != nil {
			log.Fatal(err)
		}
		newFile.Close()
	}
}
