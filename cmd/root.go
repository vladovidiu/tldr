package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tldr",
	Short: "tldr shows as much info you need on commands",
	Long:  `A very intriguing tool to let you know the must know for commands`,
}

func init() {
	rootCmd.AddCommand(cloneCmd)
	rootCmd.AddCommand(updateCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
