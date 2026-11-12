package gradle

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/helper"
)

func DownloadLib(dep dependency.Dependency, ctx dependency.DownloadContext) error {
	path := filepath.Join(helper.PathMain(), "libraries")
	os.MkdirAll(path, 0700)

	filename := dep.Filename()
	jar, err := cache.Read(
		filepath.Join("libraries", filename),
		helper.ZeroSum(),
		func() ([]byte, error) {
			fmt.Printf("flatmod: downloading %s (%s)\n", dep.Link, dep.Version)
			return dep.Vendor.DownloadJar(dep.Link, dep.Version, ctx)
		},
	)

	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(path, filename), jar, 0644)
	if err != nil {
		return err
	}
	return nil
}
