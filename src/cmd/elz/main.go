package main

import (
	"log"

	"github.com/elz-lang/elz/src/elz/cmd"
)

func main() {
	err := cmd.Compiler.Execute()
	if err != nil {
		log.Fatalf("elz compiler failed: %s", err)
	}
}
