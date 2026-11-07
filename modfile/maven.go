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

	if slices.Contains(file.Mavens, *u) {
		return fmt.Errorf("multiple maven repositories of %q", u.String())
	}

	file.Mavens = append(file.Mavens, *u)
	return nil
}
