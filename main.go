package main

import (
	"fmt"

	"github.com/JXSnack/jmod/version"
)

func main() {
	fmt.Println(version.Parse("v1.0.0"))
}
