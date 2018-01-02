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
	t.Logf("Name: %s", config.Name)
	t.Logf("SwUpdate:")
	t.Logf("  CheckForUpdate: %t", config.SwUpdate.CheckForUpdate)
	t.Logf("  DeviceTypes:")
	t.Logf("    Bridge: %t", config.SwUpdate.DeviceTypes.Bridge)
	t.Logf("    Lights (length): %d", len(config.SwUpdate.DeviceTypes.Lights))
	t.Logf("    Sensors (length): %d", len(config.SwUpdate.DeviceTypes.Sensors))
	t.Logf("  UpdateState: %d", config.SwUpdate.UpdateState)
	t.Logf("  Notify: %t", config.SwUpdate.Notify)
	t.Logf("  Url: %s", config.SwUpdate.Url)
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
	t.Logf("ApiVersion: %s", config.ApiVersion)
	t.Logf("SwVersion: %s", config.SwVersion)
	t.Logf("ProxyAddress: %s", config.ProxyAddress)
	t.Logf("ProxyPort: %d", config.ProxyPort)
	t.Logf("LinkButton: %t", config.LinkButton)
	t.Logf("IpAddress: %s", config.IpAddress)
	t.Logf("Mac: %s", config.Mac)
	t.Logf("NetMask: %s", config.NetMask)
	t.Logf("Gateway: %s", config.Gateway)
	t.Logf("Dhcp: %t", config.Dhcp)
	t.Logf("PortalServices: %t", config.PortalServices)
	t.Logf("UTC: %s", config.UTC)
	t.Logf("LocalTime: %s", config.LocalTime)
	t.Logf("TimeZone: %s", config.TimeZone)
	t.Logf("ZigbeeChannel: %d", config.ZigbeeChannel)
	t.Logf("ModelId: %s", config.ModelId)
	t.Logf("BridgeId: %s", config.BridgeId)
	t.Logf("FactoryNew: %t", config.FactoryNew)
	t.Logf("ReplacesBridgeId: %s", config.ReplacesBridgeId)
	t.Logf("DatastoreVersion: %s", config.DatastoreVersion)
	t.Logf("StarterKitId: %s", config.StarterKitId)
}

func TestCreateUser(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), "")
	u, err := hue.CreateUser("huego#tests")
	if err != nil {
		t.Error(err)
	}
	t.Logf("User created with username: %s", u)
}

func TestGetUsers(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	users, err := hue.GetUsers()
	if err != nil {
		t.Error(err)
	}
	for i, u := range users {
		t.Logf("%d:", i)
		t.Logf("  Name: %s", u.Name)
		t.Logf("  Username: %s", u.Username)
		t.Logf("  CreateDate: %s", u.CreateDate)
		t.Logf("  LastUseDate: %s", u.LastUseDate)
	}
}

func TestDeleteUser(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	uid := "Ot82OcV7FWBl1kHXOWCTY5znQXk9WpNkNDZIGYQX"
	err := hue.DeleteUser(uid)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Delete user %s", uid)
	}
}