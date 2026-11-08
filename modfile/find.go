package modfile

import (
	"fmt"
	"os"
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
		entries, err := os.ReadDir(path)
		for _, entry := range entries {
			if entry != nil && entry.IsDir() {
				continue
			}

			if !strings.HasSuffix(entry.Name(), ".fmod") {
				continue
			}

			if strings.HasSuffix(path, ".fmod") {
				return nil, fmt.Errorf("multiple .fmod files found")
			}
			path = entry.Name()
		}

		if !strings.HasSuffix(path, ".fmod") {
			err = fmt.Errorf("no .fmod file in current directory")
		}

		if err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	file, errs := Read(data, path)
	if len(errs) == 0 {
		return file, nil
	}

	bigErr := make([]string, 0)
	for _, err = range errs {
		bigErr = append(bigErr, err.Error())
	}

	return file, fmt.Errorf(strings.Join(bigErr, "\n"))
}
