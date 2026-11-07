package dependency

import "github.com/JXSnack/flatmod/version"

type Maven struct {
}

func (m Maven) Identifier() string {
	return "maven"
}

func (m Maven) DownloadJar(link string, version version.Version) []byte {
	return make([]byte, 0)
}
