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

func TestCountBytes(t *testing.T) {
	testCases := []struct {
		content  []byte
		expected int
	}{
		{[]byte("Hello, world!"), 13},
		{[]byte(""), 0},
	}

	for _, tc := range testCases {
		fileName, err := createTempFile(tc.content)
		if err != nil {
			t.Fatalf("Error creating temporary test file: %v", err)
		}
		defer os.Remove(fileName)

		byteCount := CountBytes(fileName)

		if byteCount != tc.expected {
			t.Errorf("CountBytes(%s) = %d; want %d", tc.content, byteCount, tc.expected)
		}
	}
}
