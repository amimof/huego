package huego

import (
	"encoding/json"
	"strconv"
)

type Group struct {
	Name string `json:"name,omitempty"`
	Lights []string `json:"lights,omitempty"`
	Type string `json:"type,omitempty"`
	State	GroupState `json:"state,omitempty"`
	Recycle bool `json:"recycle,omitempty"`
	Class	string `json:"class,omitempty"`
	Action State `json:"action,omitempty"`
	Id int `json:"-"`
}

type GroupState struct {
	AllOn bool `json:"all_on,omitempty"`
	AnyOn bool `json:"any_on,omitempty"`
}

// GetGroups will return all groups
// See: hhttps://developers.meethue.com/documentation/groups-api#21_get_all_groups
func (h *Hue) GetGroups() ([]Group, error) {

	var m map[string]Group

	res, err := h.GetResource(h.GetApiUrl("/groups/"))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &m)
	groups := make([]Group, 0, len(m))

	for i, g := range m {
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

	var g *Group

	url := h.GetApiUrl("/groups/", strconv.Itoa(i))
	res, err := h.GetResource(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}


// SetGroupState allows for controlling light state properties for all lights in a group with the id of i
// See: https://developers.meethue.com/documentation/groups-api#25_set_group_state
func (h *Hue) SetGroupState(i int, l *State) (*Response, error) {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/groups/", id, "/action/")

	data, err := json.Marshal(&l)
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

// Update a group
func (h *Hue) UpdateGroup(i int, l *Group) (*Response, error) {
	
	var a []*ApiResponse

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/groups/", id)

	data, err := json.Marshal(&l)
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

// CreateGroup creates a new group
// See: https://developers.meethue.com/documentation/groups-api#22_create_group
func (h *Hue) CreateGroup(g *Group) (*Response, error) {

	var a []*ApiResponse

	url := h.GetApiUrl("/groups/")

	data, err := json.Marshal(&g)
	if err != nil {
		return nil, err
	}

	res, err := h.PostResource(url, data)
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

// DeleteGroup deletes a group with the id of i
// See: https://developers.meethue.com/documentation/groups-api#26_delete_group
func (h *Hue) DeleteGroup(i int) error {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/groups/", id)

	res, err := h.DeleteResource(url)
	if err != nil {
		return err
	}

	_ = json.Unmarshal(res, &a)

	_, err = handleResponse(a)
	if err != nil {
		return err
	}

	return  nil
}
