package dependency

import (
	"path/filepath"

	"github.com/JXSnack/flatmod/version"
)

type NestedVendor struct {
	Parent string
}

func (v NestedVendor) Identifier() string {
	return "nested"
}

func (v NestedVendor) Filename(link string, version version.Version) string {
	return filepath.Join(v.Parent+"-nested", link)
}

func (v NestedVendor) FormatGradle(link string) string {
	panic("nested vendor should never be formatted")
}

func (v NestedVendor) DownloadJar(link string, version version.Version, context DownloadContext) ([]byte, error) {
	panic("nested vendor should never download")
}
