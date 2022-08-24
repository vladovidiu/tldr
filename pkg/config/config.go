package config

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func StartUp() error {
	fmt.Printf("%s\n", initialMessage())

	return nil
}

func OSName() (name string) {
	switch osname := runtime.GOOS; osname {
	case "windows":
		name = osname
	case "darwin":
		name = "osx"
	case "linux":
		name = osname
	case "solaris":
		name = "sunos"
	default:
		log.Warn("Operating System couldn't be recognized")
	}

	return name
}
