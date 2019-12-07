package huego_test

import (
	"testing"

	"github.com/amimof/huego"
	"github.com/stretchr/testify/assert"
)

func TestGetResourcelinks(t *testing.T) {
	b := huego.New(hostname, username)
	resourcelinks, err := b.GetResourcelinks()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d resourcelinks", len(resourcelinks))
	for i, resourcelink := range resourcelinks {
		t.Logf("%d", i)
		t.Logf("  Name: %s", resourcelink.Name)
		t.Logf("  Description: %s", resourcelink.Description)
		t.Logf("  Type: %s", resourcelink.Type)
		t.Logf("  ClassID: %d", resourcelink.ClassID)
		t.Logf("  Owner: %s", resourcelink.Owner)
		t.Logf("  Links: %s", resourcelink.Links)
		t.Logf("  ID: %d", resourcelink.ID)
	}

	assert.Equal(t, "Sunrise", resourcelinks[0].Name)
	assert.Equal(t, "Sunrise 2", resourcelinks[1].Name)
}

func TestGetResourcelink(t *testing.T) {
	b := huego.New(hostname, username)
	l, err := b.GetResourcelink(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Name: %s", l.Name)
	t.Logf("Description: %s", l.Description)
	t.Logf("Type: %s", l.Type)
	t.Logf("ClassID: %d", l.ClassID)
	t.Logf("Owner: %s", l.Owner)
	t.Logf("Links: %s", l.Links)
	t.Logf("ID: %d", l.ID)
}

func TestCreateResourcelink(t *testing.T) {
	b := huego.New(hostname, username)
	resourcelink := &huego.Resourcelink{
		Name:        "Huego Test Resourcelink",
		Description: "Amir's wakeup experience",
		Type:        "Link",
		ClassID:     1,
		Owner:       "78H56B12BAABCDEF",
		Links:       []string{"/schedules/2", "/schedules/3", "/scenes/ABCD"},
	}
	resp, err := b.CreateResourcelink(resourcelink)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Resourcelink created")
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}

}

func TestUpdateResourcelink(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	resp, err := b.UpdateResourcelink(id, &huego.Resourcelink{
		Name:        "New Resourcelink",
		Description: "Updated Attribute",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Resourcelink %d updated", id)
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}

}

func TestDeleteResourcelink(t *testing.T) {
	b := huego.New(hostname, username)
	id := 1
	err := b.DeleteResourcelink(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Resourcelink %d deleted", id)
}
