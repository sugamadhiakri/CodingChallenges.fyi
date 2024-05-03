package wc

import (
	"bytes"
	"io"
	"os"
	"unicode/utf8"
)

func CountTotal(filename string) (int, int, int, int) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	buf := make([]byte, 32*1024)
	totalByte := 0
	totalLines := 0
	totalChars := 0
	totalWords := 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		totalByte += len(buf[:c])
		totalLines += bytes.Count(buf[:c], lineSep)
		totalChars += utf8.RuneCountInString(string(buf[:c]))
		totalWords += len(bytes.Fields(buf[:c]))

		switch {
		case err == io.EOF:
			return totalByte, totalLines, totalChars, totalWords

		case err != nil:
			panic(err.Error())
		}
	}
}
