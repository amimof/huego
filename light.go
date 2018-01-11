package huego

// https://developers.meethue.com/documentation/lights-api
type Light struct {
	State *State `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	ModelId string `json:"modelid,omitempty"`
	ManufacturerName string `json:"manufacturername,omitempty"`
	UniqueId string `json:"uniqueid,omitempty"`
	SwVersion string `json:"swversion,omitempty"`
	SwConfigId string `json:"swconfigid,omitempty"`
	ProductId string `json:"productid,omitempty"`
	Id int `json:"-"`
}

type State struct {
	On bool	`json:"on"`
	Bri uint8	`json:"bri,omitempty"`
	Hue uint16	`json:"hue,omitempty"`
	Sat uint8 `json:"sat,omitempty"`
	Xy []float32 `json:"xy,omitempty"`
	Ct uint16 `json:"ct,omitempty"`
	Alert	string `json:"alert,omitempty"`
	Effect string `json:"effect,omitempty"`
	TransitionTime uint16 `json:"transitiontime,omitempty"`
	BriInc int `json:"bri_inc,omitempty"`
	SatInc int `json:"sat_inc,omitempty"`
	HueInc int `json:"hue_inc,omitempty"`
	CtInc int `json:"ct_inc,omitempty"`
	XyInc int `json:"xy_inc,omitempty"`
	ColorMode	string `json:"colormode,omitempty"`
	Reachable	bool `json:"reachable,omitempty"`
}

type NewLight struct {
	Lights []string
	LastScan string `json:"lastscan"`
}

// Sets the On state of one light to false, turning it off
func (l *Light) TurnOff() (*Response, error) {
	return nil, nil
}