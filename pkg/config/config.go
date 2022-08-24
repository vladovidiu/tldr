package config

import (
	"log"
	"os"
)

var (
	SourceDir = getSourceDir()
)

func getSourceDir() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return homeDir + "/.cache/tldr"
}
