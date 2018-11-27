package huego_test

import (
	"github.com/amimof/huego"
	"testing"
)

func TestGetGroups(t *testing.T) {
	b := huego.New(hostname, username)
	groups, err := b.GetGroups()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d groups", len(groups))
	for i, g := range groups {
		t.Logf("%d:", i)
		t.Logf("  Name: %s", g.Name)
		t.Logf("  Lights: %s", g.Lights)
		t.Logf("  Type: %s", g.Type)
		t.Logf("  GroupState:")
		t.Logf("    AllOn: %t", g.GroupState.AllOn)
		t.Logf("    AnyOn: %t", g.GroupState.AnyOn)
		t.Logf("  Recycle: %t", g.Recycle)
		t.Logf("  Class: %s", g.Class)
		t.Logf("  State:")
		t.Logf("    On: %t", g.State.On)
		t.Logf("    Bri: %d", g.State.Bri)
		t.Logf("    Hue: %d", g.State.Hue)
		t.Logf("    Sat: %d", g.State.Sat)
		t.Logf("    Xy: %b", g.State.Xy)
		t.Logf("    Ct: %d", g.State.Ct)
		t.Logf("    Alert: %s", g.State.Alert)
		t.Logf("    Effect: %s", g.State.Effect)
		t.Logf("    TransitionTime: %d", g.State.TransitionTime)
		t.Logf("    BriInc: %d", g.State.BriInc)
		t.Logf("    SatInc: %d", g.State.SatInc)
		t.Logf("    HueInc: %d", g.State.HueInc)
		t.Logf("    CtInc: %d", g.State.CtInc)
		t.Logf("    XyInc: %d", g.State.XyInc)
		t.Logf("    ColorMode: %s", g.State.ColorMode)
		t.Logf("    Reachable: %t", g.State.Reachable)
		t.Logf("  ID: %d", g.ID)
	}
}

func TestGetGroup(t *testing.T) {
	b := huego.New(hostname, username)
	g, err := b.GetGroup(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Name: %s", g.Name)
	t.Logf("Lights: %s", g.Lights)
	t.Logf("Type: %s", g.Type)
	t.Logf("GroupState:")
	t.Logf("  AllOn: %t", g.GroupState.AllOn)
	t.Logf("  AnyOn: %t", g.GroupState.AnyOn)
	t.Logf("Recycle: %t", g.Recycle)
	t.Logf("Class: %s", g.Class)
	t.Logf("State:")
	t.Logf("  On: %t", g.State.On)
	t.Logf("  Bri: %d", g.State.Bri)
	t.Logf("  Hue: %d", g.State.Hue)
	t.Logf("  Sat: %d", g.State.Sat)
	t.Logf("  Xy: %b", g.State.Xy)
	t.Logf("  Ct: %d", g.State.Ct)
	t.Logf("  Alert: %s", g.State.Alert)
	t.Logf("  Effect: %s", g.State.Effect)
	t.Logf("  TransitionTime: %d", g.State.TransitionTime)
	t.Logf("  BriInc: %d", g.State.BriInc)
	t.Logf("  SatInc: %d", g.State.SatInc)
	t.Logf("  HueInc: %d", g.State.HueInc)
	t.Logf("  CtInc: %d", g.State.CtInc)
	t.Logf("  XyInc: %d", g.State.XyInc)
	t.Logf("  ColorMode: %s", g.State.ColorMode)
	t.Logf("  Reachable: %t", g.State.Reachable)
	t.Logf("ID: %d", g.ID)
}

func TestCreateGroup(t *testing.T) {
	b := huego.New(hostname, username)
	resp, err := b.CreateGroup(huego.Group{
		Name:   "TestGroup",
		Type:   "Room",
		Class:  "Office",
		Lights: []string{},
	})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Group created")
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestUpdateGroup(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	resp, err := b.UpdateGroup(id, huego.Group{
		Name:  "TestGroup (Updated)",
		Class: "Office",
	})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Updated group with id %d", id)
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestSetGroupState(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	resp, err := b.SetGroupState(id, huego.State{
		On:  true,
		Bri: 150,
		Sat: 210,
	})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Group state set")
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestRenameGroup(t *testing.T) {
	bridge := huego.New(hostname, username)
	id := 1
	group, err := bridge.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.Rename("MyGroup (renamed)")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Group renamed to %s", group.Name)
}

func TestTurnOffGroup(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.Off()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Turned off group with id %d", group.ID)
	t.Logf("Group IsOn: %t", group.State.On)
}

func TestTurnOnGroup(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.On()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Turned on group with id %d", group.ID)
	t.Logf("Group IsOn: %t", group.State.On)
}

func TestIfGroupIsOn(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Is group %d on?: %t", group.ID, group.IsOn())
}

func TestSetGroupBri(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.Bri(254)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Brightness of group %d set to %d", group.ID, group.State.Bri)
}

func TestSetGroupHue(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.Hue(65535)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Hue of group %d set to %d", group.ID, group.State.Hue)
}

func TestSetGroupSat(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.Sat(254)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Sat of group %d set to %d", group.ID, group.State.Sat)
}

func TestSetGroupXy(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.Xy([]float32{0.1, 0.5})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Xy of group %d set to %v", group.ID, group.State.Xy)
}

func TestSetGroupCt(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.Ct(16)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Ct of group %d set to %d", group.ID, group.State.Ct)
}

func TestSetGroupScene(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.Scene("2hgE1nGaITvy9VQ")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Scene of group %d set to %s", group.ID, group.State.Scene)
}

func TestSetGroupTransitionTime(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.TransitionTime(10)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TransitionTime of group %d set to %d", group.ID, group.State.TransitionTime)
}

func TestSetGroupEffect(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.Effect("colorloop")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Effect of group %d set to %s", group.ID, group.State.Effect)
}

func TestSetGroupAlert(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.Alert("lselect")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Alert of group %d set to %s", group.ID, group.State.Alert)
}

func TestSetStateGroup(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.SetState(huego.State{
		On:  true,
		Bri: 254,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("State set successfully on group %d", id)
}

func TestDeleteGroup(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	err := b.DeleteGroup(id)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Deleted group with id: %d", id)
	}
}
