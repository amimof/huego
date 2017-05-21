package huego

import (
	"testing"
	"os"
)

func TestGetGroups(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
}

func TestGetGroup(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))	
}

func TestSetGroupState(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
}

func TestSetGroup(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
}

func TestCreateGroup(t *testing.T) {	
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
}

func TestDeleteGroup(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
}






