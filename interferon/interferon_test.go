package interferon

import (
	"os"
	"os/exec"
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

func TestDiffMaster(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	diff, err := DiffMaster()

	if err != nil {
		t.Error(err)
	}
	if len(diff) != 1 {
		t.Error("Expected slice of size 1")
	}
}

// helpers
func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}
