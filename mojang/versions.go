package mojang

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const VersionManifestURL = "https://piston-meta.mojang.com/mc/game/version_manifest_v2.json"

type VersionManifest struct {
	Latest struct {
		Release  string `json:"release"`
		Snapshot string `json:"snapshot"`
	} `json:"latest"`

	Versions []struct {
		Id   string `json:"id"`
		Type string `json:"type"`
		URL  string `json:"url"`
		SHA1 string `json:"sha1"`
	} `json:"versions"`
}

func GetVersion(version string) (Version, error) {
	resp, err := http.Get(VersionManifestURL)
	if err != nil {
		return Version{}, err
	}
	defer resp.Body.Close()

	var manifest VersionManifest
	d := json.NewDecoder(resp.Body)
	err = d.Decode(&manifest)
	if err != nil {
		return Version{}, fmt.Errorf("decoding version manifest: %w", err)
	}

	// find version
	for _, ver := range manifest.Versions {
		if ver.Id == version {
			// get version
			resp, err = http.Get(ver.URL)
			if err != nil {
				return Version{}, fmt.Errorf("get version: %w", err)
			}
			defer resp.Body.Close()

			var actualVer Version
			d = json.NewDecoder(resp.Body)
			err = d.Decode(&actualVer)
			if err != nil {
				return Version{}, fmt.Errorf("decoding version: %w", err)
			}

			actualVer.SHA1 = ver.SHA1
			return actualVer, nil
		}
	}

	return Version{}, fmt.Errorf("version %s does not exist", version)
}
