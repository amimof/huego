package huego_test

import (
	"github.com/amimof/huego"
	"os"
	"testing"
)

func TestGetResourcelinks(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resourcelinks, err := hue.GetResourcelinks()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d resourcelinks", len(resourcelinks))
	for i, resourcelink := range resourcelinks {
		t.Logf("%d", i)
		t.Logf("  Name: %s", resourcelink.Name)
		t.Logf("  Description: %s", resourcelink.Description)
		t.Logf("  Type: %s", resourcelink.Type)
		t.Logf("  ClassID: %s", resourcelink.ClassID)
		t.Logf("  Owner: %s", resourcelink.Owner)
		t.Logf("  Links: %s", resourcelink.Links)
		t.Logf("  ID: %s", resourcelink.ID)
	}
}

func TestGetResourcelink(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resourcelinks, err := hue.GetResourcelinks()
	if err != nil {
		t.Fatal(err)
	}
	for _, resourcelink := range resourcelinks {
		l, err := hue.GetResourcelink(resourcelink.ID)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Name: %s", l.Name)
		t.Logf("Description: %s", l.Description)
		t.Logf("Type: %s", l.Type)
		t.Logf("ClassID: %s", l.ClassID)
		t.Logf("Owner: %s", l.Owner)
		t.Logf("Links: %s", l.Links)
		t.Logf("ID: %s", l.ID)
		break
	}
}

func TestCreateResourcelink(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	resourcelink := &huego.Resourcelink{
		Name:        "Huego Test Resourcelink",
		Description: "Amir's wakeup experience",
		Type:        "Link",
		ClassID:     1,
		Owner:       "78H56B12BAABCDEF",
		Links:       []string{"/schedules/2", "/schedules/3", "/scenes/ABCD"},
	}
	resp, err := hue.CreateResourcelink(resourcelink)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Resourcelink created")
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestUpdateResourcelink(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	resp, err := hue.UpdateResourcelink(id, &huego.Resourcelink{
		Name:        "New Resourcelink",
		Description: "Updated Attribute",
	})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Resourcelink %d updated", id)
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestDeleteResourcelink(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	err := hue.DeleteResourcelink(1)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Resourcelink %d deleted", id)
	}
}
