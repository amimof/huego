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
	for _, group := range groups {
		t.Log(group)
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
		t.Log(g)
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
		response, err := hue.SetGroupState(group.Id, group.Action)
		if err != nil {
			t.Error(err)
		}
		for _, r := range response {
			t.Logf("Address: %s Value: %s Interface: %s", r.Address, r.Value, r.Interface)
		}
	}
}

func TestUpdateGroup(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	groups, err := hue.GetGroups()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d groups, updating the first one", len(groups))
	for _, group := range groups {
		_, err := hue.UpdateGroup(group.Id, group)
		if err != nil {
			t.Error(err)
		}
		break
	}
}

func TestCreateGroup(t *testing.T) {
	//hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
}

func TestDeleteGroup(t *testing.T) {
	//hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
}
