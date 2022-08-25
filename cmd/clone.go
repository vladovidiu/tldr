package cmd

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	config "github.com/vladovidiu/tldr/pkg/config"

	"github.com/spf13/cobra"
)

const (
	tldrUrl = "https://github.com/tldr-pages/tldr.git"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "clone the tldr package",
	Long:  `clone the tldr package`,
	Run: func(cmd *cobra.Command, args []string) {
		cloneTldrPages()
	},
}

func cloneTldrPages() error {
	sourceDir := config.SourceDir

	_, err := git.PlainClone(sourceDir, false, &git.CloneOptions{
		URL:               tldrUrl,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	})

	if err == nil {
		fmt.Printf("Successfully cloned tldr pages in: %s\n", sourceDir)
	}

	return err
}
