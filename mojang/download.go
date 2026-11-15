package mojang

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/remap"
)

func DownloadHead(url string, version string, sum [16]byte, name string, settings remap.Settings) error {
	// download jar
	path := filepath.Join("minecraft", version, name)
	_, err := cache.Read(path+".jar", sum, func() ([]byte, error) {
		fmt.Printf("flatmod: downloading %s (%s)\n", name, version)
		resp, err := http.Get(url)
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
	if err != nil {
		return err
	}

	// remap
	if settings.TinyJar != "" {
		settings.From = "official"

		err = remap.JarCached(filepath.Join(cache.FileDir, path+".jar"), path+"-remapped.jar", settings)
		if err != nil {
			return fmt.Errorf("remapping: %s", err)
		}
	} else {
		return fmt.Errorf("skipping remap because remap is missing key tinyJar")
	}

	return nil
}
