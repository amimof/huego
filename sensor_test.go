package huego

import (
	"testing"
	"os"
)

func TestGetSensors(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
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
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
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