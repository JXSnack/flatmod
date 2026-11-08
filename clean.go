package main

import (
	"fmt"
	"os"

	"github.com/JXSnack/flatmod/cache"
)

func Clean() {
	// init cache
	if err := cache.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	// clear cache
	if err := cache.Clean(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
