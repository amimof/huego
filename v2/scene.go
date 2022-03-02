package huego

type Scene struct {
	Actions  []*ActionGet `json:"actions,omitempty"`
	Group    *Owner   `json:"group,omitempty"`
	Palette  *ColorPalette   `json:"palette,omitempty"`
	Speed    float64      `json:"speed,omitempty"`
	BaseResource
}

// Action is the action to be executed on recall
type ActionGet struct {
	Action *Action`json:"action,omitempty"`
	Target *Owner `json:"target,omitempty"`
}

type Action struct {
	Dimming *Dimming `json:"dimming,omitempty"`
	On      *On `json:"on,omitempty"`
	Color 	*Xy
	ColorTemperature *ColorTemperature `json:"color_temperature,omitempty"`
	Gradient *Gradient	 `json:"gradient,omitempty"`
}

// ColorPalette describes the p
type ColorPalette struct {
	Color   []*Xy      `json:"color,omitempty"`
	Dimming []*Dimming `json:"dimming,omitempty"`
	ColorTemperature []*ColorTemperature `json:"color_temperature,omitempty"`
}

// type Gradient struct {
// 	Points []*GradientPoint `json:"points,omitempty"`
// }

// type GradientPoint struct {
// 	Color *Xy `json:"color,omitempty"`
// }