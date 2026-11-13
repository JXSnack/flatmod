package follow

import "fmt"

func ApplyDuplicateStrategy(strategy string, defined, user []byte) ([]byte, error) {
	switch strategy {
	case "override":
		return user, nil
	case "append":
		return append(defined, user...), nil
	case "prepend":
		return append(user, defined...), nil
	}

	return []byte{}, fmt.Errorf("unknown duplicate strategy %q", strategy)
}
