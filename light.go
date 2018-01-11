package huego

import (
	"encoding/json"
	"strconv"
)

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

// Returns all lights known to the bridge
func (b *Bridge) GetLights() ([]Light, error) {

	m := map[string]Light{}

	url, err := b.getApiPath("/lights/")
	if err != nil {
		return nil, err
	}

	res, err := b.getResource(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &m)
	if err != nil {
		return nil, err
	}

	lights := make([]Light, 0, len(m))

	for i, l := range m {
		l.Id, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		lights = append(lights, l)
	}

	return lights, nil

}

// Returns one light with the id of i
func (b *Bridge) GetLight(i int) (*Light, error) {

	var light *Light

	url, err := b.getApiPath("/lights/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.getResource(url)
	if err != nil {
		return light, err
	}

	err = json.Unmarshal(res, &light)
	if err != nil {
		return light, err
	}

	return light, nil
}

// Allows for controlling one light's state
func (b *Bridge) SetLight(i int, l *State) (*Response, error) {

	var a []*ApiResponse

	l.Reachable = false
	l.ColorMode = ""

	data, err := json.Marshal(&l)
	if err != nil {
		return nil, err
	}

	url, err := b.getApiPath("/lights/", strconv.Itoa(i), "/state")
	if err != nil {
		return nil, err
	}
	res, err := b.putResource(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &a)
	if err != nil {
		return nil, err
	}
	
	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// Starts a search for new lights on the bridge. 
// Use GetNewLights() verify if new lights have been detected. 
func (b *Bridge) FindLights() (*Response, error) {

	var a []*ApiResponse

	url, err := b.getApiPath("/lights/")
	if err != nil {
		return nil, err
	}

	res, err := b.postResource(url, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// Returns a list of lights that were discovered last time FindLights() was executed.
func (b *Bridge) GetNewLights() (*NewLight, error){

	var n map[string]interface{}
	
	url, err := b.getApiPath("/lights/new")
	if err != nil {
		return nil, err
	}

	res, err := b.getResource(url)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(res, &n)

	lights := make([]string, 0, len(n))
	var lastscan string

	for k, _ := range n {
		if k == "lastscan" {
			lastscan = n[k].(string)
		} else {
			lights = append(lights, n[k].(string))
		}
	}

	result := &NewLight{
		Lights: lights, 
		LastScan: lastscan,
	}

	return result, nil

}

// Deletes one lights from the bridge
func (b *Bridge) DeleteLight(i int) error {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/lights/", id)
	if err != nil {
		return err
	}

	res, err := b.deleteResource(url)
	if err != nil {
		return err
	}

	_ = json.Unmarshal(res, &a)

	_, err = handleResponse(a)
	if err != nil {
		return err
	}

	return nil

}

// Updates one light's attributes and state properties
func (b *Bridge) UpdateLight(i int, light *Light) (*Response, error) {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/lights/", id)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&light)
	if err != nil {
		return nil, err
	}

	res, err := b.putResource(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
