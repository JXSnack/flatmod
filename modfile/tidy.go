package modfile

import "strings"

func (f *File) Tidy() string {
	b := strings.Builder{}

	// group
	b.WriteString("group ")
	b.WriteString(f.Group)
	b.WriteString("\n")

	// follow
	b.WriteString("follow ")
	b.WriteString(f.Follow)
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

		for _, dep := range f.Dependencies {
			b.WriteString("require ")
			b.WriteString(dep.Vendor.Identifier())
			b.WriteString("@")
			b.WriteString(dep.Link)
			b.WriteString(" ")
			b.WriteString(dep.Version.String())

			if dep.Indirect {
				b.WriteString(" // indirect\n")
			}
		}
	}

	return b.String()
}
