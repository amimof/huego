package huego_test

import (
	"testing"
	"os"
	"github.com/amimof/huego"
)

func TestGetLights(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	lights, err := b.GetLights()
	if err != nil {
		t.Fatal(err)
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
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	lights, err := b.GetLights()
	if err != nil {
		t.Fatal(err)
	}
	for _, light := range lights {
		l, err := b.GetLight(light.Id)
		if err != nil {
			t.Fatal(err)
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
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	resp, err := b.SetLight(id, huego.State{
		On: true,
		Bri: 254,
	})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Light %d state updated", id)
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestFindLights(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resp, err := b.FindLights()
	if err != nil {
		t.Fatal(err)
	} else {
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}

}

func TestGetNewLights(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	newlights, err := b.GetNewLights()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d new lights", len(newlights.Lights))
	t.Logf("LastScan: %s", newlights.LastScan)
	for _, light := range newlights.Lights {
		t.Log(light)
	}

}

func TestUpdateLight(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 2
	resp, err := b.UpdateLight(id, huego.Light{
		Name: "New Light",
	})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Light %d updated", id)
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestDeleteLight(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	err := b.DeleteLight(id)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Light %d deleted")
	}
}

func TestTurnOffLight(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Off()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Turned off light with id %d", light.Id)
}

func TestTurnOffLightLazy(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, _ := b.GetLight(id)
	light.Off()
	t.Logf("Turned off light with id %d", light.Id)
}

func TestTurnOnLight(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.On()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Turned on light with id %d", light.Id)
}

func TestTurnOnLightLazy(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, _ := b.GetLight(id)
	light.On()
	t.Logf("Turned on light with id %d", light.Id)
}

func TestIfLightIsOn(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Is light %d on?: %t", light.Id, light.IsOn())
}

func TestRenameLight(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Rename("Color Lamp 1")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Renamed light to '%s'", light.Name)
}

func TestSetLightBri(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Bri(254)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Brightness of light %d set to %d", light.Id, light.State.Bri)
}

func TestSetLightHue(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Hue(65535)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Hue of light %d set to %d", light.Id, light.State.Hue)
}

func TestSetLightSat(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Sat(254)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Sat of light %d set to %d", light.Id, light.State.Sat)
}

func TestSetLightXy(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Xy([]float32{0.1, 0.5})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Xy of light %d set to %d", light.Id, light.State.Xy)
}

func TestSetLightCt(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Ct(16)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Ct of light %d set to %d", light.Id, light.State.Ct)
}