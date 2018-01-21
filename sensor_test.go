package huego_test

import (
	"github.com/amimof/huego"
	"os"
	"testing"
)

func TestGetSensors(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	sensors, err := hue.GetSensors()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d sensors", len(sensors))
	for _, sensor := range sensors {
		t.Logf("Sensor id=%d name=%s", sensor.ID, sensor.Name)
	}
}

func TestGetSensor(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	sensors, err := hue.GetSensors()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d sensors", len(sensors))
	for _, sensor := range sensors {
		t.Logf("Getting sensor %d, skipping the rest", sensor.ID)
		s, err := hue.GetSensor(sensor.ID)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Got sensor name=%s", s.Name)
		break
	}
}

func TestCreateSensor(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resp, err := hue.CreateSensor(&huego.Sensor{
		Name: "New Sensor",
	})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Sensor created")
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestFindSensors(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resp, err := hue.FindSensors()
	if err != nil {
		t.Fatal(err)
	} else {
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestGetNewSensors(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	newSensors, err := hue.GetNewSensors()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Sensors:")
	for _, sensor := range newSensors.Sensors {
		t.Logf("State:")
		t.Logf("  Daylight: %s", sensor.State.Daylight)
		t.Logf("  LastUpdated: %s", sensor.State.LastUpdated)
		t.Logf("Config:")
		t.Logf("  On: %t", sensor.Config.On)
		t.Logf("  Configured: %t", sensor.Config.Configured)
		t.Logf("  SunriseOffset: %d", sensor.Config.SunriseOffset)
		t.Logf("  SunsetOffset: %d", sensor.Config.SunsetOffset)
		t.Logf("Name: %s", sensor.Name)
		t.Logf("Type: %s", sensor.Type)
		t.Logf("ModelID: %s", sensor.ModelID)
		t.Logf("ManufacturerName: %s", sensor.ManufacturerName)
		t.Logf("SwVersion: %s", sensor.SwVersion)
		t.Logf("ID: %d", sensor.ID)
	}
}

func TestUpdateSensor(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	resp, err := hue.UpdateSensor(id, &huego.Sensor{
		Name: "New Sensor",
	})
	if err != nil {
		t.Fatal(err)
	} else {
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestDeleteSensor(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	err := hue.DeleteSensor(id)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Sensor %d deleted", id)
	}
}
