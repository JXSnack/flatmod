package dependency

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/JXSnack/flatmod/dependency/gitprovider"
	"github.com/JXSnack/flatmod/helper"
	"github.com/JXSnack/flatmod/version"
)

type Git struct{}

func (g Git) Identifier() string {
	return "git"
}

func (g Git) DownloadJar(link string, ver version.Version, ctx DownloadContext) ([]byte, error) {
	var releaseProvider gitprovider.ReleaseProvider
	if strings.HasPrefix(link, "github.com") {
		releaseProvider = gitprovider.GitHubProvider{}
	} else {
		return []byte{}, fmt.Errorf("no git provider defined for %q", link)
	}

	release, err := releaseProvider.GetRelease(link, ver.(version.Semantic))
	if err != nil {
		return []byte{}, err
	}

	asset, err := releaseProvider.GetAsset(release, ctx.Facet)
	if err != nil {
		return []byte{}, err
	}

	// download jar
	resp, err := http.Get(asset.DownloadURL)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	return data, err
}

func (g Git) FormatGradle(link string) string {
	panic("git dependency format is not implemented")
}

func (g Git) Filename(link string, version version.Version, facet string) string {
	return fmt.Sprintf(
		"%s-%s-%s",
		g.Name(link),
		facet,
		helper.ReValidId.ReplaceAllString(version.String(), "_"),
	)
}

func (g Git) Name(link string) string {
	return helper.ReValidId.ReplaceAllString(strings.SplitN(link, "/", 2)[1], "_")
}
