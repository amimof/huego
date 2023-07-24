package huego

import (
	"encoding/json"
)

type Scene struct {
	Actions []*ActionGet `json:"actions,omitempty"`
	Group   *Owner       `json:"group,omitempty"`
	Palette *Palette     `json:"palette,omitempty"`
	Speed   *float64     `json:"speed,omitempty"`
	BaseResource
}

// Action is the action to be executed on recall
type ActionGet struct {
	Action *Action `json:"action,omitempty"`
	Target *Owner  `json:"target,omitempty"`
}

type Action struct {
	Dimming          *Dimming `json:"dimming,omitempty"`
	On               *On      `json:"on,omitempty"`
	Color            *Color
	ColorTemperature *ColorTemperature `json:"color_temperature,omitempty"`
	Gradient         *Gradient         `json:"gradient,omitempty"`
}

// Palette describes the p
type Palette struct {
	Color            []*PaletteColor            `json:"color,omitempty"`
	ColorTemperature []*PaletteColorTemperature `json:"color_temperature,omitempty"`
	Dimming          []*Dimming                 `json:"dimming,omitempty"`
}

type PaletteColor struct {
	Color   *Color   `json:"color,omitempty"`
	Dimming *Dimming `json:"dimming,omitempty"`
}

type PaletteColorTemperature struct {
	ColorTemperature *ColorTemperature `json:"color_temperature,omitempty"`
	Dimming          *Dimming          `json:"dimming,omitempty"`
}

func NewDimming(bri float32) *Dimming {
	return &Dimming{Brightness: &bri}
}

// Raw marshals the scene into a byte array. Returns nil if errors occur on the way
func (s *Scene) Raw() []byte {
	d, err := json.Marshal(s)
	if err != nil {
		return nil
	}
	return d
}