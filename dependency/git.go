package dependency

import "github.com/JXSnack/flatmod/version"

type Git struct{}

func (g Git) Identifier() string {
	return "git"
}

func (g Git) DownloadJar(link string, version version.Version) []byte {
	return make([]byte, 0)
}

func (g Git) FormatGradle(link string) string {
	panic("git dependency format is not implemented")
}
