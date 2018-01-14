package huego

import (
	"fmt"
	"net/url"
	"path"
	"strings"
	"encoding/json"
	"strconv"
)

type Bridge struct {
	Host string `json:"internalipaddress,omitempty"` 
	User string
	Id string `json:"id,omitempty"`
}

func (b *Bridge) getApiPath(str ...string) (string, error) {

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

// Calls New() and passes Host on this Bridge instance
func (b *Bridge) Login(u string) *Bridge {
	return New(b.Host, u)
}

/*

	CONFIGURATION API

*/

// Returns the bridge configuration
func (b *Bridge) GetConfig() (*Config, error) {
  
  var config *Config

  url, err := b.getApiPath("/config/")
  res, err := get(url)
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(res, &config)
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

// Creates a user by adding n to the list of whitelists in the bridge
func (b *Bridge) CreateUser(n string) (string, error) {

  var a []*ApiResponse

  body := struct {
    DeviceType string `json:"devicetype,omitempty"`
    GenerateClientKey bool `json:"generateclientkey,omitempty"`
  }{
    n,
    true,
  }

  url, err := b.getApiPath("/")
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

  err = json.Unmarshal(res, &a)
  if err != nil {
    return "", err
  }

  resp, err := handleResponse(a)
  if err != nil {
    return "", err
  }

  return resp.Success["username"].(string), nil

}

// Returns a list of whitelists from the bridge
func (b *Bridge) GetUsers() ([]Whitelist, error) {
  c, err := b.GetConfig()
  if err != nil {
    return nil, err
  }
  return c.Whitelist, nil
}

// Updates the bridge configuration with c
func (b *Bridge) UpdateConfig(c *Config) (*Response, error) {

	var a []*ApiResponse

  url, err := b.getApiPath("/config/")
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

// Removes a whitelist item from whitelists on the bridge
func (b *Bridge) DeleteUser(n string) error {

  var a []*ApiResponse

  url, err := b.getApiPath("/config/whitelist/", n)
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

  return nil

}

// Get full state (datastore)
func (b *Bridge) GetFullState() (*Datastore, error) {

    var ds *Datastore

    url, err := b.getApiPath("/")
    if err != nil {
      return nil, err
    }

    res, err := get(url)
    if err != nil {
      return nil, err
    }

    err = json.Unmarshal(res, &ds)
    if err != nil {
      return nil, err
    }

    return ds, nil
}

/*

	GROUP API

*/

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
		g.bridge = b
		groups = append(groups, g)
	}

	return groups, err

}

// Returns one group known to the bridge by its id
func (b *Bridge) GetGroup(i int) (*Group, error) {

	g := &Group{
		Id: i,
	}

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

	g.bridge = b

	return g, nil
}


// Allows for setting the state of one group, controlling the state of all lights in that group.
func (b *Bridge) SetGroupState(i int, l State) (*Response, error) {

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
func (b *Bridge) UpdateGroup(i int, l Group) (*Response, error) {
	
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

/*

	LIGHT API

*/

// Returns all lights known to the bridge
func (b *Bridge) GetLights() ([]Light, error) {

	m := map[string]Light{}

	url, err := b.getApiPath("/lights/")
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

	lights := make([]Light, 0, len(m))

	for i, l := range m {
		l.Id, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		l.bridge = b
		lights = append(lights, l)
	}

	return lights, nil

}

// Returns one light with the id of i
func (b *Bridge) GetLight(i int) (*Light, error) {

	light := &Light{
		Id: i,
	}

	url, err := b.getApiPath("/lights/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return light, err
	}

	err = json.Unmarshal(res, &light)
	if err != nil {
		return light, err
	}

	light.bridge = b

	return light, nil
}

// Allows for controlling one light's state
func (b *Bridge) SetLight(i int, l State) (*Response, error) {

	var a []*ApiResponse

	l.Reachable = false
	l.ColorMode = ""

	data, err := json.Marshal(&l)
	if err != nil {
		return nil, err
	}

	url, err := b.getApiPath("/lights/", strconv.Itoa(i), "/state")
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

// Starts a search for new lights on the bridge. 
// Use GetNewLights() verify if new lights have been detected. 
func (b *Bridge) FindLights() (*Response, error) {

	var a []*ApiResponse

	url, err := b.getApiPath("/lights/")
	if err != nil {
		return nil, err
	}

	res, err := post(url, nil)
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

// Returns a list of lights that were discovered last time FindLights() was executed.
func (b *Bridge) GetNewLights() (*NewLight, error){

	var n map[string]interface{}
	
	url, err := b.getApiPath("/lights/new")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(res, &n)

	lights := make([]string, 0, len(n))
	var lastscan string

	for k, _ := range n {
		if k == "lastscan" {
			lastscan = n[k].(string)
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

// Deletes one lights from the bridge
func (b *Bridge) DeleteLight(i int) error {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/lights/", id)
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

	return nil

}

// Updates one light's attributes and state properties
func (b *Bridge) UpdateLight(i int, light Light) (*Response, error) {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/lights/", id)
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

/*

	RESOURCELINK API

*/

// Returns all resourcelinks known to the bridge
func (b *Bridge) GetResourcelinks() ([]*Resourcelink, error) {

  var r map[string]Resourcelink

  url, err := b.getApiPath("/resourcelinks/")
  if err != nil {
    return nil, err
  }

  res, err := get(url)
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(res, &r)
  if err != nil {
    return nil, err
  }

  resourcelinks := make([]*Resourcelink, 0, len(r))

  for i, s := range r {
    s.Id, err = strconv.Atoi(i)
    if err != nil {
      return nil, err
    }
    resourcelinks = append(resourcelinks, &s)
  }

  return resourcelinks, nil

}

// Returns one resourcelink by its id defined by i
func (b *Bridge) GetResourcelink(i int) (*Resourcelink, error) {

	var resourcelink *Resourcelink

  url, err := b.getApiPath("/resourcelinks/", strconv.Itoa(i))


	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &resourcelink)
	if err != nil {
		return nil, err
	}

	return resourcelink, nil

}

// Creates one new resourcelink on the bridge
func (b *Bridge) CreateResourcelink(s *Resourcelink) (*Response, error) {

  var a []*ApiResponse

  data, err := json.Marshal(&s)
  if err != nil {
    return nil, err
  }

  url, err := b.getApiPath("/resourcelinks/")
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

// Updates one resourcelink with attributes defined by resourcelink
func (b *Bridge) UpdateResourcelink(i int, resourcelink *Resourcelink) (*Response, error) {
	var a []*ApiResponse

	data, err := json.Marshal(&resourcelink)
	if err != nil {
		return nil, err
	}

  url, err := b.getApiPath("/resourcelinks/", strconv.Itoa(i))
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

// Deletes one resourcelink with the id of i
func (b *Bridge) DeleteResourcelink(i int) error {
  
  var a []*ApiResponse

	id := strconv.Itoa(i)
  url, err := b.getApiPath("/resourcelinks/", id)
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

	return nil
}

/*

	RULE API

*/

// Returns all rules known to the bridge
func (b *Bridge) GetRules() ([]*Rule, error) {

  var r map[string]Rule

  url, err := b.getApiPath("/rules/")
  if err != nil {
    return nil, err
  }

  res, err := get(url)
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(res, &r)
  if err != nil {
    return nil, err
  }

  rules := make([]*Rule, 0, len(r))

  for i, s := range r {
    s.Id, err = strconv.Atoi(i)
    if err != nil {
      return nil, err
    }
    rules = append(rules, &s)
  }

  return rules, nil

}

// Returns one rule by its id of i
func (b *Bridge) GetRule(i int) (*Rule, error) {

	var rule *Rule

  url, err := b.getApiPath("/rules/", strconv.Itoa(i))
  if err != nil {
    return nil, err
  }

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &rule)
	if err != nil {
		return nil, err
	}

	return rule, nil

}

// Creates one rule with attribues defined in s
func (b *Bridge) CreateRule(s *Rule) (*Response, error) {

  var a []*ApiResponse

  data, err := json.Marshal(&s)
  if err != nil {
    return nil, err
  }

  url, err := b.getApiPath("/rules/")
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

// Updates one rule by its id of i and rule configuration of rule
func (b *Bridge) UpdateRule(i int, rule *Rule) (*Response, error) {
  
  var a []*ApiResponse

	data, err := json.Marshal(&rule)
	if err != nil {
		return nil, err
	}

  url, err := b.getApiPath("/rules/", strconv.Itoa(i))
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

// Deletes one rule from the bridge
func (b *Bridge) DeleteRule(i int) error {
  
  var a []*ApiResponse

	id := strconv.Itoa(i)
  url, err := b.getApiPath("/rules/", id)
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

  return nil
}

/*

	SCENE API

*/

// Returns all scenes known to the bridge
func (b *Bridge) GetScenes() ([]Scene, error) {

	var m map[string]Scene

	url, err := b.getApiPath("/scenes/")
  if err != nil {
    return nil, err
  }

	res, err := get(url)
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

// Deletes one scene from the bridge
func (b *Bridge) DeleteScene(i int) error {

	var a []*ApiResponse

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/scenes/", id)
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

	return nil
}


/*

	SCHEDULE API

*/

// Returns all scehdules known to the bridge
func (b *Bridge) GetSchedules() ([]*Schedule, error) {

  var r map[string]Schedule

  url, err := b.getApiPath("/schedules/")
  if err != nil {
    return nil, err
  }

  res, err := get(url)
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(res, &r)
  if err != nil {
    return nil, err
  }

  schedules := make([]*Schedule, 0, len(r))

  for i, s := range r {
    s.Id, err = strconv.Atoi(i)
    if err != nil {
      return nil, err
    }
    schedules = append(schedules, &s)
  }

  return schedules, nil

}

// Returns one schedule by id defined in i
func (b *Bridge) GetSchedule(i int) (*Schedule, error) {

	var schedule *Schedule

  url, err := b.getApiPath("/schedules/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &schedule)
	if err != nil {
		return nil, err
	}

	return schedule, nil

}

// Creates one schedule and sets its attributes defined in s
func (b *Bridge) CreateSchedule(s *Schedule) (*Response, error) {

  var a []*ApiResponse

  data, err := json.Marshal(&s)
  if err != nil {
    return nil, err
  }

  url, err := b.getApiPath("/schedules/")
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

// Updates one schedule by its id of i and attributes by schedule
func (b *Bridge) UpdateSchedule(i int, schedule *Schedule) (*Response, error) {
  
  var a []*ApiResponse

	data, err := json.Marshal(&schedule)
	if err != nil {
		return nil, err
	}

  url, err := b.getApiPath("/schedules/", strconv.Itoa(i))
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

// Deletes one schedule from the bridge by its id of i
func (b *Bridge) DeleteSchedule(i int) error {
  
  var a []*ApiResponse

	id := strconv.Itoa(i)
  url, err := b.getApiPath("/schedules/", id)
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

	return nil
}

/*

	SENSOR API

*/

// Returns all sensors known to the bridge
func (b *Bridge) GetSensors() ([]Sensor, error) {

	s := map[string]Sensor{}
	
	url, err := b.getApiPath("/sensors/")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &s)
	if err != nil {
		return nil, err
	}

	sensors := make([]Sensor, 0, len(s))

	for i, k := range s {
		k.Id, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, k)
	}
	return sensors, err
}

// Returns one sensor by its id of i
func (b *Bridge) GetSensor(i int) (*Sensor, error) {

	var r *Sensor

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/sensors/", id)
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}

	return r, err

}

// Creates one new sensor
func (b *Bridge) CreateSensor(s *Sensor) (*Response, error) {

	var a []*ApiResponse

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getApiPath("/sensors/")
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

// Starts a search for new sensors.
// Use GetNewSensors() to verify if new sensors have been discovered in the bridge. 
func (b *Bridge) FindSensors() (*Response, error) {

	var a []*ApiResponse

	url, err := b.getApiPath("/sensors/")
	if err != nil {
		return nil, err
	}

	res, err := post(url, nil)
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

// Returns a list of sensors that were discovered last time GetNewSensors() was executed.
func (b *Bridge) GetNewSensors() (*NewSensor, error){

	var n map[string]Sensor
	var result *NewSensor

	url, err := b.getApiPath("/sensors/new")
	if err != nil {
		return nil, err
	}

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &n)
	sensors := make([]*Sensor, 0, len(n))

	for i, l := range n {
		if i != "lastscan" {
			l.Id, err = strconv.Atoi(i)
			if err != nil {
				return nil, err
			}
			sensors = append(sensors, &l)
		}
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	resu := &NewSensor{sensors, result.LastScan}

	return resu, nil

}

// Updates one sensor by its id and attributes by sensor
func (b *Bridge) UpdateSensor(i int, sensor *Sensor) (*Response, error) {
	
	var a []*ApiResponse

	data, err := json.Marshal(&sensor)
	if err != nil {
		return nil, err
	}

	url, err := b.getApiPath("/sensors/", strconv.Itoa(i))
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

// Deletes one sensor from the bridge
func (b *Bridge) DeleteSensor(i int) error {
	
	var a []*ApiResponse

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/sensors/", id)
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

	return nil
}

// Updates the configuration of one sensor. The allowed configuration parameters depend on the sensor type
func (b *Bridge) UpdateSensorConfig(i int, config *SensorConfig) (*Response, error) {
	var a []*ApiResponse

	data, err := json.Marshal(&config)
	if err != nil {
		return nil, err
	}

	url, err := b.getApiPath("/sensors/", strconv.Itoa(i), "/config")
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
