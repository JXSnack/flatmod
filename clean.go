package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/helper"
)

var CleanCommand = flag.NewFlagSet("clean", flag.ExitOnError)

func Clean() {
	CleanCommand.Parse(os.Args[1:])

	// init cache
	if err := cache.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	// clear cache
	if err := cache.Clean(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	// clear .flatmod
	helper.Vprintf("clean: deleting .flatmod directory")
	if err := os.RemoveAll(".flatmod"); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
}
