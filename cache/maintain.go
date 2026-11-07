package cache

import (
	"fmt"
	"os"
	"path/filepath"
)

var FileDir = ""
var InfoDir = ""

// InitCache ensures all cache-related directories are loaded
func InitCache() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cache: %w", err)
	}

	FileDir = filepath.Join(home, ".flatmod", "cache")
	if err = os.MkdirAll(FileDir, 0700); err != nil {
		return fmt.Errorf("cache: %w", err)
	}

	InfoDir = filepath.Join(home, ".flatmod", "cache-info")
	if err = os.MkdirAll(InfoDir, 0700); err != nil {
		return fmt.Errorf("cache: %w", err)
	}

	return nil
}
