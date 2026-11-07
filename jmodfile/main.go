package main

import (
	"net/url"

	"github.com/JXSnack/jmod/dependency"
)

type File struct {
	Group        string
	Follow       url.URL
	Dependencies []dependency.Dependency
}
