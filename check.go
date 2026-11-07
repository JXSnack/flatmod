package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/JXSnack/flatmod/modfile"
)

// TODO: multiple errors instead of breaking at the first problem

func Check() {
	path := ""
	if len(os.Args) < 3 {
		err := filepath.WalkDir(".", func(candidate string, d fs.DirEntry, err error) error {
			if d != nil && d.IsDir() {
				return nil
			}

			if strings.HasSuffix(candidate, ".fmod") {
				if path != "" {
					return fmt.Errorf("check: multiple .fmod files found")
				}

				path = candidate
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	} else {
		path = os.Args[2]
	}

	if path == "" {
		log.Fatal("check: no .fmod file found in the current directory")
	}

	if !strings.HasSuffix(path, ".fmod") {
		log.Fatal("check: invalid path, must end with .fmod")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Errorf("check: %w", err))
	}

	file, err := modfile.Read(data, path)
	if err != nil {
		log.Fatal(err)
	}

	if ok, problems := file.CheckValidity(); !ok {
		fmt.Println(strings.Join(problems, "\n"))
	}
}
