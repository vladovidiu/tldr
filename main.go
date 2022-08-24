package main

import (
	"fmt"
	"log"
	"time"

	"github.com/vladovidiu/tldr/pkg/config"
	"github.com/vladovidiu/tldr/pkg/pages"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	clone  = kingpin.Flag("clone", "clones the tldr repo").Short('c').Bool()
	update = kingpin.Flag("update", "pulls the latest commits from origin").Short('u').Bool()
	page   = kingpin.Arg("command", "Name of the command.").Required().String()
)

func main() {
	start := time.Now()
	kingpin.Version("tldr version 0.0.1")

	kingpin.Parse()

	if *clone {
		config.StartUp()
		err := config.CloneSource()
		if err != nil {
			log.Fatal(err)
		}
	}

	if *update {
		config.StartUp()
		err := config.PullSource()

		if err != nil {
			log.Fatal(err)
		}
	}

	s, err := pages.Read(*page)

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("%s\n", s.String())
	elapsed := time.Since(start)
	fmt.Printf("Query finished in: %s\n", elapsed)
}
