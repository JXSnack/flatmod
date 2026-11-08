package modfile

import (
	"fmt"

	"github.com/JXSnack/flatmod/dependency"
)

func applyRequireInstruction(file *File, body string) error {
	dep, err := dependency.Parse(body)
	if err != nil {
		return err
	}

	for _, installed := range file.Dependencies {
		if installed.Link == dep.Link {
			return fmt.Errorf("duplicate dependency %q", installed.Link)
		}
	}

	file.Dependencies = append(file.Dependencies, dep)
	return nil
}
