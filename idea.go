package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/gradle"
	"github.com/JXSnack/flatmod/idea"
)

func Idea() {
	if _, err := os.Stat(".idea"); err != nil {
		fmt.Fprintf(os.Stderr, "idea: no .idea directory found")
		os.Exit(1)
	}

	file, cfg := GetMod("idea")
	if file == nil {
		os.Exit(1)
	}

	// download jars
	downloadPath := filepath.Join(".idea", "libraries")
	os.MkdirAll(downloadPath, 0700)

	failed := false
	ctx := dependency.DownloadContext{
		Repos: file.Mavens,
	}
	for _, dep := range file.Dependencies {
		err := gradle.DownloadLib(dep, ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "download (%s): %v\n", dep.Vendor.Identifier(), err)
			failed = true
		}
	}

	if failed {
		os.Exit(1)
	}

	fmt.Println(idea.GetIML())
}
