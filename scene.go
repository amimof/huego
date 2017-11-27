package huego

import (
	"encoding/json"
	"strconv"
)

type Scene struct {
	Name 	string 		`json:"name,omitempty"`
	Scenes 	[]string 	`json:"scenes,omitempty"`
	Owner 	string 		`json:"owner,omitempty"`
	Recycle 	bool `json:"recycle,omitempty"`
	Locked bool 		`json:"locked,omitempty"`
	AppData 	interface{} 		`json:"appdata,omitempty"`
	Picture 	string 	`json:"picture,omitempty"`
  LastUpdated string `json:"lastupdated,omitempty"`
  Version int `json:"version,omitempty"`
  StoreSceneState bool `json:"storescenestate,omitempty"`
  Id 		int			`json:",omitempty"`
}

// Get all scenes
func (h *Hue) GetScenes() ([]*Scene, error) {

	var m map[string]Scene

	res, err := h.GetResource(h.GetApiUrl("/scenes/"))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &m)
	scenes := make([]*Scene, 0, len(m))

	for i, g := range m {
		g.Id, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		scenes = append(scenes, &g)
	}

	return scenes, err

}

// Get one scene
func (h *Hue) GetScene(i int) (*Scene, error) {

	var g *Scene

	url := h.GetApiUrl("/scenes/", strconv.Itoa(i))
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


// Update a scene
func (h *Hue) UpdateScene(i int, s *Scene) ([]*Response, error) {
	var r []*Response

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/scenes/", id)

	data, err := json.Marshal(&s)
	if err != nil {
		return r, err
	}

	res, err := h.PutResource(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

// CreateScene creates a new scene
// See: https://developers.meethue.com/documentation/scenes-api#22_create_scene
func (h *Hue) CreateScene(s *Scene) ([]*Response, error) {

	var r []*Response

	url := h.GetApiUrl("/scenes/")

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	res, err := h.PostResource(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

// DeleteScene deletes a scene with the id of i
// See: https://developers.meethue.com/documentation/scenes-api#26_delete_scene
func (h *Hue) DeleteScene(i int) ([]*Response, error) {

	var r []*Response

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/scenes/", id)

	res, err := h.DeleteResource(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}
