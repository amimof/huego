package huego_test

import (
	"testing"
	"os"
	"github.com/amimof/huego"
)

func TestGetSchedules(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	schedules, err := hue.GetSchedules()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d schedules", len(schedules))
	for _, schedule := range schedules {
		t.Log(schedule)
	}
}


func TestGetSchedule(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	schedules, err := hue.GetSchedules()
	if err != nil {
		t.Error(err)
	}
	for _, schedule := range schedules {
		l, err := hue.GetSchedule(schedule.Id)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(l)
		}
		break
	}
}

func TestCreateSchedule(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	command := &huego.Command{
		Address: "/api/"+os.Getenv("HUE_USERNAME")+"/lights/0",
		Body: &huego.State{On: false},
		Method: "PUT",
	}
	schedule := &huego.Schedule{
		Name: "TestSchedule",
		Description: "Huego test schedule",
		Command: command,
		LocalTime: "2017-09-22T13:37:00",
	}
	response, err := hue.CreateSchedule(schedule)
	if err != nil {
		t.Error(err)
	}
	for _, r := range response {
		t.Logf("Address: %s Value: %s Interface: %s", r.Address, r.Value, r.Interface)
	}
}

func TestUpdateSchedule(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	schedules, err := hue.GetSensors()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d schedules, setting the first one", len(schedules))
	for _, schedule := range schedules {
		response, err := hue.UpdateSensor(schedule.Id, schedule)
		if err != nil {
			t.Error(err)
		}
		for _, r := range response {
			t.Logf("Address: %s Value: %s Interface: %s", r.Address, r.Value, r.Interface)
		}
		break
	}
}

func TestDeleteSchedule(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	res, err := hue.DeleteSchedule(1)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		for _, r := range res {
			t.Logf("Address: %s Value: %s Interface: %s", r.Address, r.Value, r.Interface)
		}
	}
}
