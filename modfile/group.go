package main

import "fmt"

func applyGroupInstruction(file *File, body string) error {
	if file.Group != "" {
		return fmt.Errorf("multiple group statements in file")
	}

	file.Group = body
	return nil
}
