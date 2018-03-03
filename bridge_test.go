package huego_test

import (
	"github.com/amimof/huego"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_PASSWORD"))
	c, err := b.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Logged in and got config which means that we are authorized")
	t.Logf("Name: %s, SwVersion: %s", c.Name, c.SwVersion)
}

func TestLoginUnauthorized(t *testing.T) {
	b := huego.New(os.Getenv("HUE_HOSTNAME"), "")
	b = b.Login("invalid_password")
	_, err := b.GetLights()
	if err != nil {
		t.Fatal(err)
		/*
			We should do something here to check if the user is allowed
			the requested resource (in this case Config). A simple err
			shouldn't fail the test.
		*/
	}
	t.Logf("Bridge: %s, Username: %s", b.Host, b.User)
	t.Log("User logged in and authenticated which isn't what we want")
}
