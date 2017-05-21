package huego

import (
	//"github.com/amimof/loglevel-go"
	"net/http"
	"encoding/json"
	//"crypto/tls"
	//"net/url"
	//"path"
	"strconv"
	"io/ioutil"
	"strings"
)

type Group struct {
	Name 	string 		`json:"name,omitempty"`
	Lights 	[]string 	`json:"lights,omitempty"`
	Type 	string 		`json:"type,omitempty"`
	State 	*GroupState `json:"state,omitempty"`
	Recycle bool 		`json:"recycle,omitempty"`
	Class 	string 		`json:"class,omitempty"`
	Action 	*Action 	`json:"action,omitempty"`
	Id 		int			`json:",omitempty"`
}

type Action struct {
	On 			bool 		`json:"on,omitempty"`
	Bri 		int 		`json:"bri,omitempty"`
	Hue 		int 		`json:"hue,omitempty"`
	Sat 		int 		`json:"sat,omitempty"`
	Effect 		string 		`json:"effect,omitempty"`
	Xy 			[]float32 	`json:"xy,omitempty"`
	Ct 			int 		`json:"ct,omitempty"`
	Alert 		string 		`json:"alert,omitempty"`
	ColorMode 	string 		`json:"colormode,omitempty"`
}

type GroupState struct {
	AllOn bool `json:"all_on,omitempty"`
	AnyOn bool `json:"any_on,omitempty"`
}

// GetGroups will return all groups
// See: hhttps://developers.meethue.com/documentation/groups-api#21_get_all_groups
func (h *Hue) GetGroups() ([]Group, error) {
	
	gm := map[string]Group{}
	url := h.GetApiUrl("/groups/")

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&gm)
	groups := make([]Group, 0, len(gm))

	for i, g := range gm {
		g.Id, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}
	return groups, err
}

// GetGroup returns a group with the id of i
// See: https://developers.meethue.com/documentation/groups-api#23_get_group_attributes
func (h *Hue) GetGroup(i int) (*Group, error) {
	
	var group *Group

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/groups/", id)

	res, err := http.Get(url)
	if err != nil {
		return group, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&group)
	if err != nil {
		return group, err
	}

	return group, err
}

// SetGroupState allows for controlling light state properties for all lights in a group with the id of i
// See: https://developers.meethue.com/documentation/groups-api#25_set_group_state
func (h *Hue) SetGroupState(i int, l *Action) ([]Response, error) {
	
	var r []Response

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/groups/", id, "/action/")

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

// SetGroup sets the name, class and light members of a group with the id of i
// See: https://developers.meethue.com/documentation/groups-api#24_set_group_attributes
func (h *Hue) SetGroup(i int, l *Group) ([]Response, error) {
	var r []Response

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/groups/", id)

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

// CreateGroup creates a new group
// See: https://developers.meethue.com/documentation/groups-api#22_create_group
func (h *Hue) CreateGroup(g *Group) ([]Response, error) {
	
	var r []Response

	url := h.GetApiUrl("/groups/")

	data, err := json.Marshal(&g)
	if err != nil {
		return r, err
	}

	body := strings.NewReader(string(data))
	
	res, err := http.Post(url, "application/json", body)
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

// DeleteGroup deletes a group with the id of i
// See: https://developers.meethue.com/documentation/groups-api#26_delete_group
func (h *Hue) DeleteGroup(i int) ([]Response, error) {
	
	var r []Response

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/groups/", id)
	
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

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(b, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}






