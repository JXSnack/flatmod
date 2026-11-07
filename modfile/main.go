package modfile

import (
	"fmt"
	"strings"

	"github.com/JXSnack/jmod/dependency"
)

// TODO: multiple src directories
// TODO: multiple resource directories

type File struct {
	Group        string
	Follow       string
	Dependencies []dependency.Dependency
}

func Read(data []byte, path string) (*File, error) {
	file := &File{}
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		instruction := strings.SplitN(line, " ", 2)
		if len(instruction) != 2 {
			return nil, wrapErr(path, i+1, fmt.Errorf("unfinished instruction"))
		}

		body := strings.SplitN(instruction[1], "//", 2)[0]
		if strings.TrimSpace(body) == "" {
			continue
		}

		switch instruction[0] {
		case "group":
			err := applyGroupInstruction(file, body)
			if err != nil {
				return nil, wrapErr(path, i+1, err)
			}
		case "follow":
			err := applyFollowInstruction(file, body)
			if err != nil {
				return nil, wrapErr(path, i+1, err)
			}
		case "require":
			err := applyRequireInstruction(file, body)
			if err != nil {
				return nil, wrapErr(path, i+1, err)
			}
		}
	}

	return file, nil
}

func wrapErr(path string, line int, err error) error {
	return fmt.Errorf("%s:%d: %w", line, err)
}
