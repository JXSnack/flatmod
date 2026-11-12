package follow

import (
	"fmt"

	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/maven"
	"github.com/JXSnack/flatmod/modfile"
)

type Config struct {
	DuplicateStrategies map[string]string `json:"duplicateStrategies"`

	// Retain-ed files are the only files not deleted on Gradle workspace generation
	Retain []string `json:"retain"`

	// Indirect dependencies are always loaded but never shown in the configuration
	Indirect Dependencies `json:"indirect"`

	// Init dependencies are used on flatmod project init and never regenerated
	Init Dependencies `json:"init"`
}

type Dependencies struct {
	Maven   []string `json:"maven"`
	Require []string `json:"require"`
}

func ApplyConfigToFile(c Config, file *modfile.File) error {
	for _, indirect := range c.Indirect.Require {
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

	for _, link := range c.Indirect.Maven {
		repo := maven.Repo{Link: link}

		// no duplicates
		for _, installed := range file.Mavens {
			if installed.Link == repo.Link {
				return fmt.Errorf("maven repository %s already defined by follow repository", installed.Link)
			}
		}

		repo.Indirect = true
		file.Mavens = append(file.Mavens, repo)
	}

	return nil
}

func ApplyInitConfigToFile(c Config, file *modfile.File) error {
	for _, init := range c.Init.Require {
		dep, err := dependency.Parse(init)
		if err != nil {
			return err
		}
		file.Dependencies = append(file.Dependencies, dep)
	}

	for _, link := range c.Init.Maven {
		repo := maven.Repo{Link: link}
		file.Mavens = append(file.Mavens, repo)
	}

	return nil
}
