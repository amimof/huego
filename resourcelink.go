package huego

// Resourcelink represents a bridge resourcelink https://developers.meethue.com/documentation/resourcelinks-api
type Resourcelink struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type,omitempty"`
	ClassID     uint16   `json:"classid,omitempty"`
	Owner       string   `json:"owner,omitempty"`
	Links       []string `json:"links,omitempty"`
	ID          int      `json:",omitempty"`
}
