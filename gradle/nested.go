package gradle

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/dependency"
	"github.com/JXSnack/flatmod/helper"
	"github.com/JXSnack/flatmod/remap"
)

func ExtractNestedJars(parentName string, jar []byte, ctx dependency.DownloadContext) ([]string, error) {
	// read jar
	zr, err := zip.NewReader(bytes.NewReader(jar), int64(len(jar)))
	if err != nil {
		return make([]string, 0), nil
	}

	nested := make([]string, 0)
	for _, file := range zr.File {
		// look for jars
		if strings.HasPrefix(file.Name, filepath.Join("META-INF", "jars")) && strings.HasSuffix(file.Name, ".jar") {
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

			// write found jar
			nestedPath := filepath.Join("libraries", parentName+"-nested", name)
			cache.Write(nestedPath+".jar", helper.ZeroSum(), nest)

			// remap
			if ctx.Remap.TinyJar != "" {
				if err = remap.JarCached(
					filepath.Join(cache.FileDir, nestedPath+".jar"),
					filepath.Join(nestedPath+"-remapped.jar"),
					ctx.Remap,
				); err != nil {
					return nested, fmt.Errorf("remap: %w", err)
				}
			}

			// recursively check downloaded jars for other jars
			recNests, err := ExtractNestedJars(parentName, nest, ctx)
			if err != nil {
				return nested, err
			}

			for _, recNest := range recNests {
				nested = append(nested, recNest)
			}
		}
	}

	return nested, nil
}
