package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("fmod is used for managing Flatmod source code.")
		fmt.Println()
		fmt.Println("Usage: fmod <command> [args]")
		fmt.Println()
		fmt.Println("All commands:")
		fmt.Println("\tcheck\tread configuration files for errors")
		return
	}

	if os.Args[1] == "check" {
		Check()
	}
}
