package cmd

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/vladovidiu/tldr/pkg/config"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates the tldr pages for newer version",
	Long:  "use update once in a while to make sure tldr pages are up to date.",
	Run: func(cmd *cobra.Command, args []string) {
		updateTldrPages()
	},
}

func updateTldrPages() error {
	sourceDir := config.SourceDir

	repo, err := git.PlainOpen(sourceDir)
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	err = worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
	})

	if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("Successfully updated into: %s\n", sourceDir)
	}

	return err
}
