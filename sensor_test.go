package huego_test

import (
	"github.com/amimof/huego"
	"testing"
)

func TestGetSensors(t *testing.T) {
	b := huego.New(hostname, username)
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
}

func TestGetSensor(t *testing.T) {
	b := huego.New(hostname, username)
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
	b := huego.New(hostname, username)
	resp, err := b.CreateSensor(&huego.Sensor{
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
	b := huego.New(hostname, username)
	resp, err := b.FindSensors()
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}

}

func TestGetNewSensors(t *testing.T) {
	b := huego.New(hostname, username)
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
}

func TestUpdateSensor(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	resp, err := b.UpdateSensor(id, &huego.Sensor{
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
	b := huego.New(hostname, username)
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
	b := huego.New(hostname, username)
	id := 1
	err := b.DeleteSensor(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Sensor %d deleted", id)
}
