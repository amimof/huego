package huego_test

import (
	"testing"
	"os"
	"github.com/amimof/huego"
)

func TestGetLights(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	lights, err := hue.GetLights()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d lights", len(lights))
	for _, l := range lights {
		t.Logf("Id: %d", l.Id)
		t.Logf("  Type: %s", l.Type)
		t.Logf("  Name: %s", l.Name)
		t.Logf("  ModelId: %s", l.ModelId)
		t.Logf("  ManufacturerName: %s", l.ManufacturerName)
		t.Logf("  UniqueId: %s", l.UniqueId)
		t.Logf("  SwVersion: %s", l.SwVersion)
		t.Logf("  SwConfigId: %s", l.SwConfigId)
		t.Logf("  ProductId: %s", l.ProductId)
	}
}

func TestGetLight(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	lights, err := hue.GetLights()
	if err != nil {
		t.Error(err)
	}
	for _, light := range lights {
		l, err := hue.GetLight(light.Id)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Id: %d", l.Id)
			t.Logf("Type: %s", l.Type)
			t.Logf("Name: %s", l.Name)
			t.Logf("ModelId: %s", l.ModelId)
			t.Logf("ManufacturerName: %s", l.ManufacturerName)
			t.Logf("UniqueId: %s", l.UniqueId)
			t.Logf("SwVersion: %s", l.SwVersion)
			t.Logf("SwConfigId: %s", l.SwConfigId)
			t.Logf("ProductId: %s", l.ProductId)
			t.Logf("State:")
			t.Logf("  On: %t", l.State.On)
			t.Logf("  Bri: %d", l.State.Bri)
			t.Logf("  Hue: %d", l.State.Hue)
			t.Logf("  Sat: %d", l.State.Sat)
			t.Logf("  Xy: %b", l.State.Xy)
			t.Logf("  Ct: %d", l.State.Ct)
			t.Logf("  Alert: %s", l.State.Alert)
			t.Logf("  Effect: %s", l.State.Effect)
			t.Logf("  TransitionTime: %d", l.State.TransitionTime)
			t.Logf("  BriInc: %d", l.State.BriInc)
			t.Logf("  SatInc: %d", l.State.SatInc)
			t.Logf("  HueInc: %d", l.State.HueInc)
			t.Logf("  CtInc: %d", l.State.CtInc)
			t.Logf("  XyInc: %d", l.State.XyInc)
			t.Logf("  ColorMode: %s", l.State.ColorMode)
			t.Logf("  Reachable: %t", l.State.Reachable)
		}
		break
	}
}

func TestSetLight(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	lights, err := hue.GetLights()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d lights, setting the first one", len(lights))
	for _, light := range lights {
		light.State.On = true
		light.State.Bri = 254
		response, err := hue.SetLight(light.Id, light.State)
		if err != nil {
			t.Error(err)
		}
		for k, r := range response {
			t.Logf("%d", k)
			t.Logf("  Address: %s", r.Address)
			t.Logf("  Value: %s", r.Value)
			t.Logf("  Interface: %s", r.Interface)
		}
		break
	}
}

func TestFindLights(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	search, err := hue.FindLights()
	if err != nil {
		t.Error(err)
	}
	for _, r := range search {
		t.Logf("Address: %s Value: %s Interface: %s", r.Address, r.Value, r.Interface)
	}

}

func TestGetNewLights(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	newlights, err := hue.GetNewLights()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d new lights", len(newlights.Lights))
	t.Logf("LastScan: %s", newlights.LastScan)
	for _, light := range newlights.Lights {
		t.Log(light)
	}

}

func TestUpdateLight(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
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
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	res, err := hue.DeleteLight(3)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		for _, r := range res {
			t.Logf("Address: %s Value: %s Interface: %s", r.Address, r.Value, r.Interface)
		}
	}
}
