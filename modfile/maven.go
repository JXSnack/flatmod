package modfile

import (
	"fmt"
	"net/url"
	"slices"
)

func applyMavenInstruction(file *File, body string) error {
	u, err := url.Parse(body)
	if err != nil {
		return fmt.Errorf("illegal maven url: %w", err)
	}

	maven := u.String()
	if slices.Contains(file.Mavens, maven) {
		return fmt.Errorf("multiple maven repositories of %q", maven)
	}

	file.Mavens = append(file.Mavens, body)
	return nil
}
