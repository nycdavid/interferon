package interferon

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// read from .interfere
// run a git diff
// see if the item exists

func ReadFile(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return []string{}, err
	}

	buf := bytes.NewBuffer(data)
	files := strings.Split(buf.String(), "\n")
	files = compact(files)

	return files, nil
}

func DiffMaster() ([]string, error) {
	cmd := exec.Cmd("git", "--no-pager", "diff", "--name-only", "HEAD..master")
	err := cmd.Run
}

func compact(slice []string) []string {
	var newSlice []string

	for _, el := range slice {
		if el != "" {
			newSlice = append(newSlice, el)
		}
	}

	return newSlice
}
