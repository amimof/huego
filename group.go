package huego

import (
	"encoding/json"
	"strconv"
)

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

// Returns all groups known to the bridge
func (b *Bridge) GetGroups() ([]Group, error) {

	var m map[string]Group

	url, err := b.getApiPath("/groups/")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &m)
	if err != nil {
		return nil, err
	}
	
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

// Returns one group known to the bridge by its id
func (b *Bridge) GetGroup(i int) (*Group, error) {

	var g *Group

	url, err := b.getApiPath("/groups/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}


// Allows for setting the state of one group, controlling the state of all lights in that group.
func (b *Bridge) SetGroupState(i int, l *State) (*Response, error) {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/groups/", id, "/action/")
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&l)
	if err != nil {
		return nil, err
	}

	res, err := put(url, data)
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

// Updates one group known to the bridge 
func (b *Bridge) UpdateGroup(i int, l *Group) (*Response, error) {
	
	var a []*ApiResponse

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/groups/", id)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&l)
	if err != nil {
		return nil, err
	}

	res, err := put(url, data)
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

// Creates one new group with attributes defined by g
func (b *Bridge) CreateGroup(g Group) (*Response, error) {

	var a []*ApiResponse

	url, err := b.getApiPath("/groups/")
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&g)
	if err != nil {
		return nil, err
	}

	res, err := post(url, data)
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

// Deletes one group with the id of i
func (b *Bridge) DeleteGroup(i int) error {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/groups/", id)
	if err != nil {
		return err
	}

	res, err := delete(url)
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