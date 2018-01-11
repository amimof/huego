package huego

// https://developers.meethue.com/documentation/groups-api
type Group struct {
	Name string `json:"name,omitempty"`
	Lights []string `json:"lights,omitempty"`
	Type string `json:"type,omitempty"`
	State	*GroupState `json:"state,omitempty"`
	Recycle bool `json:"recycle,omitempty"`
	Class	string `json:"class,omitempty"`
	Action *State `json:"action,omitempty"`
	Id int `json:"-"`
}

type GroupState struct {
	AllOn bool `json:"all_on,omitempty"`
	AnyOn bool `json:"any_on,omitempty"`
}