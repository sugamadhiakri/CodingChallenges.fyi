package wc

import "os"

func readFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)

	if err != nil {
		panic(err.Error())
	}

	return data
}
