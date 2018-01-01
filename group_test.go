package huego_test

import (
	"testing"
	"os"
	"github.com/amimof/huego"
)

func TestGetGroups(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	groups, err := hue.GetGroups()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d groups", len(groups))
	for i, g := range groups {
		t.Logf("%d:", i)
		t.Logf("  Name: %s", g.Name)
		t.Logf("  Lights: %s", g.Lights)
		t.Logf("  Type: %s", g.Type)
		t.Logf("  State")
		t.Logf("    AllOn: %t", g.State.AllOn)
		t.Logf("    AnyOn: %t", g.State.AnyOn)
		t.Logf("  Recycle: %t", g.Recycle)
		t.Logf("  Class: %s", g.Class)
		t.Logf("  Action:")
		t.Logf("    On: %t", g.Action.On)
		t.Logf("    Bri: %d", g.Action.Bri)
		t.Logf("    Hue: %d", g.Action.Hue)
		t.Logf("    Sat: %d", g.Action.Sat)
		t.Logf("    Xy: %b", g.Action.Xy)
		t.Logf("    Ct: %d", g.Action.Ct)
		t.Logf("    Alert: %s", g.Action.Alert)
		t.Logf("    Effect: %s", g.Action.Effect)
		t.Logf("    TransitionTime: %d", g.Action.TransitionTime)
		t.Logf("    BriInc: %d", g.Action.BriInc)
		t.Logf("    SatInc: %d", g.Action.SatInc)
		t.Logf("    HueInc: %d", g.Action.HueInc)
		t.Logf("    CtInc: %d", g.Action.CtInc)
		t.Logf("    XyInc: %d", g.Action.XyInc)
		t.Logf("    ColorMode: %s", g.Action.ColorMode)
		t.Logf("    Reachable: %t", g.Action.Reachable)
		t.Logf("  Id: %d", g.Id)
	}
}

func TestGetGroup(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	groups, err := hue.GetGroups()
	if err != nil {
		t.Error(err)
	}
	for _, group := range groups {
		g, err := hue.GetGroup(group.Id)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Name: %s", g.Name)
		t.Logf("Lights: %s", g.Lights)
		t.Logf("Type: %s", g.Type)
		t.Logf("State")
		t.Logf("  AllOn: %t", g.State.AllOn)
		t.Logf("  AnyOn: %t", g.State.AnyOn)
		t.Logf("Recycle: %t", g.Recycle)
		t.Logf("Class: %s", g.Class)
		t.Logf("Action:")
		t.Logf("  On: %t", g.Action.On)
		t.Logf("  Bri: %d", g.Action.Bri)
		t.Logf("  Hue: %d", g.Action.Hue)
		t.Logf("  Sat: %d", g.Action.Sat)
		t.Logf("  Xy: %b", g.Action.Xy)
		t.Logf("  Ct: %d", g.Action.Ct)
		t.Logf("  Alert: %s", g.Action.Alert)
		t.Logf("  Effect: %s", g.Action.Effect)
		t.Logf("  TransitionTime: %d", g.Action.TransitionTime)
		t.Logf("  BriInc: %d", g.Action.BriInc)
		t.Logf("  SatInc: %d", g.Action.SatInc)
		t.Logf("  HueInc: %d", g.Action.HueInc)
		t.Logf("  CtInc: %d", g.Action.CtInc)
		t.Logf("  XyInc: %d", g.Action.XyInc)
		t.Logf("  ColorMode: %s", g.Action.ColorMode)
		t.Logf("  Reachable: %t", g.Action.Reachable)
		t.Logf("Id: %d", g.Id)
		break
	}
}

func TestSetGroupState(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	groups, err := hue.GetGroups()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d groups, using the first one", len(groups))
	for _, group := range groups {
		resp, err := hue.SetGroupState(group.Id, &huego.State{
			On: true,
		})
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Group state set")
			for k, v := range resp.Success {
				t.Logf("%v: %s", k, v)
			}		
		}
		break
	}
}

func TestCreateGroup(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resp, err := hue.CreateGroup(&huego.Group{
		Name: "TestGroup",
		Type: "Room",
		Class: "Office",
		Lights: []string{},
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Group created")
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestUpdateGroup(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	resp, err := hue.UpdateGroup(id, &huego.Group{
		Name: "TestGroup (Updated)",
		Class: "Office",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Updated group with id %d", id)
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestDeleteGroup(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	err := hue.DeleteGroup(id)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Deleted group with id: %d", id)
	}
}
