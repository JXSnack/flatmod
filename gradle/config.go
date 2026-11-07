package gradle

import (
	"fmt"
	"net/url"
	"slices"

	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/modfile"
)

type Config struct {
	DuplicateStrategies map[string]string `json:"duplicateStrategies"`

	// Indirect dependencies are dependencies that must be loaded
	Indirect []string `json:"indirect"`

	// Require-d dependencies exist for initial project creation, but
	// are not enforced
	Require []string `json:"require"`

	// Maven contains a list of repositories for initial project creation,
	// but are not enforced
	Maven []string `json:"maven"`
}

func ApplyConfigToFile(c Config, file *modfile.File) error {
	for _, indirect := range c.Indirect {
		dep, err := dependency.Parse(indirect)
		if err != nil {
			return err
		}

		dep.Indirect = true
		file.Dependencies = append(file.Dependencies, dep)
	}

	for _, maven := range c.Maven {
		u, err := url.Parse(maven)
		if err != nil {
			return fmt.Errorf("illegal maven url: %w", err)
		}

		if slices.Contains(file.Mavens, *u) {
			continue
		}

		file.Mavens = append(file.Mavens, *u)
	}

	return nil
}
