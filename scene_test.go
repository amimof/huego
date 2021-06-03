package huego

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetScenes(t *testing.T) {
	b := New(hostname, username)
	scenes, err := b.GetScenes()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d scenes", len(scenes))
	for i, scene := range scenes {
		t.Logf("%d", i)
		t.Logf("  Name: %s", scene.Name)
		t.Logf("  Type: %s", scene.Type)
		t.Logf("  Group: %s", scene.Group)
		t.Logf("  Lights: %s", scene.Lights)
		t.Logf("  Owner: %s", scene.Owner)
		t.Logf("  Recycle: %t", scene.Recycle)
		t.Logf("  Locked: %t", scene.Locked)
		t.Logf("  AppData: %s", scene.AppData)
		t.Logf("  Picture: %s", scene.Picture)
		t.Logf("  LastUpdated: %s", scene.LastUpdated)
		t.Logf("  Version: %d", scene.Version)
		t.Logf("  StoreLightState: %t", scene.StoreLightState)
		t.Logf("  ID: %s", scene.ID)
	}

	contains := func(name string, ss []Scene) bool {
		for _, s := range ss {
			if s.Name == name {
				return true
			}
		}
		return false
	}

	assert.True(t, contains("Kathyon1449133269486", scenes))
	assert.True(t, contains("Cozydinner", scenes))

	b.Host = badHostname
	_, err = b.GetScenes()
	assert.NotNil(t, err)

}

func TestGetScene(t *testing.T) {
	b := New(hostname, username)
	id := "4e1c6b20e-on-0"
	s, err := b.GetScene(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("  Name: %s", s.Name)
	t.Logf("  Type: %s", s.Type)
	t.Logf("  Group: %s", s.Group)
	t.Logf("  Lights: %s", s.Lights)
	t.Logf("  Owner: %s", s.Owner)
	t.Logf("  Recycle: %t", s.Recycle)
	t.Logf("  Locked: %t", s.Locked)
	t.Logf("  AppData: %s", s.AppData)
	t.Logf("  Picture: %s", s.Picture)
	t.Logf("  LastUpdated: %s", s.LastUpdated)
	t.Logf("  Version: %d", s.Version)
	t.Logf("  StoreLightState: %t", s.StoreLightState)
	t.Logf("  ID: %s", s.ID)
	t.Logf("  LightStates: %d", len(s.LightStates))
	for k := range s.LightStates {
		t.Logf("    Light %d: %+v", k, s.LightStates[k])
	}

	b.Host = badHostname
	_, err = b.GetScene(id)
	assert.NotNil(t, err)
}

func TestCreateScene(t *testing.T) {
	b := New(hostname, username)
	scene := &Scene{
		Name:    "New Scene",
		Lights:  []string{"4", "5"},
		Recycle: true,
	}
	resp, err := b.CreateScene(scene)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Scene created")
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}

	b.Host = badHostname
	_, err = b.CreateScene(scene)
	assert.NotNil(t, err)
}

func TestUpdateScene(t *testing.T) {
	b := New(hostname, username)
	scene, err := b.GetScene("4e1c6b20e-on-0")
	if err != nil {
		t.Fatal(err)
	}
	newscene := &Scene{
		Name:   "New Scene",
		Lights: []string{},
	}
	resp, err := b.UpdateScene(scene.ID, newscene)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Scene '%s' (%s) updated", scene.Name, scene.ID)
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}

	b.Host = badHostname
	_, err = b.UpdateScene(scene.ID, newscene)
	assert.NotNil(t, err)
}

func TestSetSceneLightState(t *testing.T) {
	b := New(hostname, username)
	scene, err := b.GetScene("4e1c6b20e-on-0")
	if err != nil {
		t.Fatal(err)
	}
	light := 1
	state := &State{
		On:  true,
		Bri: 255,
	}
	t.Logf("Name: %s", scene.Name)
	t.Logf("ID: %s", scene.ID)
	t.Logf("LightStates: %+v", scene.LightStates)
	_, err = b.SetSceneLightState(scene.ID, light, state)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully set the state of light %d of scene '%s'", light, scene.Name)

	b.Host = badHostname
	_, err = b.SetSceneLightState(scene.ID, light, state)
	assert.NotNil(t, err)

}

func TestDeleteScene(t *testing.T) {
	b := New(hostname, username)
	scene, err := b.GetScene("4e1c6b20e-on-0")
	if err != nil {
		t.Fatal(err)
	}
	err = b.DeleteScene(scene.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Scene %s (%s) deleted", scene.Name, scene.ID)
}

func TestRecallScene(t *testing.T) {
	b := New(hostname, username)
	scene := "4e1c6b20e-on-0"
	group := 1
	resp, err := b.RecallScene("HcK1mEcgS7gcVcT", group)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Scene %s recalled in group %d", scene, group)
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}

	b.Host = badHostname
	_, err = b.RecallScene("HcK1mEcgS7gcVcT", group)
	assert.NotNil(t, err)
}

func TestRecallScene2(t *testing.T) {
	b := New(hostname, username)
	group := 1
	scene, err := b.GetScene("4e1c6b20e-on-0")
	if err != nil {
		t.Fatal(err)
	}
	err = scene.Recall(group)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Scene %s (%s) recalled in group %d", scene.Name, scene.ID, group)

	b.Host = badHostname
	err = scene.Recall(group)
	assert.NotNil(t, err)
}
