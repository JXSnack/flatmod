package modfile

import (
	"strings"

	"github.com/JXSnack/flatmod/dependency"
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
			b.WriteString("maven ")
			b.WriteString(maven)
			b.WriteString("\n")
		}
	}

	// dependencies
	if len(f.Dependencies) > 0 {
		b.WriteString("\n")

		indirects := make([]dependency.Dependency, 0)
		for _, dep := range f.Dependencies {
			if dep.Indirect {
				indirects = append(indirects, dep)
				continue
			}

			b.WriteString("require ")
			b.WriteString(dep.String())
			b.WriteString("\n")
		}

		if len(indirects) > 0 {
			b.WriteString("\n")
			for _, dep := range indirects {
				b.WriteString("require ")
				b.WriteString(dep.String())
				b.WriteString("\n")
			}
		}
	}

	return b.String()
}
