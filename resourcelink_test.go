package huego

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetResourcelinks(t *testing.T) {
	b := New(hostname, username)
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

	contains := func(name string, ss []*Resourcelink) bool {
		for _, s := range ss {
			if s.Name == name {
				return true
			}
		}
		return false
	}

	assert.True(t, contains("Sunrise", resourcelinks))
	assert.True(t, contains("Sunrise 2", resourcelinks))

	b.Host = badHostname
	_, err = b.GetResourcelinks()
	assert.NotNil(t, err)
}

func TestGetResourcelink(t *testing.T) {
	b := New(hostname, username)
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

	b.Host = badHostname
	_, err = b.GetResourcelink(1)
	assert.NotNil(t, err)
}

func TestCreateResourcelink(t *testing.T) {
	b := New(hostname, username)
	resourcelink := &Resourcelink{
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

	b.Host = badHostname
	_, err = b.CreateResourcelink(resourcelink)
	assert.NotNil(t, err)
}

func TestUpdateResourcelink(t *testing.T) {
	b := New(hostname, username)
	id := 1
	rl := &Resourcelink{
		Name:        "New Resourcelink",
		Description: "Updated Attribute",
	}
	resp, err := b.UpdateResourcelink(id, rl)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Resourcelink %d updated", id)
	for k, v := range resp.Success {
		t.Logf("%v: %s", k, v)
	}

	b.Host = badHostname
	_, err = b.UpdateResourcelink(id, rl)
	assert.NotNil(t, err)
}

func TestDeleteResourcelink(t *testing.T) {
	b := New(hostname, username)
	id := 1
	err := b.DeleteResourcelink(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Resourcelink %d deleted", id)
}
