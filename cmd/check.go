package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vladovidiu/tldr/pkg/pages"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check the tldr for given command",
	Long:  `check the tldr for given command`,
	Run: func(cmd *cobra.Command, args []string) {
		readPage(args[0])
	},
}

func readPage(command string) {
	s, err := pages.Read(command)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	fmt.Printf("%s\n", s.String())
}
