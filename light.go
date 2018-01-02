package huego

import (
	"encoding/json"
	"strconv"
	"time"
	"fmt"
)

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
	LastScan time.Time `json:"lastscan"`
}

// GetLights will return all lights
// See: https://developers.meethue.com/documentation/lights-api#11_get_all_lights
func (h *Hue) GetLights() ([]Light, error) {

	m := map[string]Light{}

	res, err := h.GetResource(h.GetApiUrl("/lights/"))
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

// GetLight returns a light with the id of i
// See: https://developers.meethue.com/documentation/lights-api#11_get_all_lights
func (h *Hue) GetLight(i int) (*Light, error) {

	var light *Light

	res, err := h.GetResource(h.GetApiUrl("/lights/", strconv.Itoa(i)))
	if err != nil {
		return light, err
	}

	err = json.Unmarshal(res, &light)
	if err != nil {
		return light, err
	}

	return light, nil
}

// SetLight allows for controlling a light state properties.
// See: https://developers.meethue.com/documentation/lights-api#15_set_light_attributes_rename
func (h *Hue) SetLight(i int, l *State) (*Response, error) {

	var a []*ApiResponse

	l.Reachable = false
	l.ColorMode = ""

	data, err := json.Marshal(&l)
	if err != nil {
		return nil, err
	}

	url := h.GetApiUrl("/lights/", strconv.Itoa(i), "/state")
	res, err := h.PutResource(url, data)
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

// Search starts a search for new lights
// See: https://developers.meethue.com/documentation/lights-api#13_search_for_new_lights
func (h *Hue) FindLights() (*Response, error) {

	var a []*ApiResponse

	url := h.GetApiUrl("/lights/")

	res, err := h.PostResource(url, nil)
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

// See: https://developers.meethue.com/documentation/lights-api#12_get_new_lights
func (h *Hue) GetNewLights() (*NewLight, error){

	var n map[string]interface{}
	
	url := h.GetApiUrl("/lights/new")

	res, err := h.GetResource(url)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(res, &n)

	lights := make([]string, 0, len(n))
	var lastscan time.Time

	for k, _ := range n {
		if k == "lastscan" {
			lastscan = n[k].(time.Time)
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

// DeleteLight deletes a light
// See: https://developers.meethue.com/documentation/lights-api#17_delete_lights
func (h *Hue) DeleteLight(i int) error {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/lights/", id)

	res, err := h.DeleteResource(url)
	if err != nil {
		return err
	}

	fmt.Println(string(res))

	_ = json.Unmarshal(res, &a)

	_, err = handleResponse(a)
	if err != nil {
		return err
	}

	return nil

}

// Update a light
func (h *Hue) UpdateLight(i int, light *Light) (*Response, error) {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/lights/", id)

	data, err := json.Marshal(&light)
	if err != nil {
		return nil, err
	}

	res, err := h.PutResource(url, data)
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
