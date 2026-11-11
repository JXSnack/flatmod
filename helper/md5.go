package helper

import (
	"crypto/md5"
	"fmt"
	"time"
)

// ZeroSum returns a md5 sum of the string "no value". This method is used for caching.
func ZeroSum() [16]byte {
	return md5.Sum([]byte("no value"))
}

func HourlySum() [16]byte {
	return md5.Sum([]byte(fmt.Sprintf("%d-%d", time.Now().Hour(), time.Now().Day())))
}
