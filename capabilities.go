package huego

// Capabilities holds a combined model of resource capabilities on the bridge: https://developers.meethue.com/documentation/lights-api
type Capabilities struct {
	Groups        Capability `json:"groups,omitempty"`
	Lights        Capability `json:"lights,omitempty"`
	Resourcelinks Capability `json:"resourcelinks,omitempty"`
	Schedules     Capability `json:"schedules,omitempty"`
	Rules         Capability `json:"rules,omitempty"`
	Scenes        Capability `json:"scenes,omitempty"`
	Sensors       Capability `json:"sensors,omitempty"`
	Streaming     Capability `json:"streaming,omitempty"`
}

// Capability defines the resource and subresource capabilities.
type Capability struct {
	Available int `json:"available,omitempty"`
}
