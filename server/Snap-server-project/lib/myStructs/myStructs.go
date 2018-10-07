package myStructs

type snapDetails struct {
	AnonDownloadURL  string   `json:"anon_download_url,omitempty"`
	Architectures    []string `json:"architecture"`
	Channel          string   `json:"channel,omitempty"`
	DownloadSha3_384 string   `json:"download_sha3_384,omitempty"`

	DownloadSize int64  `json:"binary_filesize,omitempty"`
	DownloadURL  string `json:"download_url,omitempty"`
	// Epoch        snap.Epoch         `json:"epoch"`
	IconURL     string             `json:"icon_url"`
	LastUpdated string             `json:"last_updated,omitempty"`
	Name        string             `json:"package_name"`
	Prices      map[string]float64 `json:"prices,omitempty"`

	Publisher      string   `json:"publisher,omitempty"`
	RatingsAverage float64  `json:"ratings_average,omitempty"`
	Revision       int      `json:"revision"`
	ScreenshotURLs []string `json:"screenshot_urls,omitempty"`
	SnapID         string   `json:"snap_id"`
	License        string   `json:"license,omitempty"`
	Base           string   `json:"base,omitempty"`

	SupportURL string `json:"support_url"`
	Contact    string `json:"contact"`

	Title   string `json:"title"`
	Version string `json:"version"`

	Developer           string `json:"origin"`
	DeveloperID         string `json:"developer_id"`
	DeveloperName       string `json:"developer_name"`
	DeveloperValidation string `json:"developer_validation"`

	Private     bool   `json:"private"`
	Confinement string `json:"confinement"`

	CommonIDs []string `json:"common_ids,omitempty"`
}
