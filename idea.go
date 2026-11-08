package main

import (
	"fmt"
	"os"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/gradle"
)

func Idea() {
	file := GetMod("idea")
	if file == nil {
		os.Exit(1)
	}

	// init cache
	if err := cache.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	workspace, err := gradle.NewWorkspace(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	if err = workspace.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}
