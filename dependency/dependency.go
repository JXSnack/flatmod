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

	// OmitVendor specifies whether the vendor was omitted in the modfile.
	OmitVendor bool
}

func (d Dependency) String() string {
	b := strings.Builder{}

	if !d.OmitVendor {
		b.WriteString(d.Vendor.Identifier())
		b.WriteString("@")
	}

	b.WriteString(d.Link)
	b.WriteString(" ")
	b.WriteString(d.Version.String())

	if d.Indirect {
		b.WriteString(" // indirect")
	}

	return b.String()
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

	if vendor.Identifier() == (Git{}).Identifier() {
		if !version.IsSemantic(split[1]) {
			return Dependency{}, fmt.Errorf("git dependencies must follow semantic versioning")
		}
	}
	ver := version.Parse(split[1])

	return Dependency{
		Vendor:   vendor,
		Link:     link,
		Version:  ver,
		Indirect: false,

		OmitVendor: !strings.Contains(split[0], "@"),
	}, nil
}
