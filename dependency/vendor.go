package dependency

type Vendor interface {
	Identifier() string
	DownloadJar(link string, version Version) []byte
}
