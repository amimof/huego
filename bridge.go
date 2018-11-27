package huego

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"
)

// Bridge exposes a hardware bridge through a struct.
type Bridge struct {
	Host string `json:"internalipaddress,omitempty"`
	User string
	ID   string `json:"id,omitempty"`
}

func (b *Bridge) getAPIPath(str ...string) (string, error) {

	if strings.Index(strings.ToLower(b.Host), "http://") <= -1 && strings.Index(strings.ToLower(b.Host), "https://") <= -1 {
		b.Host = fmt.Sprintf("%s%s", "http://", b.Host)
	}

	u, err := url.Parse(b.Host)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "/api/", b.User)
	for _, p := range str {
		u.Path = path.Join(u.Path, p)
	}
	return u.String(), nil
}

// Login calls New() and passes Host on this Bridge instance.
func (b *Bridge) Login(u string) *Bridge {
	b.User = u
	return New(b.Host, u)
}

/*

	CONFIGURATION API

*/

// GetConfig returns the bridge configuration
func (b *Bridge) GetConfig() (*Config, error) {

	var config *Config

	url, err := b.getAPIPath("/config/")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &config)
	if err != nil {
		return nil, err
	}

	wl := make([]Whitelist, 0, len(config.WhitelistMap))
	for k, v := range config.WhitelistMap {
		v.Username = k
		wl = append(wl, v)
	}

	config.Whitelist = wl

	return config, nil

}

// CreateUser creates a user by adding n to the list of whitelists in the bridge
func (b *Bridge) CreateUser(n string) (string, error) {

	var a []*APIResponse

	body := struct {
		DeviceType        string `json:"devicetype,omitempty"`
		GenerateClientKey bool   `json:"generateclientkey,omitempty"`
	}{n, true}

	url, err := b.getAPIPath("/")
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(&body)
	if err != nil {
		return "", err
	}

	res, err := post(url, data)
	if err != nil {
		return "", err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return "", err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return "", err
	}

	return resp.Success["username"].(string), nil

}

// GetUsers returns a list of whitelists from the bridge
func (b *Bridge) GetUsers() ([]Whitelist, error) {
	c, err := b.GetConfig()
	if err != nil {
		return nil, err
	}
	return c.Whitelist, nil
}

// UpdateConfig updates the bridge configuration with c
func (b *Bridge) UpdateConfig(c *Config) (*Response, error) {

	var a []*APIResponse

	url, err := b.getAPIPath("/config/")
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&c)
	if err != nil {
		return nil, err
	}

	res, err := put(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteUser removes a whitelist item from whitelists on the bridge
func (b *Bridge) DeleteUser(n string) error {

	var a []*APIResponse

	url, err := b.getAPIPath("/config/whitelist/", n)
	if err != nil {
		return err
	}

	res, err := delete(url)
	if err != nil {
		return err
	}

	_ = unmarshal(res, &a)

	_, err = handleResponse(a)
	if err != nil {
		return err
	}

	return nil

}

// GetFullState returns the entire bridge configuration.
func (b *Bridge) GetFullState() (map[string]interface{}, error) {

	var n map[string]interface{}

	url, err := b.getAPIPath("/")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &n)
	if err != nil {
		return nil, err
	}

	return n, nil
}

/*

	GROUP API

*/

// GetGroups returns all groups known to the bridge
func (b *Bridge) GetGroups() ([]Group, error) {

	var m map[string]Group

	url, err := b.getAPIPath("/groups/")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &m)
	if err != nil {
		return nil, err
	}

	groups := make([]Group, 0, len(m))

	for i, g := range m {
		g.ID, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		g.bridge = b
		groups = append(groups, g)
	}

	return groups, err

}

// GetGroup returns one group known to the bridge by its id
func (b *Bridge) GetGroup(i int) (*Group, error) {

	g := &Group{
		ID: i,
	}

	url, err := b.getAPIPath("/groups/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &g)
	if err != nil {
		return nil, err
	}

	g.bridge = b

	return g, nil
}

// SetGroupState allows for setting the state of one group, controlling the state of all lights in that group.
func (b *Bridge) SetGroupState(i int, l State) (*Response, error) {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/groups/", id, "/action/")
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

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateGroup updates one group known to the bridge
func (b *Bridge) UpdateGroup(i int, l Group) (*Response, error) {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/groups/", id)
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

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateGroup creates one new group with attributes defined by g
func (b *Bridge) CreateGroup(g Group) (*Response, error) {

	var a []*APIResponse

	url, err := b.getAPIPath("/groups/")
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

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteGroup deletes one group with the id of i
func (b *Bridge) DeleteGroup(i int) error {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/groups/", id)
	if err != nil {
		return err
	}

	res, err := delete(url)
	if err != nil {
		return err
	}

	_ = unmarshal(res, &a)

	_, err = handleResponse(a)
	if err != nil {
		return err
	}

	return nil
}

/*

	LIGHT API

*/

// GetLights returns all lights known to the bridge
func (b *Bridge) GetLights() ([]Light, error) {

	m := map[string]Light{}

	url, err := b.getAPIPath("/lights/")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &m)
	if err != nil {
		return nil, err
	}

	lights := make([]Light, 0, len(m))

	for i, l := range m {
		l.ID, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		l.bridge = b
		lights = append(lights, l)
	}

	return lights, nil

}

// GetLight returns one light with the id of i
func (b *Bridge) GetLight(i int) (*Light, error) {

	light := &Light{
		ID: i,
	}

	url, err := b.getAPIPath("/lights/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return light, err
	}

	err = unmarshal(res, &light)
	if err != nil {
		return light, err
	}

	light.bridge = b

	return light, nil
}

// SetLightState allows for controlling one light's state
func (b *Bridge) SetLightState(i int, l State) (*Response, error) {

	var a []*APIResponse

	l.Reachable = false
	l.ColorMode = ""

	data, err := json.Marshal(&l)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/lights/", strconv.Itoa(i), "/state")
	if err != nil {
		return nil, err
	}
	res, err := put(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// FindLights starts a search for new lights on the bridge.
// Use GetNewLights() verify if new lights have been detected.
func (b *Bridge) FindLights() (*Response, error) {

	var a []*APIResponse

	url, err := b.getAPIPath("/lights/")
	if err != nil {
		return nil, err
	}

	res, err := post(url, nil)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// GetNewLights returns a list of lights that were discovered last time FindLights() was executed.
func (b *Bridge) GetNewLights() (*NewLight, error) {

	var n map[string]interface{}

	url, err := b.getAPIPath("/lights/new")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	_ = unmarshal(res, &n)

	lights := make([]string, 0, len(n))
	var lastscan string

	for k := range n {
		if k == "lastscan" {
			lastscan = n[k].(string)
		} else {
			lights = append(lights, k)
		}
	}

	result := &NewLight{
		Lights:   lights,
		LastScan: lastscan,
	}

	return result, nil

}

// DeleteLight deletes one lights from the bridge
func (b *Bridge) DeleteLight(i int) error {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/lights/", id)
	if err != nil {
		return err
	}

	res, err := delete(url)
	if err != nil {
		return err
	}

	_ = unmarshal(res, &a)

	_, err = handleResponse(a)
	if err != nil {
		return err
	}

	return nil

}

// UpdateLight updates one light's attributes and state properties
func (b *Bridge) UpdateLight(i int, light Light) (*Response, error) {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/lights/", id)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&light)
	if err != nil {
		return nil, err
	}

	res, err := put(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*

	RESOURCELINK API

*/

// GetResourcelinks returns all resourcelinks known to the bridge
func (b *Bridge) GetResourcelinks() ([]*Resourcelink, error) {

	var r map[string]Resourcelink

	url, err := b.getAPIPath("/resourcelinks/")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}

	resourcelinks := make([]*Resourcelink, 0, len(r))

	for i, s := range r {
		s.ID, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		resourcelinks = append(resourcelinks, &s)
	}

	return resourcelinks, nil

}

// GetResourcelink returns one resourcelink by its id defined by i
func (b *Bridge) GetResourcelink(i int) (*Resourcelink, error) {

	g := &Resourcelink{
		ID: i,
	}

	url, err := b.getAPIPath("/resourcelinks/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &g)
	if err != nil {
		return nil, err
	}

	return g, nil

}

// CreateResourcelink creates one new resourcelink on the bridge
func (b *Bridge) CreateResourcelink(s *Resourcelink) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/resourcelinks/")
	if err != nil {
		return nil, err
	}

	res, err := post(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// UpdateResourcelink updates one resourcelink with attributes defined by resourcelink
func (b *Bridge) UpdateResourcelink(i int, resourcelink *Resourcelink) (*Response, error) {
	var a []*APIResponse

	data, err := json.Marshal(&resourcelink)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/resourcelinks/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := put(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteResourcelink deletes one resourcelink with the id of i
func (b *Bridge) DeleteResourcelink(i int) error {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/resourcelinks/", id)
	if err != nil {
		return err
	}

	res, err := delete(url)
	if err != nil {
		return err
	}

	_ = unmarshal(res, &a)

	_, err = handleResponse(a)
	if err != nil {
		return err
	}

	return nil
}

/*

	RULE API

*/

// GetRules returns all rules known to the bridge
func (b *Bridge) GetRules() ([]*Rule, error) {

	var r map[string]Rule

	url, err := b.getAPIPath("/rules/")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}

	rules := make([]*Rule, 0, len(r))

	for i, s := range r {
		s.ID, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		rules = append(rules, &s)
	}

	return rules, nil

}

// GetRule returns one rule by its id of i
func (b *Bridge) GetRule(i int) (*Rule, error) {

	g := &Rule{
		ID: i,
	}

	url, err := b.getAPIPath("/rules/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &g)
	if err != nil {
		return nil, err
	}

	return g, nil

}

// CreateRule creates one rule with attribues defined in s
func (b *Bridge) CreateRule(s *Rule) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/rules/")
	if err != nil {
		return nil, err
	}

	res, err := post(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// UpdateRule updates one rule by its id of i and rule configuration of rule
func (b *Bridge) UpdateRule(i int, rule *Rule) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&rule)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/rules/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := put(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteRule deletes one rule from the bridge
func (b *Bridge) DeleteRule(i int) error {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/rules/", id)
	if err != nil {
		return err
	}

	res, err := delete(url)
	if err != nil {
		return err
	}

	_ = unmarshal(res, &a)

	_, err = handleResponse(a)
	if err != nil {
		return err
	}

	return nil
}

/*

	SCENE API

*/

// GetScenes returns all scenes known to the bridge
func (b *Bridge) GetScenes() ([]Scene, error) {

	var m map[string]Scene

	url, err := b.getAPIPath("/scenes/")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &m)
	scenes := make([]Scene, 0, len(m))

	for i, g := range m {
		g.ID = i
		g.bridge = b
		scenes = append(scenes, g)
	}

	return scenes, err

}

// GetScene returns one scene by its id of i
func (b *Bridge) GetScene(i string) (*Scene, error) {

	g := &Scene{ID: i}
	l := struct {
		LightStates map[int]State `json:"lightstates"`
	}{}

	url, err := b.getAPIPath("/scenes/", i)
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &l)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &g)
	if err != nil {
		return nil, err
	}

	g.bridge = b

	return g, nil
}

// UpdateScene updates one scene and its attributes by id of i
func (b *Bridge) UpdateScene(id string, s *Scene) (*Response, error) {

	var a []*APIResponse

	url, err := b.getAPIPath("/scenes/", id)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	res, err := put(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetSceneLightState allows for setting the state of a light in a scene.
// SetSceneLightState accepts the id of the scene, the id of a light associated with the scene and the state object.
func (b *Bridge) SetSceneLightState(id string, iid int, l *State) (*Response, error) {

	var a []*APIResponse

	lightid := strconv.Itoa(iid)
	url, err := b.getAPIPath("scenes", id, "lightstates", lightid)
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

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RecallScene will recall a scene in a group identified by both scene and group identifiers
func (b *Bridge) RecallScene(id string, gid int) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(struct {
		Scene string `json:"scene"`
	}{id})

	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/groups/", strconv.Itoa(gid), "/action")
	if err != nil {
		return nil, err
	}

	res, err := put(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// CreateScene creates one new scene with its attributes defined in s
func (b *Bridge) CreateScene(s *Scene) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/scenes/")
	if err != nil {
		return nil, err
	}

	res, err := post(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteScene deletes one scene from the bridge
func (b *Bridge) DeleteScene(id string) error {

	var a []*APIResponse

	url, err := b.getAPIPath("/scenes/", id)
	if err != nil {
		return err
	}

	res, err := delete(url)
	if err != nil {
		return err
	}

	_ = unmarshal(res, &a)

	_, err = handleResponse(a)
	if err != nil {
		return err
	}

	return nil
}

/*

	SCHEDULE API

*/

// GetSchedules returns all scehdules known to the bridge
func (b *Bridge) GetSchedules() ([]*Schedule, error) {

	var r map[string]Schedule

	url, err := b.getAPIPath("/schedules/")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}

	schedules := make([]*Schedule, 0, len(r))

	for i, s := range r {
		s.ID, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, &s)
	}

	return schedules, nil

}

// GetSchedule returns one schedule by id defined in i
func (b *Bridge) GetSchedule(i int) (*Schedule, error) {

	g := &Schedule{
		ID: i,
	}

	url, err := b.getAPIPath("/schedules/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &g)
	if err != nil {
		return nil, err
	}

	return g, nil

}

// CreateSchedule creates one schedule and sets its attributes defined in s
func (b *Bridge) CreateSchedule(s *Schedule) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/schedules/")
	if err != nil {
		return nil, err
	}

	res, err := post(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// UpdateSchedule updates one schedule by its id of i and attributes by schedule
func (b *Bridge) UpdateSchedule(i int, schedule *Schedule) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&schedule)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/schedules/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := put(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteSchedule deletes one schedule from the bridge by its id of i
func (b *Bridge) DeleteSchedule(i int) error {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/schedules/", id)
	if err != nil {
		return err
	}

	res, err := delete(url)
	if err != nil {
		return err
	}

	_ = unmarshal(res, &a)

	_, err = handleResponse(a)
	if err != nil {
		return err
	}

	return nil
}

/*

	SENSOR API

*/

// GetSensors returns all sensors known to the bridge
func (b *Bridge) GetSensors() ([]Sensor, error) {

	s := map[string]Sensor{}

	url, err := b.getAPIPath("/sensors/")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &s)
	if err != nil {
		return nil, err
	}

	sensors := make([]Sensor, 0, len(s))

	for i, k := range s {
		k.ID, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, k)
	}
	return sensors, err
}

// GetSensor returns one sensor by its id of i
func (b *Bridge) GetSensor(i int) (*Sensor, error) {

	r := &Sensor{
		ID: i,
	}

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/sensors/", id)
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return r, err
	}

	err = unmarshal(res, &r)
	if err != nil {
		return r, err
	}

	return r, err

}

// CreateSensor creates one new sensor
func (b *Bridge) CreateSensor(s *Sensor) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/sensors/")
	if err != nil {
		return nil, err
	}

	res, err := post(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// FindSensors starts a search for new sensors.
// Use GetNewSensors() to verify if new sensors have been discovered in the bridge.
func (b *Bridge) FindSensors() (*Response, error) {

	var a []*APIResponse

	url, err := b.getAPIPath("/sensors/")
	if err != nil {
		return nil, err
	}

	res, err := post(url, nil)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// GetNewSensors returns a list of sensors that were discovered last time GetNewSensors() was executed.
func (b *Bridge) GetNewSensors() (*NewSensor, error) {

	var n map[string]Sensor
	var result *NewSensor

	url, err := b.getAPIPath("/sensors/new")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	_ = unmarshal(res, &n)

	sensors := make([]*Sensor, 0, len(n))

	for i, l := range n {
		if i != "lastscan" {
			l.ID, err = strconv.Atoi(i)
			if err != nil {
				return nil, err
			}
			sensors = append(sensors, &l)
		}
	}

	err = unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	resu := &NewSensor{sensors, result.LastScan}

	return resu, nil

}

// UpdateSensor updates one sensor by its id and attributes by sensor
func (b *Bridge) UpdateSensor(i int, sensor *Sensor) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&sensor)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/sensors/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := put(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteSensor deletes one sensor from the bridge
func (b *Bridge) DeleteSensor(i int) error {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/sensors/", id)
	if err != nil {
		return err
	}

	res, err := delete(url)
	if err != nil {
		return err
	}

	_ = unmarshal(res, &a)

	_, err = handleResponse(a)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSensorConfig updates the configuration of one sensor. The allowed configuration parameters depend on the sensor type
func (b *Bridge) UpdateSensorConfig(i int, c interface{}) (*Response, error) {
	var a []*APIResponse

	data, err := json.Marshal(&c)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/sensors/", strconv.Itoa(i), "/config")
	if err != nil {
		return nil, err
	}

	res, err := put(url, data)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*

	CAPABILITIES API

*/

// GetCapabilities returns a list of capabilities of resources supported in the bridge.
func (b *Bridge) GetCapabilities() (*Capabilities, error) {

	s := &Capabilities{}

	url, err := b.getAPIPath("/capabilities/")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &s)
	if err != nil {
		return nil, err
	}

	return s, err
}
