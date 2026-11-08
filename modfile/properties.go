package modfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/JXSnack/flatmod/helper"
)

func FillProperties(file *File, text string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	text = strings.ReplaceAll(text, "%flatmod:maven%", formatMavens(file, "    "))
	text = strings.ReplaceAll(text, "%flatmod:dependencies%", formatDependencies(file, "    "))
	text = strings.ReplaceAll(text, "%flatmod:root%", filepath.Base(wd))
	text = strings.ReplaceAll(text, "%flatmod:rootid%", helper.ReValidId.ReplaceAllString(filepath.Base(wd), "_"))
	text = strings.ReplaceAll(text, "%flatmod:group%", file.Group)
	return text
}

func formatMavens(file *File, prefix string) string {
	b := strings.Builder{}

	for _, maven := range file.Mavens {
		b.WriteString(prefix)
		b.WriteString("maven { url = ")
		b.WriteString(fmt.Sprintf("%q", maven))
		b.WriteString(" }\n")
	}

	return b.String()
}

func formatDependencies(file *File, prefix string) string {
	b := strings.Builder{}

	for _, dep := range file.Dependencies {
		if dep.Indirect {
			continue
		}

		b.WriteString(prefix)
		b.WriteString(dep.Vendor.FormatGradle(dep.Link))
		b.WriteString("\n")
	}

	return b.String()
}
