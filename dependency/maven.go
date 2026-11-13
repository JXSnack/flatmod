package dependency

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/JXSnack/flatmod/helper"
	"github.com/JXSnack/flatmod/maven"
	"github.com/JXSnack/flatmod/version"
)

type Maven struct {
}

func (m Maven) Identifier() string {
	return "maven"
}

func (m Maven) DownloadJar(link string, version version.Version, context DownloadContext) ([]byte, error) {
	artifact, err := maven.NewArtifact(fmt.Sprintf("%s:%s", link, version.String()))
	if err != nil {
		return []byte{}, err
	}

	for _, repo := range context.Repos {
		url := artifact.JarURL(repo.Link)
		resp, err := http.Get(url)
		if err != nil {
			return []byte{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, err
		}

		return body, nil
	}
	return make([]byte, 0), fmt.Errorf("maven dependency \"%s %s\" was not found on any repository", link, version.String())
}

func (m Maven) FormatGradle(link string) string {
	return fmt.Sprintf(
		"modImplementation %q",
		fmt.Sprintf("%s:${project.dep_%s}", link, helper.ReValidId.ReplaceAllString(strings.SplitN(link, ":", 2)[1], "_")),
	)
}

func (m Maven) Filename(link string, version version.Version) string {
	return fmt.Sprintf(
		"%s-%s",
		m.Name(link),
		helper.ReValidId.ReplaceAllString(version.String(), "_"),
	)
}

func (m Maven) Name(link string) string {
	return helper.ReValidId.ReplaceAllString(strings.SplitN(link, ":", 2)[1], "_")
}
