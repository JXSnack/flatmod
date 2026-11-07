package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Jmod is used for managing Jmod projects.")
		fmt.Println()
		fmt.Println("Usage: jmod <command> [args]")
		fmt.Println()
		fmt.Println("All commands:")
		fmt.Println("\tcheck\tread configuration files for errors")
		return
	}

	if os.Args[1] == "check" {
		Check()
	}
}
