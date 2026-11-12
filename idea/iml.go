package idea

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func GetIML() string {
	iml := ""
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if strings.HasPrefix(path, ".") && !strings.HasPrefix(path, ".idea") {
			return nil
		}

		if strings.HasSuffix(path, ".iml") {
			iml = path
		}
		return nil
	})

	return iml
}
