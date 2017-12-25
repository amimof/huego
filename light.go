package huego

import (
	"encoding/json"
	"strconv"
)

type Light struct {
	State 			 *State `json:"state,omitempty"`
	Type 			 string `json:"type,omitempty"`
	Name 			 string `json:"name,omitempty"`
	ModelId 		 string `json:"modelid,omitempty"`
	ManufacturerName string `json:"modelid,omitempty"`
	UniqueId 		 string `json:"string,omitempty"`
	SwVersion 		 string `json:"string,omitempty"`
	SwConfigId 		 string `json:"string,omitempty"`
	ProductId 		 string `json:"productid,omitempty"`
	Id 				 int 	`json:",omitempty"`
}

type State struct {
	On 			bool 		`json:"on"`
	Bri 		int 		`json:"bri"`
	Hue 		int 		`json:"hue"`
	Sat 		int 		`json:"sat"`
	Effect 		string 		`json:"effect"`
	Xy 			[]float32 	`json:"xy"`
	Ct 			int 		`json:"ct"`
	Alert 		string 		`json:"alert"`
	ColorMode 	string 		`json:"colormode"`
	Reachable 	bool 		`json:"reachable"`
}

type NewLight struct {
	Lights []*Light
	LastScan string `json:"lastscan"`
}

type Reply struct {
	Type string
	Address string
	Value string
}

// GetLights will return all lights
// See: https://developers.meethue.com/documentation/lights-api#11_get_all_lights
func (h *Hue) GetLights() ([]*Light, error) {

	m := map[string]Light{}

	res, err := h.GetResource(h.GetApiUrl("/lights/"))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &m)
	if err != nil {
		return nil, err
	}

	lights := make([]*Light, 0, len(m))

	for i, l := range m {
		l.Id, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		lights = append(lights, &l)
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
func (h *Hue) SetLight(i int, l State) ([]*Response, error) {

	var a []*ApiResponse

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
func (h *Hue) FindLights() ([]*Response, error) {

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

	n := map[string]Light{}
	var result *NewLight

	url := h.GetApiUrl("/lights/new")

	res, err := h.GetResource(url)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(res, &n)
	lights := make([]*Light, 0, len(n))

	for i, l := range n {
		if i != "lastscan" {
			l.Id, err = strconv.Atoi(i)
			if err != nil {
				return result, err
			}
			lights = append(lights, &l)
		}
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return result, err
	}

	resu := &NewLight{lights, result.LastScan}

	return resu, nil

}

// DeleteLight deletes a light
// See: https://developers.meethue.com/documentation/lights-api#17_delete_lights
func (h *Hue) DeleteLight(i int) ([]*Response, error) {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/lights/", id)

	res, err := h.DeleteResource(url)
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

// Update a light
func (h *Hue) UpdateLight(i int, light *Light) ([]*Response, error) {

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
