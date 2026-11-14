package remap

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/helper"
)

type Settings struct {
	TinyJar  string `json:"tinyJar"`
	TinyPath string `json:"tinyPath"`
	From     string `json:"from"`
	To       string `json:"to"`
}

func HashMappings(settings Settings) string {
	sum := md5.Sum([]byte(settings.TinyJar))
	return hex.EncodeToString(sum[:])
}

// DownloadMappings downloads the given mappings. Returns the path to the mappings.
func DownloadMappings(settings Settings) (string, error) {
	path := filepath.Join("mappings", HashMappings(settings)+".tiny")

	// download if not in cache
	_, err := cache.Read(path, helper.ZeroSum(), func() ([]byte, error) {
		// download tinyJar
		resp, err := http.Get(settings.TinyJar)
		if err != nil {
			return []byte{}, err
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, err
		}

		// extract mappings
		zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
		if err != nil {
			return []byte{}, err
		}

		for _, file := range zr.File {
			if file.Name == settings.TinyPath {
				rc, err := file.Open()
				if err != nil {
					return []byte{}, fmt.Errorf("opening mapping in zip: %w", err)
				}

				data, err = io.ReadAll(rc)
				if err != nil {
					return []byte{}, fmt.Errorf("reading mapping in zip: %w", err)
				}
				return data, nil
			}
		}

		return []byte{}, fmt.Errorf("mappings do not exist")
	})

	return filepath.Join(cache.FileDir, path), err
}
