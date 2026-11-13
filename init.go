package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/follow"
	"github.com/JXSnack/flatmod/maven"
	"github.com/JXSnack/flatmod/modfile"
)

var InitCommand = flag.NewFlagSet("init", flag.ExitOnError)

func Init() {
	group := InitCommand.String("group", "com.example.mod", "the group for the initialized mod")

	InitCommand.Parse(os.Args[1:])
	args := InitCommand.Args()

	if file, _ := modfile.Find(""); file != nil {
		fmt.Fprintf(os.Stderr, "flatmod: there already exists a project in this directory\n")
		os.Exit(1)
	}

	// init cache
	if err := cache.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	// ensure follow exists
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "flatmod: require follow repository and branch\n")
		os.Exit(1)
	}

	cfg, err := follow.GetConfig(args[0], args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}

	// make base fmod file
	file := &modfile.File{
		Path: "mc.fmod",

		Group:        *group,
		Follow:       args[0],
		FollowBranch: args[1],

		Dependencies: make([]dependency.Dependency, 0),
		Mavens:       make([]maven.Repo, 0),
	}

	if err = follow.ApplyConfigToFile(cfg, file); err != nil {
		fmt.Fprintf(os.Stderr, "init: %v\n", err)
		os.Exit(1)
	}

	if err = follow.ApplyInitConfigToFile(cfg, file); err != nil {
		fmt.Fprintf(os.Stderr, "init: %v\n", err)
		os.Exit(1)
	}

	// write fmod file
	err = os.WriteFile(file.Path, []byte(file.Tidy()), 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "init: %v\n", err)
		os.Exit(1)
	}

	// make resource directory
	os.Mkdir("resources", 0700)

	// extract
	for cp, target := range cfg.Extract {
		data, err := follow.GetFile(file.Follow, file.FollowBranch, cp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "flatmod: %v", err)
			continue
		}

		target = modfile.FillProperties(file, target)
		data = []byte(modfile.FillProperties(file, string(data)))
		err = os.WriteFile(target, data, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "flatmod: %v", err)
			continue
		}
	}
}
