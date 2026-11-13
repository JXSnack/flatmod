package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JXSnack/flatmod/gradle"
)

var GradleCommand = flag.NewFlagSet("gradle", flag.ExitOnError)

func Gradle() {
	path := PathFlag(GradleCommand)
	GradleCommand.Parse(os.Args[1:])

	file, _ := GetMod(*path)
	if file == nil {
		os.Exit(1)
	}

	// make workspace
	workspace, err := gradle.NewWorkspace(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "idea: %v\n", err)
	}

	err = workspace.Generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "idea: %v\n", err)
	}

	// execute command
	//workspace.ExecGradlew(os.Args)
}
