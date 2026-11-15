package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

	// get minecraft & jars
	classpaths, err := workspace.GetClasspath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "flatmod: %v\n", err)
		os.Exit(1)
	}

	// make XML
	name := filepath.Base(strings.TrimSuffix(file.Path, ".fmod"))
	mc := idea.MakeMinecraft(name, classpaths.Minecraft, classpaths.MinecraftLibraries)
	mc.Indent(4)
	mc.WriteToFile(filepath.Join(".idea", "libraries", fmt.Sprintf("%s-minecraft.xml", name)))

	libs := idea.MakeDependencies(name, classpaths.Dependencies)
	libs.Indent(4)
	libs.WriteToFile(filepath.Join(".idea", "libraries", fmt.Sprintf("%s-dependencies.xml", name)))

	// make IML
	doc := idea.MakeIML(name, cfg.FacetPlatform, file.Group)
	doc.Indent(2)
	doc.WriteToFile(filepath.Join(".idea", "flatmod-project.iml"))
}
