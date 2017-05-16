package huego

import (
	"testing"
)

var (
	hue *Hue
)

func init() {
	hue = New("http://192.168.1.80/")
}

func TestGetSensors(t *testing.T) {
	sensors, err := hue.GetSensors()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d sensors", len(sensors))
	for _, sensor := range sensors {
		t.Logf("Sensor id=%d name=%s", sensor.Id, sensor.Name)
	}
}

func TestGetSensor(t *testing.T) {
	sensors, err := hue.GetSensors()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d sensors", len(sensors))
	for _, sensor := range sensors {
		t.Logf("Getting sensor %d, skipping the rest", sensor.Id)
		s, err := hue.GetSensor(sensor.Id)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Got sensor name=%s", s.Name)
	}
}