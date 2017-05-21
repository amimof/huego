package huego

import (
	"testing"
	"os"
)

func TestGetGroups(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
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
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
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
	}
}

func TestSetGroupState(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
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
			t.Logf("Response from put: Success=%v Error=%v", r.Success, r.Error)
		}
	}
}

func TestSetGroup(t *testing.T) {
	//hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
}

func TestCreateGroup(t *testing.T) {	
	//hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
}

func TestDeleteGroup(t *testing.T) {
	//hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
}






