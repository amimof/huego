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
	light := l[0]
	return light, nil
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
