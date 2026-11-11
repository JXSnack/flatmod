package maven

import (
	"fmt"
	"strings"
)

type Artifact struct {
	Group    string
	Artifact string
	Version  string
}

func NewArtifact(coord string) (Artifact, error) {
	parts := strings.SplitN(coord, ":", 3)
	if len(parts) != 3 {
		return Artifact{}, fmt.Errorf("invalid maven coordinate: %q", coord)
	}
	return Artifact{
		Group:    parts[0],
		Artifact: parts[1],
		Version:  parts[2],
	}, nil
}

func (a Artifact) JarURL(repo string) string {
	groupPath := strings.ReplaceAll(a.Group, ".", "/")
	return fmt.Sprintf("%s/%s/%s/%s/%s",
		strings.TrimSuffix(repo, "/"),
		groupPath, a.Artifact, a.Version, a.Filename())
}

func (a Artifact) Filename() string {
	return fmt.Sprintf("%s-%s.jar", a.Artifact, a.Version)
}
