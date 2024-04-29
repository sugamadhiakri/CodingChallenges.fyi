package main

import (
	"flag"
	"fmt"

	"github.com/sugamadhiakri/ccwc/internal/wc"
)

func main() {
	var countBytes bool

	flag.BoolVar(&countBytes, "c", false, "Count Bytes")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		panic("Error: No files were provided.")
	}

	fileName := args[0]

	if countBytes {
		fmt.Printf("%d %s\n", wc.CountBytes(fileName), fileName)
	}
}
