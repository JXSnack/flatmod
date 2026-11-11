package modfile

import (
	"fmt"
	"net/url"

	"github.com/JXSnack/flatmod/dependency"
)

func applyMavenInstruction(file *File, body string) error {
	u, err := url.Parse(body)
	if err != nil {
		return fmt.Errorf("invalid maven url: %w", err)
	}

	maven := u.String()
	for _, repo := range file.Mavens {
		if repo.Link == maven {
			return fmt.Errorf("maven repository defined twice: %q", maven)
		}
	}

	file.Mavens = append(file.Mavens, dependency.MavenRepo{Link: maven, Indirect: false})
	return nil
}
