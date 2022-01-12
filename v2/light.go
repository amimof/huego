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
	Dynamics *struct {
		// +required
		Speed *float32 `json:"speed,omitempty"`
		// +required
		Status *string `json:"status,omitempty"`
		// +required
		StatusValues []*string `json:"status_values,omitempty"`
		// +required
		SpeedValid *bool `json:"speed_valid,omitempty"`
	}
	// +optional
	Alert *AlertEffectType `json:"alert,omitempty"`
	// +required
	Mode *string `json:"mode,omitempty`
	// +optional
	Gradient *Gradient `json:"gradient,omitempty"`

	BaseResource
}

// On controlls the on/off state of a resource
type On struct {
	// +required
	On *bool `json:"on,omitempty"`
}

// Dimming controlls the dimming properties of a resource
type Dimming struct {
	// +required
	Brightness *float32 `json:"brightness,omitempty"`
	// +optional
	MinDimLevel *float32 `json:"min_dim_level"`
}

// ColorTemperature controlls the color temperature in mirek properties of a resource
type ColorTemperature struct {
	// +required
	Mirek *uint16 `json:"mirek,omitempty"`
	// +required
	MirekValid *bool `json:"mirek_valid,omitempty"`
	// +required
	MirekSchema struct {
		// +required
		MirekMinimum *uint16 `json:"mirek_minimum,omitempty"`
		// +required
		MirekMaximum *uint16 `json:"mirek_maximum,omitempty"`
	}
}

// Xy controlls the CIE XY gamut position of a resource
type Xy struct {
	// +required
	X *float32 `json:"x,omitempty"`
	// +required
	Y *float32 `json:"y,omitempty"`
}

// Color controlls the color properties of a resource
type Color struct {
	// +required
	Xy *Xy `json:"xy,omitempty"`
	// +required
	Gamut *struct {
		// +required
		Red *Xy `json:"red,omitempty"`
		// +required
		Green *Xy `json:"green,omitempty"`
		// +required
		Blue *Xy `json:"blue,omitempty"`
	}
	// +required
	GamutType *string `json:"gamut_type,omitempty"`
}

// Gradient controlls the gradient properties of a resource
type Gradient struct {
	// +required
	Points []*Color `json:"points,omitempty"`
	// +required
	PointsCapable *int `json:"points_capable,omitempty"`
}

// AlertEffectType controlls alert effects that the resource supports
type AlertEffectType struct {
	// +required
	ActionValues []*string `json:"action_values,omitempty"`
}
