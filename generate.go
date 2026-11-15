package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JXSnack/flatmod/gradle"
)

var GenerateCommand = flag.NewFlagSet("generate", flag.ExitOnError)

func Generate() {
	path := PathFlag(GenerateCommand)
	GenerateCommand.Parse(os.Args[:1])

	file, _ := GetMod(*path)
	if file == nil {
		os.Exit(1)
	}

	// make workspace
	workspace, err := gradle.NewWorkspace(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}

	err = workspace.Generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}
}
