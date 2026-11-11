package modfile

import (
	"fmt"
	"net/url"

	"github.com/JXSnack/flatmod/maven"
)

func applyMavenInstruction(file *File, body string) error {
	u, err := url.Parse(body)
	if err != nil {
		return fmt.Errorf("invalid maven url: %w", err)
	}

	url := u.String()
	for _, repo := range file.Mavens {
		if repo.Link == url {
			return fmt.Errorf("maven repository defined twice: %q", url)
		}
	}

	file.Mavens = append(file.Mavens, maven.Repo{Link: url, Indirect: false})
	return nil
}
