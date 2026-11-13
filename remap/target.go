package remap

type Mapping struct {
	TinyURL string `json:"tinyURL"`
	Name    string `json:"name"`
}

// DownloadMapping downloads the given mappings. The hash is used to ensure multiple projects
// don't overwrite each other.
func DownloadMapping(mapping Mapping, hash string) (string, error) {

}
