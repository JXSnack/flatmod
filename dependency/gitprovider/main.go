package gitprovider

import "github.com/JXSnack/flatmod/version"

type ReleaseProvider interface {
	GetRelease(link string, version version.Semantic) (Release, error)
	GetAsset(release Release, facet string) (Asset, error)
}

type Release struct {
	Tag    string
	Name   string
	Assets []Asset
}

type Asset struct {
	Name        string
	DownloadURL string
}
