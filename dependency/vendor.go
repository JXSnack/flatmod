package dependency

import "github.com/JXSnack/flatmod/version"

type Vendor interface {
	Identifier() string
	DownloadJar(link string, version version.Version) []byte
}
