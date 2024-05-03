package main

import (
	"flag"
	"fmt"

	"github.com/sugamadhiakri/ccwc/internal/wc"
)

func main() {
	var countBytes bool
	var countLines bool
	var countChars bool
	var countWords bool

	flag.BoolVar(&countBytes, "c", false, "Count Bytes")
	flag.BoolVar(&countLines, "l", false, "Count Lines")
	flag.BoolVar(&countChars, "m", false, "Count Chars")
	flag.BoolVar(&countWords, "w", false, "Count Words")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		panic("Error: No files were provided.")
	}

	fileName := args[0]

	bytes, lines, chars, words := wc.CountTotal(fileName)

	if countBytes {
		fmt.Printf("%d %s\n", bytes, fileName)
	}

	if countLines {
		fmt.Printf("%d %s\n", lines, fileName)
	}

	if countWords {
		fmt.Printf("%d %s\n", words, fileName)
	}

	if countChars {
		fmt.Printf("%d %s\n", chars, fileName)
	}
}
