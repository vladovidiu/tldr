package utils

import (
	"log"
	"runtime"
)

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
		log.Fatal("Operating System cannot be recognized.")
	}

	return name
}
