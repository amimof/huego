package huego_test

import (
	"github.com/amimof/huego"
	"os"
	"testing"
)

func TestGetLights(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	lights, err := b.GetLights()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d lights", len(lights))
	for _, l := range lights {
		t.Logf("ID: %d", l.ID)
		t.Logf("  Type: %s", l.Type)
		t.Logf("  Name: %s", l.Name)
		t.Logf("  ModelID: %s", l.ModelID)
		t.Logf("  ManufacturerName: %s", l.ManufacturerName)
		t.Logf("  UniqueID: %s", l.UniqueID)
		t.Logf("  SwVersion: %s", l.SwVersion)
		t.Logf("  SwConfigID: %s", l.SwConfigID)
		t.Logf("  ProductID: %s", l.ProductID)
	}
}

func TestGetLight(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	lights, err := b.GetLights()
	if err != nil {
		t.Fatal(err)
	}
	for _, light := range lights {
		l, err := b.GetLight(light.ID)
		if err != nil {
			t.Fatal(err)
		} else {
			t.Logf("ID: %d", l.ID)
			t.Logf("Type: %s", l.Type)
			t.Logf("Name: %s", l.Name)
			t.Logf("ModelID: %s", l.ModelID)
			t.Logf("ManufacturerName: %s", l.ManufacturerName)
			t.Logf("UniqueID: %s", l.UniqueID)
			t.Logf("SwVersion: %s", l.SwVersion)
			t.Logf("SwConfigID: %s", l.SwConfigID)
			t.Logf("ProductID: %s", l.ProductID)
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
	resp, err := b.SetLightState(id, huego.State{
		On:  true,
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
		t.Logf("Light %d deleted", id)
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
	t.Logf("Turned off light with id %d", light.ID)
}

func TestTurnOffLightLazy(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, _ := b.GetLight(id)
	light.Off()
	t.Logf("Turned off light with id %d", light.ID)
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
	t.Logf("Turned on light with id %d", light.ID)
}

func TestTurnOnLightLazy(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, _ := b.GetLight(id)
	light.On()
	t.Logf("Turned on light with id %d", light.ID)
}

func TestIfLightIsOn(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Is light %d on?: %t", light.ID, light.IsOn())
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
	t.Logf("Brightness of light %d set to %d", light.ID, light.State.Bri)
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
	t.Logf("Hue of light %d set to %d", light.ID, light.State.Hue)
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
	t.Logf("Sat of light %d set to %d", light.ID, light.State.Sat)
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
	t.Logf("Xy of light %d set to %+v", light.ID, light.State.Xy)
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
	t.Logf("Ct of light %d set to %d", light.ID, light.State.Ct)
}

func TestSetLightTransitionTime(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.TransitionTime(10)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TransitionTime of light %d set to %d", light.ID, light.State.TransitionTime)
}

func TestSetLightEffect(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Effect("colorloop")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Effect of light %d set to %s", light.ID, light.State.Effect)
}

func TestSetLightAlert(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Alert("lselect")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Alert of light %d set to %s", light.ID, light.State.Alert)
}

func TestSetStateLight(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 4
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.SetState(huego.State{
		On:  true,
		Bri: 254,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("State set successfully on light %d", id)
}
