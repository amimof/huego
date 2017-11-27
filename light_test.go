package huego

import (
	"testing"
	"os"
)

func TestGetLights(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	lights, err := hue.GetLights()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d lights", len(lights))
	for _, light := range lights {
		t.Log(light)
	}
}

func TestGetLight(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	lights, err := hue.GetLights()
	if err != nil {
		t.Error(err)
	}
	for _, light := range lights {
		l, err := hue.GetLight(light.Id)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(l)
		}
		break
	}
}

func TestSetLight(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	lights, err := hue.GetLights()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d lights, setting the first one", len(lights))
	for _, light := range lights {
		if light.State.On {
			response, err := hue.SetLight(light.Id, *light.State)
			if err != nil {
				t.Error(err)
			}
			for _, r := range response {
				t.Logf("Response from put: Success=%v Error=%v", r.Success, r.Error)
			}
		}
		break
	}
}

func TestFindLights(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	search, err := hue.FindLights()
	if err != nil {
		t.Error(err)
	}
	for _, response := range search {
		t.Logf("Response from search: Success=%v Error=%v", response.Success, response.Error)
	}

}

func TestGetNewLights(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	newlights, err := hue.GetNewLights()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d new lights. LastScan: %s", len(newlights.Lights), newlights.LastScan)
	for _, light := range newlights.Lights {
		t.Log(light)
	}

}

func TestUpdateLight(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	lights, err := hue.GetLights()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d lights, updating the first one", len(lights))
	for _, light := range lights {
		_, err := hue.UpdateLight(light.Id, light)
		if err != nil {
			t.Error(err)
		}
		break
	}
}

func TestDeleteLight(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	res, err := hue.DeleteLight(3)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		for _, r := range res {
			t.Log(r.Success, r.Error)
		}
	}
}
