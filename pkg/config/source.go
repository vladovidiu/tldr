package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
)

const (
	tldrurl   = "https://github.com/tldr-pages/tldr.git"
	directory = "/.cache/tldr"
)

var (
	SourceDir = getCloneDir()
)

func CloneSource() error {
	_, gitErr := git.PlainClone(SourceDir,
		false, &git.CloneOptions{
			URL:               tldrurl,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			Progress:          os.Stdout,
		})

	if gitErr == nil {
		fmt.Printf("Succesfully clones into: %s\n", SourceDir)
	}

	return gitErr
}

func PullSource() error {
	r, gitErr := git.PlainOpen(SourceDir)

	if gitErr != nil {
		return gitErr
	}

	w, worktreeErr := r.Worktree()
	if worktreeErr != nil {
		return worktreeErr
	}

	worktreeErr = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
	})

	if worktreeErr != nil {
		fmt.Printf(" %s\n", worktreeErr.Error())
	} else {
		fmt.Printf("Succesfully cloned into %s\n", SourceDir)
	}

	return worktreeErr
}

func getCloneDir() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return strings.Join([]string{homeDir, directory}, "")
}
