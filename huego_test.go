package huego_test

// I'm too lazy to have this elsewhere
// export HUE_USERNAME=9D6iHMbM-Bt7Kd0Cwh9Quo4tE02FMnmrNrnFAdAq
// export HUE_HOSTNAME=192.168.1.59

import (
	"github.com/amimof/huego"
	"testing"
)

func TestDiscoverAndLoginLazy(t *testing.T) {
	b, _ := huego.Discover()
	b = b.Login("n7yx6YCUvV6-CGJZ5-VuyZoc3qgi9S2WjtEeDFpO")
	t.Logf("Successfully logged in to bridge")
}

func TestDiscoverAndLogin(t *testing.T) {
	bridge, err := huego.Discover()
	if err != nil {
		t.Fatal(err)
	}
	bridge = bridge.Login("n7yx6YCUvV6-CGJZ5-VuyZoc3qgi9S2WjtEeDFpO")
	t.Logf("Successfully logged in to bridge")
	config, err := bridge.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Bridge Version %s", config.SwVersion)
}

func TestDiscoverAndCreateUser(t *testing.T) {
	bridge, err := huego.Discover()
	if err != nil {
		t.Fatal(err)
	}
	user, err := bridge.CreateUser("My AweSome App")
	if err != nil {
		t.Fatal(err)
	}
	bridge = bridge.Login(user)
	t.Logf("User created and logged in")
}

func TestDiscoverAllBridges(t *testing.T) {
	bridges, err := huego.DiscoverAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Discovered %d bridge(s)", len(bridges))
	for i, bridge := range bridges {
		t.Logf("%d: ", i)
		t.Logf("  Host: %s", bridge.Host)
		t.Logf("  User: %s", bridge.User)
		t.Logf("  ID: %s", bridge.ID)
	}
}

func TestDiscoverAllAndCreateUser(t *testing.T) {
	bridges, err := huego.DiscoverAll()
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
		t.Logf("Successfully logged in to bridge. ID: %s", config.BridgeID)
	} else {
		t.Logf("No bridges found")
	}
}
