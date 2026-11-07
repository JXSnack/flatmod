package dependency

import (
	"fmt"
	"strings"

	"github.com/JXSnack/flatmod/version"
)

type Dependency struct {
	Vendor   Vendor
	Link     string
	Version  version.Version
	Indirect bool
}

func Parse(body string) (Dependency, error) {
	split := strings.SplitN(body, " ", -1)
	if len(split) != 2 {
		return Dependency{}, fmt.Errorf("bad dependency format. expected \"require <dependency> <version>\"")
	}

	vendor, link, err := ParseLink(split[0])
	if err != nil {
		return Dependency{}, err
	}

	ver := version.Parse(split[1])

	return Dependency{
		Vendor:   vendor,
		Link:     link,
		Version:  ver,
		Indirect: false,
	}, nil
}
