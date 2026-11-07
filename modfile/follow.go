package main

import "fmt"

func applyFollowInstruction(file *File, body string) error {
	if file.Follow != "" {
		return fmt.Errorf("multiple follow statements in file")
	}

	file.Follow = body
	return nil
}
