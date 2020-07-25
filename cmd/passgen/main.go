package main

import (
	"log"

	"github.com/antolis/passgen/cmd"
)

func main() {
	root := cmd.RootCmd()
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
