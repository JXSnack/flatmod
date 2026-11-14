package mojang

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/helper"
)

type Artifact struct {
	Path string `json:"path"`
	SHA1 string `json:"sha1"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}

type Library struct {
	Downloads struct {
		Artifact Artifact `json:"artifact"`
	} `json:"downloads"`
	Name string `json:"name"`
}

// DownloadLibrary downloads the requested library or uses the cache.
func DownloadLibrary(library Library, sum [16]byte) (string, error) {
	name := strings.Split(library.Name, ":")[1]
	path := filepath.Join("libraries", helper.ReValidId.ReplaceAllString(name, "_")+".jar")
	_, err := cache.Read(path, sum, func() ([]byte, error) {
		helper.Vprintf("mojang: downloading %q\n", library.Name)

		// download artifact
		resp, err := http.Get(library.Downloads.Artifact.URL)
		if err != nil {
			return []byte{}, err
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, err
		}

		return data, nil
	})

	return filepath.Join(cache.FileDir, path), err
}
