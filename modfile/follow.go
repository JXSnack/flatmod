package modfile

import (
	"fmt"
	"strings"
)

func applyFollowInstruction(file *File, body string) error {
	if file.Follow != "" {
		return fmt.Errorf("multiple follow statements in file")
	}

	split := strings.SplitN(body, " ", 2)
	if len(split) != 2 {
		return fmt.Errorf("missing git branch for follow")
	}

	file.Follow = split[0]
	file.FollowBranch = split[1]
	return nil
}
