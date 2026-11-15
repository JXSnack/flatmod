package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JXSnack/flatmod/gradle"
)

var BuildCommand = flag.NewFlagSet("build", flag.ExitOnError)

func Build() {
	path := PathFlag(BuildCommand)
	out := BuildCommand.String("o", "", "output name for jar")
	BuildCommand.Parse(os.Args[1:])

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

	// execute build
	args := []string{"build"}
	if *out != "" {
		args = append(args, fmt.Sprintf("-PjarName=%s", *out))
	}

	err = workspace.ExecGradlew(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}
}
