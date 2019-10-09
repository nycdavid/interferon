package interferon

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	testFilePath := "test/.interfere"
	files, err := ReadFile(testFilePath)
	if err != nil {
		t.Error(err)
	}

	// expect
	if len(files) != 3 {
		t.Error("Expected slice of size 3")
	}
}
