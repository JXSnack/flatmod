package main

import (
	"os"
)

func Check() {
	file := GetMod("check")
	if file == nil {
		os.Exit(1)
	}
}
