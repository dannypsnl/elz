package main

import (
	"github.com/elz-lang/elz/src/elz/cmd"

	"github.com/sirupsen/logrus"
)

func main() {
	err := cmd.Compiler.Execute()
	if err != nil {
		logrus.Fatalf("elz compiler failed: %s", err)
	}
}
