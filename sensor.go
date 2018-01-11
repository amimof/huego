package huego

import (
	"encoding/json"
	"strconv"
)

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

// Returns all sensors known to the bridge
func (b *Bridge) GetSensors() ([]Sensor, error) {

	s := map[string]Sensor{}
	
	url, err := b.getApiPath("/sensors/")
	if err != nil {
		return nil, err
	}

	res, err := b.getResource(url)
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

// Returns one sensor by its id of i
func (b *Bridge) GetSensor(i int) (*Sensor, error) {

	var r *Sensor

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/sensors/", id)
	if err != nil {
		return nil, err
	}

	res, err := b.getResource(url)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}

	return r, err

}

// Creates one new sensor
func (b *Bridge) CreateSensor(s *Sensor) (*Response, error) {

	var a []*ApiResponse

	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	url, err := b.getApiPath("/sensors/")
	if err != nil {
		return nil, err
	}

	res, err := b.postResource(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// Starts a search for new sensors.
// Use GetNewSensors() to verify if new sensors have been discovered in the bridge. 
func (b *Bridge) FindSensors() (*Response, error) {

	var a []*ApiResponse

	url, err := b.getApiPath("/sensors/")
	if err != nil {
		return nil, err
	}

	res, err := b.postResource(url, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// Returns a list of sensors that were discovered last time GetNewSensors() was executed.
func (b *Bridge) GetNewSensors() (*NewSensor, error){

	var n map[string]Sensor
	var result *NewSensor

	url, err := b.getApiPath("/sensors/new")
	if err != nil {
		return nil, err
	}

	res, err := b.getResource(url)
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

// Updates one sensor by its id and attributes by sensor
func (b *Bridge) UpdateSensor(i int, sensor *Sensor) (*Response, error) {
	
	var a []*ApiResponse

	data, err := json.Marshal(&sensor)
	if err != nil {
		return nil, err
	}

	url, err := b.getApiPath("/sensors/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.putResource(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Deletes one sensor from the bridge
func (b *Bridge) DeleteSensor(i int) error {
	
	var a []*ApiResponse

	id := strconv.Itoa(i)
	url, err := b.getApiPath("/sensors/", id)
	if err != nil {
		return err
	}

	res, err := b.deleteResource(url)
	if err != nil {
		return err
	}

	_ = json.Unmarshal(res, &a)

	_, err = handleResponse(a)
	if err != nil {
		return err
	}

	return nil
}

// Updates the configuration of one sensor. The allowed configuration parameters depend on the sensor type
func (b *Bridge) UpdateSensorConfig(i int, config *SensorConfig) (*Response, error) {
	var a []*ApiResponse

	data, err := json.Marshal(&config)
	if err != nil {
		return nil, err
	}

	url, err := b.getApiPath("/sensors/", strconv.Itoa(i), "/config")
	if err != nil {
		return nil, err
	}
	
	res, err := b.putResource(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &a)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponse(a)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
