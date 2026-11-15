package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/JXSnack/flatmod/version"
)

type GitHubProvider struct {
}

func (p GitHubProvider) GetRelease(owner string, repo string, version version.Semantic) (Release, error) {
	// download release
	data, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo))
	if err != nil {
		return Release{}, err
	}
	defer data.Body.Close()

	var j []struct {
		Tag    string `json:"tag_name"`
		Name   string `json:"name"`
		Assets []struct {
			Name        string `json:"name"`
			DownloadURL string `json:"browser_download_url"`
		} `json:"assets"`
	}

	d := json.NewDecoder(data.Body)
	err = d.Decode(&j)
	if err != nil {
		return Release{}, err
	}

	// find release
	for _, r := range j {
		// convert to release
		release := Release{
			Tag:    r.Tag,
			Name:   r.Name,
			Assets: make([]Asset, 0),
		}

		for _, asset := range r.Assets {
			release.Assets = append(release.Assets, Asset{
				Name:        asset.Name,
				DownloadURL: asset.DownloadURL,
			})
		}
		return release, nil
	}

	return Release{}, fmt.Errorf("no release with tag %s", version.String())
}

func (p GitHubProvider) DownloadAsset(release Release, facet string) (Asset, error) {
	// filter out only jars
	jars := make([]Asset, 0)
	for _, asset := range release.Assets {
		if strings.HasSuffix(asset.Name, ".jar") {
			jars = append(jars, asset)
		}
	}

	if len(jars) == 0 {
		return Asset{}, fmt.Errorf("release contains no jars")
	}

	// if only one release, return release
	if len(jars) == 1 {
		return jars[0], nil
	}

	// filter out jars for the facet
	facet = strings.ToLower(facet)
	facetJars := make([]Asset, 0)
	for _, asset := range jars {
		if strings.Contains(asset.Name, facet) {
			facetJars = append(facetJars, asset)
		}
	}

	if len(facetJars) == 0 {
		return Asset{}, fmt.Errorf("release contains no compatible jars")
	}

	if len(facetJars) == 1 {
		return facetJars[0], nil
	}
	return Asset{}, fmt.Errorf("no compatible jars found")
}
