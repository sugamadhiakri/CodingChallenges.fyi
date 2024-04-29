package wc

func CountBytes(fileName string) int {
	return len(readFile(fileName))
}
