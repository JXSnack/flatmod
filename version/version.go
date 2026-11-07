package version

type Version interface {
	Kind() string
	IsNewerThan(version Version) bool
	String() string
}

func Parse(version string) Version {
	if sem, ok := ParseSemantic(version); ok {
		return sem
	}

	return Custom{V: version}
}
