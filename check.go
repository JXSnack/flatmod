package main

import (
	"os"
)

func Check() {
	file, _ := GetMod("check")
	if file == nil {
		os.Exit(1)
	}
}
