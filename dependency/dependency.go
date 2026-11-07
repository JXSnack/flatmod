package dependency

import "github.com/JXSnack/jmod/version"

type Dependency struct {
	Vendor  Vendor
	Version version.Version
}
