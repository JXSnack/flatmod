package mojang

type AssetIndex struct {
	Id        string `json:"id"`
	SHA1      string `json:"sha1"`
	Size      int    `json:"size"`
	TotalSize int    `json:"totalSize"`
	URL       string `json:"url"`
}

type Download struct {
	SHA1 string `json:"sha1"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}

type Version struct {
	AssetIndex      AssetIndex `json:"assetIndex"`
	Assets          string     `json:"assets"`
	ComplianceLevel int        `json:"complianceLevel"`
	Downloads       struct {
		Client         Download `json:"client"`
		ClientMappings Download `json:"client_mappings"`
		Server         Download `json:"server"`
		ServerMappings Download `json:"server_mappings"`
	} `json:"downloads"`
	JavaVersion struct {
		Component    string `json:"component"`
		MajorVersion string `json:"majorVersion"`
	} `json:"javaVersion"`
	Libraries              []Library `json:"libraries"`
	MainClass              string    `json:"mainClass"`
	MinimumLauncherVersion int       `json:"minimumLauncherVersion"`
	Type                   string    `json:"type"`
}
