package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/follow"
	"github.com/JXSnack/flatmod/helper"
	"github.com/JXSnack/flatmod/maven"
	"github.com/JXSnack/flatmod/modfile"
)

var InitCommand = flag.NewFlagSet("init", flag.ExitOnError)

func Init() {
	group := InitCommand.String("group", "com.example.mod", "the group for the initialized mod")
	merge := InitCommand.Bool("merge", false, "allow initializing an additional project in a directory that already contains one")

	InitCommand.Parse(os.Args[1:])
	args := InitCommand.Args()

	// check if project already exists
	configName := "mc"
	if file, err := modfile.Find(""); file != nil || err.Error() == "multiple .fmod files found" {
		if *merge {
			// merge permission
			ok := helper.InputBoolf("A flatmod project already exists in this directory. Merging will add the new follow's configuration alongside it, and you will be asked before any existing file is overwritten. Continue? (y/n): ")
			if !ok {
				fmt.Println("flatmod: merge cancelled, no changes were made")
				os.Exit(0)
			}

			// config name
			configName = helper.InputStringf("Name for the new config file (e.g. \"fabric\" creates fabric.fmod): ")
			configName = strings.TrimSuffix(configName, ".fmod")

			if _, err := os.Stat(configName + ".fmod"); err == nil {
				fmt.Fprintf(os.Stderr, "flatmod: a config named %q already exists.\n", configName+".fmod")
				os.Exit(1)
			}
		} else {
			fmt.Fprintf(os.Stderr, "flatmod: a project already exists in this directory. To add another project alongside it, run init again with -merge\n")
			os.Exit(1)
		}
	} else {
		// only merge with multiple projects
		if *merge {
			fmt.Fprintln(os.Stderr, "flatmod: there are no other projects in this directory.")
			os.Exit(1)
		}
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
		Path: fmt.Sprintf("%s.fmod", configName),

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

		// don't overwrite files without permission
		if *merge {
			if _, err := os.Stat(target); err == nil {
				overwrite := helper.InputBoolf("The follow wants to overwrite the existing file %q. Overwrite it? (y/n): ", target)
				if !overwrite {
					continue
				}
			}
		}

		// write
		os.MkdirAll(filepath.Dir(target), 0700)
		target = modfile.FillProperties(file, target, cfg.FacetPlatform)
		data = []byte(modfile.FillProperties(file, string(data), cfg.FacetPlatform))
		err = os.WriteFile(target, data, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "flatmod: %v", err)
			continue
		}
	}
}
