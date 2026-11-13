package main

import (
	"flag"
	"os"
)

var CheckCommand = flag.NewFlagSet("check", flag.ExitOnError)

func Check() {
	path := PathFlag(CheckCommand)
	CheckCommand.Parse(os.Args[1:])

	file, _ := GetMod(*path)
	if file == nil {
		os.Exit(1)
	}
}
