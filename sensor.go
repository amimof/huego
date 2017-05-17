package huego 

import (
	"net/http"
	"encoding/json"
	"strconv"
)

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

// GetSensors will return all sensors
// See: https://developers.meethue.com/documentation/sensors-api#51_get_all_sensors
func (h *Hue) GetSensors() ([]Sensor, error) {

	s := map[string]Sensor{}
	url := h.GetApiUrl("/sensors/")

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	sensors := make([]Sensor, 0, len(s))

	for i, k := range s {
		k.Id, err = strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, k)
	}
	return sensors, err
}

// GetSensor returns a sensor with the id of i
// See: https://developers.meethue.com/documentation/sensors-api#55_get_sensor
func (h *Hue) GetSensor(i int) (*Sensor, error) {

	var r *Sensor
	id := strconv.Itoa(i)
	url := h.GetApiUrl("/sensors/", id)

	res, err := http.Get(url)
	if err != nil {
		return r, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return r, err
	}

	return r, err

}