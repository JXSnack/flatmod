package main

import (
	"fmt"
	"os"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/follow"
)

func Tidy() {
	file := GetMod("tidy")
	if file == nil {
		os.Exit(1)
	}

	// init cache
	err := cache.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	// load follow
	config, err := follow.Get(file.Follow, file.FollowBranch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "follow: %v", err)
		os.Exit(1)
	}

	err = follow.ApplyConfigToFile(config, file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "apply follow: %v", err)
		os.Exit(1)
	}

	// tidy file
	stat, err := os.Stat(file.Path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "(unexpected) tidy: %v", err)
		os.Exit(1)
	}

	err = os.WriteFile(file.Path, []byte(file.Tidy()), stat.Mode())
	if err != nil {
		fmt.Fprintf(os.Stderr, "tidy: %v", err)
		os.Exit(1)
	}
}
