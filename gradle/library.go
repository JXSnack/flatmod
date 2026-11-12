package gradle

import (
	"fmt"
	"path/filepath"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/helper"
)

func DownloadLib(dep dependency.Dependency, ctx dependency.DownloadContext) ([]byte, error) {
	filename := dep.Filename() + ".jar"
	jar, err := cache.Read(
		filepath.Join("libraries", filename),
		helper.ZeroSum(),
		func() ([]byte, error) {
			fmt.Printf("flatmod: downloading %s (%s)\n", dep.Link, dep.Version)
			return dep.Vendor.DownloadJar(dep.Link, dep.Version, ctx)
		},
	)

	if err != nil {
		return jar, err
	}
	return jar, nil
}
