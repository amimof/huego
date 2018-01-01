package huego_test

import (
	"testing"
	"os"
	"github.com/amimof/huego"
)

func TestGetResourcelinks(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resourcelinks, err := hue.GetResourcelinks()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d resourcelinks", len(resourcelinks))
	for i, resourcelink := range resourcelinks {
		t.Logf("%d", i)
		t.Logf("  Name: %s", resourcelink.Name)
		t.Logf("  Description: %s", resourcelink.Description)
		t.Logf("  Type: %s", resourcelink.Type)
		t.Logf("  ClassId: %s", resourcelink.ClassId)
		t.Logf("  Owner: %s", resourcelink.Owner)
		t.Logf("  Links: %s", resourcelink.Links)
		t.Logf("  Id: %s", resourcelink.Id)
	}
}


func TestGetResourcelink(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resourcelinks, err := hue.GetResourcelinks()
	if err != nil {
		t.Error(err)
	}
	for _, resourcelink := range resourcelinks {
		l, err := hue.GetResourcelink(resourcelink.Id)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Name: %s", l.Name)
		t.Logf("Description: %s", l.Description)
		t.Logf("Type: %s", l.Type)
		t.Logf("ClassId: %s", l.ClassId)
		t.Logf("Owner: %s", l.Owner)
		t.Logf("Links: %s", l.Links)
		t.Logf("Id: %s", l.Id)
		break
	}
}

func TestCreateResourcelink(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resourcelink := &huego.Resourcelink{
		Name: "Huego Test Resourcelink",
    Description: "Amir's wakeup experience",
    Type: "Link",
    ClassId: 1,
    Owner: "78H56B12BAABCDEF",
    Links: []string{"/schedules/2", "/schedules/3", "/scenes/ABCD"},
	}
	response, err := hue.CreateResourcelink(resourcelink)
	if err != nil {
		t.Error(err)
	}
	for _, r := range response {
		t.Logf("Address: %s", r.Address)
		t.Logf("Value: %s", r.Value)
		t.Logf("Interface: %s", r.Interface)
		t.Logf("Json: %s", r.Json)
	}
}

func TestUpdateResourcelink(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resourcelinks, err := hue.GetSensors()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d resourcelinks, setting the first one", len(resourcelinks))
	for _, resourcelink := range resourcelinks {
		response, err := hue.UpdateSensor(resourcelink.Id, resourcelink)
		if err != nil {
			t.Error(err)
		}
		for _, r := range response {
			t.Logf("Address: %s Value: %s Interface: %s", r.Address, r.Value, r.Interface)
		}
		break
	}
}

func TestDeleteResourcelink(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	res, err := hue.DeleteResourcelink(1)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		for _, r := range res {
			t.Logf("Address: %s Value: %s Interface: %s", r.Address, r.Value, r.Interface)
		}
	}
}
