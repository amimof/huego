package huego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGroups(t *testing.T) {
	b := New(hostname, username)
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

	contains := func(name string, ss []Group) bool {
		for _, s := range ss {
			if s.Name == name {
				return true
			}
		}
		return false
	}

	assert.True(t, contains("Group 1", groups))
	assert.True(t, contains("Group 2", groups))

	b.Host = badHostname
	_, err = b.GetGroups()
	assert.NotNil(t, err)
}

func TestGetGroup(t *testing.T) {
	b := New(hostname, username)
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

	b.Host = badHostname
	_, err = b.GetGroup(1)
	assert.NotNil(t, err)
}

func TestCreateGroup(t *testing.T) {
	b := New(hostname, username)
	group := Group{
		Name:   "TestGroup",
		Type:   "Room",
		Class:  "Office",
		Lights: []string{},
	}
	resp, err := b.CreateGroup(group)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Group created")
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}

	b.Host = badHostname
	_, err = b.CreateGroup(group)
	assert.NotNil(t, err)

}

func TestUpdateGroup(t *testing.T) {
	b := New(hostname, username)
	id := 1
	group := Group{
		Name:  "TestGroup (Updated)",
		Class: "Office",
	}
	resp, err := b.UpdateGroup(id, group)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Updated group with id %d", id)
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}

	b.Host = badHostname
	_, err = b.UpdateGroup(id, group)
	assert.NotNil(t, err)

}

func TestSetGroupState(t *testing.T) {
	b := New(hostname, username)
	id := 1
	state := State{
		On:  true,
		Bri: 150,
		Sat: 210,
	}
	resp, err := b.SetGroupState(id, state)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Group state set")
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}

	b.Host = badHostname
	_, err = b.SetGroupState(id, state)
	assert.NotNil(t, err)
}

func TestRenameGroup(t *testing.T) {
	bridge := New(hostname, username)
	id := 1
	group, err := bridge.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	newName := "MyGroup (renamed)"
	err = group.Rename(newName)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Group renamed to %s", group.Name)

	bridge.Host = badHostname
	err = group.Rename(newName)
	assert.NotNil(t, err)

}

func TestTurnOffGroup(t *testing.T) {
	b := New(hostname, username)
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

	b.Host = badHostname
	err = group.Off()
	assert.NotNil(t, err)

}

func TestTurnOnGroup(t *testing.T) {
	b := New(hostname, username)
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

	b.Host = badHostname
	err = group.On()
	assert.NotNil(t, err)
}

func TestIfGroupIsOn(t *testing.T) {
	b := New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Is group %d on?: %t", group.ID, group.IsOn())
}

func TestSetGroupBri(t *testing.T) {
	b := New(hostname, username)
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

	b.Host = badHostname
	err = group.Bri(254)
	assert.NotNil(t, err)
}

func TestSetGroupHue(t *testing.T) {
	b := New(hostname, username)
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

	b.Host = badHostname
	err = group.Hue(65535)
	assert.NotNil(t, err)
}

func TestSetGroupSat(t *testing.T) {
	b := New(hostname, username)
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

	b.Host = badHostname
	err = group.Sat(254)
	assert.NotNil(t, err)
}

func TestSetGroupXy(t *testing.T) {
	b := New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	xy := []float32{0.1, 0.5}
	err = group.Xy(xy)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Xy of group %d set to %v", group.ID, group.State.Xy)

	b.Host = badHostname
	err = group.Xy(xy)
	assert.NotNil(t, err)
}

func TestSetGroupCt(t *testing.T) {
	b := New(hostname, username)
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

	b.Host = badHostname
	err = group.Ct(16)
	assert.NotNil(t, err)
}

func TestSetGroupScene(t *testing.T) {
	b := New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	scene := "2hgE1nGaITvy9VQ"
	err = group.Scene(scene)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Scene of group %d set to %s", group.ID, group.State.Scene)

	b.Host = badHostname
	err = group.Scene(scene)
	assert.NotNil(t, err)
}

func TestSetGroupTransitionTime(t *testing.T) {
	b := New(hostname, username)
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

	b.Host = badHostname
	err = group.TransitionTime(10)
	assert.NotNil(t, err)
}

func TestSetGroupEffect(t *testing.T) {
	b := New(hostname, username)
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

	b.Host = badHostname
	err = group.Effect("colorloop")
	assert.NotNil(t, err)
}

func TestSetGroupAlert(t *testing.T) {
	b := New(hostname, username)
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

	b.Host = badHostname
	err = group.Alert("lselect")
	assert.NotNil(t, err)
}

func TestSetStateGroup(t *testing.T) {
	b := New(hostname, username)
	id := 1
	group, err := b.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	state := State{
		On:  true,
		Bri: 254,
	}
	err = group.SetState(state)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("State set successfully on group %d", id)

	b.Host = badHostname
	err = group.SetState(state)
	assert.NotNil(t, err)
}

func TestDeleteGroup(t *testing.T) {
	b := New(hostname, username)
	id := 1
	err := b.DeleteGroup(id)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Deleted group with id: %d", id)
	}
}

func TestEnableStreamingGroup(t *testing.T) {
	bridge := New(hostname, username)
	id := 1
	group, err := bridge.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.EnableStreaming()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDisableStreamingGroup(t *testing.T) {
	bridge := New(hostname, username)
	id := 1
	group, err := bridge.GetGroup(id)
	if err != nil {
		t.Fatal(err)
	}
	err = group.DisableStreaming()
	if err != nil {
		t.Fatal(err)
	}
}
