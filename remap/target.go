package remap

type Settings struct {
	TinyJar  string `json:"tinyJar"`
	TinyPath string `json:"tinyPath"`
	From     string `json:"from"`
	To       string `json:"to"`
}

// DownloadMapping downloads the given mappings. The hash is used to ensure multiple projects
// don't overwrite each other.
func DownloadMapping(mapping Mapping, hash string) (string, error) {

}
