package mojang

type Artifact struct {
	Path string `json:"path"`
	SHA1 string `json:"sha1"`
	Size string `json:"size"`
	URL  string `json:"url"`
}

type Library struct {
	Downloads struct {
		Artifact Artifact `json:"artifact"`
	} `json:"downloads"`
	Name string `json:"name"`
}
