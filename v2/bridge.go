package huego

// Bridge is a resource of type bridge: https://developers.meethue.com/develop/hue-api-v2/api-reference/#resource_bridge_get
type Bridge struct {
	BridgeID *string   `json:"bridge_id"`
	TimeZone *TimeZone `json:"time_zone"`
	BaseResource
}

// TimeZone is the timezone properties of a bridge type
type TimeZone struct {
	TimeZone string `json:"time_zone"`
}
