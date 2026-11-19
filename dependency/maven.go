package dependency

import (
	"fmt"
	"strings"

	"github.com/JXSnack/flatmod/helper"
	"github.com/JXSnack/flatmod/version"
)

type Maven struct {
}

func (m Maven) Identifier() string {
	return "maven"
}

func (m Maven) DownloadJar(link string, version version.Version, context DownloadContext) ([]byte, error) {
	panic("maven vendor should never download")
}

func (m Maven) FormatGradle(link string, facet string) string {
	return fmt.Sprintf("%q", fmt.Sprintf("%s:${project.dep_%s}", link, m.Name(link)))
}

func (m Maven) Filename(link string, version version.Version, facet string) string {
	panic("should not create filename for maven vendor")
}

func (m Maven) Name(link string) string {
	return helper.ReValidId.ReplaceAllString(strings.SplitN(link, ":", 2)[1], "_")
}

func (m Maven) ShouldDownload() bool {
	return false
}
