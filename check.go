package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JXSnack/flatmod/modfile"
)

// TODO: multiple errors instead of breaking at the first problem

func Check() {
	path := ""
	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	file, err := modfile.Find(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check: %v\n", err)
		os.Exit(1)
	}

	Vprintf("using %s\n", file.Path)

	if ok, problems := file.CheckValidity(); !ok {
		fmt.Fprintln(os.Stderr, strings.Join(problems, "\n"))
		os.Exit(1)
	}
}
