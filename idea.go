package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/gradle"
	"github.com/JXSnack/flatmod/helper"
)

func Idea() {
	if _, err := os.Stat(".idea"); err != nil {
		fmt.Fprintf(os.Stderr, "idea: no .idea directory found")
		os.Exit(1)
	}

	file := GetMod("idea")
	if file == nil {
		os.Exit(1)
	}

	// generate workspace
	workspace, err := gradle.NewWorkspace(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	if err = workspace.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
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
		filename := dep.Filename()
		jar, err := cache.Read(
			filepath.Join("libraries", filename),
			helper.ZeroSum(),
			func() ([]byte, error) {
				fmt.Printf("flatmod: downloading %s (%s)\n", dep.Link, dep.Version)
				return dep.Vendor.DownloadJar(dep.Link, dep.Version, ctx)
			},
		)

		if err != nil {
			failed = true
			fmt.Fprintf(os.Stderr, "download (%s): %v\n", dep.Vendor.Identifier(), err)
			continue
		}

		err = os.WriteFile(filepath.Join(downloadPath, filename), jar, 0644)
		if err != nil {
			failed = true
			fmt.Fprintf(os.Stderr, "write (%s): %v\n", dep.Vendor.Identifier(), err)
			continue
		}
	}

	if failed {
		os.Exit(1)
	}

}
