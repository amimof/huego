package huego

import (
	"context"
	"net/http"
	"fmt"
	"encoding/json"
)

// Response represents an API response returned by a bridge
type APIResponse struct {
	Data    json.RawMessage `json:"data"`
	Errors  []APIError      `json:"errors,omitempty"`
}

// APIError represents an error returned in a response from a bridge
type APIError struct {
	Description string `json:"description"`
}

type ClientV2 struct {
	Clip *CLIPClient
}

func (a *APIResponse) Into(obj interface{}) error {
	return json.Unmarshal(a.Data, obj)
}

func (a APIError) Error() string {
	return a.Description
}

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
	light.client = c
	return light, nil
}

func (c *ClientV2) SetLightContext(ctx context.Context, id string, light *Light) (*Response, error) {
	res, err := 
		NewRequest(c.Clip).
			Verb(http.MethodPut).
			Resource(TypeLight).
			OnError(errorHandler).
			ID(id).
			Body(light.Raw()).
			Do(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}