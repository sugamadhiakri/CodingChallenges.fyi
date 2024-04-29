package main

import (
	"flag"
	"fmt"
)

func main() {
	var countBytes bool

	flag.BoolVar(&countBytes, "c", false, "Count Bytes")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Error: No files were provided.")
	}

	fileName := args[0]

	if countBytes {
		fmt.Printf("Counting Bytes for %s", fileName)
	}
}
