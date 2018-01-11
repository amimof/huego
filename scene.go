package huego

import (
	"encoding/json"
	"strconv"
	"fmt"
)

// https://developers.meethue.com/documentation/scenes-api
type Scene struct {
	Name string `json:"name,omitempty"`
	Lights []string `json:"lights,omitempty"`
	Owner string `json:"owner,omitempty"`
	Recycle bool `json:"recycle,omitempty"`
	Locked bool `json:"locked,omitempty"`
	AppData interface{} `json:"appdata,omitempty"`
	Picture	string `json:"picture,omitempty"`
  LastUpdated string `json:"lastupdated,omitempty"`
  Version int `json:"version,omitempty"`
  StoreSceneState bool `json:"storescenestate,omitempty"`
  Id string `json:"-"`
}

// Returns all scenes known to the bridge
func (b *Bridge) GetScenes() ([]Scene, error) {

	var m map[string]Scene

	url, err := b.getApiPath("/scenes/")
  if err != nil {
    return nil, err
  }

	res, err := b.getResource(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &m)
	scenes := make([]Scene, 0, len(m))

	for i, g := range m {
		g.Id = i
		scenes = append(scenes, g)
	}

	return scenes, err

}

// Returns one scene by its id of i
func (b *Bridge) GetScene(i string) (*Scene, error) {

	var g *Scene

	url, err := b.getApiPath("/scenes/", i)
  if err != nil {
    return nil, err
	}
	
	res, err := b.getResource(url)
	if err != nil {
		return nil, err
	}

	fmt.Println("Hi")
	fmt.Println(string(res))

	err = json.Unmarshal(res, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}


// Updates one scene and its attributes by id of i
func (b *Bridge) UpdateScene(i int, s *Scene) (*Response, error) {
	
	var a []*ApiResponse

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/scenes/", id)
  if err != nil {
    return nil, err
	}
	
	data, err := json.Marshal(&s)
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

// Creates one new scene with its attributes defined in s
func (b *Bridge) CreateScene(s *Scene) (*Response, error) {

	var a []*ApiResponse
	
	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getApiPath("/scenes/")
  if err != nil {
    return nil, err
	}

	res, err := b.postResource(url, data)
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

// Deletes one scene from the bridge
func (b *Bridge) DeleteScene(i int) error {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/scenes/", id)
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
