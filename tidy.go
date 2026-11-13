package main

import (
	"flag"
	"fmt"
	"os"
)

var TidyCommand = flag.NewFlagSet("tidy", flag.ExitOnError)

func Tidy() {
	path := PathFlag(TidyCommand)
	TidyCommand.Parse(os.Args[1:])

	file, _ := GetMod(*path)
	if file == nil {
		os.Exit(1)
	}

	// tidy file
	stat, err := os.Stat(file.Path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "(unexpected) tidy: %v", err)
		os.Exit(1)
	}

	err = os.WriteFile(file.Path, []byte(file.Tidy()), stat.Mode())
	if err != nil {
		fmt.Fprintf(os.Stderr, "tidy: %v", err)
		os.Exit(1)
	}
}
