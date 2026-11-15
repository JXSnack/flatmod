package helper

import (
	"fmt"
	"strings"
)

func InputBoolf(format string, a ...any) bool {
	input := InputStringf(format, a...)

	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	return input == "y" || input == "yes"
}

func InputStringf(format string, a ...any) string {
	var input string
	fmt.Printf(format, a...)
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}

	return input
}
