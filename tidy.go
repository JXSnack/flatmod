package main

import (
	"fmt"
	"os"
)

func Tidy() {
	file, _ := GetMod("tidy")
	if file == nil {
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
