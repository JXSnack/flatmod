package dependency

import "github.com/JXSnack/flatmod/version"

type Dependency struct {
	Vendor  Vendor
	Version version.Version
}
