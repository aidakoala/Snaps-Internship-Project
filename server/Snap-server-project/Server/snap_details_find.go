package main

// FIND STRUCTS

type snapsNall struct {
	Embedded snapPacks `json:"_embedded"`
	Links    linkTypes `json:"_links,omitempty"`
}

type linkTypes struct {
	SelfLink topLevelLinks `json:"self,omitempty"`
}

type topLevelLinks struct {
	Href    string `json:"href,omitempty"`
	Related string `json:"related,omitempty"`
}

type snapPacks struct {
	ClickIndex_Package []snapDetails `json:"clickindex:package"`
}

type snapDetails struct {
	AnonDownloadURL  string             `json:"anon_download_url,omitempty"`
	Architectures    []string           `json:"architecture"`
	Channel          string             `json:"channel,omitempty"`
	DownloadSha3_384 string             `json:"download_sha3_384,omitempty"`
	Summary          string             `json:"summary,omitempty"`
	Description      string             `json:"description,omitempty"`
	DownloadSize     int64              `json:"binary_filesize,omitempty"`
	DownloadURL      string             `json:"download_url,omitempty"`
	Epoch            string             `json:"epoch"`
	IconURL          string             `json:"icon_url"`
	LastUpdated      string             `json:"last_updated,omitempty"`
	Name             string             `json:"package_name"`
	Prices           map[string]float64 `json:"prices,omitempty"`

	// Note that the publisher is really the "display name" of the
	// publisher
	Publisher      string   `json:"publisher,omitempty"`
	RatingsAverage float64  `json:"ratings_average,omitempty"`
	Revision       int      `json:"revision"` // store revisions are ints starting at 1
	ScreenshotURLs []string `json:"screenshot_urls,omitempty"`
	SnapID         string   `json:"snap_id"`
	License        string   `json:"license,omitempty"`
	Base           string   `json:"base,omitempty"`

	// FIXME: the store should send "contact" here, once it does we
	//        can remove support_url
	SupportURL string `json:"support_url"`
	Contact    string `json:"contact"`

	Title   string `json:"title"`
	Type    string `json:"content,omitempty"`
	Version string `json:"version"`

	Developer           string `json:"origin"`
	DeveloperID         string `json:"developer_id"`
	DeveloperName       string `json:"developer_name"`
	DeveloperValidation string `json:"developer_validation"`

	Private     bool   `json:"private"`
	Confinement string `json:"confinement"`

	CommonIDs []string `json:"common_ids,omitempty"`
}
