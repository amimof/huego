package huego

import (
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
type NewSensor struct {
	Sensors []*Sensor
	LastScan string `json:"lastscan"`
}

// GetSensors will return all sensors
// See: https://developers.meethue.com/documentation/sensors-api#51_get_all_sensors
func (h *Hue) GetSensors() ([]Sensor, error) {

	s := map[string]Sensor{}
	url := h.GetApiUrl("/sensors/")

	res, err := h.GetResource(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &s)
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

	res, err := h.GetResource(url)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}

	return r, err

}

// Creates a sensor
func (h *Hue) CreateSensor(s *Sensor) ([]*Response, error) {

	var r []*Response

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url := h.GetApiUrl("/sensors/")
	res, err := h.PostResource(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil

}

func (h *Hue) FindSensors() ([]*Response, error) {

	var r []*Response

	url := h.GetApiUrl("/sensors/")

	res, err := h.PostResource(url, nil)

	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}

	return r, nil

}

// See: https://developers.meethue.com/documentation/lights-api#12_get_new_lights
func (h *Hue) GetNewSensors() (*NewSensor, error){

	var n map[string]Sensor
	var result *NewSensor

	url := h.GetApiUrl("/sensors/new")
	res, err := h.GetResource(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &n)
	sensors := make([]*Sensor, 0, len(n))

	for i, l := range n {
		if i != "lastscan" {
			l.Id, err = strconv.Atoi(i)
			if err != nil {
				return nil, err
			}
			sensors = append(sensors, &l)
		}
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	resu := &NewSensor{sensors, result.LastScan}

	return resu, nil

}

// Update a sensor
func (h *Hue) UpdateSensor(i int, sensor Sensor) ([]*Response, error) {
	var r []*Response

	data, err := json.Marshal(&sensor)
	if err != nil {
		return r, err
	}

	url := h.GetApiUrl("/sensors/", strconv.Itoa(i))
	res, err := h.PutResource(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (h *Hue) DeleteSensor(i int) ([]*Response, error) {
	var r []*Response

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/sensors/", id)

	res, err := h.DeleteResource(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (h *Hue) UpdateSensorConfig(i int, config *SensorConfig) ([]*Response, error) {
	var r []*Response

	data, err := json.Marshal(&config)
	if err != nil {
		return r, err
	}

	url := h.GetApiUrl("/sensors/", strconv.Itoa(i), "/config")
	res, err := h.PutResource(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
