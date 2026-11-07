package modfile

import (
	"github.com/JXSnack/flatmod/dependency"
)

func applyRequireInstruction(file *File, body string) error {
	dep, err := dependency.Parse(body)
	if err != nil {
		return err
	}

	file.Dependencies = append(file.Dependencies, dep)
	return nil
}
