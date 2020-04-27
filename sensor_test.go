package huego

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSensors(t *testing.T) {
	b := New(hostname, username)
	sensors, err := b.GetSensors()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d sensors", len(sensors))
	for _, sensor := range sensors {
		t.Logf("State:")
		t.Logf("  Interface: %+v", sensor.State)
		t.Logf("Config:")
		t.Logf("  On: %+v", sensor.Config)
		t.Logf("Name: %s", sensor.Name)
		t.Logf("Type: %s", sensor.Type)
		t.Logf("ModelID: %s", sensor.ModelID)
		t.Logf("ManufacturerName: %s", sensor.ManufacturerName)
		t.Logf("UniqueID: %s", sensor.UniqueID)
		t.Logf("SwVersion: %s", sensor.SwVersion)
		t.Logf("ID: %d", sensor.ID)
	}
	contains := func(name string, ss []Sensor) bool {
		for _, s := range ss {
			if s.Name == name {
				return true
			}
		}
		return false
	}

	assert.True(t, contains("Daylight", sensors))
	assert.True(t, contains("Tap Switch 2", sensors))
}

func TestGetSensor(t *testing.T) {
	b := New(hostname, username)
	sensor, err := b.GetSensor(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("State:")
	t.Logf("  Interface: %+v", sensor.State)
	t.Logf("Config:")
	t.Logf("  On: %+v", sensor.Config)
	t.Logf("Name: %s", sensor.Name)
	t.Logf("Type: %s", sensor.Type)
	t.Logf("ModelID: %s", sensor.ModelID)
	t.Logf("ManufacturerName: %s", sensor.ManufacturerName)
	t.Logf("UniqueID: %s", sensor.UniqueID)
	t.Logf("SwVersion: %s", sensor.SwVersion)
	t.Logf("ID: %d", sensor.ID)
}

func TestCreateSensor(t *testing.T) {
	b := New(hostname, username)
	resp, err := b.CreateSensor(&Sensor{
		Name: "New Sensor",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Sensor created")
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}

}

func TestFindSensors(t *testing.T) {
	b := New(hostname, username)
	resp, err := b.FindSensors()
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}

}

func TestGetNewSensors(t *testing.T) {
	b := New(hostname, username)
	newSensors, err := b.GetNewSensors()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Sensors:")
	for _, sensor := range newSensors.Sensors {
		t.Logf("State:")
		t.Logf("  Interface: %+v", sensor.State)
		t.Logf("Config:")
		t.Logf("  On: %+v", sensor.Config)
		t.Logf("Name: %s", sensor.Name)
		t.Logf("Type: %s", sensor.Type)
		t.Logf("ModelID: %s", sensor.ModelID)
		t.Logf("ManufacturerName: %s", sensor.ManufacturerName)
		t.Logf("UniqueID: %s", sensor.UniqueID)
		t.Logf("SwVersion: %s", sensor.SwVersion)
		t.Logf("ID: %d", sensor.ID)
	}

	contains := func(name string, ss []*Sensor) bool {
		for _, s := range ss {
			if s.Name == name {
				return true
			}
		}
		return false
	}

	assert.True(t, contains("Hue Tap 1", newSensors.Sensors))
	assert.True(t, contains("Button 3", newSensors.Sensors))
}

func TestUpdateSensor(t *testing.T) {
	b := New(hostname, username)
	id := 1
	resp, err := b.UpdateSensor(id, &Sensor{
		Name: "New Sensor",
	})
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}
}

func TestUpdateSensorConfig(t *testing.T) {
	b := New(hostname, username)
	id := 1
	resp, err := b.UpdateSensorConfig(id, "test")
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}

}

func TestDeleteSensor(t *testing.T) {
	b := New(hostname, username)
	id := 1
	err := b.DeleteSensor(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Sensor %d deleted", id)
}
