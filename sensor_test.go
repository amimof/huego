package huego_test

import (
	"testing"
	"os"
	"github.com/amimof/huego"
)

func TestGetSensors(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	sensors, err := hue.GetSensors()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d sensors", len(sensors))
	for _, sensor := range sensors {
		t.Logf("Sensor id=%d name=%s", sensor.Id, sensor.Name)
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
		t.Logf("Getting sensor %d, skipping the rest", sensor.Id)
		s, err := hue.GetSensor(sensor.Id)
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
		t.Logf("  SunriseOffset: %s", sensor.Config.SunriseOffset)
		t.Logf("  SunsetOffset: %s", sensor.Config.SunsetOffset)
		t.Logf("Name: %s", sensor.Name)
		t.Logf("Type: %s", sensor.Type)
		t.Logf("ModelId: %s", sensor.ModelId)
		t.Logf("ManufacturerName: %s", sensor.ManufacturerName)
		t.Logf("SwVersion: %s", sensor.SwVersion)
		t.Logf("Id: %d", sensor.Id)
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
		t.Logf("Sensor %d deleted")
	}
}
