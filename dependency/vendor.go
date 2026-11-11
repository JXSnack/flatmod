package dependency

import (
	"fmt"
	"strings"

	"github.com/JXSnack/flatmod/maven"
	"github.com/JXSnack/flatmod/version"
)

type DownloadContext struct {
	Repos []maven.Repo
}

type Vendor interface {
	Identifier() string
	DownloadJar(link string, version version.Version, context DownloadContext) ([]byte, error)
	FormatGradle(link string) string
	Filename(link string, version version.Version) string
}

// ParseLink takes a dependency and returns the dependency's vendor and link.
//
// Example:
//
//	maven@com.example:dependency
//	(Maven, "com.example:dependency", nil)
func ParseLink(link string) (Vendor, string, error) {
	parts := strings.SplitN(link, "@", 2)
	if len(parts) == 1 {
		parts = append(parts[:1], parts[0:]...)
		parts[0] = ""
	}

	if len(parts) != 2 {
		return nil, "", fmt.Errorf("incorrect dependency format")
	}

	var vendor Vendor

	switch parts[0] {
	case "maven":
		vendor = Maven{}
	case "git":
		vendor = Git{}
	default:
		vendor = Git{}
	}

	return vendor, parts[1], nil
}
