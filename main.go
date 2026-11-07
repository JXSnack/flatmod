package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var Opts []string

func main() {
	for i, arg := range os.Args {
		if strings.HasPrefix(arg, "-") {
			Opts = append(Opts, strings.TrimPrefix(arg, "-"))
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
		}
	}

	if len(os.Args) == 1 {
		fmt.Println("fmod is used for managing Flatmod source code.")
		fmt.Println()
		fmt.Println("Usage: fmod <command> [args]")
		fmt.Println()
		fmt.Println("All commands:")
		fmt.Println("\tcheck\tread configuration files for errors")
		return
	}

	switch os.Args[1] {
	case "check":
		Check()
	case "idea":
		Idea()
	default:
		fmt.Fprintln(os.Stderr, fmt.Sprintf("%s: unknown command", os.Args[1]))
		os.Exit(1)
	}
}

func Vprintln(a ...any) {
	if slices.Contains(Opts, "v") || slices.Contains(Opts, "-verbose") {
		fmt.Println(a...)
	}
}

func Vprintf(format string, a ...any) {
	if slices.Contains(Opts, "v") || slices.Contains(Opts, "-verbose") {
		fmt.Printf(format, a...)
	}
}
