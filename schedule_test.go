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
	for i, schedule := range schedules {
		t.Logf("%d:", i)
		t.Logf("  Name: %s", schedule.Name)
		t.Logf("  Description: %s", schedule.Description)
		t.Logf("  Command:")
		t.Logf("    Address: %s", schedule.Command.Address)
		t.Logf("    Method: %s", schedule.Command.Method)
		t.Logf("    Body: %s", schedule.Command.Body)
		t.Logf("  Time: %s", schedule.Time)
		t.Logf("  LocalTime: %s", schedule.LocalTime)
		t.Logf("  StartTime: %s", schedule.StartTime)
		t.Logf("  Status: %s", schedule.Status)
		t.Logf("  AutoDelete: %t", schedule.AutoDelete)
		t.Logf("  Id: %d", schedule.Id)
	}
}






func TestGetSchedule(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	schedules, err := hue.GetSchedules()
	if err != nil {
		t.Error(err)
	}
	for _, schedule := range schedules {
		s, err := hue.GetSchedule(schedule.Id)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Time: %s", s.Time)
		t.Logf("LocalTime: %s", s.LocalTime)
		t.Logf("StartTime: %s", s.StartTime)
		t.Logf("Status: %s", s.Status)
		t.Logf("AutoDelete: %s", s.AutoDelete)
		t.Logf("Id: %s", s.Id)
		break
	}
}







func TestCreateSchedule(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	command := &huego.Command{
		Address: "/api/"+os.Getenv("HUE_USERNAME")+"/lights/0",
		Body: &struct{
			on bool
		}{
			false,	
		},
		Method: "PUT",
	}
	schedule := &huego.Schedule{
		Name: "TestSchedule",
		Description: "Huego test schedule",
		Command: command,
		LocalTime: "2019-09-22T13:37:00",
	}
	response, err := hue.CreateSchedule(schedule)
	if err != nil {
		t.Error(err)
	}
	for _, r := range response {
		t.Logf("Address: %s", r.Address)
		t.Logf("Value: %s ", r.Value)
		t.Logf("Interface: %s", r.Interface)
		t.Logf("Json: %s", r.Json)
	}
}

func TestUpdateSchedule(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	schedules, err := hue.GetSchedules()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d schedules, setting the first one", len(schedules))
	for _, schedule := range schedules {
		response, err := hue.UpdateSchedule(schedule.Id, schedule)
		if err != nil {
			t.Error(err)
		}
		for i, r := range response {
			t.Logf("Param %d:", i)
			t.Logf("  Address: %s", r.Address)
			t.Logf("  Value: %s", r.Value)
			t.Logf("  Interface: %s", r.Interface)
			t.Logf("  Json: %s", r.Json)
		}
		break
	}
}

func TestDeleteSchedule(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id, err := hue.DeleteSchedule(2)
	if err != nil {
		t.Error(err)
	} 
	t.Logf("Deleted Schedule: %d", id)
}
