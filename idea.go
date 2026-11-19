package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/gradle"
	"github.com/JXSnack/flatmod/helper"
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

	// make workspace
	workspace, err := gradle.NewWorkspace(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}

	err = workspace.Generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}

	// download local dependencies
	ctx := dependency.DownloadContext{Facet: cfg.FacetPlatform}
	for _, dep := range workspace.File.Dependencies {
		if !dep.Vendor.ShouldDownload() {
			continue
		}

		_, err = cache.Read(filepath.Join("libraries", dep.Filename(cfg.FacetPlatform)+".jar"),
			helper.ZeroSum(), func() ([]byte, error) {
				fmt.Printf("flatmod: downloading %q (%s)\n", dep.Link, dep.Version.String())
				return dep.Vendor.DownloadJar(dep.Link, dep.Version, ctx)
			},
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "flatmod: failed to download %q: %v", dep.Link, err)
		}
	}

	// get minecraft & jars
	classpath, err := workspace.GetClasspath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}

	// make XML
	os.MkdirAll(filepath.Join(".idea", "libraries"), 0700)

	name := filepath.Base(strings.TrimSuffix(file.Path, ".fmod"))
	mc := idea.MakeMinecraft(name, classpath.Minecraft, classpath.MinecraftLibraries)
	mc.Indent(4)
	err = mc.WriteToFile(filepath.Join(".idea", "libraries", fmt.Sprintf("%s-minecraft.xml", name)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}

	libs := idea.MakeDependencies(name, classpath.Dependencies)
	libs.Indent(4)
	err = libs.WriteToFile(filepath.Join(".idea", "libraries", fmt.Sprintf("%s-dependencies.xml", name)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}

	// make IML
	doc := idea.MakeIML(name, cfg.FacetPlatform, file.Group)
	doc.Indent(2)
	err = doc.WriteToFile(filepath.Join(".idea", "flatmod-project.iml"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}
}
