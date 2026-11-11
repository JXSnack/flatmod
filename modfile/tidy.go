package modfile

import (
	"strings"
)

func (f *File) Tidy() string {
	b := strings.Builder{}

	// group
	b.WriteString("group ")
	b.WriteString(f.Group)
	b.WriteString("\n")

	// follow
	b.WriteString("follow ")
	b.WriteString(f.Follow)
	b.WriteString(" ")
	b.WriteString(f.FollowBranch)
	b.WriteString("\n")

	// mavens
	if len(f.Mavens) > 0 {
		b.WriteString("\n")

		for _, maven := range f.Mavens {
			if maven.Indirect {
				continue
			}

			b.WriteString("maven ")
			b.WriteString(maven.Link)
			b.WriteString("\n")
		}
	}

	// dependencies
	if len(f.Dependencies) > 0 {
		b.WriteString("\n")

		for _, dep := range f.Dependencies {
			if dep.Indirect {
				continue
			}

			b.WriteString("require ")
			b.WriteString(dep.FormatFlatmod())
			b.WriteString("\n")
		}
	}

	return b.String()
}
