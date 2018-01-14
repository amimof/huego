package huego

// https://developers.meethue.com/documentation/lights-api
type Light struct {
	State *State `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	ModelId string `json:"modelid,omitempty"`
	ManufacturerName string `json:"manufacturername,omitempty"`
	UniqueId string `json:"uniqueid,omitempty"`
	SwVersion string `json:"swversion,omitempty"`
	SwConfigId string `json:"swconfigid,omitempty"`
	ProductId string `json:"productid,omitempty"`
	Id int `json:"-"`
	bridge *Bridge
}

type State struct {
	On bool	`json:"on"`
	Bri uint8	`json:"bri,omitempty"`
	Hue uint16	`json:"hue,omitempty"`
	Sat uint8 `json:"sat,omitempty"`
	Xy []float32 `json:"xy,omitempty"`
	Ct uint16 `json:"ct,omitempty"`
	Alert	string `json:"alert,omitempty"`
	Effect string `json:"effect,omitempty"`
	TransitionTime uint16 `json:"transitiontime,omitempty"`
	BriInc int `json:"bri_inc,omitempty"`
	SatInc int `json:"sat_inc,omitempty"`
	HueInc int `json:"hue_inc,omitempty"`
	CtInc int `json:"ct_inc,omitempty"`
	XyInc int `json:"xy_inc,omitempty"`
	ColorMode	string `json:"colormode,omitempty"`
	Reachable	bool `json:"reachable,omitempty"`
	Scene	string `json:"scene,omitempty"`
}

type NewLight struct {
	Lights []string
	LastScan string `json:"lastscan"`
}


// Sets the On state of one light to false, turning it off
func (l *Light) Off() error {
	state := State{ On: false }
	_, err := l.bridge.SetLight(l.Id, state)
	if err != nil {
		return err
	}
	l.State.On = false
	return nil
}

// Sets the On state of one light to true, turning it on
func (l *Light) On() error {
	state := State{ On: true }
	_, err := l.bridge.SetLight(l.Id, state)
	if err != nil {
		return err
	}	
	l.State.On = true
	return nil
}

// Returns true if light state On property is true
func (l *Light) IsOn() bool {
	return l.State.On
}

// Sets the name property of the light
func (l *Light) Rename(new string) error {
	update := Light{ Name: new }
	_, err := l.bridge.UpdateLight(l.Id, update)
	if err != nil {
		return err
	}
	l.Name = new
	return nil
}

// Sets the light brightness state property
func (l *Light) Bri(new uint8) error {
	update := State{ On: true, Bri: new }
	_, err := l.bridge.SetLight(l.Id, update)
	if err != nil {
		return err
	}
	l.State.Bri = new
	return nil
}

// Sets the light hue state property (0-65535)
func (l *Light) Hue(new uint16) error {
	update := State{ On: true, Hue: new }
	_, err := l.bridge.SetLight(l.Id, update)
	if err != nil {
		return err
	}
	l.State.Hue = new
	return nil
}

// Sets the light saturation state property (0-254)
func (l *Light) Sat(new uint8) error {
	update := State{ On: true, Sat: new }
	_, err := l.bridge.SetLight(l.Id, update)
	if err != nil {
		return err
	}
	l.State.Sat = new
	return nil
}

// Sets the x and y coordinates of a color in CIE color space. (0-1 per value)
func (l *Light) Xy(new []float32) error {
	update := State{ On: true, Xy: new }
	_, err := l.bridge.SetLight(l.Id, update)
	if err != nil {
		return err
	}
	l.State.Xy = new
	return nil
}

// Sets the light color temperature state property
func (l *Light) Ct(new uint16) error {
	update := State{ On: true, Ct: new }
	_, err := l.bridge.SetLight(l.Id, update)
	if err != nil {
		return err
	}
	l.State.Ct = new
	return nil
}

// Sets the duration of the transition from the light’s current state to the new state
func (l *Light) TransitionTime(new uint16) error {
	update := State{ On: l.State.On, TransitionTime: new }
	_, err := l.bridge.SetLight(l.Id, update)
	if err != nil {
		return err
	}
	l.State.TransitionTime = new
	return nil
}

// The dynamic effect of the light, currently “none” and “colorloop” are supported
func (l *Light) Effect(new string) error {
	update := State{ On: true, Effect: new }
	_, err := l.bridge.SetLight(l.Id, update)
	if err != nil {
		return err
	}
	l.State.Effect = new
	return nil
}

// Makes the light blink in its current color. Supported values are:
// “none” – The light is not performing an alert effect.
// “select” – The light is performing one breathe cycle.
// “lselect” – The light is performing breathe cycles for 15 seconds or until alert is set to "none".
func (l *Light) Alert(new string) error {
	update := State{ On: true, Alert: new }
	_, err := l.bridge.SetLight(l.Id, update)
	if err != nil {
		return err
	}
	l.State.Effect = new
	return nil
}

