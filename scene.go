package huego

// Scene represents a bridge scene https://developers.meethue.com/documentation/scenes-api
type Scene struct {
	Name            string      `json:"name,omitempty"`
	Lights          []string    `json:"lights,omitempty"`
	Owner           string      `json:"owner,omitempty"`
	Recycle         bool        `json:"recycle,omitempty"`
	Locked          bool        `json:"locked,omitempty"`
	AppData         interface{} `json:"appdata,omitempty"`
	Picture         string      `json:"picture,omitempty"`
	LastUpdated     string      `json:"lastupdated,omitempty"`
	Version         int         `json:"version,omitempty"`
	StoreSceneState bool        `json:"storescenestate,omitempty"`
	ID              string      `json:"-"`
}
