package huego

import (
	"testing"
	"os"
)

func TestGetSchedules(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
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
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
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
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	command := &Command{
		Address: "/api/"+os.Getenv("HUE_USERNAME")+"/lights/0",
		Body: &State{On: false},
		Method: "PUT",
	}
	schedule := &Schedule{
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
		t.Logf("Response from put: Success=%v Error=%v", r.Success, r.Error)
	}
}

// func TestDeleteSchedule(t *testing.T) {
// 	res, err := hue.DeleteSchedule(1)
// 	if err != nil {
// 		t.Log(err)
// 		t.Fail()
// 	} else {
// 		for _, r := range res {
// 			t.Log(r.Success, r.Error)
// 		}
// 	}
// }
