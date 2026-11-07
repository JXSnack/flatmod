package helper

import "slices"

var Opts []string

func IsVerbose() bool {
	return slices.Contains(Opts, "v") || slices.Contains(Opts, "-verbose")
}

func ForceGitHTTPS() bool {
	return slices.Contains(Opts, "-git-https")
}
