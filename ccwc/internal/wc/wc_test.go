package wc

import (
	"os"
	"testing"
)

// Helper function to create a temporary test file with given content
func createTempFile(content []byte) (string, error) {
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	_, err = tempFile.Write(content)
	if err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}

func TestCountTotal(t *testing.T) {
	testCases := []struct {
		content       []byte
		expectedBytes int
		expectedLines int
		expectedChars int
		expectedWords int
	}{
		{[]byte("Hello, world!"), 13, 0, 13, 2},
		{[]byte(""), 0, 0, 0, 0},
		{[]byte("Hello\nworld!"), 12, 1, 12, 2},
		{[]byte("Hello\nworld!\n"), 13, 2, 13, 2},
	}

	for _, tc := range testCases {
		fileName, err := createTempFile(tc.content)
		if err != nil {
			t.Fatalf("Error creating temporary test file: %v", err)
		}
		defer os.Remove(fileName)

		bytes, lines, chars, words := CountTotal(fileName)

		if bytes != tc.expectedBytes {
			t.Errorf("CountTotal(%s) bytes = %d; want %d", tc.content, bytes, tc.expectedBytes)
		}

		if lines != tc.expectedLines {
			t.Errorf("CountTotal(%s) lines = %d; want %d", tc.content, lines, tc.expectedLines)
		}

		if chars != tc.expectedChars {
			t.Errorf("CountTotal(%s) chars = %d; want %d", tc.content, chars, tc.expectedChars)
		}

		if words != tc.expectedWords {
			t.Errorf("CountTotal(%s) words = %d; want %d", tc.content, words, tc.expectedWords)
		}
	}
}
