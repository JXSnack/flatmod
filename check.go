package main

import (
	"os"
)

// TODO: multiple errors instead of breaking at the first problem

func Check() {
	file := GetMod("check")
	if file == nil {
		os.Exit(1)
	}
}
