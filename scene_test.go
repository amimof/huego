package huego_test

import (
	"github.com/amimof/huego"
	"os"
	"testing"
)

func TestGetScenes(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	scenes, err := b.GetScenes()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d scenes", len(scenes))
	for i, scene := range scenes {
		t.Logf("%d", i)
		t.Logf("  Name: %s", scene.Name)
		t.Logf("  Lights: %s", scene.Lights)
		t.Logf("  Owner: %s", scene.Owner)
		t.Logf("  Recycle: %t", scene.Recycle)
		t.Logf("  Locked: %t", scene.Locked)
		t.Logf("  AppData: %s", scene.AppData)
		t.Logf("  Picture: %s", scene.Picture)
		t.Logf("  LastUpdated: %s", scene.LastUpdated)
		t.Logf("  Version: %d", scene.Version)
		t.Logf("  StoreSceneState: %t", scene.StoreSceneState)
		t.Logf("  ID: %s", scene.ID)
	}
}

func TestGetScene(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	scenes, err := b.GetScenes()
	if err != nil {
		t.Fatal(err)
	}
	for _, scene := range scenes {
		t.Logf("Getting scene %s", scene.ID)
		s, err := b.GetScene(scene.ID)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("  Name: %s", s.Name)
		t.Logf("  Lights: %s", s.Lights)
		t.Logf("  Owner: %s", s.Owner)
		t.Logf("  Recycle: %t", s.Recycle)
		t.Logf("  Locked: %t", s.Locked)
		t.Logf("  AppData: %s", s.AppData)
		t.Logf("  Picture: %s", s.Picture)
		t.Logf("  LastUpdated: %s", s.LastUpdated)
		t.Logf("  Version: %d", s.Version)
		t.Logf("  StoreSceneState: %t", s.StoreSceneState)
		t.Logf("  ID: %s", s.ID)
		t.Logf("  LightStates: %d", len(s.LightStates))
		break
	}
}

func TestCreateScene(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resp, err := b.CreateScene(&huego.Scene{
		Name:    "New Scene",
		Lights:  []string{"4", "5"},
		Recycle: true,
	})
	if err != nil {
		t.Fatal(err)
	} 
	t.Logf("Scene created")
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}
}

func TestUpdateScene(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	scenes, err := b.GetScenes()
	if err != nil {
		t.Fatal(err)
	}
	for _, scene := range scenes {
		if scene.Name == "New Scene" {
			resp, err := b.UpdateScene(scene.ID, &huego.Scene{
				Name:   "New Scene",
				Lights: []string{},
			})
			if err != nil {
				t.Fatal(err)
			} 
			t.Logf("Scene '%s' (%s) updated", scene.Name, scene.ID)
			for k, v := range resp.Success {
				t.Logf("%v: %s", k, v)
			}
		}
	}
}

func TestSetSceneLightState(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	scenes, err := b.GetScenes()
	if err != nil {
		t.Fatal(err)
	}
	for _, scene := range scenes {
		if scene.Name == "New Scene" {
			light := 4
			t.Logf("Name: %s", scene.Name)
			t.Logf("ID: %s", scene.ID)
			t.Logf("LightStates: %+v", scene.LightStates)
			_, err := b.SetSceneLightState(scene.ID, light, &huego.State{
				On: true,
				Bri: 255,
			})
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("Successfully set the state of light %d of scene '%s'", light, scene.Name)
		}
	}
}

func TestDeleteScene(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	scenes, err := b.GetScenes()
	if err != nil {
		t.Fatal(err)
	} 
	for _, scene := range scenes {
		if scene.Name == "New Scene" {
			err := b.DeleteScene(scene.ID)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("Scene %s (%s) deleted", scene.Name, scene.ID)
		}
	}
}

func TestRecallScene(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	scene := "HcK1mEcgS7gcVcT"
	group := 1
	resp, err := b.RecallScene("HcK1mEcgS7gcVcT", group)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Scene %s recalled in group %d", scene, group)
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}
}

func TestRecallScene2(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	group := 1
	scenes, err := b.GetScenes()
	if err != nil {
		t.Fatal(err)
	}
	for _, scene := range scenes {
		err = scene.Recall(group)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Scene %s (%s) recalled in group %d", scene.Name, scene.ID, group)
		break
	}
}