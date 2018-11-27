package huego_test

import (
	"github.com/amimof/huego"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	b := huego.New(hostname, username)
	c, err := b.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Logged in and got config which means that we are authorized")
	t.Logf("Name: %s, SwVersion: %s", c.Name, c.SwVersion)
}

func TestLoginUnauthorized(t *testing.T) {
	b := huego.New(hostname, "")
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
	b := huego.New(hostname, username)
	c, err := b.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	_, err = b.UpdateConfig(c)
	if err != nil {
		t.Fatal(err)
	}
}