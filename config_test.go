package huego_test

import (
	"testing"
	"os"
	"github.com/amimof/huego"
)

// O60ECZZJhwrTI8AkY1xjOK5ifj20igjw6R5WsWih

func TestGetConfig(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	config, err := hue.GetConfig()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Got config: %+v", config)
}

func TestCreateUser(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), "")
	u, err := hue.CreateUser("huego#tests")
	if err != nil {
		t.Error(err)
	}
	t.Logf("User created with username: %s", u)
}