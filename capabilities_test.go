package huego_test

import (
	"github.com/amimof/huego"
	"os"
	"testing"
)

func TestGetCapabilities(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	c, err := hue.GetCapabilities()
	if err != nil {
		t.Fatal(c)
	}
	t.Log("Capabilities:")
	t.Log("  Groups")
	t.Logf("    Available: %d", c.Groups.Available)
	t.Log("  Lights")
	t.Logf("    Available: %d", c.Lights.Available)
	t.Log("  Resourcelinks")
	t.Logf("    Available: %d", c.Resourcelinks.Available)
	t.Log("  Schedules")
	t.Logf("    Available: %d", c.Schedules.Available)
	t.Log("  Rules")
	t.Logf("    Available: %d", c.Rules.Available)
	t.Log("  Scenes")
	t.Logf("    Available: %d", c.Scenes.Available)
	t.Log("  Sensors")
	t.Logf("    Available: %d", c.Sensors.Available)
	t.Log("  Streaming")
	t.Logf("    Available: %d", c.Streaming.Available)
}
