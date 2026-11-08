package gradle

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

		err = cp.Copy(filepath.Join(wd, entry.Name()), filepath.Join(path, entry.Name()))
		if err != nil {
			return fmt.Errorf("cannot move %s: %w", entry.Name(), err)
		}
	}
	return nil
}
