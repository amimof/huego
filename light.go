package huego

import (
	"net/http"
	"encoding/json"
	"strconv"
	"strings"
	"io/ioutil"
)

type Light struct {
	State *State `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	ModelId string `json:"modelid,omitempty"`
	ManufacturerName string `json:"modelid,omitempty"`
	UniqueId string `json:"string,omitempty"`
	SwVersion string `json:"string,omitempty"`
	SwConfigId string `json:"string,omitempty"`
	ProductId string `json:"productid,omitempty"`
	Id int `json:",omitempty"`
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
	Lights *[]Light
	LastScan string `json:"lastscan"`
}

// GetLights will return all lights
// See: https://developers.meethue.com/documentation/lights-api#11_get_all_lights
func (h *Hue) GetLights() ([]Light, error) {
	
	lm := map[string]Light{}

	//url := GetPath("/lights/")
	url := h.GetApiUrl("/lights/")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&lm)
	lights := make([]Light, 0, len(lm))

	for i, l := range lm {
		l.Id, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		lights = append(lights, l)
	}
	return lights, err
}

// GetLight returns a light with the id of i
// See: https://developers.meethue.com/documentation/lights-api#11_get_all_lights
func (h *Hue) GetLight(i int) (*Light, error) {

	var light *Light
	id := strconv.Itoa(i)
	url := h.GetApiUrl("/lights/", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return light, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return light, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&light)
	if err != nil {
		return light, err
	}

	return light, nil
} 

// SetLight allows for controlling a light state properties.
// See: https://developers.meethue.com/documentation/lights-api#15_set_light_attributes_rename
func (h *Hue) SetLight(i int, l State) ([]Response, error) {
	
	var r []Response

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/lights/", id, "/state")

	data, err := json.Marshal(&l)
	if err != nil {
		return r, err
	}

	body := strings.NewReader(string(data))
	
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return r, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return r, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return r, err
	}

	return r, nil

}

// Search starts a search for new lights
// See: https://developers.meethue.com/documentation/lights-api#13_search_for_new_lights
func (h *Hue) Search() ([]Response, error) {

	var r []Response

	url := h.GetApiUrl("/lights/")

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return r, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return r, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return r, err
	}

	return r, nil

}

// See: https://developers.meethue.com/documentation/lights-api#12_get_new_lights
func (h *Hue) GetNewLights() (*NewLight, error){

	n := map[string]Light{}
	var result *NewLight

	url := h.GetApiUrl("/lights/new")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return result, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &n)
	newlights := make([]Light, 0, len(n))

	for i, l := range n {
		if i != "lastscan" {
			l.Id, err = strconv.Atoi(i)
			if err != nil {
				return result, err
			}
			newlights = append(newlights, l)
		}
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	resu := &NewLight{ &newlights, result.LastScan}

	return resu, nil

}

// DeleteLight deletes a light
// See: https://developers.meethue.com/documentation/lights-api#17_delete_lights
func (h *Hue) DeleteLight(i int) ([]Response, error) {

	var r []Response

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/lights/", id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return r, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return r, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return r, err
	}

	return r, nil

}

// RenameLight sets the name attribute on a light
// See: https://developers.meethue.com/documentation/lights-api#15_set_light_attributes_rename
func (h *Hue) RenameLight(i int, n string) ([]Response, error) {

	var r []Response
	var l *Light = &Light{Name: n}

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/lights/", id)

	data, err := json.Marshal(&l)
	if err != nil {
		return r, err
	}

	body := strings.NewReader(string(data))	

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return r, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return r, err 
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return r, err
	}

	return r, nil
}


