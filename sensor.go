package huego

// Sensor represents a bridge sensor https://developers.meethue.com/documentation/sensors-api
type Sensor struct {
	State            map[string]interface{} `json:"state,omitempty"`
	Config           map[string]interface{} `json:"config,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Type             string                 `json:"type,omitempty"`
	ModelID          string                 `json:"modelid,omitempty"`
	ManufacturerName string                 `json:"manufacturername,omitempty"`
	SwVersion        string                 `json:"swversion,omitemptyn"`
	ID               int                    `json:",omitempty"`
}

// NewSensor defines a list of sensors discovered the last time the bridge performed a sensor discovery.
// Also stores the timestamp the last time a discovery was performed.
type NewSensor struct {
	Sensors  []*Sensor
	LastScan string `json:"lastscan"`
}
