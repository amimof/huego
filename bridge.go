package huego

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

	client *http.Client
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
	return b.GetConfigContext(context.Background())
}

// GetConfigContext returns the bridge configuration
func (b *Bridge) GetConfigContext(ctx context.Context) (*Config, error) {

	var config *Config

	url, err := b.getAPIPath("/config/")
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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

// CreateUser creates a user by adding n to the list of whitelists in the bridge.
// The link button on the bridge must have been pressed before calling CreateUser.
func (b *Bridge) CreateUser(n string) (string, error) {
	return b.CreateUserContext(context.Background(), n)
}

// CreateUserContext creates a user by adding n to the list of whitelists in the bridge.
// The link button on the bridge must have been pressed before calling CreateUser.
func (b *Bridge) CreateUserContext(ctx context.Context, n string) (string, error) {
	wl, err := b.createUserWithContext(ctx, n, false)
	if err != nil {
		return "", err
	}

	return wl.Username, nil
}

// CreateUserWithClientKey creates a user by adding deviceType to the list of whitelisted users on the bridge.
// The link button on the bridge must have been pressed before calling CreateUser.
func (b *Bridge) CreateUserWithClientKey(deviceType string) (*Whitelist, error) {
	return b.createUserWithContext(context.Background(), deviceType, true)
}

// CreateUserWithClientKeyContext creates a user by adding deviceType to the list of whitelisted users on the bridge
// The link button on the bridge must have been pressed before calling CreateUser.
func (b *Bridge) CreateUserWithClientKeyContext(ctx context.Context, deviceType string) (*Whitelist, error) {
	return b.createUserWithContext(ctx, deviceType, true)
}

func (b *Bridge) createUserWithContext(ctx context.Context, deviceType string, generateClientKey bool) (*Whitelist, error) {

	var a []*APIResponse

	body := struct {
		DeviceType        string `json:"devicetype,omitempty"`
		GenerateClientKey bool   `json:"generateclientkey,omitempty"`
	}{deviceType, generateClientKey}

	url, err := b.getAPIPath("/")
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	res, err := b.post(ctx, url, data)
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

	wl := Whitelist{
		Name:     deviceType,
		Username: resp.Success["username"].(string),
	}

	if ck, ok := resp.Success["clientkey"]; ok {
		wl.ClientKey = ck.(string)
	} else if generateClientKey {
		return nil, errors.New("no client key was returned when requested to generate")
	}

	return &wl, nil
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
	return b.UpdateConfigContext(context.Background(), c)
}

// UpdateConfigContext updates the bridge configuration with c
func (b *Bridge) UpdateConfigContext(ctx context.Context, c *Config) (*Response, error) {

	var a []*APIResponse

	url, err := b.getAPIPath("/config/")
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&c)
	if err != nil {
		return nil, err
	}

	res, err := b.put(ctx, url, data)
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
	return b.DeleteUserContext(context.Background(), n)
}

// DeleteUserContext removes a whitelist item from whitelists on the bridge
func (b *Bridge) DeleteUserContext(ctx context.Context, n string) error {

	var a []*APIResponse

	url, err := b.getAPIPath("/config/whitelist/", n)
	if err != nil {
		return err
	}

	res, err := b.delete(ctx, url)
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
	return b.GetFullStateContext(context.Background())
}

// GetFullStateContext returns the entire bridge configuration.
func (b *Bridge) GetFullStateContext(ctx context.Context) (map[string]interface{}, error) {

	var n map[string]interface{}

	url, err := b.getAPIPath("/")
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
	return b.GetGroupsContext(context.Background())
}

// GetGroupsContext returns all groups known to the bridge
func (b *Bridge) GetGroupsContext(ctx context.Context) ([]Group, error) {

	var m map[string]Group

	url, err := b.getAPIPath("/groups/")
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
	return b.GetGroupContext(context.Background(), i)
}

// GetGroupContext returns one group known to the bridge by its id
func (b *Bridge) GetGroupContext(ctx context.Context, i int) (*Group, error) {

	g := &Group{
		ID: i,
	}

	url, err := b.getAPIPath("/groups/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
	return b.SetGroupStateContext(context.Background(), i, l)
}

// SetGroupStateContext allows for setting the state of one group, controlling the state of all lights in that group.
func (b *Bridge) SetGroupStateContext(ctx context.Context, i int, l State) (*Response, error) {

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

	res, err := b.put(ctx, url, data)
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
	return b.UpdateGroupContext(context.Background(), i, l)
}

// UpdateGroupContext updates one group known to the bridge
func (b *Bridge) UpdateGroupContext(ctx context.Context, i int, l Group) (*Response, error) {

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

	res, err := b.put(ctx, url, data)
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
	return b.CreateGroupContext(context.Background(), g)
}

// CreateGroupContext creates one new group with attributes defined by g
func (b *Bridge) CreateGroupContext(ctx context.Context, g Group) (*Response, error) {

	var a []*APIResponse

	url, err := b.getAPIPath("/groups/")
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&g)
	if err != nil {
		return nil, err
	}

	res, err := b.post(ctx, url, data)
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
	return b.DeleteGroupContext(context.Background(), i)
}

// DeleteGroupContext deletes one group with the id of i
func (b *Bridge) DeleteGroupContext(ctx context.Context, i int) error {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/groups/", id)
	if err != nil {
		return err
	}

	res, err := b.delete(ctx, url)
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
	return b.GetLightsContext(context.Background())
}

// GetLightsContext returns all lights known to the bridge
func (b *Bridge) GetLightsContext(ctx context.Context) ([]Light, error) {

	m := map[string]Light{}

	url, err := b.getAPIPath("/lights/")
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
	return b.GetLightContext(context.Background(), i)
}

// GetLightContext returns one light with the id of i
func (b *Bridge) GetLightContext(ctx context.Context, i int) (*Light, error) {

	light := &Light{
		ID: i,
	}

	url, err := b.getAPIPath("/lights/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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

// IdentifyLight allows identifying a light
func (b *Bridge) IdentifyLight(i int) (*Response, error) {
	return b.IdentifyLightContext(context.Background(), i)
}

// IdentifyLightContext allows identifying a light
func (b *Bridge) IdentifyLightContext(ctx context.Context, i int) (*Response, error) {

	var a []*APIResponse

	url, err := b.getAPIPath("/lights/", strconv.Itoa(i), "/state")
	if err != nil {
		return nil, err
	}
	res, err := b.put(ctx, url, []byte(`{"alert":"select"}`))
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

// SetLightState allows for controlling one light's state
func (b *Bridge) SetLightState(i int, l State) (*Response, error) {
	return b.SetLightStateContext(context.Background(), i, l)
}

// SetLightStateContext allows for controlling one light's state
func (b *Bridge) SetLightStateContext(ctx context.Context, i int, l State) (*Response, error) {

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
	res, err := b.put(ctx, url, data)
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
	return b.FindLightsContext(context.Background())
}

// FindLightsContext starts a search for new lights on the bridge.
// Use GetNewLights() verify if new lights have been detected.
func (b *Bridge) FindLightsContext(ctx context.Context) (*Response, error) {

	var a []*APIResponse

	url, err := b.getAPIPath("/lights/")
	if err != nil {
		return nil, err
	}

	res, err := b.post(ctx, url, nil)
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
	return b.GetNewLightsContext(context.Background())
}

// GetNewLightsContext returns a list of lights that were discovered last time FindLights() was executed.
func (b *Bridge) GetNewLightsContext(ctx context.Context) (*NewLight, error) {

	var n map[string]interface{}

	url, err := b.getAPIPath("/lights/new")
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
	return b.DeleteLightContext(context.Background(), i)
}

// DeleteLightContext deletes one lights from the bridge
func (b *Bridge) DeleteLightContext(ctx context.Context, i int) error {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/lights/", id)
	if err != nil {
		return err
	}

	res, err := b.delete(ctx, url)
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
	return b.UpdateLightContext(context.Background(), i, light)
}

// UpdateLightContext updates one light's attributes and state properties
func (b *Bridge) UpdateLightContext(ctx context.Context, i int, light Light) (*Response, error) {

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

	res, err := b.put(ctx, url, data)
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
	return b.GetResourcelinksContext(context.Background())
}

// GetResourcelinksContext returns all resourcelinks known to the bridge
func (b *Bridge) GetResourcelinksContext(ctx context.Context) ([]*Resourcelink, error) {

	var r map[string]Resourcelink

	url, err := b.getAPIPath("/resourcelinks/")
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
		sCopy := s
		resourcelinks = append(resourcelinks, &sCopy)
	}

	return resourcelinks, nil

}

// GetResourcelink returns one resourcelink by its id defined by i
func (b *Bridge) GetResourcelink(i int) (*Resourcelink, error) {
	return b.GetResourcelinkContext(context.Background(), i)
}

// GetResourcelinkContext returns one resourcelink by its id defined by i
func (b *Bridge) GetResourcelinkContext(ctx context.Context, i int) (*Resourcelink, error) {

	g := &Resourcelink{
		ID: i,
	}

	url, err := b.getAPIPath("/resourcelinks/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
	return b.CreateResourcelinkContext(context.Background(), s)
}

// CreateResourcelinkContext creates one new resourcelink on the bridge
func (b *Bridge) CreateResourcelinkContext(ctx context.Context, s *Resourcelink) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/resourcelinks/")
	if err != nil {
		return nil, err
	}

	res, err := b.post(ctx, url, data)
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
	return b.UpdateResourcelinkContext(context.Background(), i, resourcelink)
}

// UpdateResourcelinkContext updates one resourcelink with attributes defined by resourcelink
func (b *Bridge) UpdateResourcelinkContext(ctx context.Context, i int, resourcelink *Resourcelink) (*Response, error) {
	var a []*APIResponse

	data, err := json.Marshal(&resourcelink)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/resourcelinks/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.put(ctx, url, data)
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
	return b.DeleteResourcelinkContext(context.Background(), i)
}

// DeleteResourcelinkContext deletes one resourcelink with the id of i
func (b *Bridge) DeleteResourcelinkContext(ctx context.Context, i int) error {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/resourcelinks/", id)
	if err != nil {
		return err
	}

	res, err := b.delete(ctx, url)
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
	return b.GetRulesContext(context.Background())
}

// GetRulesContext returns all rules known to the bridge
func (b *Bridge) GetRulesContext(ctx context.Context) ([]*Rule, error) {

	var r map[string]Rule

	url, err := b.getAPIPath("/rules/")
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
		sCopy := s
		rules = append(rules, &sCopy)
	}

	return rules, nil

}

// GetRule returns one rule by its id of i
func (b *Bridge) GetRule(i int) (*Rule, error) {
	return b.GetRuleContext(context.Background(), i)
}

// GetRuleContext returns one rule by its id of i
func (b *Bridge) GetRuleContext(ctx context.Context, i int) (*Rule, error) {

	g := &Rule{
		ID: i,
	}

	url, err := b.getAPIPath("/rules/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
	return b.CreateRuleContext(context.Background(), s)
}

// CreateRuleContext creates one rule with attribues defined in s
func (b *Bridge) CreateRuleContext(ctx context.Context, s *Rule) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/rules/")
	if err != nil {
		return nil, err
	}

	res, err := b.post(ctx, url, data)
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
	return b.UpdateRuleContext(context.Background(), i, rule)
}

// UpdateRuleContext updates one rule by its id of i and rule configuration of rule
func (b *Bridge) UpdateRuleContext(ctx context.Context, i int, rule *Rule) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&rule)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/rules/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.put(ctx, url, data)
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
	return b.DeleteRuleContext(context.Background(), i)
}

// DeleteRuleContext deletes one rule from the bridge
func (b *Bridge) DeleteRuleContext(ctx context.Context, i int) error {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/rules/", id)
	if err != nil {
		return err
	}

	res, err := b.delete(ctx, url)
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
	return b.GetScenesContext(context.Background())
}

// GetScenesContext returns all scenes known to the bridge
func (b *Bridge) GetScenesContext(ctx context.Context) ([]Scene, error) {

	var m map[string]Scene

	url, err := b.getAPIPath("/scenes/")
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
	return b.GetSceneContext(context.Background(), i)
}

// GetSceneContext returns one scene by its id of i
func (b *Bridge) GetSceneContext(ctx context.Context, i string) (*Scene, error) {

	g := &Scene{ID: i}
	l := struct {
		LightStates map[int]State `json:"lightstates"`
	}{}

	url, err := b.getAPIPath("/scenes/", i)
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
	return b.UpdateSceneContext(context.Background(), id, s)
}

// UpdateSceneContext updates one scene and its attributes by id of i
func (b *Bridge) UpdateSceneContext(ctx context.Context, id string, s *Scene) (*Response, error) {

	var a []*APIResponse

	url, err := b.getAPIPath("/scenes/", id)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	res, err := b.put(ctx, url, data)
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
	return b.SetSceneLightStateContext(context.Background(), id, iid, l)
}

// SetSceneLightStateContext allows for setting the state of a light in a scene.
// SetSceneLightStateContext accepts the id of the scene, the id of a light associated with the scene and the state object.
func (b *Bridge) SetSceneLightStateContext(ctx context.Context, id string, iid int, l *State) (*Response, error) {

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

	res, err := b.put(ctx, url, data)
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
	return b.RecallSceneContext(context.Background(), id, gid)
}

// RecallSceneContext will recall a scene in a group identified by both scene and group identifiers
func (b *Bridge) RecallSceneContext(ctx context.Context, id string, gid int) (*Response, error) {

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

	res, err := b.put(ctx, url, data)
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
	return b.CreateSceneContext(context.Background(), s)
}

// CreateSceneContext creates one new scene with its attributes defined in s
func (b *Bridge) CreateSceneContext(ctx context.Context, s *Scene) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/scenes/")
	if err != nil {
		return nil, err
	}

	res, err := b.post(ctx, url, data)
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
	return b.DeleteSceneContext(context.Background(), id)
}

// DeleteSceneContext deletes one scene from the bridge
func (b *Bridge) DeleteSceneContext(ctx context.Context, id string) error {

	var a []*APIResponse

	url, err := b.getAPIPath("/scenes/", id)
	if err != nil {
		return err
	}

	res, err := b.delete(ctx, url)
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

// GetSchedules returns all schedules known to the bridge
func (b *Bridge) GetSchedules() ([]*Schedule, error) {
	return b.GetSchedulesContext(context.Background())
}

// GetSchedulesContext returns all schedules known to the bridge
func (b *Bridge) GetSchedulesContext(ctx context.Context) ([]*Schedule, error) {

	var r map[string]Schedule

	url, err := b.getAPIPath("/schedules/")
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
		sCopy := s
		schedules = append(schedules, &sCopy)
	}

	return schedules, nil

}

// GetSchedule returns one schedule by id defined in i
func (b *Bridge) GetSchedule(i int) (*Schedule, error) {
	return b.GetScheduleContext(context.Background(), i)
}

// GetScheduleContext returns one schedule by id defined in i
func (b *Bridge) GetScheduleContext(ctx context.Context, i int) (*Schedule, error) {

	g := &Schedule{
		ID: i,
	}

	url, err := b.getAPIPath("/schedules/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
	return b.CreateScheduleContext(context.Background(), s)
}

// CreateScheduleContext creates one schedule and sets its attributes defined in s
func (b *Bridge) CreateScheduleContext(ctx context.Context, s *Schedule) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/schedules/")
	if err != nil {
		return nil, err
	}

	res, err := b.post(ctx, url, data)
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
	return b.UpdateScheduleContext(context.Background(), i, schedule)
}

// UpdateScheduleContext updates one schedule by its id of i and attributes by schedule
func (b *Bridge) UpdateScheduleContext(ctx context.Context, i int, schedule *Schedule) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&schedule)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/schedules/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.put(ctx, url, data)
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
	return b.DeleteScheduleContext(context.Background(), i)
}

// DeleteScheduleContext deletes one schedule from the bridge by its id of i
func (b *Bridge) DeleteScheduleContext(ctx context.Context, i int) error {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/schedules/", id)
	if err != nil {
		return err
	}

	res, err := b.delete(ctx, url)
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
	return b.GetSensorsContext(context.Background())
}

// GetSensorsContext returns all sensors known to the bridge
func (b *Bridge) GetSensorsContext(ctx context.Context) ([]Sensor, error) {

	s := map[string]Sensor{}

	url, err := b.getAPIPath("/sensors/")
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
	return b.GetSensorContext(context.Background(), i)
}

// GetSensorContext returns one sensor by its id of i
func (b *Bridge) GetSensorContext(ctx context.Context, i int) (*Sensor, error) {

	r := &Sensor{
		ID: i,
	}

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/sensors/", id)
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
	return b.CreateSensorContext(context.Background(), s)
}

// CreateSensorContext creates one new sensor
func (b *Bridge) CreateSensorContext(ctx context.Context, s *Sensor) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/sensors/")
	if err != nil {
		return nil, err
	}

	res, err := b.post(ctx, url, data)
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
	return b.FindSensorsContext(context.Background())
}

// FindSensorsContext starts a search for new sensors.
// Use GetNewSensorsContext() to verify if new sensors have been discovered in the bridge.
func (b *Bridge) FindSensorsContext(ctx context.Context) (*Response, error) {

	var a []*APIResponse

	url, err := b.getAPIPath("/sensors/")
	if err != nil {
		return nil, err
	}

	res, err := b.post(ctx, url, nil)
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
	return b.GetNewSensorsContext(context.Background())
}

// GetNewSensorsContext returns a list of sensors that were discovered last time GetNewSensors() was executed.
func (b *Bridge) GetNewSensorsContext(ctx context.Context) (*NewSensor, error) {

	var n map[string]Sensor
	var result *NewSensor

	url, err := b.getAPIPath("/sensors/new")
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
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
			lCopy := l
			sensors = append(sensors, &lCopy)
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
	return b.UpdateSensorContext(context.Background(), i, sensor)
}

// UpdateSensorContext updates one sensor by its id and attributes by sensor
func (b *Bridge) UpdateSensorContext(ctx context.Context, i int, sensor *Sensor) (*Response, error) {

	var a []*APIResponse

	data, err := json.Marshal(&sensor)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/sensors/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.put(ctx, url, data)
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
	return b.DeleteSensorContext(context.Background(), i)
}

// DeleteSensorContext deletes one sensor from the bridge
func (b *Bridge) DeleteSensorContext(ctx context.Context, i int) error {

	var a []*APIResponse

	id := strconv.Itoa(i)
	url, err := b.getAPIPath("/sensors/", id)
	if err != nil {
		return err
	}

	res, err := b.delete(ctx, url)
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
	return b.UpdateSensorConfigContext(context.Background(), i, c)
}

// UpdateSensorConfigContext updates the configuration of one sensor. The allowed configuration parameters depend on the sensor type
func (b *Bridge) UpdateSensorConfigContext(ctx context.Context, i int, c interface{}) (*Response, error) {
	var a []*APIResponse

	data, err := json.Marshal(&c)
	if err != nil {
		return nil, err
	}

	url, err := b.getAPIPath("/sensors/", strconv.Itoa(i), "/config")
	if err != nil {
		return nil, err
	}

	res, err := b.put(ctx, url, data)
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
	return b.GetCapabilitiesContext(context.Background())
}

// GetCapabilitiesContext returns a list of capabilities of resources supported in the bridge.
func (b *Bridge) GetCapabilitiesContext(ctx context.Context) (*Capabilities, error) {

	s := &Capabilities{}

	url, err := b.getAPIPath("/capabilities/")
	if err != nil {
		return nil, err
	}

	res, err := b.get(ctx, url)
	if err != nil {
		return nil, err
	}

	err = unmarshal(res, &s)
	if err != nil {
		return nil, err
	}

	return s, err
}
