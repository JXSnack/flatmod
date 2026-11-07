package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/JXSnack/flatmod/modfile"
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
		fmt.Println("\tidea\tupdate IntelliJ configuration")
		fmt.Println("\ttidy\tensure .fmod matches source code")
		return
	}

	switch os.Args[1] {
	case "check":
		Check()
	case "idea":
		Idea()
	case "tidy":
		Tidy()
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

func GetMod(command string) *modfile.File {
	path := ""
	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	file, err := modfile.Find(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v", command, err)
		return nil
	}

	Vprintf("using %s\n", file.Path)

	if ok, problems := file.CheckValidity(); !ok {
		fmt.Fprintln(os.Stderr, strings.Join(problems, "\n"))
		return nil
	}

	return file
}
