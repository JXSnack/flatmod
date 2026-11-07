package version

import (
	"regexp"
	"strconv"
)

var SemanticPattern = regexp.MustCompile(`^v(\d+)\.(\d+)\.(\d+)(-.+)?$`)

type Semantic struct {
	Major int
	Minor int
	Patch int

	PreRelease string
}

func (s Semantic) Kind() string {
	return "semantic"
}

func (s Semantic) IsNewerThan(version Version) bool {
	other, ok := version.(Semantic)
	if !ok {
		return false
	}

	if s.Major != other.Major {
		return s.Major > other.Major
	}

	if s.Minor != other.Minor {
		return s.Minor > other.Minor
	}

	if s.Patch != other.Patch {
		return s.Patch > other.Patch
	}

	return false
}

func IsSemantic(version string) bool {
	match := SemanticPattern.FindAllStringSubmatch(version, -1)
	return match != nil
}

func ParseSemantic(version string) (Semantic, bool) {
	match := SemanticPattern.FindStringSubmatch(version)

	if match == nil {
		return Semantic{}, false
	}

	major, err := strconv.Atoi(match[1])
	if err != nil {
		return Semantic{}, false
	}

	minor, err := strconv.Atoi(match[2])
	if err != nil {
		return Semantic{}, false
	}

	patch, err := strconv.Atoi(match[3])
	if err != nil {
		return Semantic{}, false
	}

	return Semantic{
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		PreRelease: match[4],
	}, true
}
