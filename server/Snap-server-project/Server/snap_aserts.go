package main

type snapsAssert struct {
	Embedded snapAssertPacks `json:"_embedded"`
}

type snapAssertPacks struct {
	ClickIndex_Package []snapAssertDetails `json:"clickindex:package"`
}

type snapAssertDetails struct {
	Type         string `json:"type"`
	Authority    string `json:"authority-id,omitempty"`
	Developer    string `json:"developer-id,omitempty"`
	Publisher    string `json:"publisher-id,omitempty"`
	Revision     int    `json:"revision,omitempty"`
	Series       int    `json:"series,omitempty"`
	SnapRevision int    `json:"snap-revision,omitempty"`
	SnapId       string `json:"snap-id,omitempty"`
	SnapSize     int    `json:"snap-size,omitempty"`
	SignKeySha   string `json:"snap-size,omitempty"`
	Timestamp    string `json:"snap-size,omitempty"`
}
