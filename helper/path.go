package helper

import (
	"os"
	"path/filepath"
)

func PathMain() string {
	ud, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(ud, ".flatmod")
}
