package huego_test

import (
	"testing"
	"github.com/amimof/huego"
)

func TestFindBridges(t *testing.T) {
	bridges, err := huego.Discover()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Discovered %d bridge(s)", len(bridges))
		for i, bridge := range bridges {
			t.Logf("%d: ", i)
			t.Logf("  Host: %s", bridge.Host)
			t.Logf("  User: %s", bridge.User)
			t.Logf("  Id: %s", bridge.Id)
		}
	}
}

func TestDiscoverAndCreateUser(t *testing.T) {
	bridges, err := huego.Discover()
	if err != nil {
		t.Fatal(err)
	} 
	t.Logf("Discovered %d bridges, using the first one", len(bridges))
	if len(bridges) > 0 {
		user, err := bridges[0].CreateUser("github.com/amimof/huego#tests")
		if err != nil {
			t.Fatal(err)
		}
		bridge := bridges[0].Login(user)
		config, err := bridge.GetConfig()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Successfully logged in to bridge. Id: %s", config.BridgeId)
	} else {
		t.Logf("No bridges found")
	}
}