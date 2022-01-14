package huego

// Light is a resource of type light: https://developers.meethue.com/develop/hue-api-v2/api-reference/#resource_light
type Light struct {
	// +required
	On *On `json:"on,omitempty"`
	// +optional
	Dimming *Dimming `json:"dimming,omitempty"`
	// +optional
	ColorTemperature *ColorTemperature `json:"color_temperature,omitempty"`
	// +optional
	Color *Color `json:"color,omitempty"`
	// +optional
	Dynamics *Dynamics `json:"dynamics,omitempty"`
	// +optional
	Alert *Alert `json:"alert,omitempty"`
	// +required
	Mode *string `json:"mode,omitempty"`
	// +optional
	Gradient *Gradient `json:"gradient,omitempty"`
	// +optional
	Effects *Effects `json:"effects,omitempty"`

	BaseResource
}

// Effects controls the visual effects of a resource
type Effects struct {
	// +required
	EffectValues []string `json:"effect_values,omitempty"`
	// +required
	Status *string `json:"status,omitempty"`
	// +required
	StatusValues []string `json:"status_values,omitempty"`
}

// Dynamics is the dynamic properties of a resource
type Dynamics struct {
	// +required
	Speed *float32 `json:"speed,omitempty"`
	// +required
	Status *string `json:"status,omitempty"`
	// +required
	StatusValues []string `json:"status_values,omitempty"`
	// +required
	SpeedValid *bool `json:"speed_valid,omitempty"`
}

// On controls the on/off state of a resource
type On struct {
	// +required
	On *bool `json:"on,omitempty"`
}

// Dimming controls the dimming properties of a resource
type Dimming struct {
	// +required
	Brightness *float32 `json:"brightness,omitempty"`
	// +optional
	MinDimLevel *float32 `json:"min_dim_level"`
}

// ColorTemperature controls the color temperature in mirek properties of a resource
type ColorTemperature struct {
	// +required
	Mirek *uint16 `json:"mirek,omitempty"`
	// +required
	MirekValid *bool `json:"mirek_valid,omitempty"`
	// +required
	MirekSchema *MirekSchema `json:"mirek_schema,omitempty"`
}

// MirekSchema is the color temperature mirek schema
type MirekSchema struct {
	// +required
	MirekMinimum *uint16 `json:"mirek_minimum,omitempty"`
	// +required
	MirekMaximum *uint16 `json:"mirek_maximum,omitempty"`
}

// Xy controls the CIE XY gamut position of a resource
type Xy struct {
	// +required
	X *float32 `json:"x,omitempty"`
	// +required
	Y *float32 `json:"y,omitempty"`
}

// Color controls the color properties of a resource
type Color struct {
	// +required
	Xy *Xy `json:"xy,omitempty"`
	// +required
	Gamut *Gamut `json:"gamut,omitempty"`
	// +required
	GamutType *string `json:"gamut_type,omitempty"`
}

// Gamut is the color gamut of a color light bulb
type Gamut struct {
	// +required
	Red *Xy `json:"red,omitempty"`
	// +required
	Green *Xy `json:"green,omitempty"`
	// +required
	Blue *Xy `json:"blue,omitempty"`
}

// Gradient controls the gradient properties of a resource
type Gradient struct {
	// +required
	Points []*Color `json:"points,omitempty"`
	// +required
	PointsCapable *int `json:"points_capable,omitempty"`
}

// Alert effects that the resource supports
type Alert struct {
	// +required
	ActionValues []string `json:"action_values,omitempty"`
}
