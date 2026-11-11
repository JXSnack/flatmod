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

func (m Maven) DownloadJar(link string, version version.Version, context DownloadContext) []byte {
	return make([]byte, 0)
}

func (m Maven) FormatGradle(link string) string {
	return fmt.Sprintf(
		"modImplementation %q",
		fmt.Sprintf("%s:${project.dep_%s}", link, helper.ReValidId.ReplaceAllString(strings.SplitN(link, ":", 2)[1], "_")),
	)
}
