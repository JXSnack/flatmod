package cache

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JXSnack/flatmod/helper"
)

var FileDir = ""
var InfoDir = ""

// Init ensures all cache-related directories are loaded
func Init() error {
	FileDir = filepath.Join(helper.PathMain(), "cache")
	if err := os.MkdirAll(FileDir, 0700); err != nil {
		return fmt.Errorf("cache: %w", err)
	}

	InfoDir = filepath.Join(helper.PathMain(), "cache-info")
	if err := os.MkdirAll(InfoDir, 0700); err != nil {
		return fmt.Errorf("cache: %w", err)
	}

	return nil
}

func Clean() error {
	// ignore if cache isn't initialized
	if InfoDir == "" || FileDir == "" {
		return fmt.Errorf("cache isn't initialized, cannot clean")
	}

	// clean cache
	helper.Vprintf("cache: removing %s\n", FileDir)
	os.RemoveAll(FileDir)

	helper.Vprintf("cache: removing %s\n", InfoDir)
	os.RemoveAll(InfoDir)

	// restart
	helper.Vprintf("cache: restarting cache\n")
	err := Init()
	if err != nil {
		return fmt.Errorf("restart %w", err)
	}
	return nil
}
