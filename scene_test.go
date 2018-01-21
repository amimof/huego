package huego_test

import (
	"github.com/amimof/huego"
	"os"
	"testing"
)

func TestGetScenes(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	scenes, err := hue.GetScenes()
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
		t.Logf("  ID: %d", scene.ID)
	}
}

func TestGetScene(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	scenes, err := hue.GetScenes()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d scenes", len(scenes))
	for _, scene := range scenes {
		t.Logf("Getting scene %d, skipping the rest", scene.ID)
		s, err := hue.GetScene(scene.ID)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Got scene name=%s", s.Name)
		break
	}
}

func TestCreateScene(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resp, err := hue.CreateScene(&huego.Scene{
		Name:    "New Scene",
		Lights:  []string{},
		Recycle: true,
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

func TestUpdateScene(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	resp, err := hue.UpdateScene(id, &huego.Scene{
		Name:   "New Scene",
		Lights: []string{},
	})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Scene %d updated", id)
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestDeleteScene(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	err := hue.DeleteScene(3)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Scene %d deleted", id)
	}
}
