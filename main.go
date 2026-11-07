package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
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
	if slices.Contains(os.Args, "-v") {
		fmt.Println(a...)
	}
}

func Vprintf(format string, a ...any) {
	if slices.Contains(os.Args, "-v") {
		fmt.Printf(format, a...)
	}
}
