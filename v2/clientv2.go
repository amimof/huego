package huego

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Response represents an API response returned by a bridge
type APIResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []APIError      `json:"errors,omitempty"`
}

// APIError represents an error returned in a response from a bridge
type APIError struct {
	Description string `json:"description"`
}

// ClientV2 is a construct for interacting with the Hue API V2
type ClientV2 struct {
	Clip *CLIPClient
}

// Into
func (a *APIResponse) Into(obj interface{}) error {
	return json.Unmarshal(a.Data, obj)
}

// Error implements the error type
func (a APIError) Error() string {
	return a.Description
}

// SetClient can be used to set a custom http.Client that the ClientV2 uses for http connections
func (c *ClientV2) SetClient(client *http.Client) {
	c.Clip.SetClient(client)
}

// GetLights returns an array of light resources by using an empty context with GetLightsContext
func (c *ClientV2) GetLights() ([]*Light, error) {
	return c.GetLightsContext(context.Background())
}

// GetLightsContext accepts a context and returns an array of light resources
func (c *ClientV2) GetLightsContext(ctx context.Context) ([]*Light, error) {
	res, err :=
		NewRequest(c.Clip).
			Verb(http.MethodGet).
			Resource(TypeLight).
			OnError(errorHandler).
			Do(ctx)
	if err != nil {
		return nil, err
	}
	var l []*Light
	return l, unmarshal(res.BodyRaw, &l)
}

// GetLight returns a light resource by ID using an empty context with GetLightContext
func (c *ClientV2) GetLight(id string) (*Light, error) {
	return c.GetLightContext(context.Background(), id)
}

// GetLightContext returns a light resource by ID using the provided context
func (c *ClientV2) GetLightContext(ctx context.Context, id string) (*Light, error) {
	res, err :=
		NewRequest(c.Clip).
			Verb(http.MethodGet).
			Resource(TypeLight).
			OnError(errorHandler).
			ID(id).
			Do(ctx)
	if err != nil {
		return nil, err
	}
	var l []*Light
	err = unmarshal(res.BodyRaw, &l)
	if err != nil {
		return nil, err
	}
	if len(l) <= 0 {
		return nil, fmt.Errorf("light %s not found", id)
	}
	return l[0], nil
}

// SetLight updates a light by id using properties in the given light paramter.
func (c *ClientV2) SetLight(id string, light *Light) (*Response, error) {
	return c.SetLightContext(context.Background(), id, light)
}

// SetLightContext updates a light by id using properties in the given light paramter.
func (c *ClientV2) SetLightContext(ctx context.Context, id string, light *Light) (*Response, error) {
	l := light
	res, err :=
		NewRequest(c.Clip).
			Verb(http.MethodPut).
			Resource(TypeLight).
			OnError(errorHandler).
			ID(id).
			Body(l.Raw()).
			Do(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ClientV2) GetScenesContext(ctx context.Context) ([]*Scene, error) {
	res, err :=
		NewRequest(c.Clip).
			Verb(http.MethodGet).
			Resource(TypeScene).
			OnError(errorHandler).
			Do(ctx)
	fmt.Println(res.Response.Request.URL.String())
	//fmt.Println(string(res.BodyRaw))
	if err != nil {
		return nil, err
	}
	var s []*Scene
	return s, unmarshal(res.BodyRaw, &s)
}

func (c *ClientV2) GetScenes() ([]*Scene, error) {
	return c.GetScenesContext(context.Background())
}

func (c *ClientV2) GetSceneContext(ctx context.Context, id string) (*Scene, error) {
	res, err :=
		NewRequest(c.Clip).
			Verb(http.MethodGet).
			Resource(TypeScene).
			OnError(errorHandler).
			ID(id).
			Do(ctx)
	if err != nil {
		return nil, err
	}
	var s []*Scene
	err = unmarshal(res.BodyRaw, &s)
	if err != nil {
		return nil, err
	}
	if len(s) <= 0 {
		return nil, fmt.Errorf("scene %s not found", id)
	}
	return s[0], nil
}

func (c *ClientV2) GetScene(id string) (*Scene, error) {
	return c.GetSceneContext(context.Background(), id)
}

// SetSceneContext updates a scene by id using properties in the given scene paramter.
func (c *ClientV2) SetSceneContext(ctx context.Context, id string, scene *Scene) (*Response, error) {
	s := scene
	res, err :=
		NewRequest(c.Clip).
			Verb(http.MethodPut).
			Resource(TypeScene).
			OnError(errorHandler).
			ID(id).
			Body(s.Raw()).
			Do(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ClientV2) GetRoomsContext(ctx context.Context) ([]*Room, error) {
	res, err :=
		NewRequest(c.Clip).
			Verb(http.MethodGet).
			Resource(TypeRoom).
			OnError(errorHandler).
			Do(ctx)
	fmt.Println(res.Response.Request.URL.String())
	//fmt.Println(string(res.BodyRaw))
	if err != nil {
		return nil, err
	}
	var s []*Room
	return s, unmarshal(res.BodyRaw, &s)
}

func (c *ClientV2) GetRooms() ([]*Room, error) {
	return c.GetRoomsContext(context.Background())
}

func (c *ClientV2) GetRoomContext(ctx context.Context, id string) (*Room, error) {
	res, err :=
		NewRequest(c.Clip).
			Verb(http.MethodGet).
			Resource(TypeRoom).
			OnError(errorHandler).
			ID(id).
			Do(ctx)
	if err != nil {
		return nil, err
	}
	var s []*Room
	err = unmarshal(res.BodyRaw, &s)
	if err != nil {
		return nil, err
	}
	if len(s) <= 0 {
		return nil, fmt.Errorf("room %s not found", id)
	}
	return s[0], nil
}

func (c *ClientV2) GetRoom(id string) (*Room, error) {
	return c.GetRoomContext(context.Background(), id)
}

// SetRoomContext updates a Room by id using properties in the given Room paramter.
func (c *ClientV2) SetRoomContext(ctx context.Context, id string, room *Room) (*Response, error) {
	s := room
	res, err :=
		NewRequest(c.Clip).
			Verb(http.MethodPut).
			Resource(TypeRoom).
			OnError(errorHandler).
			ID(id).
			Body(s.Raw()).
			Do(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ClientV2) GetZonesContext(ctx context.Context) ([]*Zone, error) {
	res, err :=
		NewRequest(c.Clip).
			Verb(http.MethodGet).
			Resource(TypeZone).
			OnError(errorHandler).
			Do(ctx)
	if err != nil {
		return nil, err
	}
	var s []*Zone
	return s, unmarshal(res.BodyRaw, &s)
}

func (c *ClientV2) GetZones() ([]*Zone, error) {
	return c.GetZonesContext(context.Background())
}

func (c *ClientV2) GetZoneContext(ctx context.Context, id string) (*Zone, error) {
	res, err :=
		NewRequest(c.Clip).
			Verb(http.MethodGet).
			Resource(TypeZone).
			OnError(errorHandler).
			ID(id).
			Do(ctx)
	if err != nil {
		return nil, err
	}
	var s []*Zone
	err = unmarshal(res.BodyRaw, &s)
	if err != nil {
		return nil, err
	}
	if len(s) <= 0 {
		return nil, fmt.Errorf("zone %s not found", id)
	}
	return s[0], nil
}

func (c *ClientV2) GetZone(id string) (*Zone, error) {
	return c.GetZoneContext(context.Background(), id)
}

// SetZoneContext updates a Zone by id using properties in the given Zone paramter.
func (c *ClientV2) SetZoneContext(ctx context.Context, id string, zone *Zone) (*Response, error) {
	s := zone
	res, err :=
		NewRequest(c.Clip).
			Verb(http.MethodPut).
			Resource(TypeZone).
			OnError(errorHandler).
			ID(id).
			Body(s.Raw()).
			Do(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}