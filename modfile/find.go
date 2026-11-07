package modfile

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Find(path string) (*File, error) {
	if path == "" {
		path = "."
	}

	stat, err := os.Stat(path)
	if err != nil {
		if !strings.HasSuffix(path, ".fmod") {
			path += ".fmod"
			return Find(path)
		}

		return nil, fmt.Errorf("no file at %s", path)
	}

	if stat.IsDir() {
		err = filepath.WalkDir(path, func(candidate string, d fs.DirEntry, err error) error {
			if d != nil && d.IsDir() {
				return nil
			}

			if !strings.HasSuffix(candidate, ".fmod") {
				return nil
			}

			if strings.HasSuffix(path, ".fmod") {
				return fmt.Errorf("multiple .fmod files found")
			}
			path = candidate
			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return Read(data, path)
}
