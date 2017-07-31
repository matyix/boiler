package main

import (
	"log"

	"github.com/matyix/boiler/command"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}


