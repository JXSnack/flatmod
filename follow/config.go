package follow

import (
	"fmt"
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

		// no duplicates
		for _, installed := range file.Dependencies {
			if installed.Link == dep.Link {
				return fmt.Errorf("dependency %s already defined by follow repository", installed.Link)
			}
		}

		dep.Indirect = true
		file.Dependencies = append(file.Dependencies, dep)
	}

	for _, maven := range c.Maven {
		if slices.Contains(file.Mavens, maven) {
			continue
		}

		file.Mavens = append(file.Mavens, maven)
	}

	return nil
}
