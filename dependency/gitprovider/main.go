package main

import "github.com/JXSnack/flatmod/version"

type ReleaseProvider interface {
	GetRelease(owner string, repo string, version version.Semantic) (Release, error)
	DownloadAsset(release Release, facet string) (Asset, error)
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
