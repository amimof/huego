package huego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	b := New(hostname, username)
	config, err := b.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Name: %s", config.Name)
	t.Logf("SwUpdate:")
	t.Logf("  CheckForUpdate: %t", config.SwUpdate.CheckForUpdate)
	t.Logf("  DeviceTypes:")
	t.Logf("    Bridge: %t", config.SwUpdate.DeviceTypes.Bridge)
	t.Logf("    Lights (length): %d", len(config.SwUpdate.DeviceTypes.Lights))
	t.Logf("    Sensors (length): %d", len(config.SwUpdate.DeviceTypes.Sensors))
	t.Logf("  UpdateState: %d", config.SwUpdate.UpdateState)
	t.Logf("  Notify: %t", config.SwUpdate.Notify)
	t.Logf("  URL: %s", config.SwUpdate.URL)
	t.Logf("  Text: %s", config.SwUpdate.Text)
	t.Logf("SwUpdate2:")
	t.Logf("  Bridge: %s", config.SwUpdate2.Bridge)
	t.Logf("    State: %s", config.SwUpdate2.Bridge.State)
	t.Logf("    LastInstall: %s", config.SwUpdate2.Bridge.LastInstall)
	t.Logf("  CheckForUpdate: %t", config.SwUpdate2.CheckForUpdate)
	t.Logf("  State: %s", config.SwUpdate2.State)
	t.Logf("  Install: %t", config.SwUpdate2.Install)
	t.Logf("  AutoInstall:")
	t.Logf("    On: %t", config.SwUpdate2.AutoInstall.On)
	t.Logf("    UpdateTime: %s", config.SwUpdate2.AutoInstall.UpdateTime)
	t.Logf("  LastChange: %s", config.SwUpdate2.LastChange)
	t.Logf("  LastInstall: %s", config.SwUpdate2.LastInstall)
	t.Logf("Whitelist (length): %d", len(config.Whitelist))
	t.Logf("PortalState:")
	t.Logf("  SignedOn: %t", config.PortalState.SignedOn)
	t.Logf("  Incoming: %t", config.PortalState.Incoming)
	t.Logf("  Outgoing: %t", config.PortalState.Outgoing)
	t.Logf("  Communication: %s", config.PortalState.Communication)
	t.Logf("APIVersion: %s", config.APIVersion)
	t.Logf("SwVersion: %s", config.SwVersion)
	t.Logf("ProxyAddress: %s", config.ProxyAddress)
	t.Logf("ProxyPort: %d", config.ProxyPort)
	t.Logf("LinkButton: %t", config.LinkButton)
	t.Logf("IPAddress: %s", config.IPAddress)
	t.Logf("Mac: %s", config.Mac)
	t.Logf("NetMask: %s", config.NetMask)
	t.Logf("Gateway: %s", config.Gateway)
	t.Logf("Dhcp: %t", config.Dhcp)
	t.Logf("PortalServices: %t", config.PortalServices)
	t.Logf("UTC: %s", config.UTC)
	t.Logf("LocalTime: %s", config.LocalTime)
	t.Logf("TimeZone: %s", config.TimeZone)
	t.Logf("ZigbeeChannel: %d", config.ZigbeeChannel)
	t.Logf("ModelID: %s", config.ModelID)
	t.Logf("BridgeID: %s", config.BridgeID)
	t.Logf("FactoryNew: %t", config.FactoryNew)
	t.Logf("ReplacesBridgeID: %s", config.ReplacesBridgeID)
	t.Logf("DatastoreVersion: %s", config.DatastoreVersion)
	t.Logf("StarterKitID: %s", config.StarterKitID)
}

func TestGetConfigError(t *testing.T) {
	b := New(badHostname, username)
	_, err := b.GetConfig()
	if err == nil {
		t.Fatal("Expected error not to be nil")
	}
}

func TestCreateUser(t *testing.T) {
	b := New(hostname, "")
	u, err := b.CreateUser("github.com/amimof/huego#go test")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("User created with username: %s", u)
	}
}

func TestCreateUserError(t *testing.T) {
	b := New(badHostname, username)
	_, err := b.CreateUser("github.com/amimof/huego#go test")
	if err == nil {
		t.Fatal("Expected error not to be nil")
	}
}

func TestCreateUser2(t *testing.T) {
	b := New(hostname, "")
	u, err := b.CreateUser2("github.com/amimof/huego#go test", true)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("User created with username: %s", u)
	}
}

func TestCreateUser2Error(t *testing.T) {
	b := New(badHostname, username)
	_, err := b.CreateUser2("github.com/amimof/huego#go test", true)
	if err == nil {
		t.Fatal("Expected error not to be nil")
	}
}

func TestGetUsers(t *testing.T) {
	b := New(hostname, username)
	users, err := b.GetUsers()
	if err != nil {
		t.Fatal(err)
	}
	for i, u := range users {
		t.Logf("%d:", i)
		t.Logf("  Name: %s", u.Name)
		t.Logf("  Username: %s", u.Username)
		t.Logf("  CreateDate: %s", u.CreateDate)
		t.Logf("  LastUseDate: %s", u.LastUseDate)
	}
	contains := func(name string, ss []Whitelist) bool {
		for _, s := range ss {
			if s.Name == name {
				return true
			}
		}
		return false
	}

	assert.True(t, contains("PhilipsHueAndroidApp#TCTALCATELONETOU", users))
	assert.True(t, contains("MyApplication", users))
}

func TestGetUsersError(t *testing.T) {
	b := New(badHostname, username)
	_, err := b.GetUsers()
	if err == nil {
		t.Fatal("Expected error not to be nil")
	}
}

func TestDeleteUser(t *testing.T) {
	b := New(hostname, username)
	err := b.DeleteUser("ffffffffe0341b1b376a2389376a2389")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Deleted user '%s'", "ffffffffe0341b1b376a2389376a2389")
}

func TestGetFullState(t *testing.T) {
	b := New(hostname, username)
	_, err := b.GetFullState()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFullStateError(t *testing.T) {
	b := New(badHostname, username)
	_, err := b.GetFullState()
	if err == nil {
		t.Fatal("Expected error not to be nil")
	}
}
