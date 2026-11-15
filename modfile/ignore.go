package modfile

import (
	"fmt"
	"path/filepath"
)

func applyIgnoreInstruction(file *File, body string) error {
	if body == "" {
		return fmt.Errorf("cannot ignore empty path")
	}

	path := filepath.Clean(body)
	file.IgnoredPaths = append(file.IgnoredPaths, path)

	return nil
}
