package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/gradle"
	"github.com/JXSnack/flatmod/idea"
	"github.com/JXSnack/flatmod/mojang"
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

	// download Minecraft
	downloadPath := filepath.Join(".idea", "libraries")
	os.MkdirAll(downloadPath, 0700)

	mcVersion, err := mojang.GetVersion(cfg.Minecraft)
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}

	// download Minecraft client
	err = mojang.DownloadHead(
		mcVersion.Downloads.Client.URL, mcVersion.Id,
		md5.Sum([]byte(mcVersion.Downloads.Client.SHA1)),
		"client", cfg.Remap,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: failed to download client: %v\n", err)
		os.Exit(1)
	}

	// download Minecraft server
	err = mojang.DownloadHead(
		mcVersion.Downloads.Server.URL, mcVersion.Id,
		md5.Sum([]byte(mcVersion.Downloads.Server.SHA1)),
		"server", cfg.Remap,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: failed to download server: %v\n", err)
		os.Exit(1)
	}

	// download Minecraft libraries
	mcLibs := make([]string, 0)
	for _, lib := range mcVersion.Libraries {
		sum := md5.Sum([]byte(mcVersion.SHA1))
		libPath, err := mojang.DownloadLibrary(lib, sum)
		if err != nil {
			fmt.Fprintf(os.Stderr, "flatmod: failed downloading %q: %v\n", lib.Name, err)
			os.Exit(1)
		}

		mcLibs = append(mcLibs, filepath.Base(libPath))
	}

	// download jars
	failed := false
	ctx := dependency.DownloadContext{
		Repos: file.Mavens,
		Remap: cfg.Remap,
		Facet: cfg.FacetPlatform,
	}
	dependencies := make(map[dependency.Dependency][]dependency.Dependency, 0)
	for _, dep := range file.Dependencies {
		jar, err := gradle.DownloadLib(dep, ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "flatmod: failed to download (%s): %v\n", dep.Vendor.Identifier(), err)
			failed = true
		}
		dependencies[dep] = make([]dependency.Dependency, 0)

		// extract nested jars
		nests, err := gradle.ExtractNestedJars(dep.Filename(cfg.FacetPlatform), jar, ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "flatmod: failed to extract %q: %v\n", dep.Link, err)
			failed = true
		}

		for _, nest := range nests {
			dependencies[dep] = append(dependencies[dep], dependency.Dependency{
				Vendor:     dependency.NestedVendor{Parent: dep.Filename(cfg.FacetPlatform)},
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
	err = doc.WriteToFile(filepath.Join(".idea", "flatmod-project.iml"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: failed to create IML file: %v\n", err)
		os.Exit(1)
	}

	// make jar suffixes
	var suffix string
	if cfg.Remap.Enabled {
		suffix = "-remapped"
	} else {
		suffix = ""
	}

	// minecraft dependency
	mc := idea.MakeMinecraft(mcVersion.Id, mcLibs, "-remapped")
	mc.Indent(2)

	err = mc.WriteToFile(filepath.Join(".idea", "libraries", fmt.Sprintf("minecraft.xml")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}

	// generate libraries
	for _, dep := range file.Dependencies {
		doc := idea.MakeLibrary(dep.Filename(cfg.FacetPlatform), dependencies[dep], cfg.FacetPlatform, suffix)
		doc.Indent(2)

		err := doc.WriteToFile(filepath.Join(".idea", "libraries", dep.Filename(cfg.FacetPlatform)+".xml"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
			os.Exit(1)
		}
	}
}
