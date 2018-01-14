package huego

// https://developers.meethue.com/documentation/groups-api
type Group struct {
	Name string `json:"name,omitempty"`
	Lights []string `json:"lights,omitempty"`
	Type string `json:"type,omitempty"`
	GroupState	*GroupState `json:"state,omitempty"`
	Recycle bool `json:"recycle,omitempty"`
	Class	string `json:"class,omitempty"`
	State *State `json:"action,omitempty"`
	Id int `json:"-"`
	bridge *Bridge
}

type GroupState struct {
	AllOn bool `json:"all_on,omitempty"`
	AnyOn bool `json:"any_on,omitempty"`
}

// Sets the name property of the group
func (g *Group) Rename(new string) error {
	update := Group{ Name: new }
	_, err := g.bridge.UpdateGroup(g.Id, update)
	if err != nil {
		return err
	}
	g.Name = new
	return nil
}

// Sets the On state of one group to false, turning all lights in the group off
func (g *Group) Off() error {
	state := State{ On: false }
	_, err := g.bridge.SetGroupState(g.Id, state)
	if err != nil {
		return err
	}
	g.State.On = false
	return nil
}

// Sets the On state of one group to true, turning all lights in the group on
func (g *Group) On() error {
	state := State{ On: true }
	_, err := g.bridge.SetGroupState(g.Id, state)
	if err != nil {
		return err
	}	
	g.State.On = true
	return nil
}

// Returns true if light state On property is true
func (g *Group) IsOn() bool {
	return g.State.On
}

// Sets the light brightness state property
func (g *Group) Bri(new uint8) error {
	update := State{ On: true, Bri: new }
	_, err := g.bridge.SetGroupState(g.Id, update)
	if err != nil {
		return err
	}
	g.State.Bri = new
	return nil
}

// Sets the light hue state property (0-65535)
func (g *Group) Hue(new uint16) error {
	update := State{ On: true, Hue: new }
	_, err := g.bridge.SetGroupState(g.Id, update)
	if err != nil {
		return err
	}
	g.State.Hue = new
	return nil
}

// Sets the light saturation state property (0-254)
func (g *Group) Sat(new uint8) error {
	update := State{ On: true, Sat: new }
	_, err := g.bridge.SetGroupState(g.Id, update)
	if err != nil {
		return err
	}
	g.State.Sat = new
	return nil
}

// Sets the x and y coordinates of a color in CIE color space. (0-1 per value)
func (g *Group) Xy(new []float32) error {
	update := State{ On: true, Xy: new }
	_, err := g.bridge.SetGroupState(g.Id, update)
	if err != nil {
		return err
	}
	g.State.Xy = new
	return nil
}

// Sets the light color temperature state property
func (g *Group) Ct(new uint16) error {
	update := State{ On: true, Ct: new }
	_, err := g.bridge.SetGroupState(g.Id, update)
	if err != nil {
		return err
	}
	g.State.Ct = new
	return nil
}

// Sets the scene by it's identifier of the scene you wish to recall
func (g *Group) Scene(scene string) error {
	update := State{ On: true, Scene: scene }
	_, err := g.bridge.SetGroupState(g.Id, update)
	if err != nil {
		return err
	}
	g.State.Scene = scene
	return nil
}


// Sets the duration of the transition from the light’s current state to the new state
func (g *Group) TransitionTime(new uint16) error {
	update := State{ On: g.State.On, TransitionTime: new }
	_, err := g.bridge.SetGroupState(g.Id, update)
	if err != nil {
		return err
	}
	g.State.TransitionTime = new
	return nil
}


// The dynamic effect of the lights in the group, currently “none” and “colorloop” are supported
func (g *Group) Effect(new string) error {
	update := State{ On: true, Effect: new }
	_, err := g.bridge.SetGroupState(g.Id, update)
	if err != nil {
		return err
	}
	g.State.Effect = new
	return nil
}

// Makes the lights in the group blink in its current color. Supported values are:
// “none” – The light is not performing an alert effect.
// “select” – The light is performing one breathe cycle.
// “lselect” – The light is performing breathe cycles for 15 seconds or until alert is set to "none".
func (g *Group) Alert(new string) error {
	update := State{ On: true, Alert: new }
	_, err := g.bridge.SetGroupState(g.Id, update)
	if err != nil {
		return err
	}
	g.State.Effect = new
	return nil
}

