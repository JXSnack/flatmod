package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/follow"
	"github.com/JXSnack/flatmod/helper"
	"github.com/JXSnack/flatmod/modfile"
)

func main() {
	// global flags
	globalFlags := flag.NewFlagSet("fmod", flag.ExitOnError)

	verbose := globalFlags.Bool("v", false, "print more information")
	gitHTTPS := globalFlags.Bool("git-https", false, "use https instead of ssh for git connections")
	noCache := globalFlags.Bool("no-cache", false, "disable cache reading and writing")

	// find command beginning
	args := os.Args[1:]

	commandIndex := -1
	for i, arg := range args {
		if len(arg) > 0 && arg[0] != '-' {
			commandIndex = i
			break
		}
	}

	if commandIndex == -1 {
		fmt.Println("fmod is used for managing Flatmod source code.")
		fmt.Println()
		fmt.Println("Usage: fmod [global args] <command> [command args]")
		fmt.Println()
		fmt.Println("All commands:")
		fmt.Println("\tcheck\tread configuration files for errors")
		fmt.Println("\tclean\tremove cached files")
		fmt.Println("\tgradle\trun gradle command")
		fmt.Println("\tidea\tupdate IntelliJ configuration")
		fmt.Println("\tinit\tcreate a new project")
		fmt.Println("\ttidy\tensure .fmod matches source code")
		return
	}

	// parse global flags
	if err := globalFlags.Parse(args[:commandIndex]); err != nil {
		os.Exit(1)
	}

	helper.Verbose = *verbose
	helper.ForceGitHTTPS = *gitHTTPS
	helper.CacheEnabled = !*noCache

	command := os.Args[commandIndex+1]
	os.Args = append([]string{os.Args[0]}, args[commandIndex+1:]...)

	switch command {
	case "check":
		Check()
	case "clean":
		Clean()
	case "gradle":
		Gradle()
	case "idea":
		Idea()
	case "init":
		Init()
	case "tidy":
		Tidy()
	default:
		fmt.Fprintln(os.Stderr, fmt.Sprintf("%s: unknown command", command))
		os.Exit(1)
	}
}

func GetMod(path string) (*modfile.File, follow.Config) {
	// load .fmod
	file, err := modfile.Find(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return nil, follow.Config{}
	}

	helper.Vprintf("flatmod: using %s\n", file.Path)

	if ok, problems := file.CheckValidity(); !ok {
		fmt.Fprintln(os.Stderr, strings.Join(problems, "\n"))
		return nil, follow.Config{}
	}

	// init cache
	if err := cache.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	// load follow
	config, err := follow.GetConfig(file.Follow, file.FollowBranch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v", err)
		os.Exit(1)
	}

	err = follow.ApplyConfigToFile(config, file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v", err)
		os.Exit(1)
	}

	return file, config
}

func PathFlag(set *flag.FlagSet) *string {
	return set.String("p", "", "the path to the flatmod file in case of multiple files")
}
