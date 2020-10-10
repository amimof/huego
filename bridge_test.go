package huego

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func ExampleBridge_CreateUser() {
	bridge, _ := Discover()
	user, err := bridge.CreateUser("my awesome hue app") // Link button needs to be pressed
	if err != nil {
		fmt.Printf("Error creating user: %s", err.Error())
	}
	bridge = bridge.Login(user)
	light, _ := bridge.GetLight(1)
	light.Off()
}

func TestLogin(t *testing.T) {
	b := New(hostname, username)
	c, err := b.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Logged in and got config which means that we are authorized")
	t.Logf("Name: %s, SwVersion: %s", c.Name, c.SwVersion)
}

func TestLoginUnauthorized(t *testing.T) {
	b := New(hostname, "")
	b = b.Login("invalid_password")
	_, err := b.GetLights()
	if err != nil {
		if strings.Contains(err.Error(), "unauthorized user") == false {
			t.Fatal(err)
		}
	}
	t.Logf("Bridge: %s, Username: %s", b.Host, b.User)
	t.Log("User logged in and authenticated which isn't what we want")
}

func TestUpdateBridgeConfig(t *testing.T) {
	b := New(hostname, username)
	c, err := b.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	_, err = b.UpdateConfig(c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateBridgeConfigError(t *testing.T) {
	b := New(badHostname, username)
	_, err := b.GetConfig()
	if err == nil {
		t.Fatal("Expected error not to be nil")
	}
}

func TestBridge_getAPIPathError(t *testing.T) {
	b := New("invalid hostname", "")
	expected := `parse "http://invalid hostname": invalid character " " in host name`
	_, err := b.getAPIPath("/")
	if err.Error() != expected {
		t.Fatalf("Expected error %s but got %s", expected, err.Error())
	}
}

func TestBridge_getError(t *testing.T) {
	httpmock.Deactivate()
	defer httpmock.Activate()
	expected := `Get "invalid%20hostname": unsupported protocol scheme ""`
	_, err := get(context.Background(), "invalid hostname")
	if err.Error() != expected {
		t.Fatalf("Expected error %s but got %s", expected, err.Error())
	}
}

func TestBridge_putError(t *testing.T) {
	httpmock.Deactivate()
	defer httpmock.Activate()
	expected := `Put "invalid%20hostname": unsupported protocol scheme ""`
	_, err := put(context.Background(), "invalid hostname", []byte("huego"))
	if err.Error() != expected {
		t.Fatalf("Expected error %s but got %s", expected, err.Error())
	}
}

func TestBridge_postError(t *testing.T) {
	httpmock.Deactivate()
	defer httpmock.Activate()
	expected := `Post "invalid%20hostname": unsupported protocol scheme ""`
	_, err := post(context.Background(), "invalid hostname", []byte("huego"))
	if err.Error() != expected {
		t.Fatalf("Expected error %s but got %s", expected, err.Error())
	}
}

func TestBridge_deleteError(t *testing.T) {
	httpmock.Deactivate()
	defer httpmock.Activate()
	expected := `Delete "invalid%20hostname": unsupported protocol scheme ""`
	_, err := delete(context.Background(), "invalid hostname")
	if err.Error() != expected {
		t.Fatalf("Expected error %s but got %s", expected, err.Error())
	}
}
