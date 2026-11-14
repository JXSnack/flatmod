package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/gradle"
	"github.com/JXSnack/flatmod/idea"
)

var IdeaCommand = flag.NewFlagSet("idea", flag.ExitOnError)

func Idea() {
	if _, err := os.Stat(".idea"); err != nil {
		fmt.Fprintf(os.Stderr, "idea: no .idea directory found")
		os.Exit(1)
	}

	// command
	path := PathFlag(IdeaCommand)
	IdeaCommand.Parse(os.Args[1:])

	file, cfg := GetMod(*path)
	if file == nil {
		os.Exit(1)
	}

	// download jars
	downloadPath := filepath.Join(".idea", "libraries")
	os.MkdirAll(downloadPath, 0700)

	failed := false
	ctx := dependency.DownloadContext{
		Repos: file.Mavens,
		Remap: cfg.Remap,
	}
	dependencies := make(map[dependency.Dependency][]dependency.Dependency, 0)
	for _, dep := range file.Dependencies {
		jar, err := gradle.DownloadLib(dep, ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "download (%s): %v\n", dep.Vendor.Identifier(), err)
			failed = true
		}
		dependencies[dep] = make([]dependency.Dependency, 0)

		for _, nest := range gradle.ExtractNestedJars(dep.Filename(), jar) {
			dependencies[dep] = append(dependencies[dep], dependency.Dependency{
				Vendor:     dependency.NestedVendor{Parent: dep.Filename()},
				Link:       nest,
				Version:    nil,
				Indirect:   dep.Indirect,
				OmitVendor: false,
			})
		}
	}

	if failed {
		os.Exit(1)
	}

	// generate IML
	doc := idea.MakeIML(
		cfg.FacetPlatform,
		file.Group,
		file.Dependencies,
	)
	doc.Indent(2)
	err := doc.WriteToFile(filepath.Join(".idea", "flatmod-project.iml"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: failed to create IML file: %v\n", err)
		os.Exit(1)
	}

	// generate libraries
	for _, dep := range file.Dependencies {
		doc := idea.MakeLibrary(dep.Filename(), dependencies[dep])
		doc.Indent(2)

		err := doc.WriteToFile(filepath.Join(".idea", "libraries", dep.Filename()+".xml"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
			os.Exit(1)
		}
	}
}
