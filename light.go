package huego

// Light represents a bridge light https://developers.meethue.com/documentation/lights-api
type Light struct {
	State            *State `json:"state,omitempty"`
	Type             string `json:"type,omitempty"`
	Name             string `json:"name,omitempty"`
	ModelID          string `json:"modelid,omitempty"`
	ManufacturerName string `json:"manufacturername,omitempty"`
	UniqueID         string `json:"uniqueid,omitempty"`
	SwVersion        string `json:"swversion,omitempty"`
	SwConfigID       string `json:"swconfigid,omitempty"`
	ProductID        string `json:"productid,omitempty"`
	ID               int    `json:"-"`
	bridge           *Bridge
}

// State defines the attributes and properties of a light
type State struct {
	On             bool      `json:"on"`
	Bri            uint8     `json:"bri,omitempty"`
	Hue            uint16    `json:"hue,omitempty"`
	Sat            uint8     `json:"sat,omitempty"`
	Xy             []float32 `json:"xy,omitempty"`
	Ct             uint16    `json:"ct,omitempty"`
	Alert          string    `json:"alert,omitempty"`
	Effect         string    `json:"effect,omitempty"`
	TransitionTime uint16    `json:"transitiontime,omitempty"`
	BriInc         int       `json:"bri_inc,omitempty"`
	SatInc         int       `json:"sat_inc,omitempty"`
	HueInc         int       `json:"hue_inc,omitempty"`
	CtInc          int       `json:"ct_inc,omitempty"`
	XyInc          int       `json:"xy_inc,omitempty"`
	ColorMode      string    `json:"colormode,omitempty"`
	Reachable      bool      `json:"reachable,omitempty"`
	Scene          string    `json:"scene,omitempty"`
}

// NewLight defines a list of lights discovered the last time the bridge performed a light discovery.
// Also stores the timestamp the last time a discovery was performed.
type NewLight struct {
	Lights   []string
	LastScan string `json:"lastscan"`
}

// SetState sets the state of the light to s.
func (l *Light) SetState(s State) error {
	_, err := l.bridge.SetLightState(l.ID, s)
	if err != nil {
		return err
	}
	l.State = &s
	return nil
}

// Off sets the On state of one light to false, turning it off
func (l *Light) Off() error {
	state := State{On: false}
	_, err := l.bridge.SetLightState(l.ID, state)
	if err != nil {
		return err
	}
	l.State.On = false
	return nil
}

// On sets the On state of one light to true, turning it on
func (l *Light) On() error {
	state := State{On: true}
	_, err := l.bridge.SetLightState(l.ID, state)
	if err != nil {
		return err
	}
	l.State.On = true
	return nil
}

// IsOn returns true if light state On property is true
func (l *Light) IsOn() bool {
	return l.State.On
}

// Rename sets the name property of the light
func (l *Light) Rename(new string) error {
	update := Light{Name: new}
	_, err := l.bridge.UpdateLight(l.ID, update)
	if err != nil {
		return err
	}
	l.Name = new
	return nil
}

// Bri sets the light brightness state property
func (l *Light) Bri(new uint8) error {
	update := State{On: true, Bri: new}
	_, err := l.bridge.SetLightState(l.ID, update)
	if err != nil {
		return err
	}
	l.State.Bri = new
	return nil
}

// Hue sets the light hue state property (0-65535)
func (l *Light) Hue(new uint16) error {
	update := State{On: true, Hue: new}
	_, err := l.bridge.SetLightState(l.ID, update)
	if err != nil {
		return err
	}
	l.State.Hue = new
	return nil
}

// Sat sets the light saturation state property (0-254)
func (l *Light) Sat(new uint8) error {
	update := State{On: true, Sat: new}
	_, err := l.bridge.SetLightState(l.ID, update)
	if err != nil {
		return err
	}
	l.State.Sat = new
	return nil
}

// Xy sets the x and y coordinates of a color in CIE color space. (0-1 per value)
func (l *Light) Xy(new []float32) error {
	update := State{On: true, Xy: new}
	_, err := l.bridge.SetLightState(l.ID, update)
	if err != nil {
		return err
	}
	l.State.Xy = new
	return nil
}

// Ct sets the light color temperature state property
func (l *Light) Ct(new uint16) error {
	update := State{On: true, Ct: new}
	_, err := l.bridge.SetLightState(l.ID, update)
	if err != nil {
		return err
	}
	l.State.Ct = new
	return nil
}

// TransitionTime sets the duration of the transition from the light’s current state to the new state
func (l *Light) TransitionTime(new uint16) error {
	update := State{On: l.State.On, TransitionTime: new}
	_, err := l.bridge.SetLightState(l.ID, update)
	if err != nil {
		return err
	}
	l.State.TransitionTime = new
	return nil
}

// Effect the dynamic effect of the light, currently “none” and “colorloop” are supported
func (l *Light) Effect(new string) error {
	update := State{On: true, Effect: new}
	_, err := l.bridge.SetLightState(l.ID, update)
	if err != nil {
		return err
	}
	l.State.Effect = new
	return nil
}

// Alert makes the light blink in its current color. Supported values are:
// “none” – The light is not performing an alert effect.
// “select” – The light is performing one breathe cycle.
// “lselect” – The light is performing breathe cycles for 15 seconds or until alert is set to "none".
func (l *Light) Alert(new string) error {
	update := State{On: true, Alert: new}
	_, err := l.bridge.SetLightState(l.ID, update)
	if err != nil {
		return err
	}
	l.State.Effect = new
	return nil
}
