package dependency

import (
	"fmt"
	"strings"

	"github.com/JXSnack/flatmod/helper"
	"github.com/JXSnack/flatmod/version"
)

type Git struct{}

func (g Git) Identifier() string {
	return "git"
}

func (g Git) DownloadJar(link string, version version.Version, context DownloadContext) ([]byte, error) {
	return []byte{}, fmt.Errorf("git dependency downloading is not implemented")
}

func (g Git) FormatGradle(link string) string {
	panic("git dependency format is not implemented")
}

func (g Git) Filename(link string, version version.Version) string {
	return fmt.Sprintf(
		"%s-%s",
		helper.ReValidId.ReplaceAllString(strings.SplitN(link, "/", 2)[1], "_"),
		helper.ReValidId.ReplaceAllString(version.String(), "_"),
	)
}
