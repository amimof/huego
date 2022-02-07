package huego

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	// TypeLight is a Hue resource of type light
	TypeLight = "light"
	// TypeScene is a Hue resource of type scene
	TypeScene = "scene"
	// TypeRoom is a Hue resource of type room
	TypeRoom = "room"
	// TypeZone is a Hue resource of type zone
	TypeZone = "zone"
	// TypeBridge is a Hue resource of type bridge
	TypeBridge = "bridge"
)

// DiscoveredBridge is a type i used for discovering bridges
type DiscoveredBridge struct {
	ID                string `json:"id"`
	InternalIPAddress string `json:"internalipaddress"`
	Port              int    `json:"port"`
}

// Discover uses DiscoverAll but returns the first bridge if any. Returns an error if no bridges are found
func Discover() (*DiscoveredBridge, error) {
	b, err := DiscoverAll()
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		return nil, fmt.Errorf("no bridges found during discovery")
	}
	return &b[0], nil
}

// DiscoverAll returns many discovered bridges
func DiscoverAll() ([]DiscoveredBridge, error) {
	c, err := NewClient("https://discovery.meethue.com", "")
	if err != nil {
		return nil, err
	}
	res, err := NewRequest(c).
		Path("/").
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	var discovered []DiscoveredBridge
	err = json.Unmarshal(res.BodyRaw, &discovered)
	if err != nil {
		return nil, err
	}

	return discovered, nil
}
