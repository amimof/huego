package huego

// https://developers.meethue.com/documentation/sensors-api
type Sensor struct {
	State 			 *SensorState 	`json:"state,omitempty"`
	Config 			 *SensorConfig 	`json:"config,omitempty"`
	Name 			 string 		`json:"name,omitempty"`
	Type 			 string 		`json:"type,omitempty"`
	ModelId 		 string 		`json:"modelid,omitempty"`
	ManufacturerName string 		`json:"manufacturername,omitempty"`
	SwVersion 		 string 		`json:"swversion,omitemptyn"`
	Id 				 int 			`json:",omitempty"`
}

type SensorState struct {
	Daylight 	string `json:"daylight,omitempty"`
	LastUpdated string `json:"lastupdated,omitempty"`
}

type SensorConfig struct {
	On 				bool `json:"on,omitempty"`
	Configured 		bool `json:"configured,omitempty"`
	SunriseOffset 	int  `json:"sunriseoffset,omitempty"`
	SunsetOffset  	int  `json:"sunsetoffset,omitempty"`
}

type NewSensor struct {
	Sensors []*Sensor
	LastScan string `json:"lastscan"`
}