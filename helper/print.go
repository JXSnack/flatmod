package helper

import "fmt"

func Vprintln(a ...any) {
	if Verbose {
		fmt.Println(a...)
	}
}

func Vprintf(format string, a ...any) {
	if Verbose {
		fmt.Printf(format, a...)
	}
}
