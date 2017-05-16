package huego

import (
	"testing"
)

var (
	hue *Hue
)

func init() {
	hue = New("http://192.168.1.80/")
}

func TestGetLights(t *testing.T) {
	lights, err := hue.GetLights()
	if err != nil {
		t.Fail()
	}
	for _, light := range lights {
		t.Log(light)
	}
}

func TestGetLight(t *testing.T) {
	lights, err := hue.GetLights()
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		for _, light := range lights {
			l, err := hue.GetLight(light.Id)
			if err != nil {
				t.Log(err)
				t.Fail()
			} else {
				t.Log(l)
			}
		}
	}
}

func TestSetLight(t *testing.T) {
	// state := State{
	// 	On: true,
	// 	Bri: 254,
	// 	Hue: 65515,
	// 	Sat: 254,
	// 	Xy: []float32{0.6909, 0.308},
	// 	Ct: 153,
	// }

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
			break
		}
		
	}
	
}

func TestSearch(t *testing.T) {
	search, err := hue.Search() 
	if err != nil {
		t.Error(err) 
	}
	for _, response := range search {
		t.Logf("Response from search: Success=%v Error=%v", response.Success, response.Error)
	}
	
}

func TestGetNewLights(t *testing.T) {
	newlights, err := hue.GetNewLights()
	if err != nil {
		t.Error(err)
	} 
	t.Logf("Found %d new lights. LastScan: %s", len(*newlights.Lights), newlights.LastScan)
	for _, light := range *newlights.Lights {
		t.Log(light)		
	}
	
}

func TestRenameLight(t *testing.T) {
	
	lights, err := hue.GetLights()
	if err != nil {
		t.Error(err)
	}
	
	t.Logf("Found %d lights, renaming the first one", len(lights))

	for _, light := range lights {
		
		oriName := light.Name
		newName := "Huego Test Lamp Name"
		
		_, err := hue.RenameLight(light.Id, newName)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Renamed light %d to %s", light.Id, newName)

		_, err = hue.RenameLight(light.Id, oriName)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Renamed light %d back to %s", light.Id, oriName)
		break
	}	
	
}

// func TestDeleteLight(t *testing.T) {
// 	res, err := hue.DeleteLight(3)
// 	if err != nil {
// 		t.Log(err)
// 		t.Fail()
// 	} else {
// 		for _, r := range res {
// 			t.Log(r.Success, r.Error)
// 		}
// 	}
// }