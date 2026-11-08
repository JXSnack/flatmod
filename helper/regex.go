package helper

import "regexp"

var ReValidId = regexp.MustCompile(`[^a-zA-Z0-9_]`)
