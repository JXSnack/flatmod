package helper

import "fmt"

func Vprintln(a ...any) {
	if IsVerbose() {
		fmt.Println(a...)
	}
}

func Vprintf(format string, a ...any) {
	if IsVerbose() {
		fmt.Printf(format, a...)
	}
}
