package version

type Custom struct {
	V string
}

func (c Custom) Kind() string {
	return "custom"
}

func (c Custom) IsNewerThan(version Version) bool {
	return false
}

func (c Custom) String() string {
	return c.V
}
