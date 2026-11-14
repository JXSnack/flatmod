package remap

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/helper"
)

// trVersion is the version of Tiny Remapper
const trVersion = "0.15.0"

// DownloadTinyRemapper downloads tiny remapper and returns its path.
func DownloadTinyRemapper() (string, error) {
	path := filepath.Join("mappings", fmt.Sprintf("tiny-remapper-%s.jar", trVersion))

	_, err := cache.Read(path, helper.ZeroSum(), func() ([]byte, error) {
		resp, err := http.Get(fmt.Sprintf("https://maven.fabricmc.net/net/fabricmc/tiny-remapper/%s/tiny-remapper-%s-fat.jar", trVersion, trVersion))
		if err != nil {
			return []byte{}, err
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, err
		}
		return data, nil
	})
	return filepath.Join(cache.FileDir, path), err
}
