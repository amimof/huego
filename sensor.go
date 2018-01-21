package huego

// Sensor represents a bridge sensor https://developers.meethue.com/documentation/sensors-api
type Sensor struct {
	State            *SensorState  `json:"state,omitempty"`
	Config           *SensorConfig `json:"config,omitempty"`
	Name             string        `json:"name,omitempty"`
	Type             string        `json:"type,omitempty"`
	ModelID          string        `json:"modelid,omitempty"`
	ManufacturerName string        `json:"manufacturername,omitempty"`
	SwVersion        string        `json:"swversion,omitemptyn"`
	ID               int           `json:",omitempty"`
}

// SensorState defines the state a sensor has
type SensorState struct {
	Daylight    string `json:"daylight,omitempty"`
	LastUpdated string `json:"lastupdated,omitempty"`
}

// SensorConfig defines the configuration of a sensor
type SensorConfig struct {
	On            bool `json:"on,omitempty"`
	Configured    bool `json:"configured,omitempty"`
	SunriseOffset int  `json:"sunriseoffset,omitempty"`
	SunsetOffset  int  `json:"sunsetoffset,omitempty"`
}

// NewSensor defines a list of sensors discovered the last time the bridge performed a sensor discovery.
// Also stores the timestamp the last time a discovery was performed.
type NewSensor struct {
	Sensors  []*Sensor
	LastScan string `json:"lastscan"`
}
