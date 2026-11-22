package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/JXSnack/flatmod/updater"
)

var VersionCommand = flag.NewFlagSet("version", flag.ExitOnError)

func Version() {
	VersionCommand.Parse(os.Args[1:])

	fmt.Printf("flatmod version %s (patch %d) %s/%s\n", updater.Display, updater.Current, runtime.GOOS, runtime.GOARCH)
}
