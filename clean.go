package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/helper"
)

func Clean() {
	// init cache
	if err := cache.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	// clear cache
	if err := cache.Clean(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	// clear .flatmod
	helper.Vprintf("clean: deleting .flatmod directory")
	if err := os.RemoveAll(".flatmod"); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	// clear libraries
	helper.Vprintf("clean: deleting downloaded libraries")
	if err := os.RemoveAll(filepath.Join(helper.PathMain(), "libraries")); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
}
