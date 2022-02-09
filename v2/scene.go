package huego

// Scene is a resource of type scene https://developers.meethue.com/develop/hue-api-v2/api-reference/#resource_scene_get
type Scene struct {
	Group   *Owner        `json:"group,omitempty"`
	Actions []*ActionList `json:"actions,omitempty"`
	Palette *ColorPalette `json:"palette,omitempty"`
	Speed   *float32      `json:"speed,omitempty"`
}

// ActionList is a list of actions to be executed synchronously on recall
type ActionList struct {
	Target *Owner  `json:"target,omitempty"`
	Action *Action `json:"action,omitempty"`
}

// Action is the action to be executed on recall
type Action struct {
	On               *On               `json:"on,omitempty"`
	Dimming          *Dimming          `json:"dimming,omitempty"`
	Color            *Color            `json:"color,omitempty"`
	ColorTemperature *ColorTemperature `json:"color_temperature,omitempty"`
	Gradient         *Gradient         `json:"gradient,omitempty"`
}

// Palette is a group of colors that describe the palette of colors to be used when playing dynamics
type Palette struct {
	Color *ColorPalette `json:"color,omitempty"`
}

type ColorPalette struct {
	Color   *Xy      `json:"color,omitempty"`
	Dimming *Dimming `json:"dimming,omitempty"`
}
