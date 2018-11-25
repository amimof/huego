package huego_test

import (
	"github.com/amimof/huego"
	"os"
	"testing"
)

// I'm too lazy to have this elsewhere
// export HUE_USERNAME=9D6iHMbM-Bt7Kd0Cwh9Quo4tE02FMnmrNrnFAdAq
// export HUE_HOSTNAME=192.168.1.59

var username string
var hostname string

func init() {
	hostname = os.Getenv("HUE_HOSTNAME")
	username = os.Getenv("HUE_USERNAME")
}

func TestDiscoverAndLogin(t *testing.T) {
	bridge, err := huego.Discover()
	if err != nil {
		t.Fatal(err)
	}
	bridge = bridge.Login(username)
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
