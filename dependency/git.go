package dependency

import (
	"fmt"

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
