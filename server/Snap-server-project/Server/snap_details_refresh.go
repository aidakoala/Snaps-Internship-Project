package main

// REFRESH + INSTALL STRUCTS
type briefSnapDetails struct {
	SnapID          string `json:"snap_id"`
	InstanceKEY     string `json:"instance_key"`
	Revision        int    `json:"revision"` // store revisions are ints starting at 1
	TrackingChannel string `json:"tracking-channel"`
	RefreshedDate   string `json:"refreshed-date"`
}

type action struct {
	Action      string `json:"action"`
	InstanceKEY string `json:"instance-key"`
	Name        string `json:"name"`
	Channel     string `json:"channel"`
}

type refreshDetails struct {
	Context []briefSnapDetails `json:"context"`
	Actions []action           `json:"actions"`
	Fileds  []string           `json:"fields"`
}

// RESPONSE
type Epoch struct {
	Read  []uint32 `yaml:"read"`
	Write []uint32 `yaml:"write"`
}

type StoreAccount struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"display-name"`
	Validation  string `json:"validation,omitempty"`
}

type storeSnapDownload struct {
	Sha3_384 string           `json:"sha3-384"`
	Size     int64            `json:"size"`
	URL      string           `json:"url"`
	Deltas   []storeSnapDelta `json:"deltas"`
}

type storeSnapDelta struct {
	Format   string `json:"format"`
	Sha3_384 string `json:"sha3-384"`
	Size     int64  `json:"size"`
	Source   int    `json:"source"`
	Target   int    `json:"target"`
	URL      string `json:"url"`
}

type storeSnapMedia struct {
	Type   string `json:"type"` // icon/screenshot
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type storeSnap struct {
	Architectures []string          `json:"architectures"`
	Base          string            `json:"base"`
	CommonIDs     []string          `json:"common-ids"`
	Confinement   string            `json:"confinement"`
	Contact       string            `json:"contact"`
	CreatedAt     string            `json:"created-at"` // revision timestamp
	Description   string            `json:"description"`
	Download      storeSnapDownload `json:"download"`
	Epoch         Epoch             `json:"epoch"`
	License       string            `json:"license"`
	Media         []storeSnapMedia  `json:"media"`
	Name          string            `json:"name"`
	Prices        map[string]string `json:"prices"` // currency->price,  free: {"USD": "0"}
	Private       bool              `json:"private"`
	Publisher     StoreAccount      `json:"publisher"`
	Revision      int               `json:"revision"` // store revisions are ints starting at 1
	SnapID        string            `json:"snap-id"`
	SnapYAML      string            `json:"snap-yaml"` // optional
	Summary       string            `json:"summary"`
	Title         string            `json:"title"`
	Type          string            `json:"type"`
	Version       string            `json:"version"` // TODO: not yet defined: channel map
}

type result struct {
	Channel     string    `json:"effective-channel"`
	InstanceKEY string    `json:"instance-key"`
	Name        string    `json:"name"`
	Res         string    `json:"result"`
	Snap        storeSnap `json:"snap"`
	SnapID      string    `json:"snap-id"`
}

type snapActionResultList struct {
	ErrorList []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error-list"`
	Results []result `json:"results"`
}
