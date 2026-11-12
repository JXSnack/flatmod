package gradle

import (
	"archive/zip"
	"bytes"
	"io"
	"path/filepath"
	"strings"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/helper"
)

func DownloadNestedJars(parentName string, jar []byte) []string {
	zr, err := zip.NewReader(bytes.NewReader(jar), int64(len(jar)))
	if err != nil {
		return make([]string, 0)
	}

	nested := make([]string, 0)
	for _, file := range zr.File {
		if strings.HasPrefix(file.Name, filepath.Join("META-INF", "jars")) {
			r, err := file.Open()
			if err != nil {
				continue
			}
			defer r.Close()

			nest, err := io.ReadAll(r)
			if err != nil {
				continue
			}
			name := helper.ReValidId.ReplaceAllString(filepath.Base(strings.TrimSuffix(file.Name, ".jar")), "_")
			nested = append(nested, name)

			cache.Write(filepath.Join("libraries", parentName+"-nested", name+".jar"), helper.ZeroSum(), nest)
			for _, recNest := range DownloadNestedJars(parentName, nest) {
				nested = append(nested, recNest)
			}
		}
	}

	return nested
}
