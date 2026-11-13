package follow

import (
	"fmt"

	"github.com/magiconair/properties"
)

func ApplyDuplicateStrategy(strategy string, defined, user []byte) ([]byte, error) {
	switch strategy {
	case "override":
		return user, nil
	case "append":
		return append(defined, user...), nil
	case "prepend":
		return append(user, defined...), nil
	case "properties":
		p, err := appendToProperties(string(defined), string(user))
		return []byte(p), err
	}

	return []byte{}, fmt.Errorf("unknown duplicate strategy %q", strategy)
}

func appendToProperties(defined string, user string) (string, error) {
	followP, err := properties.LoadString(defined)
	if err != nil {
		return "", fmt.Errorf("follow properties: %w", err)
	}

	userP, err := properties.LoadString(user)
	if err != nil {
		return "", fmt.Errorf("user properties: %w", err)
	}

	for key, value := range userP.Map() {
		_, _, err = followP.Set(key, value)
		if err != nil {
			return "", fmt.Errorf("merge properties: %w", err)
		}
	}
	return followP.String(), nil
}
