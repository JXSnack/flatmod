package gradle

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/JXSnack/flatmod/follow"
	cp "github.com/otiai10/copy"
)

func moveFiles(w *Workspace, target string) error {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	entries, err := os.ReadDir(wd)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		os.MkdirAll(target, 0700)
		srcPath := filepath.Join(target, "src", "main")
		groupPath := filepath.Join(srcPath, "java", strings.ReplaceAll(w.File.Group, ".", string(filepath.Separator)))

		var path string
		if entry.IsDir() {
			switch entry.Name() {
			case "resources":
				path = srcPath
			default:
				path = groupPath

			}
		} else {
			if strings.HasSuffix(entry.Name(), ".java") {
				path = groupPath
			} else {
				path = target
			}
		}

		// if file exists, do duplicate strategy
		if _, err = os.Stat(filepath.Join(path, entry.Name())); err == nil {
			defined, err := os.ReadFile(filepath.Join(path, entry.Name()))
			if err != nil {
				return fmt.Errorf("cannot read (definition) %s: %w", entry.Name(), err)
			}

			user, err := os.ReadFile(filepath.Join(wd, entry.Name()))
			if err != nil {
				return fmt.Errorf("cannot read %s: %w", entry.Name(), err)
			}

			strategy, ok := w.Follow.Config.DuplicateStrategies[entry.Name()]
			if !ok {
				return fmt.Errorf("no duplicate strategy defined for %s", entry.Name())
			}

			data, err := follow.ApplyDuplicateStrategy(strategy, defined, user)
			if err != nil {
				return err
			}

			err = os.WriteFile(filepath.Join(path, entry.Name()), data, 0755)
			if err != nil {
				return fmt.Errorf("cannot dupl move %s: %w", entry.Name(), err)
			}
		} else {
			err = cp.Copy(filepath.Join(wd, entry.Name()), filepath.Join(path, entry.Name()))
			if err != nil {
				return fmt.Errorf("cannot move %s: %w", entry.Name(), err)
			}
		}
	}
	return nil
}
