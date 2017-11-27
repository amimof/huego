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
		break
	}
}

func TestCreateSensor(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	newSensor := &Sensor{Name: "TestSensor"}
	response, err := hue.CreateSensor(newSensor)
	if err != nil {
		t.Error(err)
	}
	for _, r := range response {
		t.Logf("Response from put: Success=%v Error=%v", r.Success, r.Error)
	}
}

func TestFindSensors(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	result, err := hue.FindSensors()
	if err != nil {
		t.Error(err)
	}
	for _, r := range result {
		t.Logf("Response from search: Success=%v Error=%v", r.Success, r.Error)
	}
}


func TestGetNewSensors(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	newSensors, err := hue.GetNewSensors()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d new sensors. LastScan: %s", len(newSensors.Sensors), newSensors.LastScan)
	for _, sensor := range newSensors.Sensors {
		t.Log(sensor)
	}
}


func TestUpdateSensor(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	sensors, err := hue.GetSensors()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d sensors, setting the first one", len(sensors))
	for _, sensor := range sensors {
		response, err := hue.UpdateSensor(sensor.Id, sensor)
		if err != nil {
			t.Error(err)
		}
		for _, r := range response {
			t.Logf("Response from put: Success=%v Error=%v", r.Success, r.Error)
		}
		break
	}
}

// func TestDeleteSensor(t *testing.T) {
// 	res, err := hue.DeleteSensor(1)
// 	if err != nil {
// 		t.Log(err)
// 		t.Fail()
// 	} else {
// 		for _, r := range res {
// 			t.Log(r.Success, r.Error)
// 		}
// 	}
// }
