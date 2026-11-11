package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/follow"
	"github.com/JXSnack/flatmod/helper"
	"github.com/JXSnack/flatmod/modfile"
)

func main() {
	for i, arg := range os.Args {
		if strings.HasPrefix(arg, "-") {
			helper.Opts = append(helper.Opts, strings.TrimPrefix(arg, "-"))
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
		fmt.Println("\tclean\tremove cached files")
		fmt.Println("\tidea\tupdate IntelliJ configuration")
		fmt.Println("\ttidy\tensure .fmod matches source code")
		return
	}

	switch os.Args[1] {
	case "check":
		Check()
	case "clean":
		Clean()
	case "idea":
		Idea()
	case "tidy":
		Tidy()
	default:
		fmt.Fprintln(os.Stderr, fmt.Sprintf("%s: unknown command", os.Args[1]))
		os.Exit(1)
	}
}

func GetMod(command string) *modfile.File {
	// load .fmod
	path := ""
	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	file, err := modfile.Find(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return nil
	}

	helper.Vprintf("flatmod: using %s\n", file.Path)

	if ok, problems := file.CheckValidity(); !ok {
		fmt.Fprintln(os.Stderr, strings.Join(problems, "\n"))
		return nil
	}

	// init cache
	if err := cache.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	// load follow
	config, err := follow.GetConfig(file.Follow, file.FollowBranch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v", command, err)
		os.Exit(1)
	}

	err = follow.ApplyConfigToFile(config, file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v", command, err)
		os.Exit(1)
	}

	return file
}
