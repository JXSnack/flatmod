package maven

import (
	"fmt"
	"strings"
)

type Artifact struct {
	Group      string
	Artifact   string
	Version    string
	Classifier string
}

func NewArtifact(coord string) (Artifact, error) {
	parts := strings.Split(coord, ":")
	if len(parts) != 3 && len(parts) != 4 {
		return Artifact{}, fmt.Errorf("invalid maven coordinate: %q", coord)
	}
	classifier := ""
	if len(parts) == 4 {
		classifier = parts[3]
	}

	return Artifact{
		Group:      parts[0],
		Artifact:   parts[1],
		Version:    parts[2],
		Classifier: classifier,
	}, nil
}

func (a Artifact) JarURL(repo string) string {
	groupPath := strings.ReplaceAll(a.Group, ".", "/")
	return fmt.Sprintf("%s/%s/%s/%s/%s",
		strings.TrimSuffix(repo, "/"),
		groupPath, a.Artifact, a.Version, a.Filename())
}

func (a Artifact) Filename() string {
	if a.Classifier == "" {
		return fmt.Sprintf("%s-%s.jar", a.Artifact, a.Version)
	}
	return fmt.Sprintf("%s-%s-%s.jar", a.Artifact, a.Version, a.Classifier)
}
