package interferon

import (
	"bytes"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

// read from .interfere
// run a git diff
// see if the item exists

var execCommand = exec.Command

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
	cmd := execCommand("git", "--no-pager", "diff", "--name-only", "HEAD..master")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	files := strings.Split(out.String(), "\n")

	return files, nil
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
