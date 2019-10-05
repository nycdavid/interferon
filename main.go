package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
)

func main() {
	Info("git remotes")

	// git.PlainOpen(path/to/.git)
	cwd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	gitPath := fmt.Sprintf("%s/.git", cwd)
	repo, err := git.PlainOpen(gitPath)
	CheckIfError(err)

	remotes, err := repo.Remotes()
	CheckIfError(err)

	for _, remote := range remotes {
		fmt.Println(remote.String())
	}
}
