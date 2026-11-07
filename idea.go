package main

import "os"

func Idea() {
	file := GetMod("idea")
	if file == nil {
		os.Exit(1)
	}

}
