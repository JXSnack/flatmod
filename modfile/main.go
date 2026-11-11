package modfile

import (
	"fmt"
	"strings"

	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/maven"
)

// TODO: multiple resource directories

type File struct {
	Path string

	Group        string
	Follow       string
	FollowBranch string
	Dependencies []dependency.Dependency
	Mavens       []maven.Repo
}

func Read(data []byte, path string) (*File, []error) {
	file := &File{Path: path}
	lines := strings.Split(string(data), "\n")
	errs := make([]error, 0)

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//") {
			continue
		}

		instruction := strings.SplitN(line, " ", 2)
		name := instruction[0]
		if name == "" {
			continue
		}

		if len(instruction) == 1 {
			errs = append(errs, wrapErr(path, i+1, fmt.Errorf("unfinished instruction")))
			continue
		}

		body := strings.SplitN(instruction[1], " //", 2)[0]
		body = strings.TrimSpace(body)
		if body == "" {
			continue
		}

		var err error
		switch name {
		case "group":
			err = applyGroupInstruction(file, body)
		case "follow":
			err = applyFollowInstruction(file, body)
		case "maven":
			err = applyMavenInstruction(file, body)
		case "require":
			err = applyRequireInstruction(file, body)
		default:
			errs = append(errs, fmt.Errorf("unknown statement: %s", name))
			continue
		}

		if err != nil {
			errs = append(errs, wrapErr(path, i+1, err))
			continue
		}
	}

	return file, errs
}

func wrapErr(path string, line int, err error) error {
	return fmt.Errorf("%s:%d: %w", path, line, err)
}

func (f *File) CheckValidity() (bool, []string) {
	problems := make([]string, 0)

	if f.Group == "" {
		problems = append(problems, fmt.Sprintf("%s: missing group", f.Path))
	}

	if f.Follow == "" {
		problems = append(problems, fmt.Sprintf("%s: missing follow", f.Path))
	}

	return len(problems) == 0, problems
}
