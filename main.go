package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var SharedPaths []string = []string{
	"shared/",
}

func main() {
	cmd := exec.Command("git", "--no-pager", "diff", "--name-only", "HEAD..master")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	files := strings.Split(out.String(), "\n")
	possibleConflicts := []string{}

	for _, sharedPath := range SharedPaths {
		filepath.Walk(sharedPath, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && exists(files, path) {
				possibleConflicts = append(possibleConflicts, path)
			}

			return nil
		})
	}

	fmt.Println("Possible conflicts:")

	for _, conflict := range possibleConflicts {
		fmt.Println(conflict)
	}
}

func exists(slc []string, el string) bool {
	for _, slcEl := range slc {
		if el == slcEl {
			return true
		}
	}

	return false
}
