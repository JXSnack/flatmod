package gradle

import (
	"fmt"
	"path/filepath"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/helper"
	"github.com/JXSnack/flatmod/remap"
)

// DownloadLib downloads a dependency and remaps it to the given mappings
func DownloadLib(dep dependency.Dependency, ctx dependency.DownloadContext) ([]byte, error) {
	// download jar
	filename := dep.Filename()
	jar, err := cache.Read(
		filepath.Join("libraries", filename+".jar"),
		helper.ZeroSum(),
		func() ([]byte, error) {
			fmt.Printf("flatmod: downloading %s (%s)\n", dep.Link, dep.Version)
			return dep.Vendor.DownloadJar(dep.Link, dep.Version, ctx)
		},
	)

	if err != nil {
		return jar, err
	}

	// remap jar
	if ctx.Remap.TinyPath != "" {
		if err = remap.JarCached(
			filepath.Join(cache.FileDir, "libraries", filename+".jar"),
			filepath.Join("libraries", filename+"-remapped.jar"),
			ctx.Remap,
		); err != nil {
			return jar, fmt.Errorf("remap: %w", err)
		}
	}
	return jar, nil
}
